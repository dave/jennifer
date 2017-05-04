package jen_test

import (
	"testing"

	. "github.com/dave/jennifer/jen"
)

var gencases = []tc{
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
