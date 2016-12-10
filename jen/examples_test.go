package jen_test

import (
	"fmt"

	. "github.com/davelondon/jennifer/jen"
)

func ExampleAppend() {
	c := Id("a").Call(
		Append(Id("b"), Id("c")),
	)
	fmt.Printf("%#v", c)
	// Output: a(append(b, c))
}

func ExampleGroup_Append() {
	c := Id("a").Op("=").Append(Id("a"), Id("b"))
	fmt.Printf("%#v", c)
	// Output: a = append(a, b)
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

func ExampleGroup_Call() {
	c := Id("a").Call(Id("b"), Id("c"))
	fmt.Printf("%#v", c)
	// Output: a(b, c)
}

//func Cap(c ...Code) *Group
//func Case() *Group
//func Chan() *Group
//func Close(c ...Code) *Group
//func Comment(comments ...string) *Group
//func Commentf(format string, a ...interface{}) *Group
//func Complex(c ...Code) *Group
//func Complex128() *Group
//func Complex64() *Group
//func Const() *Group
//func Continue() *Group
//func Copy(c ...Code) *Group
//func Decls(c ...Code) *Group
//func Default() *Group
//func Defer() *Group
//func Delete(c ...Code) *Group
//func Do(f func(*Group)) *Group
//func Else() *Group
//func Empty() *Group
//func Error() *Group
//func Fallthrough() *Group
//func False() *Group
//func Float32() *Group
//func Float64() *Group
//func For() *Group
//func Func() *Group
//func Go() *Group
//func Goto() *Group
//func Id(names ...string) *Group
//func If() *Group
//func Imag(c ...Code) *Group
//func Import() *Group
//func Index(c ...Code) *Group
//func Int() *Group
//func Int16() *Group
//func Int32() *Group
//func Int64() *Group
//func Int8() *Group
//func Interface() *Group
//func Iota() *Group
//func Len(c ...Code) *Group
//func List(c ...Code) *Group
//func Lit(v interface{}) *Group
//func Make(c ...Code) *Group
//func Map() *Group
//func New(c ...Code) *Group
//func NewFile(name string) *Group
//func NewFilePath(name, path string) *Group
//func Nil() *Group
//func Null() *Group
//func Op(op string) *Group
//func Package() *Group
//func Panic(c ...Code) *Group
//func Params(c ...Code) *Group
//func Parens(c ...Code) *Group
//func Print(c ...Code) *Group
//func Println(c ...Code) *Group
//func Range() *Group
//func Real(c ...Code) *Group
//func Recover(c ...Code) *Group
//func Return(c ...Code) *Group
//func Rune() *Group
//func Select() *Group
//func String() *Group
//func Struct() *Group
//func Switch() *Group
//func True() *Group
//func Type() *Group
//func Uint() *Group
//func Uint16() *Group
//func Uint32() *Group
//func Uint64() *Group
//func Uint8() *Group
//func Uintptr() *Group
//func Values(c ...Code) *Group
//func Var() *Group

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
