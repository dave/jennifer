# jen
--
    import "github.com/davelondon/jennifer/jen"

go:generate genjen2

## Usage

#### func  Context

```go
func Context(ctx context.Context) context.Context
```

#### func  FromContext

```go
func FromContext(ctx context.Context) *global
```

#### func  RenderFile

```go
func RenderFile(ctx context.Context, g *Group, w io.Writer) error
```

#### func  WriteFile

```go
func WriteFile(ctx context.Context, g *Group, filename string) error
```

#### type Code

```go
type Code interface {
	// contains filtered or unexported methods
}
```


#### type Group

```go
type Group struct {
}
```


#### func  Append

```go
func Append(c ...Code) *Group
```
Append inserts the append built in function

#### func  Block

```go
func Block(c ...Code) *Group
```
Block inserts curly braces containing a statements list

#### func  Bool

```go
func Bool() *Group
```
Bool inserts the bool identifier

#### func  Break

```go
func Break() *Group
```
Break inserts the break keyword

#### func  Byte

```go
func Byte() *Group
```
Byte inserts the byte identifier

#### func  Call

```go
func Call(c ...Code) *Group
```
Call inserts parenthesis containing a comma separated list

#### func  Cap

```go
func Cap(c ...Code) *Group
```
Cap inserts the cap built in function

#### func  Case

```go
func Case() *Group
```
Case inserts the case keyword

#### func  Chan

```go
func Chan() *Group
```
Chan inserts the chan keyword

#### func  Clause

```go
func Clause(c ...Code) *Group
```
Clause inserts a semicolon separated list

#### func  Close

```go
func Close(c ...Code) *Group
```
Close inserts the close built in function

#### func  Comment

```go
func Comment(comments ...string) *Group
```

#### func  Commentf

```go
func Commentf(format string, a ...interface{}) *Group
```

#### func  Complex

```go
func Complex(c ...Code) *Group
```
Complex inserts the complex built in function

#### func  Complex128

```go
func Complex128() *Group
```
Complex128 inserts the complex128 identifier

#### func  Complex64

```go
func Complex64() *Group
```
Complex64 inserts the complex64 identifier

#### func  Const

```go
func Const() *Group
```
Const inserts the const keyword

#### func  Continue

```go
func Continue() *Group
```
Continue inserts the continue keyword

#### func  Copy

```go
func Copy(c ...Code) *Group
```
Copy inserts the copy built in function

#### func  Decls

```go
func Decls(c ...Code) *Group
```
Decls inserts parenthesis containing a statement list

#### func  Default

```go
func Default() *Group
```
Default inserts the default keyword

#### func  Defer

```go
func Defer() *Group
```
Defer inserts the defer keyword

#### func  Delete

```go
func Delete(c ...Code) *Group
```
Delete inserts the delete built in function

#### func  Do

```go
func Do(f func(*Group)) *Group
```
Do creates a new statement and calls the provided function with it as a
parameter

#### func  Else

```go
func Else() *Group
```
Else inserts the else keyword

#### func  Empty

```go
func Empty() *Group
```
Empty token produces no output but is followed by a separator in a list.

#### func  Error

```go
func Error() *Group
```
Error inserts the error identifier

#### func  Fallthrough

```go
func Fallthrough() *Group
```
Fallthrough inserts the fallthrough keyword

#### func  False

```go
func False() *Group
```
False inserts the false identifier

#### func  Float32

```go
func Float32() *Group
```
Float32 inserts the float32 identifier

#### func  Float64

```go
func Float64() *Group
```
Float64 inserts the float64 identifier

#### func  For

```go
func For() *Group
```
For inserts the for keyword

#### func  Func

```go
func Func() *Group
```
Func inserts the func keyword

#### func  Go

```go
func Go() *Group
```
Go inserts the go keyword

#### func  Goto

```go
func Goto() *Group
```
Goto inserts the goto keyword

#### func  Id

```go
func Id(names ...string) *Group
```

#### func  If

```go
func If() *Group
```
If inserts the if keyword

#### func  Imag

```go
func Imag(c ...Code) *Group
```
Imag inserts the imag built in function

#### func  Import

```go
func Import() *Group
```
Import inserts the import keyword

#### func  Index

```go
func Index(c ...Code) *Group
```
Index inserts square brackets containing a colon separated list

#### func  Int

```go
func Int() *Group
```
Int inserts the int identifier

#### func  Int16

```go
func Int16() *Group
```
Int16 inserts the int16 identifier

#### func  Int32

```go
func Int32() *Group
```
Int32 inserts the int32 identifier

#### func  Int64

```go
func Int64() *Group
```
Int64 inserts the int64 identifier

#### func  Int8

```go
func Int8() *Group
```
Int8 inserts the int8 identifier

#### func  Interface

```go
func Interface() *Group
```
Interface inserts the interface keyword

#### func  Iota

```go
func Iota() *Group
```
Iota inserts the iota identifier

#### func  Len

```go
func Len(c ...Code) *Group
```
Len inserts the len built in function

#### func  List

```go
func List(c ...Code) *Group
```
List inserts a comma separated list

#### func  Lit

```go
func Lit(v interface{}) *Group
```

#### func  Make

```go
func Make(c ...Code) *Group
```
Make inserts the make built in function

#### func  Map

```go
func Map() *Group
```
Map inserts the map keyword

#### func  New

```go
func New(c ...Code) *Group
```
New inserts the new built in function

#### func  NewFile

```go
func NewFile(name string) *Group
```

#### func  NewFilePath

```go
func NewFilePath(name, path string) *Group
```

#### func  Nil

```go
func Nil() *Group
```
Nil inserts the nil identifier

#### func  Null

```go
func Null() *Group
```
Null token produces no output but also no separator in a list.

#### func  Op

```go
func Op(op string) *Group
```

#### func  Package

```go
func Package() *Group
```
Package inserts the package keyword

#### func  Panic

```go
func Panic(c ...Code) *Group
```
Panic inserts the panic built in function

#### func  Params

```go
func Params(c ...Code) *Group
```
Params inserts parenthesis containing a comma separated list

#### func  Parens

```go
func Parens(c ...Code) *Group
```
Parens inserts parenthesis

#### func  Print

```go
func Print(c ...Code) *Group
```
Print inserts the print built in function

#### func  Println

```go
func Println(c ...Code) *Group
```
Println inserts the println built in function

#### func  Range

```go
func Range() *Group
```
Range inserts the range keyword

#### func  Real

```go
func Real(c ...Code) *Group
```
Real inserts the real built in function

#### func  Recover

```go
func Recover(c ...Code) *Group
```
Recover inserts the recover built in function

#### func  Return

```go
func Return(c ...Code) *Group
```
Return inserts the return keyword

#### func  Rune

```go
func Rune() *Group
```
Rune inserts the rune identifier

#### func  Select

```go
func Select() *Group
```
Select inserts the select keyword

#### func  String

```go
func String() *Group
```
String inserts the string identifier

#### func  Struct

```go
func Struct() *Group
```
Struct inserts the struct keyword

#### func  Switch

```go
func Switch() *Group
```
Switch inserts the switch keyword

#### func  True

```go
func True() *Group
```
True inserts the true identifier

#### func  Type

```go
func Type() *Group
```
Type inserts the type keyword

#### func  Uint

```go
func Uint() *Group
```
Uint inserts the uint identifier

#### func  Uint16

```go
func Uint16() *Group
```
Uint16 inserts the uint16 identifier

#### func  Uint32

```go
func Uint32() *Group
```
Uint32 inserts the uint32 identifier

#### func  Uint64

```go
func Uint64() *Group
```
Uint64 inserts the uint64 identifier

#### func  Uint8

```go
func Uint8() *Group
```
Uint8 inserts the uint8 identifier

#### func  Uintptr

```go
func Uintptr() *Group
```
Uintptr inserts the uintptr identifier

#### func  Values

```go
func Values(c ...Code) *Group
```
Values inserts curly braces containing a comma separated list

#### func  Var

```go
func Var() *Group
```
Var inserts the var keyword

#### func (*Group) Add

```go
func (g *Group) Add(code ...Code) *Group
```
Add appends the provided code to the group.

#### func (*Group) Append

```go
func (g *Group) Append(c ...Code) *Group
```
Append inserts the append built in function

#### func (*Group) Block

```go
func (g *Group) Block(c ...Code) *Group
```
Block inserts curly braces containing a statements list

#### func (*Group) Bool

```go
func (g *Group) Bool() *Group
```
Bool inserts the bool identifier

#### func (*Group) Break

```go
func (g *Group) Break() *Group
```
Break inserts the break keyword

#### func (*Group) Byte

```go
func (g *Group) Byte() *Group
```
Byte inserts the byte identifier

#### func (*Group) Call

```go
func (g *Group) Call(c ...Code) *Group
```
Call inserts parenthesis containing a comma separated list

#### func (*Group) Cap

```go
func (g *Group) Cap(c ...Code) *Group
```
Cap inserts the cap built in function

#### func (*Group) Case

```go
func (g *Group) Case() *Group
```
Case inserts the case keyword

#### func (*Group) Chan

```go
func (g *Group) Chan() *Group
```
Chan inserts the chan keyword

#### func (*Group) Clause

```go
func (g *Group) Clause(c ...Code) *Group
```
Clause inserts a semicolon separated list

#### func (*Group) Close

```go
func (g *Group) Close(c ...Code) *Group
```
Close inserts the close built in function

#### func (*Group) Comment

```go
func (g *Group) Comment(comments ...string) *Group
```

#### func (*Group) Commentf

```go
func (g *Group) Commentf(format string, a ...interface{}) *Group
```

#### func (*Group) Complex

```go
func (g *Group) Complex(c ...Code) *Group
```
Complex inserts the complex built in function

#### func (*Group) Complex128

```go
func (g *Group) Complex128() *Group
```
Complex128 inserts the complex128 identifier

#### func (*Group) Complex64

```go
func (g *Group) Complex64() *Group
```
Complex64 inserts the complex64 identifier

#### func (*Group) Const

```go
func (g *Group) Const() *Group
```
Const inserts the const keyword

#### func (*Group) Continue

```go
func (g *Group) Continue() *Group
```
Continue inserts the continue keyword

#### func (*Group) Copy

```go
func (g *Group) Copy(c ...Code) *Group
```
Copy inserts the copy built in function

#### func (*Group) Decls

```go
func (g *Group) Decls(c ...Code) *Group
```
Decls inserts parenthesis containing a statement list

#### func (*Group) Default

```go
func (g *Group) Default() *Group
```
Default inserts the default keyword

#### func (*Group) Defer

```go
func (g *Group) Defer() *Group
```
Defer inserts the defer keyword

#### func (*Group) Delete

```go
func (g *Group) Delete(c ...Code) *Group
```
Delete inserts the delete built in function

#### func (*Group) Do

```go
func (g *Group) Do(f func(*Group)) *Group
```
Do calls the provided function with the group as a parameter

#### func (*Group) Else

```go
func (g *Group) Else() *Group
```
Else inserts the else keyword

#### func (*Group) Empty

```go
func (g *Group) Empty() *Group
```
Empty token produces no output but is followed by a separator in a list.

#### func (*Group) Error

```go
func (g *Group) Error() *Group
```
Error inserts the error identifier

#### func (*Group) Fallthrough

```go
func (g *Group) Fallthrough() *Group
```
Fallthrough inserts the fallthrough keyword

#### func (*Group) False

```go
func (g *Group) False() *Group
```
False inserts the false identifier

#### func (*Group) Float32

```go
func (g *Group) Float32() *Group
```
Float32 inserts the float32 identifier

#### func (*Group) Float64

```go
func (g *Group) Float64() *Group
```
Float64 inserts the float64 identifier

#### func (*Group) For

```go
func (g *Group) For() *Group
```
For inserts the for keyword

#### func (*Group) Func

```go
func (g *Group) Func() *Group
```
Func inserts the func keyword

#### func (*Group) Go

```go
func (g *Group) Go() *Group
```
Go inserts the go keyword

#### func (*Group) GoString

```go
func (g *Group) GoString() string
```

#### func (*Group) Goto

```go
func (g *Group) Goto() *Group
```
Goto inserts the goto keyword

#### func (*Group) Id

```go
func (g *Group) Id(names ...string) *Group
```

#### func (*Group) If

```go
func (g *Group) If() *Group
```
If inserts the if keyword

#### func (*Group) Imag

```go
func (g *Group) Imag(c ...Code) *Group
```
Imag inserts the imag built in function

#### func (*Group) Import

```go
func (g *Group) Import() *Group
```
Import inserts the import keyword

#### func (*Group) Index

```go
func (g *Group) Index(c ...Code) *Group
```
Index inserts square brackets containing a colon separated list

#### func (*Group) Int

```go
func (g *Group) Int() *Group
```
Int inserts the int identifier

#### func (*Group) Int16

```go
func (g *Group) Int16() *Group
```
Int16 inserts the int16 identifier

#### func (*Group) Int32

```go
func (g *Group) Int32() *Group
```
Int32 inserts the int32 identifier

#### func (*Group) Int64

```go
func (g *Group) Int64() *Group
```
Int64 inserts the int64 identifier

#### func (*Group) Int8

```go
func (g *Group) Int8() *Group
```
Int8 inserts the int8 identifier

#### func (*Group) Interface

```go
func (g *Group) Interface() *Group
```
Interface inserts the interface keyword

#### func (*Group) Iota

```go
func (g *Group) Iota() *Group
```
Iota inserts the iota identifier

#### func (*Group) Len

```go
func (g *Group) Len(c ...Code) *Group
```
Len inserts the len built in function

#### func (*Group) List

```go
func (g *Group) List(c ...Code) *Group
```
List inserts a comma separated list

#### func (*Group) Lit

```go
func (g *Group) Lit(v interface{}) *Group
```

#### func (*Group) Make

```go
func (g *Group) Make(c ...Code) *Group
```
Make inserts the make built in function

#### func (*Group) Map

```go
func (g *Group) Map() *Group
```
Map inserts the map keyword

#### func (*Group) New

```go
func (g *Group) New(c ...Code) *Group
```
New inserts the new built in function

#### func (*Group) Nil

```go
func (g *Group) Nil() *Group
```
Nil inserts the nil identifier

#### func (*Group) Null

```go
func (g *Group) Null() *Group
```
Null token produces no output but also no separator in a list.

#### func (*Group) Op

```go
func (g *Group) Op(op string) *Group
```

#### func (*Group) Package

```go
func (g *Group) Package() *Group
```
Package inserts the package keyword

#### func (*Group) Panic

```go
func (g *Group) Panic(c ...Code) *Group
```
Panic inserts the panic built in function

#### func (*Group) Params

```go
func (g *Group) Params(c ...Code) *Group
```
Params inserts parenthesis containing a comma separated list

#### func (*Group) Parens

```go
func (g *Group) Parens(c ...Code) *Group
```
Parens inserts parenthesis

#### func (*Group) Print

```go
func (g *Group) Print(c ...Code) *Group
```
Print inserts the print built in function

#### func (*Group) Println

```go
func (g *Group) Println(c ...Code) *Group
```
Println inserts the println built in function

#### func (*Group) Range

```go
func (g *Group) Range() *Group
```
Range inserts the range keyword

#### func (*Group) Real

```go
func (g *Group) Real(c ...Code) *Group
```
Real inserts the real built in function

#### func (*Group) Recover

```go
func (g *Group) Recover(c ...Code) *Group
```
Recover inserts the recover built in function

#### func (*Group) Return

```go
func (g *Group) Return(c ...Code) *Group
```
Return inserts the return keyword

#### func (*Group) Rune

```go
func (g *Group) Rune() *Group
```
Rune inserts the rune identifier

#### func (*Group) Select

```go
func (g *Group) Select() *Group
```
Select inserts the select keyword

#### func (*Group) String

```go
func (g *Group) String() *Group
```
String inserts the string identifier

#### func (*Group) Struct

```go
func (g *Group) Struct() *Group
```
Struct inserts the struct keyword

#### func (*Group) Switch

```go
func (g *Group) Switch() *Group
```
Switch inserts the switch keyword

#### func (*Group) True

```go
func (g *Group) True() *Group
```
True inserts the true identifier

#### func (*Group) Type

```go
func (g *Group) Type() *Group
```
Type inserts the type keyword

#### func (*Group) Uint

```go
func (g *Group) Uint() *Group
```
Uint inserts the uint identifier

#### func (*Group) Uint16

```go
func (g *Group) Uint16() *Group
```
Uint16 inserts the uint16 identifier

#### func (*Group) Uint32

```go
func (g *Group) Uint32() *Group
```
Uint32 inserts the uint32 identifier

#### func (*Group) Uint64

```go
func (g *Group) Uint64() *Group
```
Uint64 inserts the uint64 identifier

#### func (*Group) Uint8

```go
func (g *Group) Uint8() *Group
```
Uint8 inserts the uint8 identifier

#### func (*Group) Uintptr

```go
func (g *Group) Uintptr() *Group
```
Uintptr inserts the uintptr identifier

#### func (*Group) Values

```go
func (g *Group) Values(c ...Code) *Group
```
Values inserts curly braces containing a comma separated list

#### func (*Group) Var

```go
func (g *Group) Var() *Group
```
Var inserts the var keyword

#### type Token

```go
type Token struct {
	*Group
}
```
