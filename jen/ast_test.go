package jen

import (
	"bytes"
	"testing"
)

const testReuseGeneratedFile = `package main

import (
	"fmt"
	"time"
)

func main() {
	if time.Now() == time.Now() {
		panic(time.Now())
	}
	fmt.Println("example function")
}

func randomFunc() {
	panic("implement me")
}
`

func TestReuseGeneratedFile(t *testing.T) {
	f := NewFile("main")
	f.Line()
	f.Func().Id("main").Params().Block(
		If(Qual("time", "Now").Call().Op("==").Qual("time", "Now").Call()).Block(
			Panic(Qual("time", "Now").Call()),
		),
		Qual("fmt", "Println").Call(Lit("example function")),
	)

	var b bytes.Buffer
	err := f.Render(&b)
	if err != nil {
		t.Fatal(err)
	}
	f2, err := NewFileFromSource(b.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	var b2 bytes.Buffer
	if err := f2.Render(&b2); err != nil {
		t.Fatal(err)
	}
	if b.String() != b2.String() {
		t.Log("=======error======")
		t.Fatal("content of original and parsed files is not identical")
		t.Log("=======want=======")
		t.Log(b.String())
		t.Log("========got=======")
		t.Log(b2.String())
	}
	f2.Line()
	f2.Line().Func().Id("randomFunc").Params().Block(
		Panic(Lit("implement me")),
	)
	b2.Reset()
	if err := f2.Render(&b2); err != nil {
		t.Fatal(err)
	}
	if b2.String() != testReuseGeneratedFile {
		t.Log("=======error======")
		t.Fatal("content of the second file is not identical to test case")
		t.Log("=======want=======")
		t.Log(testReuseGeneratedFile)
		t.Log("========got=======")
		t.Log(b2.String())
	}
}
