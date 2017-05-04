package jen_test

import (
	"testing"

	. "github.com/dave/jennifer/jen"
)

var gencases = []tc{
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
