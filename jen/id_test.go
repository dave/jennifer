package jen

import "fmt"

func ExampleId_clone() {
	id := Id("foo")
	c := id.Clone().Id("Type")
	fmt.Printf("%#v", c)
	// Output:
	// foo.Type
}

func ExampleId_add() {
	id := Id("foo")
	c := Id("v").Add(id)
	fmt.Printf("%#v", c)
	// Output:
	// v.foo
}

func ExampleId_assert() {
	c := Id("a").Op(":=").Id("b").Assert(Id("foo")).Id("c").Call()
	fmt.Printf("%#v", c)
	// Output:
	// a := b.(foo).c()
}

func ExampleId_need_second_id() {
	c := Id("a").Op(":=").Index().Id("b").Values(Lit("c"))
	fmt.Printf("%#v", c)
	// Output:
	// a := []b{"c"}
}

func ExampleId_params() {
	c := Func().Params(
		Id("a").Id("A"),
	).Id("foo").Params(
		Id("b"),
		Id("c").Id("B"),
	).String().Block()
	fmt.Printf("%#v", c)
	// Output:
	// func (a A) foo(b, c B) string {}
}

func ExampleQual_parens() {
	c := Parens(Op("*").Id("a")).Id("b")
	fmt.Printf("%#v", c)
	// Output:
	// (*a).b
}

func ExampleQual_sel() {
	c := Qual("a.b/c", "Foo").Id("Bar")
	fmt.Printf("%#v", c)
	// Output:
	// c.Foo.Bar
}

func ExampleId_selectors_test() {
	c := Id("a").Id("b")
	fmt.Printf("%#v", c)
	// Output:
	// a.b
}

func ExampleId_selectors_call() {
	c := Id("a").Call().Id("b")
	fmt.Printf("%#v", c)
	// Output:
	// a().b
}

func ExampleId_selectors_index() {
	c := Id("a").Index(Lit(0)).Id("b")
	fmt.Printf("%#v", c)
	// Output:
	// a[0].b
}
