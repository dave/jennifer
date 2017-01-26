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

Creates a new file

# Identifiers 

Identifiers are simple methods with no parameters. They simply output the 
identifier token the code stream:

```go
c := Break()
fmt.Printf("%#v", c)
// Output: break
```

Keywords: "break", "default", "func", "interface", "select", "case", "defer", "go", "struct", "chan", "else", "goto", "switch", "const", "fallthrough", "range", "type", "continue", "var"

Built-in types: "bool", "byte", "complex64", "complex128", "error", "float32", "float64", "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"

Constants: "true", "false", "iota", "nil"

We also include Err() for the commonly used variable "err".

Note: "map", "return", "for" and "if" are special cases, and treated as blocks

Note: "import" and "package" are always handled automatically, so not included

# Built-in functions

Built in functions take a variadic slice of code items, and render the function
name followed by the items as a comma seperated list of parameters in 
parenthesis:

```go
c := Append(Id("foo"), Id("bar"))
fmt.Printf("%#v", c)
// Output: append(foo, bar)
```

Functions: "append", "cap", "close", "complex", "copy", "delete", "imag", "len", "make", "new", "panic", "print", "println", "real", "recover"

# Blocks

Blocks take either a single code item or a varidic list of code items. The 
items are rendered between open and closing toekns. Multiple items are 
seperated by a seperator token:

| Block  | Seperator | Opening | Closing |
| ------ | --------- | ------- | ------- |
| Parens | n/a       | (       | )       |
| List   | ,         |         |         |
| Values | ,         | {       | }       |
| Slice  | ,         | {       | }       |
| Index  | :         | [       | ]       |
| Block  | \n        | {\n     | }       |
| Call   | ,         | (       | )       |
| Params | ,         | (       | )       |
| Decls  | ;         | (       | )       |
| Case   | \n        | :\n     |         |
| Assert | n/a       | .(      | )       |
| Map    | n/a       | map[    | ]       |
| If     | ;         | if      |         |
| Return | ,         | return  |         |
| For    | ;         | for     |         |

### Parens
Parens should be used for enclosing a single code item in parenthesis, usually 
for type conversion or logical grouping:

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

### List
List should be used for a comma seperated list with no open or closing tokens. 
Usually for multiple return methods:

```go
c := List(Id("a"), Id("b")).Op(":=").Id("c").Call()
fmt.Printf("%#v", c)
// Output: a, b := c()
```

### Values
Values renders a comma seperated list enclosed by curly braces. Use for slice 
literals:

```go
c := Index().String().Values(Lit("a"), Lit("b"))
fmt.Printf("%#v", c)
// Output: []string{"a", "b"}
```

