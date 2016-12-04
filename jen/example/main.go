package main

import (
	"context"
	"os"

	"github.com/davelondon/jennifer/jen"
)

func main() {
	ctx := jen.Context(context.Background(), "main")

	f := jen.NewFile()
	f.Func().Id("main").Params().Block(
		jen.Id("fmt.Println").Call(
			jen.Lit("Hello, world"),
		),
	)

	if err := jen.Render(ctx, f, os.Stdout); err != nil {
		panic(err)
	}
}
