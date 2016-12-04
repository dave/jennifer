# Jennifer

Jennifer is a code generator for go:

```go
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
```

Output:

```go
package main

import fmt "fmt"

func main() { fmt.Println("Hello, world") }
```

# Imports

Jennifer manages your imports and aliases:

```go
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

	if err := jen.Render(ctx, f, os.Stdout); err != nil {
		panic(err)
	}
}
```

Output:

```go
package baz

import (
	qux1 "bar.foo/qux"
	qux "foo.bar/qux"
)

func init() {
	LocalFunction()
	qux.RemoteFunction()
	qux1.PackageCollision()
}
```
