package jen_test

import (
	"fmt"

	. "github.com/davelondon/jennifer/jen"
)

func ExampleFunc() {
	c := Func().Id("a").Params().Block()
	fmt.Printf("%#v", c)
	// Output: func a() {}
}

func ExampleNewFile() {
	f := NewFile("a")
	f.Var().Id("b").Op("=").Lit(1)
	fmt.Printf("%#v", f)
	// Output: package a
	//
	// var b = 1
}

func ExampleId() {
	f := NewFile("a")
	f.Func().Id("main").Params().Block(
		Id("fmt.Println").Call(
			Lit("Hello, world"),
		),
	)
	fmt.Printf("%#v", f)
	// Output: package a
	//
	// import fmt "fmt"
	//
	// func main() { fmt.Println("Hello, world") }
}

func ExampleNewFilePath() {
	f := NewFilePath("c", "a.b/c")
	f.Func().Id("main").Params().Block(
		Id("a.b/c.Local").Call(),
		Id("d.e/f.Remote").Call(),
		Id("g.h/f.Collision").Call(),
	)
	fmt.Printf("%#v", f)
	// Output: package c
	//
	// import (
	// 	f "d.e/f"
	// 	f1 "g.h/f"
	// )
	//
	// func main() {
	// 	Local()
	// 	f.Remote()
	// 	f1.Collision()
	// }
}
