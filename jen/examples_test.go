package jen_test

import (
	"fmt"

	. "github.com/davelondon/jennifer/jen"
)

/*
var Keywords = []string{"break", "default", "func", "interface", "select", "case", "defer", "go", "map", "struct", "chan", "else", "goto", "package", "switch", "const", "fallthrough", "if", "range", "type", "continue", "import", "var"}

 "return" and "for" are special cases
*/

func ExampleBreak() {
	c := For(
		Id("i").Op(":=").Lit(0),
		Id("i").Op("<").Lit(10),
		Id("i").Op("++"),
	).Block(
		If().Id("i").Op(">").Lit(5).Block(
			Break(),
		),
	)
	fmt.Printf("%#v", c)
	// Output: for i := 0; i < 10; i++ {
	// 	if i > 5 {
	// 		break
	// 	}
	// }
}

func ExampleFunc() {
	c := Func().Id("a").Params().Block()
	fmt.Printf("%#v", c)
	// Output: func a() {}
}

func ExampleGroup_Func() {
	c := Id("a").Op(":=").Func().Params().Block()
	fmt.Printf("%#v", c)
	// Output: a := func() {}
}

/*
var Types = []string{"bool", "byte", "complex64", "complex128", "error", "float32", "float64", "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"}
*/

func ExampleBool() {
	c := List(Id("b"), Id("ok")).Op(":=").Id("a").Op(".").Parens(Bool())
	fmt.Printf("%#v", c)
	// Output: b, ok := a.(bool)
}

func ExampleGroup_Bool() {
	c := Var().Id("a").Bool().Op("=").Lit(true)
	fmt.Printf("%#v", c)
	// Output: var a bool = true
}

func ExampleByte() {
	c := Id("b").Op(":=").Id("a").Op(".").Parens(Byte())
	fmt.Printf("%#v", c)
	// Output: b := a.(byte)
}

func ExampleGroup_Byte() {
	c := Id("b").Op(":=").Index().Byte().Parens(Id("s"))
	fmt.Printf("%#v", c)
	// Output: b := []byte(s)
}

/*
var Constants = []string{"true", "false", "iota"}
*/

/*
var Zero = []string{"nil"}
*/

/*
var Functions = []string{"append", "cap", "close", "complex", "copy", "delete", "imag", "len", "make", "new", "panic", "print", "println", "real", "recover"}
*/
func ExampleAppend() {
	c := Id("a").Call(
		Append(Id("b"), Id("c")),
	)
	fmt.Printf("%#v", c)
	// Output: a(append(b, c))
}

func ExampleGroup_Append() {
	c := Id("a").Op("=").Append(Id("a"), Id("b").Op("..."))
	fmt.Printf("%#v", c)
	// Output: a = append(a, b...)
}

/*
Blocks: "Parens", "List", "Values", "Index", "Block","Call", "Params", "Decls"
*/

func ExampleBlock() {
	c := Block(Id("a").Op("=").Id("b"))
	fmt.Printf("%#v", c)
	// Output: {
	// 	a = b
	// }
}

func ExampleGroup_Block() {
	c := If().Id("a").Op("==").Lit(1).Block(
		Return(),
	)
	fmt.Printf("%#v", c)
	// Output: if a == 1 {
	// 	return
	// }
}

func ExampleGroup_Call() {
	c := Id("a").Call(Id("b"), Id("c"))
	fmt.Printf("%#v", c)
	// Output: a(b, c)
}

//func Comment(comments ...string) *Group
//func Commentf(format string, a ...interface{}) *Group

func ExampleComment() {
	c := Comment("a")
	fmt.Printf("%#v", c)
	// Output: // a
}

func ExampleComment_multiline() {
	c := Comment("a", "b")
	fmt.Printf("%#v", c)
	// Output: /*
	// a
	// b
	// */
}

func ExampleGroup_Comment() {
	c := Id("a").Call().Comment("b")
	fmt.Printf("%#v", c)
	// Output: a() // b
}

func ExampleCommentf() {
	c := Commentf("a %d", 1)
	fmt.Printf("%#v", c)
	// Output: // a 1
}

func ExampleGroup_Commentf() {
	c := Id("a").Call().Commentf("b %d", 1)
	fmt.Printf("%#v", c)
	// Output: a() // b 1
}

//func Do(f func(*Group)) *Group
//func Empty() *Group
//func For() *Group

//func Id(names ...string) *Group
//func Lit(v interface{}) *Group

func ExampleId_local() {
	c := Id("a").Op(":=").Lit(1)
	fmt.Printf("%#v", c)
	// Output: a := 1
}

func ExampleId_select() {
	c := Id("a", "b", "c").Call()
	fmt.Printf("%#v", c)
	// Output: a.b.c()
}

func ExampleId_remote() {
	f := NewFile("main")
	f.Func().Id("main").Params().Block(
		Id("fmt.Println").Call(
			Lit("Hello, world"),
		),
	)
	fmt.Printf("%#v", f)
	// Output: package main
	//
	// import fmt "fmt"
	//
	// func main() { fmt.Println("Hello, world") }
}

//func NewFile(name string) *Group
//func NewFilePath(name, path string) *Group
//func Null() *Group
//func Op(op string) *Group
//func Return(c ...Code) *Group

func ExampleNewFile() {
	f := NewFile("main")
	f.Func().Id("main").Params().Block(
		Id("fmt.Println").Call(
			Lit("Hello, world"),
		),
	)
	fmt.Printf("%#v", f)
	// Output: package main
	//
	// import fmt "fmt"
	//
	// func main() { fmt.Println("Hello, world") }
}
func ExampleNewFilePath() {
	f := NewFilePath("c", "a.b/c")
	f.Func().Id("init").Params().Block(
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
	// func init() {
	// 	Local()
	// 	f.Remote()
	// 	f1.Collision()
	// }
}
