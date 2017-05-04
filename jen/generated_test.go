package jen_test

import (
	"testing"

	. "github.com/dave/jennifer/jen"
)

var gencases = []tc{
	{
		desc: `bool group`,
		code: BlockFunc(func(g *Group) {
			g.Bool()
		}),
		expect: `{
		bool
		}`,
	},
	{
		desc:   `recover func`,
		code:   Recover(),
		expect: `recover()`,
	},
	{
		desc:   `recover statement`,
		code:   Null().Recover(),
		expect: `recover()`,
	},
	{
		desc: `recover group`,
		code: BlockFunc(func(g *Group) {
			g.Recover()
		}),
		expect: `{
		recover()
		}`,
	},
	{
		desc:   `real func`,
		code:   Real(Id("a")),
		expect: `real(a)`,
	},
	{
		desc:   `real statement`,
		code:   Null().Real(Id("a")),
		expect: `real(a)`,
	},
	{
		desc: `real group`,
		code: BlockFunc(func(g *Group) {
			g.Real(Id("a"))
		}),
		expect: `{
		real(a)
		}`,
	},
	{
		desc:   `printlnfunc func`,
		code:   PrintlnFunc(func(g *Group) { g.Id("a") }),
		expect: `println(a)`,
	},
	{
		desc:   `printlnfunc statement`,
		code:   Null().PrintlnFunc(func(g *Group) { g.Id("a") }),
		expect: `println(a)`,
	},
	{
		desc: `printlnfunc group`,
		code: BlockFunc(func(bg *Group) {
			bg.PrintlnFunc(func(pg *Group) { pg.Id("a") })
		}),
		expect: `{
		println(a)
		}`,
	},
	{
		desc:   `println func`,
		code:   Println(Id("a")),
		expect: `println(a)`,
	},
	{
		desc:   `println statement`,
		code:   Null().Println(Id("a")),
		expect: `println(a)`,
	},
	{
		desc: `println group`,
		code: BlockFunc(func(g *Group) {
			g.Println(Id("a"))
		}),
		expect: `{
		println(a)
		}`,
	},
	{
		desc:   `printfunc func`,
		code:   PrintFunc(func(g *Group) { g.Id("a") }),
		expect: `print(a)`,
	},
	{
		desc:   `printfunc statement`,
		code:   Null().PrintFunc(func(g *Group) { g.Id("a") }),
		expect: `print(a)`,
	},
	{
		desc: `printfunc group`,
		code: BlockFunc(func(bg *Group) {
			bg.PrintFunc(func(pg *Group) { pg.Id("a") })
		}),
		expect: `{
		print(a)
		}`,
	},
	{
		desc:   `print func`,
		code:   Print(Id("a")),
		expect: `print(a)`,
	},
	{
		desc:   `print statement`,
		code:   Null().Print(Id("a")),
		expect: `print(a)`,
	},
	{
		desc: `print group`,
		code: BlockFunc(func(g *Group) {
			g.Print(Id("a"))
		}),
		expect: `{
		print(a)
		}`,
	},
	{
		desc:   `panic func`,
		code:   Panic(Id("a")),
		expect: `panic(a)`,
	},
	{
		desc:   `panic statement`,
		code:   Null().Panic(Id("a")),
		expect: `panic(a)`,
	},
	{
		desc: `panic group`,
		code: BlockFunc(func(g *Group) {
			g.Panic(Id("a"))
		}),
		expect: `{
		panic(a)
		}`,
	},
	{
		desc:   `new func`,
		code:   New(Id("a")),
		expect: `new(a)`,
	},
	{
		desc:   `new statement`,
		code:   Id("a").Op(":=").New(Id("a")),
		expect: `a := new(a)`,
	},
	{
		desc: `new group`,
		code: BlockFunc(func(g *Group) {
			g.New(Id("a"))
		}),
		expect: `{
		new(a)
		}`,
	},
	{
		desc:   `make func`,
		code:   Make(Id("a")),
		expect: `make(a)`,
	},
	{
		desc:   `make statement`,
		code:   Id("a").Op(":=").Make(Id("a")),
		expect: `a := make(a)`,
	},
	{
		desc: `make group`,
		code: BlockFunc(func(g *Group) {
			g.Make(Id("a"))
		}),
		expect: `{
		make(a)
		}`,
	},
	{
		desc:   `len func`,
		code:   Len(Id("a")),
		expect: `len(a)`,
	},
	{
		desc:   `len statement`,
		code:   Id("a").Op(":=").Len(Id("a")),
		expect: `a := len(a)`,
	},
	{
		desc: `len group`,
		code: BlockFunc(func(g *Group) {
			g.Len(Id("a"))
		}),
		expect: `{
		len(a)
		}`,
	},
	{
		desc:   `imag func`,
		code:   Imag(Id("a")),
		expect: `imag(a)`,
	},
	{
		desc:   `imag statement`,
		code:   Id("a").Op(":=").Imag(Id("a")),
		expect: `a := imag(a)`,
	},
	{
		desc: `imag group`,
		code: BlockFunc(func(g *Group) {
			g.Imag(Id("a"))
		}),
		expect: `{
		imag(a)
		}`,
	},
	{
		desc:   `delete func`,
		code:   Delete(Id("a"), Id("b")),
		expect: `delete(a, b)`,
	},
	{
		desc:   `delete statement`,
		code:   Null().Delete(Id("a"), Id("b")),
		expect: `delete(a, b)`,
	},
	{
		desc: `delete group`,
		code: BlockFunc(func(g *Group) {
			g.Delete(Id("a"), Id("b"))
		}),
		expect: `{
		delete(a, b)
		}`,
	},
	{
		desc:   `copy func`,
		code:   Copy(Id("a"), Id("b")),
		expect: `copy(a, b)`,
	},
	{
		desc:   `copy statement`,
		code:   Id("a").Op(":=").Copy(Id("a"), Id("b")),
		expect: `a := copy(a, b)`,
	},
	{
		desc: `copy group`,
		code: BlockFunc(func(g *Group) {
			g.Copy(Id("a"), Id("b"))
		}),
		expect: `{
		copy(a, b)
		}`,
	},
	{
		desc:   `complex func`,
		code:   Complex(Id("a"), Id("b")),
		expect: `complex(a, b)`,
	},
	{
		desc:   `complex statement`,
		code:   Id("a").Op(":=").Complex(Id("a"), Id("b")),
		expect: `a := complex(a, b)`,
	},
	{
		desc: `complex group`,
		code: BlockFunc(func(g *Group) {
			g.Complex(Id("a"), Id("b"))
		}),
		expect: `{
		complex(a, b)
		}`,
	},
	{
		desc: `close group`,
		code: BlockFunc(func(g *Group) { g.Close(Id("a")) }),
		expect: `{
		close(a)
		}`,
	},
	{
		desc:   `cap func`,
		code:   Cap(Id("a")),
		expect: `cap(a)`,
	},
	{
		desc:   `cap statement`,
		code:   Id("a").Op(":=").Cap(Id("b")),
		expect: `a := cap(b)`,
	},
	{
		desc: `cap group`,
		code: BlockFunc(func(g *Group) {
			g.Cap(Id("a"))
		}),
		expect: `{
		cap(a)
		}`,
	},
	{
		desc: `append group`,
		code: BlockFunc(func(g *Group) {
			g.Append(Id("a"))
		}),
		expect: `{
		append(a)
		}`,
	},
	{
		desc:   `appendfunc statement`,
		code:   Id("a").Op("=").AppendFunc(func(ag *Group) { ag.Id("a") }),
		expect: `a = append(a)`,
	},
	{
		desc:   `appendfunc func`,
		code:   AppendFunc(func(ag *Group) { ag.Id("a") }),
		expect: `append(a)`,
	},
	{
		desc: `appendfunc group`,
		code: BlockFunc(func(bg *Group) {
			bg.AppendFunc(func(ag *Group) { ag.Id("a") })
		}),
		expect: `{
		append(a)
		}`,
	},
	{
		desc: `casefunc group`,
		code: Switch().BlockFunc(func(g *Group) {
			g.CaseFunc(func(g *Group) { g.Id("a") }).Block()
		}),
		expect: `switch {
		case a:
		}`,
	},
	{
		desc: `case group`,
		code: Switch().BlockFunc(func(g *Group) {
			g.Case(Id("a")).Block()
		}),
		expect: `switch {
		case a:
		}`,
	},
	{
		desc:   `structfunc statement`,
		code:   Id("a").Op(":=").StructFunc(func(g *Group) {}).Values(),
		expect: `a := struct{}{}`,
	},
	{
		desc: `structfunc group`,
		// Don't do this! ListFunc used to kludge Group.Struct usage
		// without syntax error.
		code:   Id("a").Op(":=").ListFunc(func(g *Group) { g.StructFunc(func(g *Group) {}) }).Values(),
		expect: `a := struct{}{}`,
	},
	{
		desc:   `structfunc func`,
		code:   Id("a").Op(":=").Add(StructFunc(func(g *Group) {})).Values(),
		expect: `a := struct{}{}`,
	},
	{
		desc: `struct group`,
		// Don't do this! ListFunc used to kludge Group.Struct usage
		// without syntax error.
		code:   Id("a").Op(":=").ListFunc(func(g *Group) { g.Struct() }).Values(),
		expect: `a := struct{}{}`,
	},
	{
		desc:   `struct func`,
		code:   Id("a").Op(":=").Add(Struct()).Values(),
		expect: `a := struct{}{}`,
	},
	{
		desc: `interfacefunc func`,
		code: Id("a").Assert(InterfaceFunc(func(g *Group) {
			g.Id("a").Call().Int()
			g.Id("b").Call().Int()
		})),
		expect: `a.(interface{
		a() int
		b() int
		})`,
	},
	{
		desc: `interfacefunc statement`,
		code: Id("a").Assert(Null().InterfaceFunc(func(g *Group) {
			g.Id("a").Call().Int()
			g.Id("b").Call().Int()
		})),
		expect: `a.(interface{
		a() int
		b() int
		})`,
	},
	{
		desc: `interfacefunc group`,
		// Don't do this! ListFunc used to kludge Group.InterfaceFunc usage
		// without syntax error.
		code: Id("a").Assert(ListFunc(func(lg *Group) {
			lg.InterfaceFunc(func(ig *Group) {
				ig.Id("a").Call().Int()
				ig.Id("b").Call().Int()
			})
		})),
		expect: `a.(interface{
		a() int
		b() int
		})`,
	},
	{
		desc:   `interface func`,
		code:   Interface().Parens(Id("a")),
		expect: `interface{}(a)`,
	},
	{
		desc: `interface group`,
		code: BlockFunc(func(g *Group) {
			g.Interface().Parens(Id("a"))
		}),
		expect: `{
		interface{}(a)
		}`,
	},
	{
		desc:   `interface statement`,
		code:   Null().Interface().Parens(Id("a")),
		expect: `interface{}(a)`,
	},
	{
		desc: `switchfunc func`,
		code: SwitchFunc(func(rg *Group) {
			rg.Id("a")
		}).Block(),
		expect: `switch a {}`,
	},
	{
		desc: `switchfunc statement`,
		code: Null().SwitchFunc(func(rg *Group) {
			rg.Id("a")
		}).Block(),
		expect: `switch a {
		}`,
	},
	{
		desc: `switchfunc group`,
		code: BlockFunc(func(bg *Group) {
			bg.SwitchFunc(func(rg *Group) {
				rg.Id("a")
			}).Block()
		}),
		expect: `{
			switch a {
			}
		}`,
	},
	{
		desc: `switch group`,
		code: BlockFunc(func(bg *Group) {
			bg.Switch().Block()
		}),
		expect: `{
			switch {
			}
		}`,
	},
	{
		desc: `forfunc func`,
		code: ForFunc(func(rg *Group) {
			rg.Id("a")
		}).Block(),
		expect: `for a {}`,
	},
	{
		desc: `forfunc statement`,
		code: Null().ForFunc(func(rg *Group) {
			rg.Id("a")
		}).Block(),
		expect: `for a {
		}`,
	},
	{
		desc: `forfunc group`,
		code: BlockFunc(func(bg *Group) {
			bg.ForFunc(func(rg *Group) {
				rg.Id("a")
			}).Block()
		}),
		expect: `{
			for a {
			}
		}`,
	},
	{
		desc: `for group`,
		code: BlockFunc(func(g *Group) {
			g.For(Id("a")).Block()
		}),
		expect: `{
		for a {}
		}`,
	},
	{
		desc: `returnfunc func`,
		code: ReturnFunc(func(rg *Group) {
			rg.Lit(1)
			rg.Lit(2)
		}),
		expect: `return 1, 2`,
	},
	{
		desc: `returnfunc statement`,
		code: Empty().ReturnFunc(func(rg *Group) {
			rg.Lit(1)
			rg.Lit(2)
		}),
		expect: `return 1, 2`,
	},
	{
		desc: `returnfunc group`,
		code: BlockFunc(func(bg *Group) {
			bg.ReturnFunc(func(rg *Group) {
				rg.Lit(1)
				rg.Lit(2)
			})
		}),
		expect: `{
		return 1, 2
		}`,
	},
	{
		desc: `return group`,
		code: BlockFunc(func(g *Group) {
			g.Return()
		}),
		expect: `{
		return
		}`,
	},
	{
		desc: `iffunc group`,
		code: BlockFunc(func(bg *Group) {
			bg.IfFunc(func(ig *Group) {
				ig.Id("a")
			}).Block()
		}),
		expect: `{
		if a {} 
		}`,
	},
	{
		desc: `iffunc func`,
		code: IfFunc(func(ig *Group) {
			ig.Id("a")
		}).Block(),
		expect: `if a {}`,
	},
	{
		desc: `iffunc statement`,
		code: Null().IfFunc(func(ig *Group) {
			ig.Id("a")
		}).Block(),
		expect: `if a {}`,
	},
	{
		desc: `if group`,
		code: BlockFunc(func(g *Group) { g.If(Id("a")).Block() }),
		expect: `{
		if a {}
		}`,
	},
	{
		desc: `map group`,
		code: BlockFunc(func(g *Group) { g.Map(Int()).Int().Values(Dict{Lit(1): Lit(1)}) }),
		expect: `{
		map[int]int{1:1}
		}`,
	},
	{
		desc: `assert group`,
		// Don't do this! ListFunc used to kludge Group.Assert usage without
		// syntax error.
		code:   Id("a").ListFunc(func(g *Group) { g.Assert(Id("b")) }),
		expect: `a.(b)`,
	},
	{
		desc:   `assert func`,
		code:   Id("a").Add(Assert(Id("b"))),
		expect: `a.(b)`,
	},
	{
		desc: `paramsfunc group`,
		// Don't do this! ListFunc used to kludge Group.ParamsFunc usage without
		// syntax error.
		code:   Id("a").ListFunc(func(lg *Group) { lg.ParamsFunc(func(cg *Group) { cg.Lit(1) }) }),
		expect: `a(1)`,
	},
	{
		desc:   `paramsfunc func`,
		code:   Id("a").Add(ParamsFunc(func(g *Group) { g.Lit(1) })),
		expect: `a(1)`,
	},
	{
		desc:   `paramsfunc statement`,
		code:   Id("a").ParamsFunc(func(g *Group) { g.Lit(1) }),
		expect: `a(1)`,
	},
	{
		desc: `params group`,
		// Don't do this! ListFunc used to kludge Group.Params usage without
		// syntax error.
		code:   Id("a").ListFunc(func(g *Group) { g.Params(Lit(1)) }),
		expect: `a(1)`,
	},
	{
		desc:   `params func`,
		code:   Id("a").Add(Params(Lit(1))),
		expect: `a(1)`,
	},
	{
		desc: `callfunc group`,
		// Don't do this! ListFunc used to kludge Group.CallFunc usage without
		// syntax error.
		code:   Id("a").ListFunc(func(lg *Group) { lg.CallFunc(func(cg *Group) { cg.Lit(1) }) }),
		expect: `a(1)`,
	},
	{
		desc:   `callfunc func`,
		code:   Id("a").Add(CallFunc(func(g *Group) { g.Lit(1) })),
		expect: `a(1)`,
	},
	{
		desc: `call group`,
		// Don't do this! ListFunc used to kludge Group.Call usage without
		// syntax error.
		code:   Id("a").ListFunc(func(g *Group) { g.Call(Lit(1)) }),
		expect: `a(1)`,
	},
	{
		desc:   `call func`,
		code:   Id("a").Add(Call(Lit(1))),
		expect: `a(1)`,
	},
	{
		desc: `defsfunc statement`,
		code: Const().DefsFunc(func(g *Group) { g.Id("a").Op("=").Lit(1) }),
		expect: `const (
		a = 1
		)`,
	},
	{
		desc: `defsfunc func`,
		code: Const().Add(DefsFunc(func(g *Group) { g.Id("a").Op("=").Lit(1) })),
		expect: `const (
		a = 1
		)`,
	},
	{
		desc: `defsfunc group`,
		// Don't do this! ListFunc used to kludge Group.DefsFunc usage without
		// syntax error.
		code: Const().ListFunc(func(lg *Group) { lg.DefsFunc(func(dg *Group) { dg.Id("a").Op("=").Lit(1) }) }),
		expect: `const (
		a = 1
		)`,
	},
	{
		desc: `defs group`,
		// Don't do this! ListFunc used to kludge Group.Defs usage without
		// syntax error.
		code: Const().ListFunc(func(g *Group) { g.Defs(Id("a").Op("=").Lit(1)) }),
		expect: `const (
		a = 1
		)`,
	},
	{
		desc: `defs func`,
		code: Const().Add(Defs(Id("a").Op("=").Lit(1))),
		expect: `const (
		a = 1
		)`,
	},
	{
		desc:   `blockfunc group`,
		code:   BlockFunc(func(g *Group) { g.BlockFunc(func(g *Group) {}) }),
		expect: `{{}}`,
	},
	{
		desc:   `block group`,
		code:   BlockFunc(func(g *Group) { g.Block() }),
		expect: `{{}}`,
	},
	{
		desc:   `indexfunc group`,
		code:   BlockFunc(func(g *Group) { g.IndexFunc(func(g *Group) { g.Lit(1) }).Int().Values(Lit(1)) }),
		expect: `{[1]int{1}}`,
	},
	{
		desc:   `indexfunc statement`,
		code:   Id("a").IndexFunc(func(g *Group) { g.Lit(1) }),
		expect: `a[1]`,
	},
	{
		desc:   `indexfunc func`,
		code:   Id("a").Add(IndexFunc(func(g *Group) { g.Lit(1) })),
		expect: `a[1]`,
	},
	{
		desc:   `index group`,
		code:   BlockFunc(func(g *Group) { g.Index(Lit(1)).Int().Values(Lit(1)) }),
		expect: `{[1]int{1}}`,
	},
	{
		desc:   `index func`,
		code:   Id("a").Add(Index(Lit(1))),
		expect: `a[1]`,
	},
	{
		desc: `valuesfunc func`,
		code: ValuesFunc(func(vg *Group) {
			vg.Lit(1)
		}),
		expect: `{1}`,
	},
	{
		desc: `valuesfunc group`,
		code: BlockFunc(func(bg *Group) {
			bg.ValuesFunc(func(vg *Group) {
				vg.Lit(1)
			})
		}),
		expect: `{
		{1}
		}`,
	},
	{
		desc: `values group`,
		code: BlockFunc(func(g *Group) {
			g.Values(Lit(1))
		}),
		expect: `{
		{1}
		}`,
	},
	{
		desc: `listfunc statement`,
		code: Add(Null()).ListFunc(func(lg *Group) {
			lg.Id("a")
			lg.Id("b")
		}).Op("=").Id("c"),
		expect: `a, b = c`,
	},
	{
		desc: `listfunc func`,
		code: ListFunc(func(lg *Group) {
			lg.Id("a")
			lg.Id("b")
		}).Op("=").Id("c"),
		expect: `a, b = c`,
	},
	{
		desc: `listfunc group`,
		code: BlockFunc(func(bg *Group) {
			bg.ListFunc(func(lg *Group) {
				lg.Id("a")
				lg.Id("b")
			}).Op("=").Id("c")
		}),
		expect: `{
		a, b = c
		}`,
	},
	{
		desc: `list group`,
		code: BlockFunc(func(g *Group) { g.List(Id("a"), Id("b")).Op("=").Id("c") }),
		expect: `{
		a, b = c
		}`,
	},
	{
		desc:   `parens func`,
		code:   Parens(Lit(1)),
		expect: `(1)`,
	},
	{
		desc: `parens group`,
		code: BlockFunc(func(g *Group) { g.Parens(Lit(1)) }),
		expect: `{
		(1)
		}`,
	},
}

func TestGen(t *testing.T) {
	caseTester(t, gencases)
}
