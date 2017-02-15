package rebecca

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/printer"
	"go/token"
	"regexp"
	"strconv"
	"strings"
)

func NewCodeMap(pkg string, dir string) (*CodeMap, error) {
	m := &CodeMap{
		pkg:      pkg,
		dir:      dir,
		Examples: map[string]*doc.Example{},
		Comments: map[string]string{},
	}
	if err := m.scanDir(); err != nil {
		return nil, err
	}
	return m, nil
}

type CodeMap struct {
	pkg      string
	dir      string
	fset     *token.FileSet
	Examples map[string]*doc.Example
	Comments map[string]string
}

func (m *CodeMap) ExampleFunc(plain bool) func(in string) string {
	return func(in string) string {
		e, ok := m.Examples[in]
		if !ok {
			return fmt.Sprintf("[example %s not found.]", in)
		}
		buf := &bytes.Buffer{}
		if plain {
			printer.Fprint(buf, m.fset, e.Code)
			out := buf.String()
			if strings.HasSuffix(out, "\n\n}") {
				// fix annoying line-feed before end brace
				out = out[:len(out)-2] + "}"
			}
			return out
		}
		if bs, ok := e.Code.(*ast.BlockStmt); ok {
			for _, s := range bs.List {
				printer.Fprint(buf, m.fset, s)
				buf.WriteString("\n")
			}
		} else {
			printer.Fprint(buf, m.fset, e.Code)
		}
		quotes := "```"
		return fmt.Sprintf(`[Example](https://godoc.org/%s#example-%s):
%sgo
%s
// Output:
// %s
%s`,
			m.pkg,
			strings.Replace(in, "_", "-", -1)[len("Example"):],
			quotes,
			strings.Trim(buf.String(), "\n"),
			strings.Replace(strings.Trim(e.Output, "\n"), "\n", "\n// ", -1),
			quotes)
	}
}

func (m *CodeMap) LinkFunc(in string) string {
	return fmt.Sprintf(
		"[Example](https://godoc.org/%s#example-%s):",
		m.pkg,
		strings.Replace(in, "_", "-", -1)[len("Example"):])
}

func (m *CodeMap) OutputFunc(in string) string {
	e, ok := m.Examples[in]
	if !ok {
		return fmt.Sprintf("[example %s not found.]", in)
	}
	return strings.Trim(e.Output, "\n")
}

var docRegex = regexp.MustCompile(`(\w+)\[([0-9:, ]+)\]`)
var sectionRegex = regexp.MustCompile(`(\d+)(:?)(\d*)`)

func (m *CodeMap) DocFunc(in string) string {

	if matches := docRegex.FindStringSubmatch(in); matches != nil {
		id := matches[1]
		c, ok := m.Comments[id]
		if !ok {
			return fmt.Sprintf("[comment %s not found in.]", id, in)
		}
		sentances := strings.Split(c, ".")
		sections := strings.Split(strings.Replace(matches[2], " ", "", -1), ",")
		out := ""
		for _, sectionDef := range sections {
			var start, end int
			parts := sectionRegex.FindStringSubmatch(sectionDef)
			switch {
			case parts[2] == "":
				// single sentance index, of the form "i"
				start, _ = strconv.Atoi(parts[1])
				end = start + 1
			case parts[1] == "" && parts[2] == ":":
				// of the form: "-i"
				start = 0
				end, _ = strconv.Atoi(parts[2])
			case parts[2] == ":" && parts[3] == "":
				// of the form: "i-"
				start, _ = strconv.Atoi(parts[1])
				end = len(sentances) - 1
			default:
				// of the form: "i-j"
				start, _ = strconv.Atoi(parts[1])
				end, _ = strconv.Atoi(parts[3])
			}
			if start >= end {
				return fmt.Sprintf("[start must be > end in %s]", in)
			}
			if end >= len(sentances) {
				return fmt.Sprintf("[only %d sentances in comment %s]", len(sentances), in)
			}
			for _, s := range sentances[start:end] {
				if s != "\n" {
					out += s + "."
				}
			}
		}
		return out
	}

	c, ok := m.Comments[in]
	if !ok {
		return fmt.Sprintf("[comment %s not found.]", in)
	}
	return strings.Trim(c, "\n")
}

func (m *CodeMap) scanTests(name string, p *ast.Package) error {
	for _, f := range p.Files {
		examples := doc.Examples(f)
		for _, ex := range examples {
			m.Examples["Example"+ex.Name] = ex
		}
	}
	return nil
}

func (m *CodeMap) scanPkg(name string, p *ast.Package) error {
	for _, f := range p.Files {
		for _, d := range f.Decls {
			switch d := d.(type) {
			case *ast.FuncDecl:
				if d.Doc.Text() == "" {
					continue
				}
				if d.Recv == nil {
					// function
					//fmt.Println(d.Name, d.Doc.Text())
					name := fmt.Sprint(d.Name)
					m.Comments[name] = d.Doc.Text()
				} else {
					// method
					e := d.Recv.List[0].Type
					if se, ok := e.(*ast.StarExpr); ok {
						// if the method receiver has a *, discard it.
						e = se.X
					}
					b := &bytes.Buffer{}
					printer.Fprint(b, m.fset, e)
					//fmt.Printf("%s.%s %s", b.String(), d.Name, d.Doc.Text())
					name := fmt.Sprintf("%s.%s", b.String(), d.Name)
					m.Comments[name] = d.Doc.Text()
				}
			case *ast.GenDecl:
				if d.Doc.Text() == "" {
					continue
				}
				switch s := d.Specs[0].(type) {
				case *ast.TypeSpec:
					//fmt.Println(s.Name, d.Doc.Text())
					name := fmt.Sprint(s.Name)
					m.Comments[name] = d.Doc.Text()
					if t, ok := s.Type.(*ast.StructType); ok {
						for _, f := range t.Fields.List {
							if f.Doc.Text() == "" {
								continue
							}
							if f.Names[0].IsExported() {
								fieldName := fmt.Sprint(name, ".", f.Names[0])
								m.Comments[fieldName] = f.Doc.Text()
							}
						}
					}
				case *ast.ValueSpec:
					//fmt.Println(s.Names[0], d.Doc.Text())
					if len(s.Names) == 0 {
						continue
					}
					name := fmt.Sprint(s.Names[0])
					m.Comments[name] = d.Doc.Text()
				}
			}
		}
	}
	return nil
}

func (m *CodeMap) scanDir() error {
	// Create the AST by parsing src.
	m.fset = token.NewFileSet() // positions are relative to fset
	pkgs, err := parser.ParseDir(m.fset, m.dir, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	for name, p := range pkgs {
		if strings.HasSuffix(name, "_test") {
			if err := m.scanTests(name, p); err != nil {
				return err
			}
		}
		if err := m.scanPkg(name, p); err != nil {
			return err
		}
	}

	return nil
}
