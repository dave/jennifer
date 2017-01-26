# Jennifer

Jennifer is a code generator for go:

```go
package main

import (
	"fmt"

	. "github.com/davelondon/jennifer/jen"
)

func main() {
	f := NewFile("main")
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
package main

import fmt "fmt"

func main() {
    fmt.Println("Hello, world")
}
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
	f := NewFilePath("a.b/c")
    f.Func().Id("init").Params().Block(
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

func init() {
    Local()
    f.Remote()
    f1.Collision()
}
```

# NewFile

NewFile Creates a new file. NewFilePath creates a new file while specifying the 
package path - the package name is inferred from the path. NewFilePathName 
additionally specifies the package name.

```go
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
```

# Identifiers 

Identifiers are simple methods with no parameters. They render as the 
identifier token:

```go
c := Break()
fmt.Printf("%#v", c)
// Output: break
```

Keywords: `Break`, `Default`, `Func`, `Interface`, `Select`, `Case`, `Defer`, `Go`, `Struct`, `Chan`, `Else`, `Goto`, `Switch`, `Const`, `Fallthrough`, `Range`, `Type`, `Continue`, `Var`

Built-in types: `Bool`, `Byte`, `Complex64`, `Complex128`, `Error`, `Float32`, `Float64`, `Int`, `Int8`, `Int16`, `Int32`, `Int64`, `Rune`, `String`, `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`, `Uintptr`

Constants: `True`, `False`, `Iota`, `Nil`

Also included is `Err` for the commonly used `err` variable.

Note: `Map`, `Return`, `For` and `If` are special cases, and treated as 
blocks - see below.

Note: The `import` and `package` keywords are always rendered automatically, so 
not included.

# Built-in functions

Built in functions take a variadic slice of code items, and render the function
name followed by the items as a comma seperated list of parameters in 
parenthesis:

```go
c := Append(Id("foo"), Id("bar"))
fmt.Printf("%#v", c)
// Output: append(foo, bar)
```

Functions: `Append`, `Cap`, `Close`, `Complex`, `Copy`, `Delete`, `Imag`, `Len`, `Make`, `New`, `Panic`, `Print`, `Println`, `Real`, `Recover`

# Blocks

Blocks take either a single code item or a varidic list of code items. The 
items are rendered between open and closing toekns. Multiple items are 
seperated by a seperator token:

| Block  | Seperator | Opening  | Closing |
| ------ | --------- | -------- | ------- |
| List   | `,`       |          |         |
| Parens | n/a       | `(`      | `)`     |
| Call   | `,`       | `(`      | `)`     |
| Params | `,`       | `(`      | `)`     |
| Decls  | `;`       | `(`      | `)`     |
| Values | `,`       | `{`      | `}`     |
| Slice  | `,`       | `{`      | `}`     |
| Index  | `:`       | `[`      | `]`     |
| Block  | `\n`      | `{\n`    | `}`     |
| Case   | `\n`      | `:\n`    |         |
| Assert | n/a       | `.(`     | `)`     |
| Map    | n/a       | `map[`   | `]`     |
| If     | `;`       | `if`     |         |
| Return | `,`       | `return` |         |
| For    | `;`       | `for`    |         |

### List
List renders a comma seperated list with no open or closing tokens. Use for 
multiple return functions:

```go
c := List(Id("a"), Id("b")).Op(":=").Id("c").Call()
fmt.Printf("%#v", c)
// Output: a, b := c()
```

### Parens
Parens renders a single code item in parenthesis. Use for type conversion or 
logical grouping:

```go
c := Id("b").Op(":=").Index().Byte().Parens(Id("s"))
fmt.Printf("%#v", c)
// Output: b := []byte(s)
```

```go
c := Parens(Id("a").Op("/").Id("b")).Op("*").Id("c")
fmt.Printf("%#v", c)
// Output: (a / b) * c
```

### Values
Values renders a comma seperated list enclosed by curly braces. Use for slice 
literals:

```go
c := Index().String().Values(Lit("a"), Lit("b"))
fmt.Printf("%#v", c)
// Output: []string{"a", "b"}
```

