[![Build Status](https://travis-ci.org/dave/jennifer.svg?branch=master)](https://travis-ci.org/dave/jennifer) [![Go Report Card](https://goreportcard.com/badge/github.com/dave/jennifer)](https://goreportcard.com/report/github.com/dave/jennifer) [![codecov](https://img.shields.io/badge/codecov-100%25-brightgreen.svg)](https://codecov.io/gh/dave/jennifer) ![stability-stable](https://img.shields.io/badge/stability-stable-brightgreen.svg) [![Sourcegraph](https://sourcegraph.com/github.com/dave/jennifer/jen/-/badge.svg)](https://sourcegraph.com/github.com/dave/jennifer?badge) <a href="https://patreon.com/davebrophy" title="Help with my hosting bills using Patreon"><img src="https://img.shields.io/badge/patreon-donate-yellow.svg" style="max-width:100%;"></a>

# Jennifer
Jennifer is a code generator for Go.

```go
package main

import (
    "fmt"

    . "github.com/dave/jennifer/jen"
)

func main() {{ "ExampleNewFile" | code }}
```
Output:
```go
{{ "ExampleNewFile" | output }}
```

### Install
```
go get -u github.com/dave/jennifer/jen
```

### Need help? 
If you get stuck, have a question, would like a code review, or just want a 
chat: I'm happy to help! Feel free to open an issue, email me or mention @dave 
in your PR.

### Examples
Jennifer has a comprehensive suite of examples - see [godoc](https://godoc.org/github.com/dave/jennifer/jen#pkg-examples) for an index. Here's some examples of jennifer being used in the real-world:

* [genjen](genjen/render.go) (which generates much of jennifer, using data in [data.go](genjen/data.go))
* [zerogen](https://github.com/mrsinham/zerogen/blob/master/generator.go)
* [go-contentful-generator](https://github.com/nicolai86/go-contentful-generator)

### Rendering
For testing, a File or Statement can be rendered with the fmt package 
using the %#v verb.

{{ "ExampleCall_fmt" | example }}

This is not recommended for use in production because any error will cause a 
panic. For production use, [File.Render](#render) or [File.Save](#save) are 
preferred.

# Identifiers
**Identifiers** [Keywords](#keywords) [Operators](#operators) [Braces](#braces) [Parentheses](#parentheses) [Control flow](#control-flow) [Collections](#collections) [Literals](#literals) [Comments](#comments) [Helpers](#helpers) [Misc](#misc) [File](#file)

### Id
{{ "Id" | doc }}

{{ "ExampleId" | example }}

### Dot
{{ "Dot" | doc }} 

{{ "ExampleDot" | example }}

### Qual
{{ "Qual[0]" | doc }}

{{ "ExampleQual" | example }}

{{ "Qual[1:4]" | doc }}

{{ "ExampleQual_file" | example }}

{{ "Qual[4:]" | doc }}

### List
{{ "List" | doc }}

{{ "ExampleList" | example }}

# Keywords
[Identifiers](#identifiers) **Keywords** [Operators](#operators) [Braces](#braces) [Parentheses](#parentheses) [Control flow](#control-flow) [Collections](#collections) [Literals](#literals) [Comments](#comments) [Helpers](#helpers) [Misc](#misc) [File](#file)

Simple keywords, predeclared identifiers and built-in functions are self 
explanatory:

| Construct        | Name |
| ---------------- | ---- |
| Keywords         | Break, Chan, Const, Continue, Default, Defer, Else, Fallthrough, Func, Go, Goto, Range, Select, Type, Var |
| Functions        | Append, Cap, Close, Complex, Copy, Delete, Imag, Len, Make, New, Panic, Print, Println, Real, Recover |
| Types            | Bool, Byte, Complex64, Complex128, Error, Float32, Float64, Int, Int8, Int16, Int32, Int64, Rune, String, Uint, Uint8, Uint16, Uint32, Uint64, Uintptr |
| Constants        | True, False, Iota, Nil |
| Helpers          | Err |

Built-in functions take a list of parameters and render them appropriately:

{{ "ExampleAppend_more" | example }}

Special cases for [If, For](#if-for), [Interface, Struct](#interface-struct), [Switch, Case](#switch-select), [Return](#return) and [Map](#map) are explained below.

# Operators
[Identifiers](#identifiers) [Keywords](#keywords) **Operators** [Braces](#braces) [Parentheses](#parentheses) [Control flow](#control-flow) [Collections](#collections) [Literals](#literals) [Comments](#comments) [Helpers](#helpers) [Misc](#misc) [File](#file)

{{ "Op" | doc }}

{{ "ExampleOp" | example }}

{{ "ExampleOp_star" | example }}

{{ "ExampleOp_variadic" | example }}

{{ "ExampleOp_complex_conditions" | example }}

# Braces
[Identifiers](#identifiers) [Keywords](#keywords) [Operators](#operators) **Braces** [Parentheses](#parentheses) [Control flow](#control-flow) [Collections](#collections) [Literals](#literals) [Comments](#comments) [Helpers](#helpers) [Misc](#misc) [File](#file)

Several methods render curly braces, summarized below: 

| Name                           | Prefix       | Separator | Example                              |
| ------------------------------ | ------------ | --------- | -------------------------------------|
| [Block](#block)                |              | `\n`      | `func a() { ... }` or `if a { ... }` |
| [Interface](#interface-struct) | `interface`  | `\n`      | `interface { ... }`                  |
| [Struct](#interface-struct)    | `struct`     | `\n`      | `struct { ... }`                     |
| [Values](#values)              |              | `,`       | `[]int{1, 2}` or `A{B: "c"}`         |

### Block
{{ "Block[:2]" | doc }}

{{ "ExampleBlock" | example }}

{{ "ExampleBlock_if" | example }}

{{ "Block[2:]" | doc }} [See example](#switch-select).

### Interface, Struct
Interface and Struct render the keyword followed by a statement list enclosed 
by curly braces.

{{ "ExampleInterface_empty" | example }}

{{ "ExampleInterface" | example }}

{{ "ExampleStruct_empty" | example }}

{{ "ExampleStruct" | example }}

# Parentheses
[Identifiers](#identifiers) [Keywords](#keywords) [Operators](#operators) [Braces](#braces) **Parentheses** [Control flow](#control-flow) [Collections](#collections) [Literals](#literals) [Comments](#comments) [Helpers](#helpers) [Misc](#misc) [File](#file)

Several methods output parenthesis, summarized below:

| Name              | Prefix | Separator | Example                           |
| ----------------- | ------ | --------- | --------------------------------- |
| [Call](#call)     |        | `,`       | `fmt.Println(b, c)`               |
| [Params](#params) |        | `,`       | `func (a *A) Foo(i int) { ... }`  |
| [Defs](#defs)     |        | `\n`      | `const ( ... )`                   |
| [Parens](#parens) |        |           | `[]byte(s)` or `a / (b + c)`      |
| [Assert](#assert) | `.`    |           | `s, ok := i.(string)`             |

### Call
{{ "Call" | doc }}

{{ "ExampleCall" | example }}

### Params
{{ "Params" | doc }}

{{ "ExampleParams" | example }}

### Defs
{{ "Defs" | doc }}

{{ "ExampleDefs" | example }}

### Parens
{{ "Parens" | doc }}

{{ "ExampleParens" | example }}

{{ "ExampleParens_order" | example }}

### Assert
{{ "Assert" | doc }}

{{ "ExampleAssert" | example }}

# Control flow
[Identifiers](#identifiers) [Keywords](#keywords) [Operators](#operators) [Braces](#braces) [Parentheses](#parentheses) **Control flow** [Collections](#collections) [Literals](#literals) [Comments](#comments) [Helpers](#helpers) [Misc](#misc) [File](#file)

### If, For
If and For render the keyword followed by a semicolon separated list.

{{ "ExampleIf" | example }}

{{ "ExampleFor" | example }}

### Switch, Select
Switch, Select, Case and Block are used to build switch or select statements:

{{ "ExampleSwitch" | example }}

### Return
{{ "Return" | doc }}

{{ "ExampleReturn" | example }}

# Collections
[Identifiers](#identifiers) [Keywords](#keywords) [Operators](#operators) [Braces](#braces) [Parentheses](#parentheses) [Control flow](#control-flow) **Collections** [Literals](#literals) [Comments](#comments) [Helpers](#helpers) [Misc](#misc) [File](#file)

### Map
{{ "Map" | doc }}

{{ "ExampleMap" | example }}

### Index
{{ "Index" | doc }}

{{ "ExampleIndex" | example }}

{{ "ExampleIndex_index" | example }}

{{ "ExampleIndex_empty" | example }}

### Values
{{ "Values" | doc }}

{{ "ExampleValues" | example }}

{{ "Dict" | doc }}

{{ "ExampleValues_dict_multiple" | example }}

{{ "ExampleValues_dict_composite" | example }}

{{ "DictFunc[0]" | doc }}

{{ "ExampleDictFunc" | example }}

Note: the items are ordered by key when rendered to ensure repeatable code.

# Literals
[Identifiers](#identifiers) [Keywords](#keywords) [Operators](#operators) [Braces](#braces) [Parentheses](#parentheses) [Control flow](#control-flow) [Collections](#collections) **Literals** [Comments](#comments) [Helpers](#helpers) [Misc](#misc) [File](#file)

### Lit
{{ "Lit" | doc }}

{{ "ExampleLit" | example }}

{{ "ExampleLit_float" | example }}

{{ "LitFunc[1:2]" | doc }}

{{ "ExampleLitFunc" | example }}

For the default constant types (bool, int, float64, string, complex128), Lit 
will render the untyped constant.

| Code          | Output     |
| ------------- | ---------- |
| `Lit(true)`   | `true`     |
| `Lit(1)`      | `1`        |
| `Lit(1.0)`    | `1.0`      |
| `Lit("foo")`  | `"foo"`    |
| `Lit(0 + 1i)` | `(0 + 1i)` |

For all other built-in types (float32, int8, int16, int32, int64, uint, uint8, 
uint16, uint32, uint64, uintptr, complex64), Lit will also render the type.

| Code                     | Output              |
| ------------------------ | ------------------- |
| `Lit(float32(1))`        | `float32(1)`        |
| `Lit(int16(1))`          | `int16(1)`          |
| `Lit(uint8(0x1))`        | `uint8(0x1)`        |
| `Lit(complex64(0 + 1i))` | `complex64(0 + 1i)` |

The built-in alias types byte and rune need a special case. LitRune and LitByte 
render rune and byte literals.

| Code                     | Output      |
| ------------------------ | ----------- |
| `LitRune('x')`           | `'x'`       |
| `LitByte(byte(0x1))`     | `byte(0x1)` |

# Comments
[Identifiers](#identifiers) [Keywords](#keywords) [Operators](#operators) [Braces](#braces) [Parentheses](#parentheses) [Control flow](#control-flow) [Collections](#collections) [Literals](#literals) **Comments** [Helpers](#helpers) [Misc](#misc) [File](#file)

### Comment
{{ "Comment[:2]" | doc }}

{{ "ExampleComment" | example }}

{{ "ExampleComment_multiline" | example }}

{{ "Comment[2:]" | doc }}

{{ "ExampleComment_formatting_disabled" | example }}

### Commentf
{{ "Commentf[0]" | doc }}

{{ "ExampleCommentf" | example }}

# Helpers
[Identifiers](#identifiers) [Keywords](#keywords) [Operators](#operators) [Braces](#braces) [Parentheses](#parentheses) [Control flow](#control-flow) [Collections](#collections) [Literals](#literals) [Comments](#comments) **Helpers** [Misc](#misc) [File](#file)

### Func methods
All constructs that accept a variadic list of items are paired with GroupFunc 
functions that accept a func(*Group). Use for embedding logic.

{{ "ExampleValuesFunc" | example }}

{{ "ExampleBlockFunc" | example }}

### Add
{{ "Add" | doc }}

{{ "ExampleAdd" | example }}

{{ "ExampleAdd_var" | example }}

### Do
{{ "Do" | doc }}

{{ "ExampleDo" | example }}

# Misc
[Identifiers](#identifiers) [Keywords](#keywords) [Operators](#operators) [Braces](#braces) [Parentheses](#parentheses) [Control flow](#control-flow) [Collections](#collections) [Literals](#literals) [Comments](#comments) [Helpers](#helpers) **Misc** [File](#file)

### Tag
{{ "Tag" | doc }}

{{ "ExampleTag" | example }}

Note: the items are ordered by key when rendered to ensure repeatable code.

### Null
{{ "Null" | doc }}

In lists, nil will produce the same effect.

{{ "ExampleNull_and_nil" | example }}

### Empty
{{ "Empty" | doc }}

{{ "ExampleEmpty" | example }}

### Line
{{ "Line" | doc }}

### Clone
Be careful when passing *Statement. Consider the following... 

{{ "ExampleStatement_Clone_broken" | example }}

Id("a") returns a *Statement, which the Call() method appends to twice. To 
avoid this, use Clone. {{ "Statement.Clone" | doc }}  

{{ "ExampleStatement_Clone_fixed" | example }}

### Cgo
The cgo "C" pseudo-package is a special case, and always renders without a package alias. The 
import can be added with `Qual`, `Anon` or by supplying a preamble. The preamble is added with 
`File.CgoPreamble` which has the same semantics as [Comment](#comments). If a preamble is provided, 
the import is separated, and preceded by the preamble. 

{{ "ExampleFile_CgoPreamble" | example }}  

# File
[Identifiers](#identifiers) [Keywords](#keywords) [Operators](#operators) [Braces](#braces) [Parentheses](#parentheses) [Control flow](#control-flow) [Collections](#collections) [Literals](#literals) [Comments](#comments) [Helpers](#helpers) [Misc](#misc) **File**

{{ "File" | doc }}

### NewFile
{{ "NewFile" | doc }}

### NewFilePath
{{ "NewFilePath" | doc }}

### NewFilePathName
{{ "NewFilePathName" | doc }}

{{ "ExampleNewFilePathName" | example }}

### Save
{{ "File.Save" | doc }}

### Render
{{ "File.Render" | doc }}

{{ "ExampleFile_Render" | example }}

### Anon
{{ "File.Anon" | doc }}

{{ "ExampleFile_Anon" | example }}

### ImportName
{{ "File.ImportName" | doc }}

{{ "ExampleFile_ImportName" | example }}

### ImportNames
{{ "File.ImportNames" | doc }} 

### ImportAlias
{{ "File.ImportAlias" | doc }}

{{ "ExampleFile_ImportAlias" | example }}

### Comments
{{ "File.PackageComment" | doc }}

{{ "File.HeaderComment" | doc }}

{{ "File.CanonicalPath" | doc }}

{{ "ExampleFile_HeaderAndPackageComments" | example }}

{{ "File.CgoPreamble" | doc }}

### PackagePrefix
{{ "File.PackagePrefix" | doc }}

{{ "ExampleFile_PackagePrefix" | example }}
