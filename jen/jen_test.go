package jen_test

import (
	"fmt"
	"go/format"
	"strings"
	"testing"

	. "github.com/dave/jennifer/jen"
)

var o1 = Options{
	Close:     ")",
	Multi:     true,
	Open:      "(",
	Separator: ",",
}

var o2 = Options{
	Close:     "",
	Multi:     false,
	Open:      "",
	Separator: ",",
}

var cases = []tc{
	{
		desc:   `scientific notation`,
		code:   Lit(1e3),
		expect: `1000.0`,
	},
	{
		desc:   `big float`,
		code:   Lit(1000000.0),
		expect: `1e+06`,
	},
	{
		desc:   `lit float whole numbers`,
		code:   Index().Float64().Values(Lit(-10.0), Lit(-2.0), Lit(-1.0), Lit(0.0), Lit(1.0), Lit(2.0), Lit(10.0)),
		expect: "[]float64{-10.0, -2.0, -1.0, 0.0, 1.0, 2.0, 10.0}",
	},
	{
		desc: `custom func group`,
		code: ListFunc(func(g *Group) {
			g.CustomFunc(o2, func(g *Group) {
				g.Id("a")
				g.Id("b")
				g.Id("c")
			})
		}).Op("=").Id("foo").Call(),
		expect: `a, b, c = foo()`,
	},
	{
		desc:   `custom group`,
		code:   ListFunc(func(g *Group) { g.Custom(o2, Id("a"), Id("b"), Id("c")) }).Op("=").Id("foo").Call(),
		expect: `a, b, c = foo()`,
	},
	{
		desc: `custom function`,
		code: Id("foo").Add(Custom(o1, Lit("a"), Lit("b"), Lit("c"))),
		expect: `foo(
			"a",
			"b",
			"c",
		)`,
	},
	{
		desc: `custom function`,
		code: Id("foo").Add(Custom(o1, Lit("a"), Lit("b"), Lit("c"))),
		expect: `foo(
			"a",
			"b",
			"c",
		)`,
	},
	{
		desc: `line statement`,
		code: Block(Lit(1).Line(), Lit(2)),
		expect: `{
		1
		
		2
		}`,
	},
	{
		desc: `line func`,
		code: Block(Lit(1), Line(), Lit(2)),
		expect: `{
		1
		
		2
		}`,
	},
	{
		desc: `line group`,
		code: BlockFunc(func(g *Group) {
			g.Id("a")
			g.Line()
			g.Id("b")
		}),
		expect: `{
		a
		
		b
		}`,
	},
	{
		desc: `op group`,
		code: BlockFunc(func(g *Group) {
			g.Op("*").Id("a")
		}),
		expect: `{*a}`,
	},
	{
		desc: `empty group`,
		code: BlockFunc(func(g *Group) {
			g.Empty()
		}),
		expect: `{
		
		}`,
	},
	{
		desc: `null group`,
		code: BlockFunc(func(g *Group) {
			g.Null()
		}),
		expect: `{}`,
	},
	{
		desc:   `tag no backquote`,
		code:   Tag(map[string]string{"a": "`b`"}),
		expect: "\"a:\\\"`b`\\\"\"",
	},
	{
		desc:   `tag null`,
		code:   Tag(map[string]string{}),
		expect: ``,
	},
	{
		desc: `litrunefunc group`,
		code: BlockFunc(func(g *Group) {
			g.LitByteFunc(func() byte { return byte(0xab) })
		}),
		expect: `{byte(0xab)}`,
	},
	{
		desc: `litbyte group`,
		code: BlockFunc(func(g *Group) {
			g.LitByte(byte(0xab))
		}),
		expect: `{byte(0xab)}`,
	},
	{
		desc: `litrunefunc group`,
		code: BlockFunc(func(g *Group) {
			g.LitRuneFunc(func() rune { return 'a' })
		}),
		expect: `{'a'}`,
	},
	{
		desc: `litrune group`,
		code: BlockFunc(func(g *Group) {
			g.LitRune('a')
		}),
		expect: `{'a'}`,
	},
	{
		desc: `litfunc group`,
		code: BlockFunc(func(g *Group) {
			g.LitFunc(func() interface{} {
				return 1 + 1
			})
		}),
		expect: `{2}`,
	},
	{
		desc: `litfunc func`,
		code: LitFunc(func() interface{} {
			return 1 + 1
		}),
		expect: `2`,
	},
	{
		desc:   `group all null`,
		code:   List(Null(), Null()),
		expect: ``,
	},
	{
		desc:   `do group`,
		code:   BlockFunc(func(g *Group) { g.Do(func(s *Statement) { s.Lit(1) }) }),
		expect: `{1}`,
	},
	{
		desc:   `do func`,
		code:   Do(func(s *Statement) { s.Lit(1) }),
		expect: `1`,
	},
	{
		desc:   `dict empty`,
		code:   Values(Dict{}),
		expect: `{}`,
	},
	{
		desc:   `dict null`,
		code:   Values(Dict{Null(): Null()}),
		expect: `{}`,
	},
	{
		desc: `commentf group`,
		code: BlockFunc(func(g *Group) { g.Commentf("%d", 1) }),
		expect: `{
		// 1
		}`,
	},
	{
		desc:   `commentf func`,
		code:   Commentf("%d", 1),
		expect: `// 1`,
	},
	{
		desc:   `add func`,
		code:   Add(Lit(1)),
		expect: `1`,
	},
	{
		desc:   `add group`,
		code:   BlockFunc(func(g *Group) { g.Add(Lit(1)) }),
		expect: `{1}`,
	},
	{
		desc:   `empty block`,
		code:   Block(),
		expect: `{}`,
	},
	{
		desc:   `string literal`,
		code:   Lit("a"),
		expect: `"a"`,
	},
	{
		desc:   `int literal`,
		code:   Lit(1),
		expect: `1`,
	},
	{
		desc:   `simple id`,
		code:   Id("a"),
		expect: `a`,
	},
	{
		desc:   `foreign id`,
		code:   Qual("x.y/z", "a"),
		expect: `z.a`,
		expectImports: map[string]string{
			"x.y/z": "z",
		},
	},
	{
		desc:   `var decl`,
		code:   Var().Id("a").Op("=").Lit("b"),
		expect: `var a = "b"`,
	},
	{
		desc:   `short var decl`,
		code:   Id("a").Op(":=").Lit("b"),
		expect: `a := "b"`,
	},
	{
		desc:   `simple if`,
		code:   If(Id("a").Op("==").Lit("b")).Block(),
		expect: `if a == "b" {}`,
	},
	{
		desc: `simple if`,
		code: If(Id("a").Op("==").Lit("b")).Block(
			Id("a").Op("++"),
		),
		expect: `if a == "b" { a++ }`,
	},
	{
		desc:   `pointer`,
		code:   Op("*").Id("a"),
		expect: `*a`,
	},
	{
		desc:   `address`,
		code:   Op("&").Id("a"),
		expect: `&a`,
	},
	{
		desc: `simple call`,
		code: Id("a").Call(
			Lit("b"),
			Lit("c"),
		),
		expect: `a("b", "c")`,
	},
	{
		desc: `call fmt.Sprintf`,
		code: Qual("fmt", "Sprintf").Call(
			Lit("b"),
			Id("c"),
		),
		expect: `fmt.Sprintf("b", c)`,
	},
	{
		desc: `slices`,
		code: Id("a").Index(
			Lit(1),
			Empty(),
		),
		expect: `a[1:]`,
	},
	{
		desc:   `return`,
		code:   Return(Id("a")),
		expect: `return a`,
	},
	{
		desc:   `double return`,
		code:   Return(Id("a"), Id("b")),
		expect: `return a, b`,
	},
	{
		desc: `func`,
		code: Func().Id("a").Params(
			Id("a").String(),
		).Block(
			Return(Id("a")),
		),
		expect: `func a(a string){
			return a
		}`,
	},
	{
		desc:   `built in func`,
		code:   New(Id("a")),
		expect: `new(a)`,
	},
	{
		desc:   `multip`,
		code:   Id("a").Op("*").Id("b"),
		expect: `a * b`,
	},
	{
		desc:   `multip ptr`,
		code:   Id("a").Op("*").Op("*").Id("b"),
		expect: `a * *b`,
	},
	{
		desc:   `field`,
		code:   Id("a").Dot("b"),
		expect: `a.b`,
	},
	{
		desc:   `method`,
		code:   Id("a").Dot("b").Call(Id("c"), Id("d")),
		expect: `a.b(c, d)`,
	},
	{
		desc: `if else`,
		code: If(Id("a").Op("==").Lit(1)).Block(
			Id("b").Op("=").Lit(1),
		).Else().If(Id("a").Op("==").Lit(2)).Block(
			Id("b").Op("=").Lit(2),
		).Else().Block(
			Id("b").Op("=").Lit(3),
		),
		expect: `if a == 1 { b = 1 } else if a == 2 { b = 2 } else { b = 3 }`,
	},
	{
		desc:   `literal array`,
		code:   Index().String().Values(Lit("a"), Lit("b")),
		expect: `[]string{"a", "b"}`,
	},
	{
		desc:   `comment`,
		code:   Comment("a"),
		expect: `// a`,
	},
	{
		desc:   `null`,
		code:   Id("a").Params(Id("b"), Null(), Id("c")),
		expect: `a(b, c)`,
	},
	{
		desc: `map literal single`,
		code: Id("a").Values(Dict{
			Id("b"): Id("c"),
		}),
		expect: `a{b: c}`,
	},
	{
		desc: `map literal null`,
		code: Id("a").Values(Dict{
			Null():  Id("c"),
			Id("b"): Null(),
			Id("b"): Id("c"),
		}),
		expect: `a{b: c}`,
	},
	{
		desc: `map literal multiple`,
		code: Id("a").Values(Dict{
			Id("b"): Id("c"),
			Id("d"): Id("e"),
		}),
		expect: `a{
			b: c,
			d: e,
		}`,
	},
	{
		desc: `map literal func single`,
		code: Id("a").Values(DictFunc(func(d Dict) {
			d[Id("b")] = Id("c")
		})),
		expect: `a{b: c}`,
	},
	{
		desc: `map literal func single null`,
		code: Id("a").Values(DictFunc(func(d Dict) {
			d[Null()] = Id("c")
			d[Id("b")] = Null()
			d[Id("b")] = Id("c")
		})),
		expect: `a{b: c}`,
	},
	{
		desc: `map literal func multiple`,
		code: Id("a").Values(DictFunc(func(d Dict) {
			d[Id("b")] = Id("c")
			d[Id("d")] = Id("e")
		})),
		expect: `a{
			b: c,
			d: e,
		}`,
	},
	{
		desc: `literal func`,
		code: Id("a").Op(":=").LitFunc(func() interface{} {
			return "b"
		}),
		expect: `a := "b"`,
	},
	{
		desc:   `dot`,
		code:   Id("a").Dot("b").Dot("c"),
		expect: `a.b.c`,
	},
	{
		desc:   `do`,
		code:   Id("a").Do(func(s *Statement) { s.Dot("b") }),
		expect: `a.b`,
	},
	{
		desc:   `tags should be ordered`,
		code:   Tag(map[string]string{"z": "1", "a": "2"}),
		expect: "`a:\"2\" z:\"1\"`",
	},
	{
		desc: `dict should be ordered`,
		code: Map(String()).Int().Values(Dict{Id("z"): Lit(1), Id("a"): Lit(2)}),
		expect: `map[string]int{
		a:2, 
		z:1,
		}`,
	},
}

func TestJen(t *testing.T) {
	caseTester(t, cases)
}

func caseTester(t *testing.T, cases []tc) {
	for i, c := range cases {
		onlyTest := ""
		if onlyTest != "" && c.desc != onlyTest {
			continue
		}
		rendered := fmt.Sprintf("%#v", c.code)

		expected, err := format.Source([]byte(c.expect))
		if err != nil {
			panic(fmt.Sprintf("Error formatting expected source in test case %d. Description: %s\nError:\n%s", i, c.desc, err))
		}

		if strings.TrimSpace(string(rendered)) != strings.TrimSpace(string(expected)) {
			t.Errorf("Test case %d failed. Description: %s\nExpected:\n%s\nOutput:\n%s", i, c.desc, expected, rendered)
		}

		//if c.expectImports != nil {
		//	f := FromContext(ctx)
		//	if !reflect.DeepEqual(f.Imports, c.expectImports) {
		//		t.Errorf("Test case %d failed. Description: %s\nImports expected:\n%s\nOutput:\n%s", i, c.desc, c.expectImports, f.Imports)
		//	}
		//}
	}
}

// a test case
type tc struct {
	// path
	path string
	// description for locating the test case
	desc string
	// code to generate
	code Code
	// expected generated source
	expect string
	// expected imports
	expectImports map[string]string
}

func TestNilStatement(t *testing.T) {
	var s *Statement
	c := Func().Id("a").Params(
		s,
	)
	got := fmt.Sprintf("%#v", c)
	expect := "func a()"
	if got != expect {
		t.Fatalf("Got: %s, expect: %s", got, expect)
	}
}

func TestNilGroup(t *testing.T) {
	var g *Group
	c := Func().Id("a").Params(
		g,
	)
	got := fmt.Sprintf("%#v", c)
	expect := "func a()"
	if got != expect {
		t.Fatalf("Got: %s, expect: %s", got, expect)
	}
}

func TestGroup_GoString(t *testing.T) {
	BlockFunc(func(g *Group) {
		g.Lit(1)
		got := fmt.Sprintf("%#v", g)
		expect := "{\n\t1\n}"
		if got != expect {
			t.Fatalf("Got: %s, expect: %s", got, expect)
		}
	})
}
