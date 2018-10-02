package jen

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/printer"
	asttoken "go/token"
	"strings"
)

func NewFileFromAst(file *ast.File, set *asttoken.FileSet) (*File, error) {
	if file.Name == nil {
		return nil, errors.New("empty file name")
	}
	packageName := file.Name.Name
	imports, hints := fillImportsAndHints(file)
	f := &File{
		Group: &Group{
			multi: true,
		},
		name:    packageName,
		imports: imports,
		hints:   hints,
	}
	addPackageComments(f, file)
	removeImportsFromAstFile(file)
	tail, err := sourceTail(file, set)
	if err != nil {
		return nil, fmt.Errorf("get source tail: %v", err)
	}
	f.addRaw(tail)
	return f, nil
}

func fillImportsAndHints(file *ast.File) (imports, hints map[string]importdef) {
	imports = make(map[string]importdef)
	hints = make(map[string]importdef)
	for i := range file.Imports {
		if file.Imports[i].Path == nil {
			continue // invalid package
		}
		path := strings.Trim(file.Imports[i].Path.Value, `"`)
		if file.Imports[i].Name != nil {
			imports[path] = importdef{name: file.Imports[i].Name.Name, alias: true}
		} else {
			imports[path] = importdef{name: "", alias: false}
		}
	}
	return
}

func addPackageComments(f *File, file *ast.File) {
	if file.Doc == nil {
		return
	}
	for _, c := range file.Doc.List {
		if c != nil {
			f.PackageComment(c.Text)
		}
	}
	file.Doc = nil
	return
}

func removeImportsFromAstFile(f *ast.File) {
	resDecls := make([]ast.Decl, 0, len(f.Decls))
	for i := range f.Decls {
		d, ok := f.Decls[i].(*ast.GenDecl)
		if !ok || d.Tok != asttoken.IMPORT {
			resDecls = append(resDecls, f.Decls[i])
			continue
		}
		// ignore imports decls
	}
	f.Decls = resDecls
}

// returns []byte tail of the source file: all header comments, package name, package comment will be removed.
func sourceTail(file *ast.File, set *asttoken.FileSet) ([]byte, error) {
	var b bytes.Buffer
	err := printer.Fprint(&b, set, file)
	if err != nil {
		return nil, err
	}
	raw := b.Bytes()
	raw = append(raw[:file.Package-1], raw[file.Name.End():]...)
	return raw, nil
}
