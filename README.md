# Jennifer

Jennifer is a code generator for go:

```go
package main

import (
	"fmt"

	. "github.com/davelondon/jennifer/jen"
)

func main() {
	f := NewFile("a")
    f.Func().Id("main").Params().Block(
        Id("fmt.Println").Call(
            Lit("Hello, world"),
        ),
    )
    fmt.Printf("%#v", f)
}
```

Output:

```go
package a

import fmt "fmt"

func main() { fmt.Println("Hello, world") }
```

# Imports

Jennifer manages your imports and aliases:

```go
package main

import (
	"fmt"

	. "github.com/davelondon/jennifer/jen"
)

func main() {
	f := NewFilePath("c", "a.b/c")
    f.Func().Id("main").Params().Block(
        Id("a.b/c.Local").Call(),
        Id("d.e/f.Remote").Call(),
        Id("g.h/f.Collision").Call(),
    )
    fmt.Printf("%#v", f)
}
```

Output:

```go
package c

import (
    f "d.e/f"
    f1 "g.h/f"
)

func main() {
    Local()
    f.Remote()
    f1.Collision()
}
```
