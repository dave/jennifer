package main

import (
	"context"
	"os"

	"github.com/davelondon/jennifer/jen"
)

func main() {
	ctx := jen.ContextPath(
		context.Background(),
		"baz",
		"foo.bar/baz",
	)

	f := jen.NewFile()
	f.Func().Id("init").Params().Block(
		jen.Id("foo.bar/baz.LocalFunction").Call(),
		jen.Id("foo.bar/qux.RemoteFunction").Call(),
		jen.Id("bar.foo/qux.PackageCollision").Call(),
	)

	if err := jen.RenderFile(ctx, f, os.Stdout); err != nil {
		panic(err)
	}
}
