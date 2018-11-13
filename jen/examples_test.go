package jen_test

import (
	"fmt"

	"bytes"

	. "github.com/dave/jennifer/jen"
)

func ExampleCaseBug() {
	c := Switch(Id("a")).Block(
		Case(Lit(1)).Block(
			Var().Id("i").Int(),
			Var().Id("j").Int(),
		),
	)
	fmt.Printf("%#v", c)
	// Output:
	// switch a {
	// case 1:
	// 	var i int
	// 	var j int
	// }
}

func ExampleCustom() {
	multiLineCall := Options{
		Close:     ")",
		Multi:     true,
		Open:      "(",
		Separator: ",",
	}
	c := Id("foo").Custom(multiLineCall, Lit("a"), Lit("b"), Lit("c"))
	fmt.Printf("%#v", c)
	// Output:
	// foo(
	// 	"a",
	// 	"b",
	// 	"c",
	// )
}

func ExampleCustomFunc() {
	multiLineCall := Options{
		Close:     ")",
		Multi:     true,
		Open:      "(",
		Separator: ",",
	}
	c := Id("foo").CustomFunc(multiLineCall, func(g *Group) {
		g.Lit("a")
		g.Lit("b")
		g.Lit("c")
	})
	fmt.Printf("%#v", c)
	// Output:
	// foo(
	// 	"a",
	// 	"b",
	// 	"c",
	// )
}

func ExampleFile_ImportName_conflict() {
	f := NewFile("main")

	// We provide a hint that package foo/a should use name "a", but because package bar/a already
	// registers the required name, foo/a is aliased.
	f.ImportName("github.com/foo/a", "a")

	f.Func().Id("main").Params().Block(
		Qual("github.com/bar/a", "Bar").Call(),
		Qual("github.com/foo/a", "Foo").Call(),
	)
	fmt.Printf("%#v", f)

	// Output:
	// package main
	//
	// import (
	// 	a "github.com/bar/a"
	// 	a1 "github.com/foo/a"
	// )
	//
	// func main() {
	// 	a.Bar()
	// 	a1.Foo()
	// }
}

func ExampleFile_ImportAlias_conflict() {
	f := NewFile("main")

	// We provide a hint that package foo/a should use alias "b", but because package bar/b already
	// registers the required name, foo/a is aliased using the requested alias as a base.
	f.ImportName("github.com/foo/a", "b")

	f.Func().Id("main").Params().Block(
		Qual("github.com/bar/b", "Bar").Call(),
		Qual("github.com/foo/a", "Foo").Call(),
	)
	fmt.Printf("%#v", f)

	// Output:
	// package main
	//
	// import (
	// 	b "github.com/bar/b"
	// 	b1 "github.com/foo/a"
	// )
	//
	// func main() {
	// 	b.Bar()
	// 	b1.Foo()
	// }
}

func ExampleFile_ImportName() {
	f := NewFile("main")

	// package a should use name "a"
	f.ImportName("github.com/foo/a", "a")

	// package b is not used in the code so will not be included
	f.ImportName("github.com/foo/b", "b")

	f.Func().Id("main").Params().Block(
		Qual("github.com/foo/a", "A").Call(),
	)
	fmt.Printf("%#v", f)

	// Output:
	// package main
	//
	// import "github.com/foo/a"
	//
	// func main() {
	// 	a.A()
	// }
}

func ExampleFile_ImportNames() {

	// package a should use name "a", package b is not used in the code so will not be included
	names := map[string]string{
		"github.com/foo/a": "a",
		"github.com/foo/b": "b",
	}

	f := NewFile("main")
	f.ImportNames(names)
	f.Func().Id("main").Params().Block(
		Qual("github.com/foo/a", "A").Call(),
	)
	fmt.Printf("%#v", f)

	// Output:
	// package main
	//
	// import "github.com/foo/a"
	//
	// func main() {
	// 	a.A()
	// }
}

func ExampleFile_ImportAlias() {
	f := NewFile("main")

	// package a should be aliased to "b"
	f.ImportAlias("github.com/foo/a", "b")

	// package c is not used in the code so will not be included
	f.ImportAlias("github.com/foo/c", "c")

	f.Func().Id("main").Params().Block(
		Qual("github.com/foo/a", "A").Call(),
	)
	fmt.Printf("%#v", f)

	// Output:
	// package main
	//
	// import b "github.com/foo/a"
	//
	// func main() {
	// 	b.A()
	// }
}

func ExampleFile_ImportAliasDot() {
	f := NewFile("main")

	// package a should be a dot-import
	f.ImportAlias("github.com/foo/a", ".")

	// package b should be a dot-import
	f.ImportAlias("github.com/foo/b", ".")

	// package c is not used in the code so will not be included
	f.ImportAlias("github.com/foo/c", ".")

	f.Func().Id("main").Params().Block(
		Qual("github.com/foo/a", "A").Call(),
		Qual("github.com/foo/b", "B").Call(),
	)
	fmt.Printf("%#v", f)

	// Output:
	// package main
	//
	// import (
	// 	. "github.com/foo/a"
	// 	. "github.com/foo/b"
	// )
	//
	// func main() {
	// 	A()
	// 	B()
	// }
}

func ExampleFile_CgoPreamble() {
	f := NewFile("a")
	f.CgoPreamble(`#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
	printf("%s\n", s);
}
`)
	f.Func().Id("init").Params().Block(
		Id("cs").Op(":=").Qual("C", "CString").Call(Lit("Hello from stdio\n")),
		Qual("C", "myprint").Call(Id("cs")),
		Qual("C", "free").Call(Qual("unsafe", "Pointer").Parens(Id("cs"))),
	)
	fmt.Printf("%#v", f)
	// Output:
	// package a
	//
	// import "unsafe"
	//
	// /*
	// #include <stdio.h>
	// #include <stdlib.h>
	//
	// void myprint(char* s) {
	// 	printf("%s\n", s);
	// }
	// */
	// import "C"
	//
	// func init() {
	// 	cs := C.CString("Hello from stdio\n")
	// 	C.myprint(cs)
	// 	C.free(unsafe.Pointer(cs))
	// }
}

func ExampleFile_CgoPreamble_anon() {
	f := NewFile("a")
	f.CgoPreamble(`#include <stdio.h>`)
	f.Func().Id("init").Params().Block(
		Qual("foo.bar/a", "A"),
		Qual("foo.bar/b", "B"),
	)
	fmt.Printf("%#v", f)
	// Output:
	// package a
	//
	// import (
	// 	a "foo.bar/a"
	// 	b "foo.bar/b"
	// )
	//
	// // #include <stdio.h>
	// import "C"
	//
	// func init() {
	// 	a.A
	// 	b.B
	// }
}

func ExampleFile_CgoPreamble_no_preamble() {
	f := NewFile("a")
	f.Func().Id("init").Params().Block(
		Qual("C", "Foo").Call(),
		Qual("fmt", "Print").Call(),
	)
	fmt.Printf("%#v", f)
	// Output:
	// package a
	//
	// import (
	// 	"C"
	// 	"fmt"
	// )
	//
	// func init() {
	// 	C.Foo()
	// 	fmt.Print()
	// }
}

func ExampleFile_CgoPreamble_no_preamble_single() {
	f := NewFile("a")
	f.Func().Id("init").Params().Block(
		Qual("C", "Foo").Call(),
	)
	fmt.Printf("%#v", f)
	// Output:
	// package a
	//
	// import "C"
	//
	// func init() {
	// 	C.Foo()
	// }
}

func ExampleFile_CgoPreamble_no_preamble_single_anon() {
	f := NewFile("a")
	f.Anon("C")
	f.Func().Id("init").Params().Block()
	fmt.Printf("%#v", f)
	// Output:
	// package a
	//
	// import "C"
	//
	// func init() {}
}

func ExampleFile_CgoPreamble_no_preamble_anon() {
	f := NewFile("a")
	f.Anon("C")
	f.Func().Id("init").Params().Block(
		Qual("fmt", "Print").Call(),
	)
	fmt.Printf("%#v", f)
	// Output:
	// package a
	//
	// import (
	// 	"C"
	// 	"fmt"
	// )
	//
	// func init() {
	// 	fmt.Print()
	// }
}

func ExampleOp_complex_conditions() {
	c := If(Parens(Id("a").Op("||").Id("b")).Op("&&").Id("c")).Block()
	fmt.Printf("%#v", c)
	// Output:
	// if (a || b) && c {
	// }
}

func ExampleLit_bool_true() {
	c := Lit(true)
	fmt.Printf("%#v", c)
	// Output:
	// true
}

func ExampleLit_bool_false() {
	c := Lit(false)
	fmt.Printf("%#v", c)
	// Output:
	// false
}

func ExampleLit_byte() {
	// Lit can't tell the difference between byte and uint8. Use LitByte to
	// render byte literals.
	c := Lit(byte(0x1))
	fmt.Printf("%#v", c)
	// Output:
	// uint8(0x1)
}

func ExampleLit_complex64() {
	c := Lit(complex64(0 + 0i))
	fmt.Printf("%#v", c)
	// Output:
	// complex64(0 + 0i)
}

func ExampleLit_complex128() {
	c := Lit(0 + 0i)
	fmt.Printf("%#v", c)
	// Output:
	// (0 + 0i)
}

func ExampleLit_float32() {
	c := Lit(float32(1))
	fmt.Printf("%#v", c)
	// Output:
	// float32(1)
}

func ExampleLit_float64_one_point_zero() {
	c := Lit(1.0)
	fmt.Printf("%#v", c)
	// Output:
	// 1.0
}

func ExampleLit_float64_zero() {
	c := Lit(0.0)
	fmt.Printf("%#v", c)
	// Output:
	// 0.0
}

func ExampleLit_float64_negative() {
	c := Lit(-0.1)
	fmt.Printf("%#v", c)
	// Output:
	// -0.1
}

func ExampleLit_float64_negative_whole() {
	c := Lit(-1.0)
	fmt.Printf("%#v", c)
	// Output:
	// -1.0
}

func ExampleLit_int() {
	c := Lit(1)
	fmt.Printf("%#v", c)
	// Output:
	// 1
}

func ExampleLit_int8() {
	c := Lit(int8(1))
	fmt.Printf("%#v", c)
	// Output:
	// int8(1)
}

func ExampleLit_int16() {
	c := Lit(int16(1))
	fmt.Printf("%#v", c)
	// Output:
	// int16(1)
}

func ExampleLit_int32() {
	c := Lit(int32(1))
	fmt.Printf("%#v", c)
	// Output:
	// int32(1)
}

func ExampleLit_int64() {
	c := Lit(int64(1))
	fmt.Printf("%#v", c)
	// Output:
	// int64(1)
}

func ExampleLit_uint() {
	c := Lit(uint(0x1))
	fmt.Printf("%#v", c)
	// Output:
	// uint(0x1)
}

func ExampleLit_uint8() {
	c := Lit(uint8(0x1))
	fmt.Printf("%#v", c)
	// Output:
	// uint8(0x1)
}

func ExampleLit_uint16() {
	c := Lit(uint16(0x1))
	fmt.Printf("%#v", c)
	// Output:
	// uint16(0x1)
}

func ExampleLit_uint32() {
	c := Lit(uint32(0x1))
	fmt.Printf("%#v", c)
	// Output:
	// uint32(0x1)
}

func ExampleLit_uint64() {
	c := Lit(uint64(0x1))
	fmt.Printf("%#v", c)
	// Output:
	// uint64(0x1)
}

func ExampleLit_uintptr() {
	c := Lit(uintptr(0x1))
	fmt.Printf("%#v", c)
	// Output:
	// uintptr(0x1)
}

func ExampleLit_rune() {
	// Lit can't tell the difference between rune and int32. Use LitRune to
	// render rune literals.
	c := Lit('x')
	fmt.Printf("%#v", c)
	// Output:
	// int32(120)
}

func ExampleLitRune() {
	c := LitRune('x')
	fmt.Printf("%#v", c)
	// Output:
	// 'x'
}

func ExampleLitRuneFunc() {
	c := LitRuneFunc(func() rune {
		return '\t'
	})
	fmt.Printf("%#v", c)
	// Output:
	// '\t'
}

func ExampleLitByte() {
	c := LitByte(byte(1))
	fmt.Printf("%#v", c)
	// Output:
	// byte(0x1)
}

func ExampleLitByteFunc() {
	c := LitByteFunc(func() byte {
		return byte(2)
	})
	fmt.Printf("%#v", c)
	// Output:
	// byte(0x2)
}

func ExampleLit_string() {
	c := Lit("foo")
	fmt.Printf("%#v", c)
	// Output:
	// "foo"
}

func ExampleValues_dict_single() {
	c := Map(String()).String().Values(Dict{
		Lit("a"): Lit("b"),
	})
	fmt.Printf("%#v", c)
	// Output:
	// map[string]string{"a": "b"}
}

func ExampleValues_dict_multiple() {
	c := Map(String()).String().Values(Dict{
		Lit("a"): Lit("b"),
		Lit("c"): Lit("d"),
	})
	fmt.Printf("%#v", c)
	// Output:
	// map[string]string{
	// 	"a": "b",
	// 	"c": "d",
	// }
}

func ExampleValues_dict_composite() {
	c := Op("&").Id("Person").Values(Dict{
		Id("Age"):  Lit(1),
		Id("Name"): Lit("a"),
	})
	fmt.Printf("%#v", c)
	// Output:
	// &Person{
	// 	Age:  1,
	// 	Name: "a",
	// }
}

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

func ExampleValuesFunc() {
	c := Id("numbers").Op(":=").Index().Int().ValuesFunc(func(g *Group) {
		for i := 0; i <= 5; i++ {
			g.Lit(i)
		}
	})
	fmt.Printf("%#v", c)
	// Output:
	// numbers := []int{0, 1, 2, 3, 4, 5}
}

func ExampleBlockFunc() {
	increment := true
	name := "a"
	c := Func().Id("a").Params().BlockFunc(func(g *Group) {
		g.Id(name).Op("=").Lit(1)
		if increment {
			g.Id(name).Op("++")
		} else {
			g.Id(name).Op("--")
		}
	})
	fmt.Printf("%#v", c)
	// Output:
	// func a() {
	// 	a = 1
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

func ExampleDot() {
	c := Qual("a.b/c", "Foo").Call().Dot("Bar").Index(Lit(0)).Dot("Baz")
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
	c := Switch(Id("value").Dot("Kind").Call()).Block(
		Case(Qual("reflect", "Float32"), Qual("reflect", "Float64")).Block(
			Return(Lit("float")),
		),
		Case(Qual("reflect", "Bool")).Block(
			Return(Lit("bool")),
		),
		Case(Qual("reflect", "Uintptr")).Block(
			Fallthrough(),
		),
		Default().Block(
			Return(Lit("none")),
		),
	)
	fmt.Printf("%#v", c)
	// Output:
	// switch value.Kind() {
	// case reflect.Float32, reflect.Float64:
	// 	return "float"
	// case reflect.Bool:
	// 	return "bool"
	// case reflect.Uintptr:
	// 	fallthrough
	// default:
	// 	return "none"
	// }
}

func ExampleTag() {
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
	c := Id("a").Op(":=").Map(String()).String().Values(Dict{
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
	c := Id("a").Op(":=").Map(String()).String().Values()
	fmt.Printf("%#v", c)
	// Output:
	// a := map[string]string{}
}

func ExampleDictFunc() {
	c := Id("a").Op(":=").Map(String()).String().Values(DictFunc(func(d Dict) {
		d[Lit("a")] = Lit("b")
		d[Lit("c")] = Lit("d")
	}))
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
	c := Id("a").Dot("b").Dot("c").Call()
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
	// import "fmt"
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
		Qual("fmt", "Println").Call(Lit("Hello, world")),
	)
	fmt.Printf("%#v", f)
	// Output:
	// package main
	//
	// import "fmt"
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

func ExampleFile_HeaderAndPackageComments() {
	f := NewFile("c")
	f.CanonicalPath = "d.e/f"
	f.HeaderComment("Code generated by...")
	f.PackageComment("Package c implements...")
	f.Func().Id("init").Params().Block()
	fmt.Printf("%#v", f)
	// Output:
	// // Code generated by...
	//
	// // Package c implements...
	// package c // import "d.e/f"
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
	f := NewFile("a")
	f.PackagePrefix = "pkg"
	f.Func().Id("main").Params().Block(
		Qual("b.c/d", "E").Call(),
	)
	fmt.Printf("%#v", f)
	// Output:
	// package a
	//
	// import pkg_d "b.c/d"
	//
	// func main() {
	// 	pkg_d.E()
	// }
}
