package main

import (
	"context"
	"os"

	. "github.com/davelondon/jennifer/jen"
	"github.com/davelondon/jennifer/jen/data"
)

func main() {

	file := NewFile()
	for _, b := range data.Blocks {
		b := b // b used in closures
		/*
			// {{ .Name }} inserts {{ .Desc }}
			func {{ .Name }}(c ...Code) *Statement {
				s := new(Statement)
				return s.{{ .Name }}(c...)
			}
		*/
		comment := Commentf("%s inserts %s", b.Name, b.Desc)
		file.Add(comment)
		file.Func().Id(b.Name).Params(
			Id("c").Vari().Id("Code"),
		).Ptr().Id("Statement").Block(
			Id("s").Sas().New(Id("Statement")),
			Return().Id("s").Method(b.Name, Id("c").Vari()),
		)

		/*
			// {{ .Name }} inserts {{ .Desc }}
			func (l *StatementList) {{ .Name }}(c ...Code) *Statement {
				s := {{ .Name }}(c...)
				*l = append(*l, s)
				return s
			}
		*/
		file.Add(comment)
		file.Func().Params(
			Id("l").Ptr().Id("StatementList"),
		).Id(b.Name).Params(
			Id("c").Vari().Id("Code"),
		).Ptr().Id("Statement").Block(
			Id("s").Sas().Id(b.Name).Call(Id("c").Vari()),
			Ptr().Id("l").As().Append(Ptr().Id("l"), Id("s")),
			Return().Id("s"),
		)

		/*
			// {{ .Name }} inserts {{ .Desc }}
			func (s *Statement) {{ .Name }}(c ...Code) *Statement {
				b := block{
					Statement: s,
					code:      c,
					{{- if ne .Open "" }}
					open:      "{{ .Open }}",
					{{- end -}}
					{{- if ne .Close "" }}
					close:     "{{ .Close }}",
					{{- end -}}
					{{- if ne .Seperator "" }}
					seperator: "{{ .Seperator }}",
					{{- end }}
				}
				*s = append(*s, b)
				return s
			}
		*/
		file.Add(comment)
		file.Func().Params(
			Id("l").Ptr().Id("Statement"),
		).Id(b.Name).Params(
			Id("c").Vari().Id("Code"),
		).Ptr().Id("Statement").Block(
			Id("b").Sas().Id("block").MapLitFunc(func(m map[Code]Code) {
				m[Id("Statement")] = Id("s")
				m[Id("code")] = Id("c")
				if b.Open != "" {
					m[Id("open")] = Lit(b.Open)
				}
				if b.Close != "" {
					m[Id("close")] = Lit(b.Close)
				}
				if b.Seperator != "" {
					m[Id("seperator")] = Lit(b.Seperator)
				}
			}),
			Ptr().Id("s").As().Append(Ptr().Id("s"), Id("b")),
			Return().Id("s"),
		)

	}
	/*
		{{ range .Identifiers }}
		func {{ . | capital }}() *Statement {
			s := new(Statement)
			return s.{{ . | capital }}()
		}

		func (l *StatementList) {{ . | capital }}() *Statement {
			s := {{ . | capital }}()
			*l = append(*l, s)
			return s
		}

		func (s *Statement) {{ . | capital }}() *Statement {
			t := Token{
				Statement: s,
				typ:       identifierToken,
				content:   "{{ . }}",
			}
			*s = append(*s, t)
			return s
		}
		{{ end }}

		{{ range .Functions }}
		func {{ . | capital }}(c ...Code) *Statement {
			s := new(Statement)
			return s.{{ . | capital }}(c...)
		}

		func (l *StatementList) {{ . | capital }}(c ...Code) *Statement {
			s := {{ . | capital }}(c...)
			*l = append(*l, s)
			return s
		}

		func (s *Statement) {{ . | capital }}(c ...Code) *Statement {
			t := Token{
				Statement: s,
				typ:       identifierToken,
				content:   "{{ . }}",
			}
			ca := Call(c...)
			*s = append(*s, t, ca)
			return s
		}
		{{ end }}

		{{ range .Keywords }}
		func {{ . | capital }}() *Statement {
			s := new(Statement)
			return s.{{ . | capital }}()
		}

		func (l *StatementList) {{ . | capital }}() *Statement {
			s := {{ . | capital }}()
			*l = append(*l, s)
			return s
		}

		func (s *Statement) {{ . | capital }}() *Statement {
			t := Token{
				Statement: s,
				typ:       keywordToken,
				content:   "{{ . }}",
			}
			*s = append(*s, t)
			return s
		}

		{{ end }}

		{{ range .Operators }}
		// {{ .Name }} inserts the {{ .Desc }} operator ({{ .Op }})
		func {{ .Name }}() *Statement {
			s := new(Statement)
			return s.{{ .Name }}()
		}

		// {{ .Name }} inserts the {{ .Desc }} operator ({{ .Op }})
		func (l *StatementList) {{ .Name }}() *Statement {
			s := {{ .Name }}()
			*l = append(*l, s)
			return s
		}

		// {{ .Name }} inserts the {{ .Desc }} operator ({{ .Op }})
		func (s *Statement) {{ .Name }}() *Statement {
			t := Token{
				Statement: s,
				typ:       operatorToken,
				content:   "{{ .Op }}",
			}
			*s = append(*s, t)
			return s
		}

		{{ end }}
	*/
	/*
		for _, keyword := range data.Keywords {
			name := strings.ToUpper(keyword[:1]) + keyword[1:]
			// func Foo(c ...Code) *Statement {
			file.Comment("Foo")
			file.Func().Id(name).Params(
				Id("c").Vari().Id("Code"),
			).Ptr().Id("Statement").Block(
				// s := new(Statement)
				Id("s").Sas().New(Id("Statement")),
				// return s.Foo(c...)
				Return().Id("s").Method(name, Id("c").Vari()),
			)
		}*/

	ctx := Context(context.Background(), "jen")
	err := Render(ctx, file, os.Stdout)
	if err != nil {
		panic(err)
	}

}
