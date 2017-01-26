package jen_test

import (
	"fmt"

	. "github.com/davelondon/jennifer/jen"
)

func ExampleGotchaPointer() {
	caller := func(s *Statement) *Statement {
		return s.Call()
	}
	a := Id("a")
	c := Block(
		caller(a),
		caller(a),
	)
	fmt.Printf("%#v", c)
	// Output: {
	// 	a()()
	// 	a()()
	// }
}

func ExampleGotchaPointerFixed() {
	caller := func(s Statement) *Statement {
		return s.Call()
	}
	a := *Id("a")
	c := Block(
		caller(a),
		caller(a),
	)
	fmt.Printf("%#v", c)
	// Output: {
	// 	a()
	// 	a()
	// }
}

func ExampleGotchaAdd() {
	a := Id("a")
	c := Block(
		a.Call(),
		a.Call(),
	)
	fmt.Printf("%#v", c)
	// Output: {
	// 	a()()
	// 	a()()
	// }
}

func ExampleGotchaAddFixed() {
	a := Id("a")
	c := Block(
		Add(a).Call(),
		Add(a).Call(),
	)
	fmt.Printf("%#v", c)
	// Output: {
	// 	a()
	// 	a()
	// }
}

/*
var Keywords = []string{"break", "default", "func", "interface", "select", "case", "defer", "go", "map", "struct", "chan", "else", "goto", "package", "switch", "const", "fallthrough", "if", "range", "type", "continue", "import", "var"}

 "return" and "for" are special cases
*/

func ExampleId2() {
	id := Id("foo.Bar", "Baz")
	c := Id(id, "Qux").Call()
	fmt.Printf("%#v", c)
	// Output: foo.Bar.Baz.Qux()
}

func ExampleErr() {
	c := If(
		Err().Op(":=").Id("foo").Call(),
		Err().Op("!=").Nil(),
	).Block(
		Return(Err()),
	)
	fmt.Printf("%#v", c)
	// Output: if err := foo(); err != nil {
	// 	return err
	// }
}

func ExampleCaseBlock() {
	c := Switch().Id("foo").Block(
		Case().Lit("a").CaseBlock(
			Return(Lit(1)),
		),
		Case().Lit("b").CaseBlock(
			Return(Lit(2)),
		),
		Default().CaseBlock(
			Return(Lit(3)),
		),
	)
	fmt.Printf("%#v", c)
	// Output: switch foo {
	// case "a":
	// 	return 1
	// case "b":
	// 	return 2
	// default:
	// 	return 3
	// }
}

func ExampleTag() {
	c := Type().Id("foo").Struct().Block(
		Id("A").String().Tag(map[string]string{"json": "a"}),
		Id("B").Int().Tag(map[string]string{"json": "b", "bar": "baz"}),
	)
	fmt.Printf("%#v", c)
	// Output: type foo struct {
	// 	A string `json:"a"`
	// 	B int    `json:"b" bar:"baz"`
	// }
}

func ExampleNull() {
	c := Func().Id("foo").Params(
		nil,
		Id("s").String(),
		Null(),
		Id("i").Int(),
	).Block()
	fmt.Printf("%#v", c)
	// Output: func foo(s string, i int) {
	// }
}

func ExampleComplex() {
	collection := func(name string, key Code, value Code) *Statement {
		if key == nil {
			// slice
			return Var().Id(name).Index().Add(value)
		} else {
			// map
			return Var().Id(name).Map(key).Add(value)
		}
	}
	c := Func().Id("main").Params().Block(
		collection("foo", nil, String()),
		collection("bar", String(), Int()),
	)
	fmt.Printf("%#v", c)
	// Output: func main() {
	// 	var foo []string
	// 	var bar map[string]int
	// }
}

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
	// Output: func a() {
	// }
}

func ExampleGroup_Func() {
	c := Id("a").Op(":=").Func().Params().Block()
	fmt.Printf("%#v", c)
	// Output: a := func() {
	// }
}

/*
var Types = []string{"bool", "byte", "complex64", "complex128", "error", "float32", "float64", "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"}
*/

func ExampleBool() {
	c := List(Id("b"), Id("ok")).Op(":=").Id("a").Assert(Bool())
	fmt.Printf("%#v", c)
	// Output: b, ok := a.(bool)
}

func ExampleGroup_Bool() {
	c := Var().Id("a").Bool().Op("=").Lit(true)
	fmt.Printf("%#v", c)
	// Output: var a bool = true
}

func ExampleByte() {
	c := Id("b").Op(":=").Id("a").Assert(Byte())
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
Blocks: "Parens", "List", "Values", "Index", "Block","Call", "Params"
*/

func ExampleParens2() {
	c := Id("a").Op("/").Parens(Id("b").Op("+").Id("c"))
	fmt.Printf("%#v", c)
	// Output: a / (b + c)
}

func ExampleValues2() {
	c := Index().String().Values(Lit("a"), Lit("b"))
	fmt.Printf("%#v", c)
	// Output: []string{"a", "b"}
}

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

func ExampleGroup_BlockFunc() {
	c := Func().Id("a").Params().BlockFunc(func(g *Group) {
		g.Id("a").Op("++")
		g.Id("b").Op("--")
	})
	fmt.Printf("%#v", c)
	// Output: func a() {
	// 	a++
	//	b--
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
	c := Comment("a\nb")
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

//func Lit(v interface{}) *Group

func ExampleGroup_Dict() {
	c := Id("a").Op(":=").Map(String()).String().Dict(map[Code]Code{
		Lit("a"): Lit("b"),
	})
	fmt.Printf("%#v", c)
	// Output: a := map[string]string{
	// 	"a": "b",
	// }
}

func ExampleGroup_DictFunc() {
	c := Id("a").Op(":=").Map(String()).String().DictFunc(func(m map[Code]Code) {
		m[Lit("a")] = Lit("b")
	})
	fmt.Printf("%#v", c)
	// Output: a := map[string]string{
	// 	"a": "b",
	// }
}

//func Id(names ...string) *Group

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
	// func main() {
	// 	fmt.Println("Hello, world")
	// }
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
	// func main() {
	// 	fmt.Println("Hello, world")
	// }
}

func ExampleNewFilePath() {
	f := NewFilePath("a.b/c")
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

func ExampleNewFilePathName() {
	f := NewFilePathName("a.b/c", "main")
	f.Func().Id("main").Params().Block(
		Id("a.b/c.Foo").Call(),
	)
	fmt.Printf("%#v", f)
	// Output: package main
	//
	// func main() {
	// 	Foo()
	// }
}

func ExampleFile_PackageComment() {
	f := NewFile("c")
	f.PackageComment("a")
	f.PackageComment("b")
	f.Func().Id("init").Params().Block()
	fmt.Printf("%#v", f)
	// Output: // a
	// // b
	// package c
	//
	// func init() {
	// }
}
