package jen_test

import (
	"fmt"

	"bytes"

	. "github.com/davelondon/jennifer/jen"
)

func ExampleAdd() {
	ptr := Op("*")
	c := Id("a").Op("=").Add(ptr).Id("b")
	fmt.Printf("%#v", c)
	// Output:
	// a = *b
}

func ExampleAdd_var() {
	a := Id("a")
	i := Int()
	c := Var().Add(a, i)
	fmt.Printf("%#v", c)
	// Output:
	// var a int
}

func ExampleAppend() {
	c := Append(Id("a"), Id("b"))
	fmt.Printf("%#v", c)
	// Output:
	// append(a, b)
}

func ExampleAppend_more() {
	c := Id("a").Op("=").Append(Id("a"), Id("b").Op("..."))
	fmt.Printf("%#v", c)
	// Output:
	// a = append(a, b...)
}

func ExampleAssert() {
	c := List(Id("b"), Id("ok")).Op(":=").Id("a").Assert(Bool())
	fmt.Printf("%#v", c)
	// Output:
	// b, ok := a.(bool)
}

func ExampleBlock() {
	c := Func().Id("foo").Params().String().Block(
		Id("a").Op("=").Id("b"),
		Id("b").Op("++"),
		Return(Id("b")),
	)
	fmt.Printf("%#v", c)
	// Output:
	// func foo() string {
	// 	a = b
	// 	b++
	// 	return b
	// }
}

func ExampleBlock_if() {
	c := If(Id("a").Op(">").Lit(10)).Block(
		Id("a").Op("=").Id("a").Op("/").Lit(2),
	)
	fmt.Printf("%#v", c)
	// Output:
	// if a > 10 {
	// 	a = a / 2
	// }
}

func ExampleBlockFunc() {
	increment := true
	c := Func().Id("a").Params().BlockFunc(func(g *Group) {
		if increment {
			g.Id("a").Op("++")
		} else {
			g.Id("a").Op("--")
		}
	})
	fmt.Printf("%#v", c)
	// Output:
	// func a() {
	// 	a++
	// }
}

func ExampleBool() {
	c := Var().Id("b").Bool()
	fmt.Printf("%#v", c)
	// Output:
	// var b bool
}

func ExampleBreak() {
	c := For(
		Id("i").Op(":=").Lit(0),
		Id("i").Op("<").Lit(10),
		Id("i").Op("++"),
	).Block(
		If(Id("i").Op(">").Lit(5)).Block(
			Break(),
		),
	)
	fmt.Printf("%#v", c)
	// Output:
	// for i := 0; i < 10; i++ {
	// 	if i > 5 {
	// 		break
	// 	}
	// }
}

func ExampleByte() {
	c := Id("b").Op(":=").Id("a").Assert(Byte())
	fmt.Printf("%#v", c)
	// Output:
	// b := a.(byte)
}

func ExampleCall() {
	c := Qual("fmt", "Printf").Call(
		Lit("%#v: %T\n"),
		Id("a"),
		Id("b"),
	)
	fmt.Printf("%#v", c)
	// Output:
	// fmt.Printf("%#v: %T\n", a, b)
}

func ExampleCall_fmt() {
	c := Id("a").Call(Lit("b"))
	fmt.Printf("%#v", c)
	// Output:
	// a("b")
}

func ExampleCallFunc() {
	f := func(name string, second string) {
		c := Id("foo").CallFunc(func(g *Group) {
			g.Id(name)
			if second != "" {
				g.Lit(second)
			}
		})
		fmt.Printf("%#v\n", c)
	}
	f("a", "b")
	f("c", "")
	// Output:
	// foo(a, "b")
	// foo(c)
}

func ExampleCap() {
	c := Id("i").Op(":=").Cap(Id("v"))
	fmt.Printf("%#v", c)
	// Output:
	// i := cap(v)
}

func ExampleCase() {
	c := Switch(Id("person")).Block(
		Case(Id("John"), Id("Peter")).Block(
			Return(Lit("male")),
		),
		Case(Id("Gill")).Block(
			Return(Lit("female")),
		),
	)
	fmt.Printf("%#v", c)
	// Output:
	// switch person {
	// case John, Peter:
	// 	return "male"
	// case Gill:
	// 	return "female"
	// }
}

func ExampleBlock_case() {
	c := Select().Block(
		Case(Op("<-").Id("done")).Block(
			Return(Nil()),
		),
		Case(List(Err(), Id("open")).Op(":=").Op("<-").Id("fail")).Block(
			If(Op("!").Id("open")).Block(
				Return(Err()),
			),
		),
	)
	fmt.Printf("%#v", c)
	// Output:
	// select {
	// case <-done:
	// 	return nil
	// case err, open := <-fail:
	// 	if !open {
	// 		return err
	// 	}
	// }
}

func ExampleBlockFunc_case() {
	preventExitOnError := true
	c := Select().Block(
		Case(Op("<-").Id("done")).Block(
			Return(Nil()),
		),
		Case(Err().Op(":=").Op("<-").Id("fail")).BlockFunc(func(g *Group) {
			if !preventExitOnError {
				g.Return(Err())
			} else {
				g.Qual("fmt", "Println").Call(Err())
			}
		}),
	)
	fmt.Printf("%#v", c)
	// Output:
	// select {
	// case <-done:
	// 	return nil
	// case err := <-fail:
	// 	fmt.Println(err)
	// }
}

func ExampleCaseFunc() {
	samIsMale := false
	c := Switch(Id("person")).Block(
		CaseFunc(func(g *Group) {
			g.Id("John")
			g.Id("Peter")
			if samIsMale {
				g.Id("Sam")
			}
		}).Block(
			Return(Lit("male")),
		),
		CaseFunc(func(g *Group) {
			g.Id("Gill")
			if !samIsMale {
				g.Id("Sam")
			}
		}).Block(
			Return(Lit("female")),
		),
	)
	fmt.Printf("%#v", c)
	// Output:
	// switch person {
	// case John, Peter:
	// 	return "male"
	// case Gill, Sam:
	// 	return "female"
	// }
}

func ExampleChan() {
	c := Func().Id("init").Params().Block(
		Id("c").Op(":=").Make(Chan().Qual("os", "Signal"), Lit(1)),
		Qual("os/signal", "Notify").Call(Id("c"), Qual("os", "Interrupt")),
		Qual("os/signal", "Notify").Call(Id("c"), Qual("syscall", "SIGTERM")),
		Go().Func().Params().Block(
			Op("<-").Id("c"),
			Id("cancel").Call(),
		).Call(),
	)
	fmt.Printf("%#v", c)
	// Output:
	// func init() {
	// 	c := make(chan os.Signal, 1)
	// 	signal.Notify(c, os.Interrupt)
	// 	signal.Notify(c, syscall.SIGTERM)
	// 	go func() {
	// 		<-c
	// 		cancel()
	// 	}()
	// }
}

func ExampleClose() {
	c := Block(
		Id("ch").Op(":=").Make(Chan().Struct()),
		Go().Func().Params().Block(
			Op("<-").Id("ch"),
			Qual("fmt", "Println").Call(Lit("done.")),
		).Call(),
		Close(Id("ch")),
	)
	fmt.Printf("%#v", c)
	// Output:
	// {
	// 	ch := make(chan struct{})
	// 	go func() {
	// 		<-ch
	// 		fmt.Println("done.")
	// 	}()
	// 	close(ch)
	// }
}

func ExampleComment() {
	f := NewFile("a")
	f.Comment("Foo returns the string \"foo\"")
	f.Func().Id("Foo").Params().String().Block(
		Return(Lit("foo")).Comment("return the string foo"),
	)
	fmt.Printf("%#v", f)
	// Output:
	// package a
	//
	// // Foo returns the string "foo"
	// func Foo() string {
	// 	return "foo" // return the string foo
	// }
}

func ExampleComment_multiline() {
	c := Comment("a\nb")
	fmt.Printf("%#v", c)
	// Output:
	// /*
	// a
	// b
	// */
}

func ExampleComment_formatting_disabled() {
	c := Id("foo").Call(Comment("/* inline */")).Comment("//no-space")
	fmt.Printf("%#v", c)
	// Output:
	// foo( /* inline */ ) //no-space
}

func ExampleCommentf() {
	name := "foo"
	val := "bar"
	c := Id(name).Op(":=").Lit(val).Commentf("%s is the string \"%s\"", name, val)
	fmt.Printf("%#v", c)
	// Output:
	// foo := "bar" // foo is the string "bar"
}

func ExampleComplex() {
	c := Func().Id("main").Params().Block(
		Id("c1").Op(":=").Lit(1+3.75i),
		Id("c2").Op(":=").Complex(Lit(1), Lit(3.75)),
		Qual("fmt", "Println").Call(Id("c1").Op("==").Id("c2")),
	)
	fmt.Printf("%#v", c)
	// Output:
	// func main() {
	// 	c1 := (1 + 3.75i)
	// 	c2 := complex(1, 3.75)
	// 	fmt.Println(c1 == c2)
	// }
}

func ExampleComplex128() {
	c := Func().Id("main").Params().Block(
		Var().Id("c").Complex128(),
		Id("c").Op("=").Lit(1+2i),
		Qual("fmt", "Println").Call(Id("c")),
	)
	fmt.Printf("%#v", c)
	// Output:
	// func main() {
	// 	var c complex128
	// 	c = (1 + 2i)
	// 	fmt.Println(c)
	// }
}

func ExampleComplex64() {
	c := Func().Id("main").Params().Block(
		Var().Id("c64").Complex64(),
		Id("c64").Op("=").Complex(Lit(5), Float32().Parens(Lit(2))),
		Qual("fmt", "Printf").Call(Lit("%T\n"), Id("c64")),
	)
	fmt.Printf("%#v", c)
	// Output:
	// func main() {
	// 	var c64 complex64
	// 	c64 = complex(5, float32(2))
	// 	fmt.Printf("%T\n", c64)
	// }
}

/*
func Const() *Statement
func Continue() *Statement
func Copy(c ...Code) *Statement
func Default() *Statement
func Defer() *Statement
func Defs(c ...Code) *Statement
func DefsFunc(f func(*Group)) *Statement
func Delete(c ...Code) *Statement
func Dict(m map[Code]Code) *Statement
func DictFunc(f func(map[Code]Code)) *Statement
func Do(f func(*Statement)) *Statement
func Else() *Statement
func Empty() *Statement
func Err() *Statement
func Error() *Statement
func Fallthrough() *Statement
func False() *Statement
func Float32() *Statement
func Float64() *Statement
func For(c ...Code) *Statement
func ForFunc(f func(*Group)) *Statement
func Func() *Statement
func Go() *Statement
func Goto() *Statement
func Id(name string) *Statement
func If(c ...Code) *Statement
func IfFunc(f func(*Group)) *Statement
func Imag(c ...Code) *Statement
func Index(c ...Code) *Statement
func IndexFunc(f func(*Group)) *Statement
func Int() *Statement
func Int16() *Statement
func Int32() *Statement
func Int64() *Statement
func Int8() *Statement
func Interface(c ...Code) *Statement
func InterfaceFunc(f func(*Group)) *Statement
func Iota() *Statement
func Len(c ...Code) *Statement
func Line() *Statement
func List(c ...Code) *Statement
func ListFunc(f func(*Group)) *Statement
func Lit(v interface{}) *Statement
func LitFunc(f func() interface{}) *Statement
func Make(c ...Code) *Statement
func Map(c Code) *Statement
func New(c ...Code) *Statement
func Nil() *Statement
func Null() *Statement
func Op(op string) *Statement
func Panic(c ...Code) *Statement
func Params(c ...Code) *Statement
func ParamsFunc(f func(*Group)) *Statement
func Parens(c Code) *Statement
func Print(c ...Code) *Statement
func Println(c ...Code) *Statement
func Qual(path, name string) *Statement
func Range() *Statement
func Real(c ...Code) *Statement
func Recover(c ...Code) *Statement
func Return(c ...Code) *Statement
func ReturnFunc(f func(*Group)) *Statement
func Rune() *Statement
func Sel(c ...Code) *Statement
func SelFunc(f func(*Group)) *Statement
func Select() *Statement
func String() *Statement
func Struct() *Statement
func Switch(c ...Code) *Statement
func SwitchFunc(f func(*Group)) *Statement
func Tag(items map[string]string) *Statement
func True() *Statement
func Type() *Statement
func Uint() *Statement
func Uint16() *Statement
func Uint32() *Statement
func Uint64() *Statement
func Uint8() *Statement
func Uintptr() *Statement
func Values(c ...Code) *Statement
func ValuesFunc(f func(*Group)) *Statement
func Var() *Statement
*/

func ExampleParams() {
	c := Func().Params(
		Id("a").Id("A"),
	).Id("foo").Params(
		Id("b"),
		Id("c").String(),
	).String().Block(
		Return(Id("b").Op("+").Id("c")),
	)
	fmt.Printf("%#v", c)
	// Output:
	// func (a A) foo(b, c string) string {
	// 	return b + c
	// }
}

func ExampleIndex() {
	c := Var().Id("a").Index().String()
	fmt.Printf("%#v", c)
	// Output:
	// var a []string
}

func ExampleIndex_index() {
	c := Id("a").Op(":=").Id("b").Index(Lit(0), Lit(1))
	fmt.Printf("%#v", c)
	// Output:
	// a := b[0:1]
}

func ExampleIndex_empty() {
	c := Id("a").Op(":=").Id("b").Index(Lit(1), Empty())
	fmt.Printf("%#v", c)
	// Output:
	// a := b[1:]
}

func ExampleOp() {
	c := Id("a").Op(":=").Id("b").Call()
	fmt.Printf("%#v", c)
	// Output:
	// a := b()
}

func ExampleOp_star() {
	c := Id("a").Op("=").Op("*").Id("b")
	fmt.Printf("%#v", c)
	// Output:
	// a = *b
}

func ExampleOp_variadic() {
	c := Id("a").Call(Id("b").Op("..."))
	fmt.Printf("%#v", c)
	// Output:
	// a(b...)
}

func ExampleNewFilePath() {
	f := NewFilePath("a.b/c")
	f.Func().Id("init").Params().Block(
		Qual("a.b/c", "Foo").Call().Comment("Local package - alias is omitted."),
		Qual("d.e/f", "Bar").Call().Comment("Import is automatically added."),
		Qual("g.h/f", "Baz").Call().Comment("Colliding package name is automatically renamed."),
	)
	fmt.Printf("%#v", f)
	// Output:
	// package c
	//
	// import (
	// 	f "d.e/f"
	// 	f1 "g.h/f"
	// )
	//
	// func init() {
	// 	Foo()    // Local package - alias is omitted.
	// 	f.Bar()  // Import is automatically added.
	// 	f1.Baz() // Colliding package name is automatically renamed.
	// }
}

func ExampleStruct_empty() {
	c := Id("c").Op(":=").Make(Chan().Struct())
	fmt.Printf("%#v", c)
	// Output:
	// c := make(chan struct{})
}

func ExampleStruct() {
	c := Type().Id("foo").Struct(
		List(Id("x"), Id("y")).Int(),
		Id("u").Float32(),
	)
	fmt.Printf("%#v", c)
	// Output:
	// type foo struct {
	// 	x, y int
	// 	u    float32
	// }
}

func ExampleDefer() {
	c := Defer().Id("foo").Call()
	fmt.Printf("%#v", c)
	// Output:
	// defer foo()
}

func ExampleGoto() {
	c := Goto().Id("Outer")
	fmt.Printf("%#v", c)
	// Output:
	// goto Outer
}

func ExampleStatement_Clone_broken() {
	a := Id("a")
	c := Block(
		a.Call(),
		a.Call(),
	)
	fmt.Printf("%#v", c)
	// Output:
	// {
	// 	a()()
	// 	a()()
	// }
}

func ExampleStatement_Clone_fixed() {
	a := Id("a")
	c := Block(
		a.Clone().Call(),
		a.Clone().Call(),
	)
	fmt.Printf("%#v", c)
	// Output:
	// {
	// 	a()
	// 	a()
	// }
}

func ExampleFile_Render() {
	f := NewFile("a")
	f.Func().Id("main").Params().Block()
	buf := &bytes.Buffer{}
	err := f.Render(buf)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(buf.String())
	}
	// Output:
	// package a
	//
	// func main() {}
}

func ExampleLit() {
	c := Id("a").Op(":=").Lit("a")
	fmt.Printf("%#v", c)
	// Output:
	// a := "a"
}

func ExampleLit_float() {
	c := Id("a").Op(":=").Lit(1.5)
	fmt.Printf("%#v", c)
	// Output:
	// a := 1.5
}

func ExampleLitFunc() {
	c := Id("a").Op(":=").LitFunc(func() interface{} { return 1 + 1 })
	fmt.Printf("%#v", c)
	// Output:
	// a := 2
}

func ExampleSel() {
	c := Sel(
		Qual("a.b/c", "Foo").Call(),
		Id("Bar").Index(Lit(0)),
		Id("Baz"),
	)
	fmt.Printf("%#v", c)
	// Output:
	// c.Foo().Bar[0].Baz
}

func ExampleList() {
	c := List(Id("a"), Err()).Op(":=").Id("b").Call()
	fmt.Printf("%#v", c)
	// Output:
	// a, err := b()
}

func ExampleQual() {
	c := Qual("encoding/gob", "NewEncoder").Call()
	fmt.Printf("%#v", c)
	// Output:
	// gob.NewEncoder()
}

func ExampleQual_file() {
	f := NewFilePath("a.b/c")
	f.Func().Id("init").Params().Block(
		Qual("a.b/c", "Foo").Call().Comment("Local package - name is omitted."),
		Qual("d.e/f", "Bar").Call().Comment("Import is automatically added."),
		Qual("g.h/f", "Baz").Call().Comment("Colliding package name is renamed."),
	)
	fmt.Printf("%#v", f)
	// Output:
	// package c
	//
	// import (
	// 	f "d.e/f"
	// 	f1 "g.h/f"
	// )
	//
	// func init() {
	// 	Foo()    // Local package - name is omitted.
	// 	f.Bar()  // Import is automatically added.
	// 	f1.Baz() // Colliding package name is renamed.
	// }
}

func ExampleQual_local() {
	f := NewFilePath("a.b/c")
	f.Func().Id("main").Params().Block(
		Qual("a.b/c", "D").Call(),
	)
	fmt.Printf("%#v", f)
	// Output:
	// package c
	//
	// func main() {
	// 	D()
	// }
}

func ExampleId() {
	c := If(Id("i").Op("==").Id("j")).Block(
		Return(Id("i")),
	)
	fmt.Printf("%#v", c)
	// Output:
	// if i == j {
	// 	return i
	// }
}

func ExampleErr() {
	c := If(
		Err().Op(":=").Id("foo").Call(),
		Err().Op("!=").Nil(),
	).Block(
		Return(Err()),
	)
	fmt.Printf("%#v", c)
	// Output:
	// if err := foo(); err != nil {
	// 	return err
	// }
}

func ExampleSwitch() {
	c := Switch(Id("a")).Block(
		Case(Lit("1")).Block(
			Return(Lit(1)),
		),
		Case(Lit("2"), Lit("3")).Block(
			Return(Lit(2)),
		),
		Case(Lit("4")).Block(
			Fallthrough(),
		),
		Default().Block(
			Return(Lit(3)),
		),
	)
	fmt.Printf("%#v", c)
	// Output:
	// switch a {
	// case "1":
	// 	return 1
	// case "2", "3":
	// 	return 2
	// case "4":
	// 	fallthrough
	// default:
	// 	return 3
	// }
}

func ExampleTag() {
	// Note: Tags are ordered by key when rendered
	c := Type().Id("foo").Struct(
		Id("A").String().Tag(map[string]string{"json": "a"}),
		Id("B").Int().Tag(map[string]string{"json": "b", "bar": "baz"}),
	)
	fmt.Printf("%#v", c)
	// Output:
	// type foo struct {
	// 	A string `json:"a"`
	// 	B int    `bar:"baz" json:"b"`
	// }
}

func ExampleNull_and_nil() {
	c := Func().Id("foo").Params(
		nil,
		Id("s").String(),
		Null(),
		Id("i").Int(),
	).Block()
	fmt.Printf("%#v", c)
	// Output:
	// func foo(s string, i int) {}
}

func ExampleNull_index() {
	c := Id("a").Op(":=").Id("b").Index(Lit(1), Null())
	fmt.Printf("%#v", c)
	// Output:
	// a := b[1]
}

func ExampleEmpty() {
	c := Id("a").Op(":=").Id("b").Index(Lit(1), Empty())
	fmt.Printf("%#v", c)
	// Output:
	// a := b[1:]
}

func ExampleBlock_complex() {
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
	// Output:
	// func main() {
	// 	var foo []string
	// 	var bar map[string]int
	// }
}

func ExampleFunc_declaration() {
	c := Func().Id("a").Params().Block()
	fmt.Printf("%#v", c)
	// Output:
	// func a() {}
}

func ExampleFunc_literal() {
	c := Id("a").Op(":=").Func().Params().Block()
	fmt.Printf("%#v", c)
	// Output:
	// a := func() {}
}

func ExampleInterface() {
	c := Type().Id("a").Interface(
		Id("b").Params().String(),
	)
	fmt.Printf("%#v", c)
	// Output:
	// type a interface {
	// 	b() string
	// }
}

func ExampleInterface_empty() {
	c := Var().Id("a").Interface()
	fmt.Printf("%#v", c)
	// Output:
	// var a interface{}
}

func ExampleParens() {
	c := Id("b").Op(":=").Index().Byte().Parens(Id("s"))
	fmt.Printf("%#v", c)
	// Output:
	// b := []byte(s)
}

func ExampleParens_order() {
	c := Id("a").Op("/").Parens(Id("b").Op("+").Id("c"))
	fmt.Printf("%#v", c)
	// Output:
	// a / (b + c)
}

func ExampleValues() {
	c := Index().String().Values(Lit("a"), Lit("b"))
	fmt.Printf("%#v", c)
	// Output:
	// []string{"a", "b"}
}

func ExampleDo() {
	f := func(name string, isMap bool) *Statement {
		return Id(name).Op(":=").Do(func(s *Statement) {
			if isMap {
				s.Map(String()).String()
			} else {
				s.Index().String()
			}
		}).Values()
	}
	fmt.Printf("%#v\n%#v", f("a", true), f("b", false))
	// Output:
	// a := map[string]string{}
	// b := []string{}
}

func ExampleReturn() {
	c := Return(Id("a"), Id("b"))
	fmt.Printf("%#v", c)
	// Output:
	// return a, b
}

func ExampleMap() {
	c := Id("a").Op(":=").Map(String()).String().Values()
	fmt.Printf("%#v", c)
	// Output:
	// a := map[string]string{}
}

func ExampleDict() {
	c := Id("a").Op(":=").Map(String()).String().Dict(map[Code]Code{
		Lit("a"): Lit("b"),
		Lit("c"): Lit("d"),
	})
	fmt.Printf("%#v", c)
	// Output:
	// a := map[string]string{
	// 	"a": "b",
	// 	"c": "d",
	// }
}

func ExampleDict_nil() {
	c := Id("a").Op(":=").Map(String()).String().Dict(nil)
	fmt.Printf("%#v", c)
	// Output:
	// a := map[string]string{}
}

func ExampleDictFunc() {
	c := Id("a").Op(":=").Map(String()).String().DictFunc(func(m map[Code]Code) {
		m[Lit("a")] = Lit("b")
		m[Lit("c")] = Lit("d")
	})
	fmt.Printf("%#v", c)
	// Output:
	// a := map[string]string{
	// 	"a": "b",
	// 	"c": "d",
	// }
}

func ExampleDefs() {
	c := Const().Defs(
		Id("a").Op("=").Lit("a"),
		Id("b").Op("=").Lit("b"),
	)
	fmt.Printf("%#v", c)
	// Output:
	// const (
	// 	a = "a"
	// 	b = "b"
	// )
}

func ExampleIf() {
	c := If(
		Err().Op(":=").Id("a").Call(),
		Err().Op("!=").Nil(),
	).Block(
		Return(Err()),
	)
	fmt.Printf("%#v", c)
	// Output:
	// if err := a(); err != nil {
	// 	return err
	// }
}

func ExampleId_local() {
	c := Id("a").Op(":=").Lit(1)
	fmt.Printf("%#v", c)
	// Output:
	// a := 1
}

func ExampleId_select() {
	c := Sel(Id("a"), Id("b"), Id("c")).Call()
	fmt.Printf("%#v", c)
	// Output:
	// a.b.c()
}

func ExampleId_remote() {
	f := NewFile("main")
	f.Func().Id("main").Params().Block(
		Qual("fmt", "Println").Call(
			Lit("Hello, world"),
		),
	)
	fmt.Printf("%#v", f)
	// Output:
	// package main
	//
	// import fmt "fmt"
	//
	// func main() {
	// 	fmt.Println("Hello, world")
	// }
}

func ExampleFor() {
	c := For(
		Id("i").Op(":=").Lit(0),
		Id("i").Op("<").Lit(10),
		Id("i").Op("++"),
	).Block(
		Qual("fmt", "Println").Call(Id("i")),
	)
	fmt.Printf("%#v", c)
	// Output:
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
}

func ExampleNewFile() {
	f := NewFile("main")
	f.Func().Id("main").Params().Block(
		Qual("fmt", "Println").Call(
			Lit("Hello, world"),
		),
	)
	fmt.Printf("%#v", f)
	// Output:
	// package main
	//
	// import fmt "fmt"
	//
	// func main() {
	// 	fmt.Println("Hello, world")
	// }
}

func ExampleNewFilePathName() {
	f := NewFilePathName("a.b/c", "main")
	f.Func().Id("main").Params().Block(
		Qual("a.b/c", "Foo").Call(),
	)
	fmt.Printf("%#v", f)
	// Output:
	// package main
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
	// Output:
	// // a
	// // b
	// package c
	//
	// func init() {}
}

func ExampleFile_Anon() {
	f := NewFile("c")
	f.Anon("a")
	f.Func().Id("init").Params().Block()
	fmt.Printf("%#v", f)
	// Output:
	// package c
	//
	// import _ "a"
	//
	// func init() {}
}

func ExampleFile_PackagePrefix() {
	f := NewFile("c")
	f.PackagePrefix = "pkg"
	f.Func().Id("main").Params().Block(
		Qual("fmt", "Println").Call(),
	)
	fmt.Printf("%#v", f)
	// Output:
	// package c
	//
	// import pkg_fmt "fmt"
	//
	// func main() {
	// 	pkg_fmt.Println()
	// }
}
