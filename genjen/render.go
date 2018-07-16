package main

import (
	"io"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func render(w io.Writer) error {
	file := NewFile("jen")

	file.HeaderComment("This file is generated - do not edit.")
	file.Line()

	for _, b := range groups {
		b := b // b used in closures
		comment := Commentf("%s %s", b.name, b.comment)

		if b.variadic && len(b.parameters) > 1 {
			panic("should not have variadic function with multiple params")
		}

		var variadic Code
		if b.variadic {
			variadic = Op("...")
		}
		var funcParams []Code
		var callParams []Code
		for _, name := range b.parameters {
			funcParams = append(funcParams, Id(name).Add(variadic).Id("Code"))
			callParams = append(callParams, Id(name).Add(variadic))
		}

		addFunctionAndGroupMethod(
			file,
			b.name,
			comment,
			funcParams,
			callParams,
			false,
		)

		/*
			// <comment>
			func (s *Statement) <name>(<funcParams>) *Statement {
				g := &Group{
					items:     []Code{<paramNames>}|<paramNames[0]>,
					name:      "<name>",
					open:      "<opening>",
					close:     "<closing>",
					separator: "<separator>",
					multi:     <multi>,
				}
				*s = append(*s, g)
				return s
			}
		*/
		file.Add(comment)
		file.Func().Params(
			Id("s").Op("*").Id("Statement"),
		).Id(b.name).Params(
			funcParams...,
		).Op("*").Id("Statement").Block(
			Id("g").Op(":=").Op("&").Id("Group").Values(Dict{
				Id("items"): Do(func(s *Statement) {
					if b.variadic {
						s.Id(b.parameters[0])
					} else {
						s.Index().Id("Code").ValuesFunc(func(g *Group) {
							for _, name := range b.parameters {
								g.Id(name)
							}
						})
					}
				}),
				Id("name"):      Lit(strings.ToLower(b.name)),
				Id("open"):      Lit(b.opening),
				Id("close"):     Lit(b.closing),
				Id("separator"): Lit(b.separator),
				Id("multi"):     Lit(b.multi),
			}),
			Op("*").Id("s").Op("=").Append(Op("*").Id("s"), Id("g")),
			Return(Id("s")),
		)

		if b.variadic && !b.preventFunc {

			funcName := b.name + "Func"
			funcComment := Commentf("%sFunc %s", b.name, b.comment)
			funcFuncParams := []Code{Id("f").Func().Params(Op("*").Id("Group"))}
			funcCallParams := []Code{Id("f")}

			addFunctionAndGroupMethod(
				file,
				funcName,
				funcComment,
				funcFuncParams,
				funcCallParams,
				false,
			)

			/*
				// <funcComment>
				func (s *Statement) <funcName>(f func(*Group)) *Statement {
					g := &Group{
						name:      "<name>",
						open:      "<opening>",
						close:     "<closing>",
						separator: "<separator>",
						multi:     <multi>,
					}
					f(g)
					*s = append(*s, g)
					return s
				}
			*/
			file.Add(funcComment)
			file.Func().Params(
				Id("s").Op("*").Id("Statement"),
			).Id(funcName).Params(
				funcFuncParams...,
			).Op("*").Id("Statement").Block(
				Id("g").Op(":=").Op("&").Id("Group").Values(Dict{
					Id("name"):      Lit(strings.ToLower(b.name)),
					Id("open"):      Lit(b.opening),
					Id("close"):     Lit(b.closing),
					Id("separator"): Lit(b.separator),
					Id("multi"):     Lit(b.multi),
				}),
				Id("f").Call(Id("g")),
				Op("*").Id("s").Op("=").Append(Op("*").Id("s"), Id("g")),
				Return(Id("s")),
			)
		}
	}

	type tkn struct {
		token     string
		name      string
		tokenType string
		tokenDesc string
	}
	tokens := []tkn{}
	for _, v := range identifiers {
		tokens = append(tokens, tkn{
			token:     v,
			name:      strings.ToUpper(v[:1]) + v[1:],
			tokenType: "identifierToken",
			tokenDesc: "identifier",
		})
	}
	for _, v := range keywords {
		tokens = append(tokens, tkn{
			token:     v,
			name:      strings.ToUpper(v[:1]) + v[1:],
			tokenType: "keywordToken",
			tokenDesc: "keyword",
		})
	}

	for i, t := range tokens {
		t := t // used in closures
		comment := Commentf(
			"%s renders the %s %s.",
			t.name,
			t.token,
			t.tokenDesc,
		)
		addFunctionAndGroupMethod(
			file,
			t.name,
			comment,
			nil,
			nil,
			i != 0, // only enforce test coverage on one item
		)

		/*
			// <comment>
			func (s *Statement) <name>() *Statement {
				t := token{
					typ:     <tokenType>,
					content: "<token>",
				}
				*s = append(*s, t)
				return s
			}
		*/
		file.Add(comment)
		file.Func().Params(
			Id("s").Op("*").Id("Statement"),
		).Id(t.name).Params().Op("*").Id("Statement").Block(
			Do(func(s *Statement) {
				if i != 0 {
					// only enforce test coverage on one item
					s.Comment("notest")
				}
			}),
			Id("t").Op(":=").Id("token").Values(Dict{
				Id("typ"):     Id(t.tokenType),
				Id("content"): Lit(t.token),
			}),
			Op("*").Id("s").Op("=").Append(Op("*").Id("s"), Id("t")),
			Return(Id("s")),
		)
	}

	return file.Render(w)
}

// For each method on *Statement, this generates a package level
// function and a method on *Group, both with the same name.
func addFunctionAndGroupMethod(
	file *File,
	name string,
	comment *Statement,
	funcParams []Code,
	callParams []Code,
	notest bool,
) {
	/*
		// <comment>
		func <name>(<funcParams>) *Statement {
			return newStatement().<name>(<callParams>)
		}
	*/
	file.Add(comment)
	file.Func().Id(name).Params(funcParams...).Op("*").Id("Statement").Block(
		Do(func(s *Statement) {
			if notest {
				// only enforce test coverage on one item
				s.Comment("notest")
			}
		}),
		Return(Id("newStatement").Call().Dot(name).Call(callParams...)),
	)
	/*
		// <comment>
		func (g *Group) <name>(<funcParams>) *Statement {
			s := <name>(<callParams>)
			g.items = append(g.items, s)
			return s
		}
	*/
	file.Add(comment)
	file.Func().Params(
		Id("g").Op("*").Id("Group"),
	).Id(name).Params(funcParams...).Op("*").Id("Statement").Block(
		Do(func(s *Statement) {
			if notest {
				// only enforce test coverage on one item
				s.Comment("notest")
			}
		}),
		Id("s").Op(":=").Id(name).Params(callParams...),
		Id("g").Dot("items").Op("=").Append(Id("g").Dot("items"), Id("s")),
		Return(Id("s")),
	)
}
