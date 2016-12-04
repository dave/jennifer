package main

import (
	"strings"

	"context"
	"os"

	. "github.com/davelondon/jennifer/jen"
	"github.com/davelondon/jennifer/jen/data"
)

func main() {

	file := NewFile()
	for _, keyword := range data.Keywords {
		name := strings.ToUpper(keyword[:1]) + keyword[1:]
		c := "c"
		s := "s"
		code := "Code"
		statement := "Statement"
		// func Foo(c ...Code) *Statement {
		file.Func().Id(name).Params(
			Id(c).Vari().Id(code),
		).Ptr().Id(statement).Block(
			// s := new(Statement)
			Id(s).Sas().New(Id(statement)),
			// return s.Foo(c...)
			Return().Id(s).Method(name, Id(c).Vari()),
		)
	}

	ctx := Context(context.Background(), "jen")
	err := Render(ctx, file, os.Stdout)
	if err != nil {
		panic(err)
	}

}
