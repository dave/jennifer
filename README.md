# Jennifer

Jennifer is a code generator for Go:

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

# Examples
The tests are written mostly as examples - [see godoc.org](https://godoc.org/github.com/davelondon/jennifer/jen#pkg-examples) for an index.

Most of the code is generated using jennifer itself, see the [genjen package](https://github.com/davelondon/jennifer/tree/master/genjen) for a real-world example of usage - it generates [generated.go](https://github.com/davelondon/jennifer/blob/master/jen/generated.go).

# Rendering
For testing, a `File` or `Statement` can be rendered with the `fmt` package:

```go
c := Id("a").Call(Lit("b"))
fmt.Printf("%#v", c)
// Output: a("b")
```

This is not recommended for use in production because any error will cause a 
panic. For production use, `File.Render` or `File.Save` are preferred.

# Id
`Id` renders an identifier. For a local identifier, simply use a string:
 
```go
c := Id("a")
fmt.Printf("%#v", c)
// Output: a
```

For a remote identifier, prefix with the full package path:

```go
c := Id("encoding/gob.NewEncoder").Call()
fmt.Printf("%#v", c)
// Output: gob.NewEncoder()
```

The imports are automatically handled when used with a `File`.

To access fields, more items may be added to the `Id` method:

```go
c := Id("a", "b", "c")
fmt.Printf("%#v", c)
// Output: a.b.c
```

This can be combined with the remote syntax:

```go
c := Id("a.b/c.Foo", "Bar", "Baz")
fmt.Printf("%#v", c)
// Output: c.Foo.Bar.Baz
```

More complex chains can be formed by using Code items instead of strings:

```go
c := Id("a.b/c.Foo", Id("Bar").Call(), "Baz")
fmt.Printf("%#v", c)
// Output: c.Foo.Bar().Baz
```

More control over the package import can be gained by using the `Alias` method 
to specify the remote package:
 
```go
c := Id(Alias("a.b/c"), Id("Foo").Call(), "Bar")
fmt.Printf("%#v", c)
// Output: c.Foo().Bar
```

# Op
`Op` renders the provided string. Use for operators and tokens:

```go
c := Id("a").Op(":=").Id("b").Call()
fmt.Printf("%#v", c)
// Output: a := b()
```

```go
c := Op("*").Id("a")
fmt.Printf("%#v", c)
// Output: *a
```

```go
c := Id("a").Call(Id("b").Op("..."))
fmt.Printf("%#v", c)
// Output: a(b...)
```

# Identifiers 

Identifiers are simple methods with no parameters. They render as the 
identifier token:

```go
c := Break()
fmt.Printf("%#v", c)
// Output: break
```

Keywords: `Break`, `Default`, `Func`, `Select`, `Case`, `Defer`, `Go`, `Struct`, `Chan`, `Else`, `Goto`, `Const`, `Fallthrough`, `Range`, `Type`, `Continue`, `Var`

Built-in types: `Bool`, `Byte`, `Complex64`, `Complex128`, `Error`, `Float32`, `Float64`, `Int`, `Int8`, `Int16`, `Int32`, `Int64`, `Rune`, `String`, `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`, `Uintptr`

Constants: `True`, `False`, `Iota`, `Nil`

Also included is `Err` for the commonly used `err` variable.

Note: `Interface`, `Map`, `Return`, `Switch`, `For` and `If` are special cases, 
and treated as blocks - see below.

Note: The `import` and `package` keywords are always rendered automatically, so 
not included.

# Built-in functions

Built in functions take a variadic list of code items, and render the function
name followed by the items as a comma seperated list of parameters in 
parenthesis:

```go
c := Append(Id("a"), Id("b"))
fmt.Printf("%#v", c)
// Output: append(a, b)
```

Functions: `Append`, `Cap`, `Close`, `Complex`, `Copy`, `Delete`, `Imag`, `Len`, `Make`, `New`, `Panic`, `Print`, `Println`, `Real`, `Recover`

# Blocks

Blocks take either a single code item or a varidic list of code items. The 
items are rendered between open and closing tokens. Multiple items are 
seperated by a separator token.

### Blocks accepting a list of items:

| Block     | Opening       | Separator | Closing | Usage                             |
| --------- | ------------- | --------- | ------- | --------------------------------- |
| List      |               | `,`       |         | `a, b := c()`                     |
| Call      | `(`           | `,`       | `)`     | `fmt.Println(b, c)`               |
| Params    | `(`           | `,`       | `)`     | `func (a *A) Foo(i int) { ... }`  |
| Values    | `{`           | `,`       | `}`     | `[]int{1, 2}` or `interface{}`    |
| Index     | `[`           | `:`       | `]`     | `a[1:2]` or `[]int{}`             |
| Block     | `{`           | `\n`      | `}`     | `func a() { ... }`                |
| CaseBlock | `:`           | `\n`      |         | `switch i {case 1: ... }`         |
| Return    | `return`      | `,`       |         | `return a, b`                     |
| If        | `if`          | `;`       |         | `if a, ok := b(); ok { ... }`     |
| For       | `for`         | `;`       |         | `for i := 0; i < 10; i++ { ... }` |
| Switch    | `switch`      | `;`       |         | `switch a { ... }`                |
| Interface | `interface {` | `\n`      | `}`     | `interface { ... }`               |

### Blocks accepting a single item:

| Block  | Opening  | Closing | Usage                        |
| ------ | -------- | ------- | ---------------------------- |
| Parens | `(`      | `)`     | `[]byte(s)` or `a / (b + c)` |
| Assert | `.(`     | `)`     | `s, ok := i.(string)`        |
| Map    | `map[`   | `]`     | `map[int]string`             |

### List
`List` renders a comma seperated list with no open or closing tokens. Use for 
multiple return functions:

```go
c := List(Id("a"), Id("b")).Op(":=").Id("c").Call()
fmt.Printf("%#v", c)
// Output: a, b := c()
```

### Parens
`Parens` renders a single code item in parenthesis. Use for type conversion or 
to specify evaluation order:

```go
c := Id("b").Op(":=").Index().Byte().Parens(Id("s"))
fmt.Printf("%#v", c)
// Output: b := []byte(s)
```

```go
c := Id("a").Op("/").Parens(Id("b").Op("+").Id("c"))
fmt.Printf("%#v", c)
// Output: a / (b + c)
```

### Values
`Values` renders a comma seperated list enclosed by curly braces. Use for slice 
literals:

```go
c := Index().String().Values(Lit("a"), Lit("b"))
fmt.Printf("%#v", c)
// Output: []string{"a", "b"}
```

### Call

### Params

### Index

### Block

### CaseBlock

### Assert

### Map

### Return

### If

### For

### Switch

### Interface
`Interface` renders the interface keyword followed by a statement block:

```go
c := Var().Id("a").Interface()
fmt.Printf("%#v", c)
// Output: var a interface{}
```

```go
c := Type().Id("a").Interface(
    Id("b").Params().String(),
)
fmt.Printf("%#v", c)
// Output: type a interface {
// 	b() string
// }
```

### Alternate FooFunc methods

# Add
`Add` adds the provided Code to the Statement. When the `Add` function ic 
called, a new Statement is created. This is useful for cloning the contents of 
an existing Statement. See "Pointers" below.
 
```go
ptr := Op("*")
c := Id("a").Op("=").Add(ptr).Id("b")
fmt.Printf("%#v", c)
// Output: a = *b
```

```go
a := Id("a")
c := Block(
    Add(a).Call(),
    Add(a).Call(),
)
fmt.Printf("%#v", c)
// Output: {
// 	a()
// 	a()
// }
```

# Do
`Do` takes a `func(*Statement)` and executes it on the current statement. This 
is useful for embedding logic:

```go
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
// Output: a := map[string]string{}
// b := []string{}
```

# Lit, LitFunc
`Lit` renders a literal, using the format provided by `fmt.Sprintf("%#v", ...)`.

TODO: This probably isn't good enough for all cases. 

# Dict, DictFunc

# Tag

# Null, Empty

# Line

# Comment, Commentf

# File

### NewFile
`NewFile` Creates a new file, with the specified package name. 

### NewFilePath
`NewFilePath` creates a new file while specifying 
the package path - the package name is inferred from the path.

### NewFilePathName
`NewFilePathName` 
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

### PackageComment

### Anon

### PackagePrefix

### Save, Render

# Pointers
Be careful when passing `*Statement` around. Consider the following example:

```go
caller := func(s *Statement) *Statement {
    return s.Call()
}
a := Id("a")
c := Block(
    caller(a),
    caller(a),
)
fmt.Printf("%#v", c)
// Output: {
// 	a()()
// 	a()()
// }
```

`Id("a")` returns a `*Statement`, which the `Call()` method appends to twice. To
avoid this, pass `Statement` instead of `*Statement`:

```go
caller := func(s Statement) *Statement {
    return s.Call()
}
a := *Id("a")
c := Block(
    caller(a),
    caller(a),
)
fmt.Printf("%#v", c)
// Output: {
// 	a()
// 	a()
// }
```

Here is another variation, which can't be solved by pointer indirection:

```go
a := Id("a")
c := Block(
    a.Call(),
    a.Call(),
)
fmt.Printf("%#v", c)
// Output: {
// 	a()()
// 	a()()
// }
```

Here we can prevent the double call by using `Add` to create a new `*Statement`:  

```go
a := Id("a")
c := Block(
    Add(a).Call(),
    Add(a).Call(),
)
fmt.Printf("%#v", c)
// Output: {
// 	a()
// 	a()
// }
```