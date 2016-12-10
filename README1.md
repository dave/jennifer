

# jen
`import "github.com/davelondon/jennifer/jen"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
go:generate genjen2




## <a name="pkg-index">Index</a>
* [func Context(ctx context.Context) context.Context](#Context)
* [func FromContext(ctx context.Context) *global](#FromContext)
* [func RenderFile(ctx context.Context, g *Group, w io.Writer) error](#RenderFile)
* [func WriteFile(ctx context.Context, g *Group, filename string) error](#WriteFile)
* [type Code](#Code)
* [type Group](#Group)
  * [func Append(c ...Code) *Group](#Append)
  * [func Block(c ...Code) *Group](#Block)
  * [func Bool() *Group](#Bool)
  * [func Break() *Group](#Break)
  * [func Byte() *Group](#Byte)
  * [func Call(c ...Code) *Group](#Call)
  * [func Cap(c ...Code) *Group](#Cap)
  * [func Case() *Group](#Case)
  * [func Chan() *Group](#Chan)
  * [func Clause(c ...Code) *Group](#Clause)
  * [func Close(c ...Code) *Group](#Close)
  * [func Comment(comments ...string) *Group](#Comment)
  * [func Commentf(format string, a ...interface{}) *Group](#Commentf)
  * [func Complex(c ...Code) *Group](#Complex)
  * [func Complex128() *Group](#Complex128)
  * [func Complex64() *Group](#Complex64)
  * [func Const() *Group](#Const)
  * [func Continue() *Group](#Continue)
  * [func Copy(c ...Code) *Group](#Copy)
  * [func Decls(c ...Code) *Group](#Decls)
  * [func Default() *Group](#Default)
  * [func Defer() *Group](#Defer)
  * [func Delete(c ...Code) *Group](#Delete)
  * [func Do(f func(*Group)) *Group](#Do)
  * [func Else() *Group](#Else)
  * [func Empty() *Group](#Empty)
  * [func Error() *Group](#Error)
  * [func Fallthrough() *Group](#Fallthrough)
  * [func False() *Group](#False)
  * [func Float32() *Group](#Float32)
  * [func Float64() *Group](#Float64)
  * [func For() *Group](#For)
  * [func Func() *Group](#Func)
  * [func Go() *Group](#Go)
  * [func Goto() *Group](#Goto)
  * [func Id(names ...string) *Group](#Id)
  * [func If() *Group](#If)
  * [func Imag(c ...Code) *Group](#Imag)
  * [func Import() *Group](#Import)
  * [func Index(c ...Code) *Group](#Index)
  * [func Int() *Group](#Int)
  * [func Int16() *Group](#Int16)
  * [func Int32() *Group](#Int32)
  * [func Int64() *Group](#Int64)
  * [func Int8() *Group](#Int8)
  * [func Interface() *Group](#Interface)
  * [func Iota() *Group](#Iota)
  * [func Len(c ...Code) *Group](#Len)
  * [func List(c ...Code) *Group](#List)
  * [func Lit(v interface{}) *Group](#Lit)
  * [func Make(c ...Code) *Group](#Make)
  * [func Map() *Group](#Map)
  * [func New(c ...Code) *Group](#New)
  * [func NewFile(name string) *Group](#NewFile)
  * [func NewFilePath(name, path string) *Group](#NewFilePath)
  * [func Nil() *Group](#Nil)
  * [func Null() *Group](#Null)
  * [func Op(op string) *Group](#Op)
  * [func Package() *Group](#Package)
  * [func Panic(c ...Code) *Group](#Panic)
  * [func Params(c ...Code) *Group](#Params)
  * [func Parens(c ...Code) *Group](#Parens)
  * [func Print(c ...Code) *Group](#Print)
  * [func Println(c ...Code) *Group](#Println)
  * [func Range() *Group](#Range)
  * [func Real(c ...Code) *Group](#Real)
  * [func Recover(c ...Code) *Group](#Recover)
  * [func Return(c ...Code) *Group](#Return)
  * [func Rune() *Group](#Rune)
  * [func Select() *Group](#Select)
  * [func String() *Group](#String)
  * [func Struct() *Group](#Struct)
  * [func Switch() *Group](#Switch)
  * [func True() *Group](#True)
  * [func Type() *Group](#Type)
  * [func Uint() *Group](#Uint)
  * [func Uint16() *Group](#Uint16)
  * [func Uint32() *Group](#Uint32)
  * [func Uint64() *Group](#Uint64)
  * [func Uint8() *Group](#Uint8)
  * [func Uintptr() *Group](#Uintptr)
  * [func Values(c ...Code) *Group](#Values)
  * [func Var() *Group](#Var)
  * [func (g *Group) Add(code ...Code) *Group](#Group.Add)
  * [func (g *Group) Append(c ...Code) *Group](#Group.Append)
  * [func (g *Group) Block(c ...Code) *Group](#Group.Block)
  * [func (g *Group) Bool() *Group](#Group.Bool)
  * [func (g *Group) Break() *Group](#Group.Break)
  * [func (g *Group) Byte() *Group](#Group.Byte)
  * [func (g *Group) Call(c ...Code) *Group](#Group.Call)
  * [func (g *Group) Cap(c ...Code) *Group](#Group.Cap)
  * [func (g *Group) Case() *Group](#Group.Case)
  * [func (g *Group) Chan() *Group](#Group.Chan)
  * [func (g *Group) Clause(c ...Code) *Group](#Group.Clause)
  * [func (g *Group) Close(c ...Code) *Group](#Group.Close)
  * [func (g *Group) Comment(comments ...string) *Group](#Group.Comment)
  * [func (g *Group) Commentf(format string, a ...interface{}) *Group](#Group.Commentf)
  * [func (g *Group) Complex(c ...Code) *Group](#Group.Complex)
  * [func (g *Group) Complex128() *Group](#Group.Complex128)
  * [func (g *Group) Complex64() *Group](#Group.Complex64)
  * [func (g *Group) Const() *Group](#Group.Const)
  * [func (g *Group) Continue() *Group](#Group.Continue)
  * [func (g *Group) Copy(c ...Code) *Group](#Group.Copy)
  * [func (g *Group) Decls(c ...Code) *Group](#Group.Decls)
  * [func (g *Group) Default() *Group](#Group.Default)
  * [func (g *Group) Defer() *Group](#Group.Defer)
  * [func (g *Group) Delete(c ...Code) *Group](#Group.Delete)
  * [func (g *Group) Do(f func(*Group)) *Group](#Group.Do)
  * [func (g *Group) Else() *Group](#Group.Else)
  * [func (g *Group) Empty() *Group](#Group.Empty)
  * [func (g *Group) Error() *Group](#Group.Error)
  * [func (g *Group) Fallthrough() *Group](#Group.Fallthrough)
  * [func (g *Group) False() *Group](#Group.False)
  * [func (g *Group) Float32() *Group](#Group.Float32)
  * [func (g *Group) Float64() *Group](#Group.Float64)
  * [func (g *Group) For() *Group](#Group.For)
  * [func (g *Group) Func() *Group](#Group.Func)
  * [func (g *Group) Go() *Group](#Group.Go)
  * [func (g *Group) GoString() string](#Group.GoString)
  * [func (g *Group) Goto() *Group](#Group.Goto)
  * [func (g *Group) Id(names ...string) *Group](#Group.Id)
  * [func (g *Group) If() *Group](#Group.If)
  * [func (g *Group) Imag(c ...Code) *Group](#Group.Imag)
  * [func (g *Group) Import() *Group](#Group.Import)
  * [func (g *Group) Index(c ...Code) *Group](#Group.Index)
  * [func (g *Group) Int() *Group](#Group.Int)
  * [func (g *Group) Int16() *Group](#Group.Int16)
  * [func (g *Group) Int32() *Group](#Group.Int32)
  * [func (g *Group) Int64() *Group](#Group.Int64)
  * [func (g *Group) Int8() *Group](#Group.Int8)
  * [func (g *Group) Interface() *Group](#Group.Interface)
  * [func (g *Group) Iota() *Group](#Group.Iota)
  * [func (g *Group) Len(c ...Code) *Group](#Group.Len)
  * [func (g *Group) List(c ...Code) *Group](#Group.List)
  * [func (g *Group) Lit(v interface{}) *Group](#Group.Lit)
  * [func (g *Group) Make(c ...Code) *Group](#Group.Make)
  * [func (g *Group) Map() *Group](#Group.Map)
  * [func (g *Group) New(c ...Code) *Group](#Group.New)
  * [func (g *Group) Nil() *Group](#Group.Nil)
  * [func (g *Group) Null() *Group](#Group.Null)
  * [func (g *Group) Op(op string) *Group](#Group.Op)
  * [func (g *Group) Package() *Group](#Group.Package)
  * [func (g *Group) Panic(c ...Code) *Group](#Group.Panic)
  * [func (g *Group) Params(c ...Code) *Group](#Group.Params)
  * [func (g *Group) Parens(c ...Code) *Group](#Group.Parens)
  * [func (g *Group) Print(c ...Code) *Group](#Group.Print)
  * [func (g *Group) Println(c ...Code) *Group](#Group.Println)
  * [func (g *Group) Range() *Group](#Group.Range)
  * [func (g *Group) Real(c ...Code) *Group](#Group.Real)
  * [func (g *Group) Recover(c ...Code) *Group](#Group.Recover)
  * [func (g *Group) Return(c ...Code) *Group](#Group.Return)
  * [func (g *Group) Rune() *Group](#Group.Rune)
  * [func (g *Group) Select() *Group](#Group.Select)
  * [func (g *Group) String() *Group](#Group.String)
  * [func (g *Group) Struct() *Group](#Group.Struct)
  * [func (g *Group) Switch() *Group](#Group.Switch)
  * [func (g *Group) True() *Group](#Group.True)
  * [func (g *Group) Type() *Group](#Group.Type)
  * [func (g *Group) Uint() *Group](#Group.Uint)
  * [func (g *Group) Uint16() *Group](#Group.Uint16)
  * [func (g *Group) Uint32() *Group](#Group.Uint32)
  * [func (g *Group) Uint64() *Group](#Group.Uint64)
  * [func (g *Group) Uint8() *Group](#Group.Uint8)
  * [func (g *Group) Uintptr() *Group](#Group.Uintptr)
  * [func (g *Group) Values(c ...Code) *Group](#Group.Values)
  * [func (g *Group) Var() *Group](#Group.Var)
* [type Token](#Token)

#### <a name="pkg-examples">Examples</a>
* [Append](#example_Append)
* [Block](#example_Block)
* [Bool](#example_Bool)
* [Break](#example_Break)
* [Byte](#example_Byte)
* [Func](#example_Func)
* [Group.Append](#example_Group_Append)
* [Group.Block](#example_Group_Block)
* [Group.Bool](#example_Group_Bool)
* [Group.Byte](#example_Group_Byte)
* [Group.Func](#example_Group_Func)
* [Id](#example_Id)
* [NewFile](#example_NewFile)
* [NewFilePath](#example_NewFilePath)

#### <a name="pkg-files">Package files</a>
[comments.go](/src/github.com/davelondon/jennifer/jen/comments.go) [context.go](/src/github.com/davelondon/jennifer/jen/context.go) [generated.go](/src/github.com/davelondon/jennifer/jen/generated.go) [group.go](/src/github.com/davelondon/jennifer/jen/group.go) [jen.go](/src/github.com/davelondon/jennifer/jen/jen.go) [lit.go](/src/github.com/davelondon/jennifer/jen/lit.go) [syntax.go](/src/github.com/davelondon/jennifer/jen/syntax.go) [tokens.go](/src/github.com/davelondon/jennifer/jen/tokens.go) 





## <a name="Context">func</a> [Context](/src/target/context.go?s=104:153#L3)
``` go
func Context(ctx context.Context) context.Context
```


## <a name="FromContext">func</a> [FromContext](/src/target/context.go?s=259:304#L10)
``` go
func FromContext(ctx context.Context) *global
```


## <a name="RenderFile">func</a> [RenderFile](/src/target/jen.go?s=988:1053#L59)
``` go
func RenderFile(ctx context.Context, g *Group, w io.Writer) error
```


## <a name="WriteFile">func</a> [WriteFile](/src/target/jen.go?s=714:782#L47)
``` go
func WriteFile(ctx context.Context, g *Group, filename string) error
```



## <a name="Code">type</a> [Code](/src/target/jen.go?s=128:214#L6)
``` go
type Code interface {
    // contains filtered or unexported methods
}
```









## <a name="Group">type</a> [Group](/src/target/group.go?s=64:115#L1)
``` go
type Group struct {
    // contains filtered or unexported fields
}
```






### <a name="Append">func</a> [Append](/src/target/generated.go?s=22083:22112#L1090)
``` go
func Append(c ...Code) *Group
```
Append inserts the append built in function


### <a name="Block">func</a> [Block](/src/target/generated.go?s=2244:2272#L98)
``` go
func Block(c ...Code) *Group
```
Block inserts curly braces containing a statements list


### <a name="Bool">func</a> [Bool](/src/target/generated.go?s=3997:4015#L178)
``` go
func Bool() *Group
```
Bool inserts the bool identifier


### <a name="Break">func</a> [Break](/src/target/generated.go?s=13164:13183#L634)
``` go
func Break() *Group
```
Break inserts the break keyword


### <a name="Byte">func</a> [Byte](/src/target/generated.go?s=4367:4385#L197)
``` go
func Byte() *Group
```
Byte inserts the byte identifier


### <a name="Call">func</a> [Call](/src/target/generated.go?s=2688:2715#L118)
``` go
func Call(c ...Code) *Group
```
Call inserts parenthesis containing a comma separated list


### <a name="Cap">func</a> [Cap](/src/target/generated.go?s=22420:22446#L1103)
``` go
func Cap(c ...Code) *Group
```
Cap inserts the cap built in function


### <a name="Case">func</a> [Case](/src/target/generated.go?s=15066:15084#L729)
``` go
func Case() *Group
```
Case inserts the case keyword


### <a name="Chan">func</a> [Chan](/src/target/generated.go?s=16871:16889#L824)
``` go
func Chan() *Group
```
Chan inserts the chan keyword


### <a name="Clause">func</a> [Clause](/src/target/generated.go?s=897:926#L38)
``` go
func Clause(c ...Code) *Group
```
Clause inserts a semicolon separated list


### <a name="Close">func</a> [Close](/src/target/generated.go?s=22740:22768#L1116)
``` go
func Close(c ...Code) *Group
```
Close inserts the close built in function


### <a name="Comment">func</a> [Comment](/src/target/comments.go?s=49:88#L1)
``` go
func Comment(comments ...string) *Group
```

### <a name="Commentf">func</a> [Commentf](/src/target/comments.go?s=396:449#L17)
``` go
func Commentf(format string, a ...interface{}) *Group
```

### <a name="Complex">func</a> [Complex](/src/target/generated.go?s=23078:23108#L1129)
``` go
func Complex(c ...Code) *Group
```
Complex inserts the complex built in function


### <a name="Complex128">func</a> [Complex128](/src/target/generated.go?s=5164:5188#L235)
``` go
func Complex128() *Group
```
Complex128 inserts the complex128 identifier


### <a name="Complex64">func</a> [Complex64](/src/target/generated.go?s=4747:4770#L216)
``` go
func Complex64() *Group
```
Complex64 inserts the complex64 identifier


### <a name="Const">func</a> [Const](/src/target/generated.go?s=18723:18742#L919)
``` go
func Const() *Group
```
Const inserts the const keyword


### <a name="Continue">func</a> [Continue](/src/target/generated.go?s=20597:20619#L1014)
``` go
func Continue() *Group
```
Continue inserts the continue keyword


### <a name="Copy">func</a> [Copy](/src/target/generated.go?s=23424:23451#L1142)
``` go
func Copy(c ...Code) *Group
```
Copy inserts the copy built in function


### <a name="Decls">func</a> [Decls](/src/target/generated.go?s=3581:3609#L158)
``` go
func Decls(c ...Code) *Group
```
Decls inserts parenthesis containing a statement list


### <a name="Default">func</a> [Default](/src/target/generated.go?s=13538:13559#L653)
``` go
func Default() *Group
```
Default inserts the default keyword


### <a name="Defer">func</a> [Defer](/src/target/generated.go?s=15429:15448#L748)
``` go
func Defer() *Group
```
Defer inserts the defer keyword


### <a name="Delete">func</a> [Delete](/src/target/generated.go?s=23753:23782#L1155)
``` go
func Delete(c ...Code) *Group
```
Delete inserts the delete built in function


### <a name="Do">func</a> [Do](/src/target/group.go?s=591:621#L23)
``` go
func Do(f func(*Group)) *Group
```
Do creates a new statement and calls the provided function with it as a
parameter


### <a name="Else">func</a> [Else](/src/target/generated.go?s=17232:17250#L843)
``` go
func Else() *Group
```
Else inserts the else keyword


### <a name="Empty">func</a> [Empty](/src/target/tokens.go?s=1812:1831#L81)
``` go
func Empty() *Group
```
Empty token produces no output but is followed by a
separator in a list.


### <a name="Error">func</a> [Error](/src/target/generated.go?s=5578:5597#L254)
``` go
func Error() *Group
```
Error inserts the error identifier


### <a name="Fallthrough">func</a> [Fallthrough](/src/target/generated.go?s=19105:19130#L938)
``` go
func Fallthrough() *Group
```
Fallthrough inserts the fallthrough keyword


### <a name="False">func</a> [False](/src/target/generated.go?s=12057:12076#L577)
``` go
func False() *Group
```
False inserts the false identifier


### <a name="Float32">func</a> [Float32](/src/target/generated.go?s=5961:5982#L273)
``` go
func Float32() *Group
```
Float32 inserts the float32 identifier


### <a name="Float64">func</a> [Float64](/src/target/generated.go?s=6358:6379#L292)
``` go
func Float64() *Group
```
Float64 inserts the float64 identifier


### <a name="For">func</a> [For](/src/target/generated.go?s=20984:21001#L1033)
``` go
func For() *Group
```
For inserts the for keyword


### <a name="Func">func</a> [Func](/src/target/generated.go?s=13920:13938#L672)
``` go
func Func() *Group
```
Func inserts the func keyword


### <a name="Go">func</a> [Go](/src/target/generated.go?s=15793:15809#L767)
``` go
func Go() *Group
```
Go inserts the go keyword


### <a name="Goto">func</a> [Goto](/src/target/generated.go?s=17593:17611#L862)
``` go
func Goto() *Group
```
Goto inserts the goto keyword


### <a name="Id">func</a> [Id](/src/target/tokens.go?s=2495:2526#L121)
``` go
func Id(names ...string) *Group
```

### <a name="If">func</a> [If](/src/target/generated.go?s=19511:19527#L957)
``` go
func If() *Group
```
If inserts the if keyword


### <a name="Imag">func</a> [Imag](/src/target/generated.go?s=24092:24119#L1168)
``` go
func Imag(c ...Code) *Group
```
Imag inserts the imag built in function


### <a name="Import">func</a> [Import](/src/target/generated.go?s=21342:21362#L1052)
``` go
func Import() *Group
```
Import inserts the import keyword


### <a name="Index">func</a> [Index](/src/target/generated.go?s=1795:1823#L78)
``` go
func Index(c ...Code) *Group
```
Index inserts square brackets containing a colon separated list


### <a name="Int">func</a> [Int](/src/target/generated.go?s=6747:6764#L311)
``` go
func Int() *Group
```
Int inserts the int identifier


### <a name="Int16">func</a> [Int16](/src/target/generated.go?s=7482:7501#L349)
``` go
func Int16() *Group
```
Int16 inserts the int16 identifier


### <a name="Int32">func</a> [Int32](/src/target/generated.go?s=7861:7880#L368)
``` go
func Int32() *Group
```
Int32 inserts the int32 identifier


### <a name="Int64">func</a> [Int64](/src/target/generated.go?s=8240:8259#L387)
``` go
func Int64() *Group
```
Int64 inserts the int64 identifier


### <a name="Int8">func</a> [Int8](/src/target/generated.go?s=7110:7128#L330)
``` go
func Int8() *Group
```
Int8 inserts the int8 identifier


### <a name="Interface">func</a> [Interface](/src/target/generated.go?s=14291:14314#L691)
``` go
func Interface() *Group
```
Interface inserts the interface keyword


### <a name="Iota">func</a> [Iota](/src/target/generated.go?s=12434:12452#L596)
``` go
func Iota() *Group
```
Iota inserts the iota identifier


### <a name="Len">func</a> [Len](/src/target/generated.go?s=24415:24441#L1181)
``` go
func Len(c ...Code) *Group
```
Len inserts the len built in function


### <a name="List">func</a> [List](/src/target/generated.go?s=495:522#L18)
``` go
func List(c ...Code) *Group
```
List inserts a comma separated list


### <a name="Lit">func</a> [Lit](/src/target/lit.go?s=42:72#L1)
``` go
func Lit(v interface{}) *Group
```

### <a name="Make">func</a> [Make](/src/target/generated.go?s=24733:24760#L1194)
``` go
func Make(c ...Code) *Group
```
Make inserts the make built in function


### <a name="Map">func</a> [Map](/src/target/generated.go?s=16138:16155#L786)
``` go
func Map() *Group
```
Map inserts the map keyword


### <a name="New">func</a> [New](/src/target/generated.go?s=25056:25082#L1207)
``` go
func New(c ...Code) *Group
```
New inserts the new built in function


### <a name="NewFile">func</a> [NewFile](/src/target/jen.go?s=216:248#L11)
``` go
func NewFile(name string) *Group
```

### <a name="NewFilePath">func</a> [NewFilePath](/src/target/jen.go?s=332:374#L20)
``` go
func NewFilePath(name, path string) *Group
```

### <a name="Nil">func</a> [Nil](/src/target/generated.go?s=12802:12819#L615)
``` go
func Nil() *Group
```
Nil inserts the nil identifier


### <a name="Null">func</a> [Null](/src/target/tokens.go?s=1394:1412#L59)
``` go
func Null() *Group
```
Null token produces no output but also no separator
in a list.


### <a name="Op">func</a> [Op](/src/target/tokens.go?s=2188:2213#L102)
``` go
func Op(op string) *Group
```

### <a name="Package">func</a> [Package](/src/target/generated.go?s=17960:17981#L881)
``` go
func Package() *Group
```
Package inserts the package keyword


### <a name="Panic">func</a> [Panic](/src/target/generated.go?s=25376:25404#L1220)
``` go
func Panic(c ...Code) *Group
```
Panic inserts the panic built in function


### <a name="Params">func</a> [Params](/src/target/generated.go?s=3132:3161#L138)
``` go
func Params(c ...Code) *Group
```
Params inserts parenthesis containing a comma separated list


### <a name="Parens">func</a> [Parens](/src/target/generated.go?s=98:127#L1)
``` go
func Parens(c ...Code) *Group
```
Parens inserts parenthesis


### <a name="Print">func</a> [Print](/src/target/generated.go?s=25710:25738#L1233)
``` go
func Print(c ...Code) *Group
```
Print inserts the print built in function


### <a name="Println">func</a> [Println](/src/target/generated.go?s=26048:26078#L1246)
``` go
func Println(c ...Code) *Group
```
Println inserts the println built in function


### <a name="Range">func</a> [Range](/src/target/generated.go?s=19860:19879#L976)
``` go
func Range() *Group
```
Range inserts the range keyword


### <a name="Real">func</a> [Real](/src/target/generated.go?s=26394:26421#L1259)
``` go
func Real(c ...Code) *Group
```
Real inserts the real built in function


### <a name="Recover">func</a> [Recover](/src/target/generated.go?s=26725:26755#L1272)
``` go
func Recover(c ...Code) *Group
```
Recover inserts the recover built in function


### <a name="Return">func</a> [Return](/src/target/generated.go?s=27065:27094#L1285)
``` go
func Return(c ...Code) *Group
```
Return inserts the return keyword


### <a name="Rune">func</a> [Rune](/src/target/generated.go?s=8617:8635#L406)
``` go
func Rune() *Group
```
Rune inserts the rune identifier


### <a name="Select">func</a> [Select](/src/target/generated.go?s=14691:14711#L710)
``` go
func Select() *Group
```
Select inserts the select keyword


### <a name="String">func</a> [String](/src/target/generated.go?s=8991:9011#L425)
``` go
func String() *Group
```
String inserts the string identifier


### <a name="Struct">func</a> [Struct](/src/target/generated.go?s=16496:16516#L805)
``` go
func Struct() *Group
```
Struct inserts the struct keyword


### <a name="Switch">func</a> [Switch](/src/target/generated.go?s=18346:18366#L900)
``` go
func Switch() *Group
```
Switch inserts the switch keyword


### <a name="True">func</a> [True](/src/target/generated.go?s=11685:11703#L558)
``` go
func True() *Group
```
True inserts the true identifier


### <a name="Type">func</a> [Type](/src/target/generated.go?s=20228:20246#L995)
``` go
func Type() *Group
```
Type inserts the type keyword


### <a name="Uint">func</a> [Uint](/src/target/generated.go?s=9375:9393#L444)
``` go
func Uint() *Group
```
Uint inserts the uint identifier


### <a name="Uint16">func</a> [Uint16](/src/target/generated.go?s=10128:10148#L482)
``` go
func Uint16() *Group
```
Uint16 inserts the uint16 identifier


### <a name="Uint32">func</a> [Uint32](/src/target/generated.go?s=10516:10536#L501)
``` go
func Uint32() *Group
```
Uint32 inserts the uint32 identifier


### <a name="Uint64">func</a> [Uint64](/src/target/generated.go?s=10904:10924#L520)
``` go
func Uint64() *Group
```
Uint64 inserts the uint64 identifier


### <a name="Uint8">func</a> [Uint8](/src/target/generated.go?s=9747:9766#L463)
``` go
func Uint8() *Group
```
Uint8 inserts the uint8 identifier


### <a name="Uintptr">func</a> [Uintptr](/src/target/generated.go?s=11294:11315#L539)
``` go
func Uintptr() *Group
```
Uintptr inserts the uintptr identifier


### <a name="Values">func</a> [Values](/src/target/generated.go?s=1335:1364#L58)
``` go
func Values(c ...Code) *Group
```
Values inserts curly braces containing a comma separated list


### <a name="Var">func</a> [Var](/src/target/generated.go?s=21715:21732#L1071)
``` go
func Var() *Group
```
Var inserts the var keyword





### <a name="Group.Add">func</a> (\*Group) [Add](/src/target/group.go?s=303:343#L11)
``` go
func (g *Group) Add(code ...Code) *Group
```
Add appends the provided code to the group.




### <a name="Group.Append">func</a> (\*Group) [Append](/src/target/generated.go?s=22200:22240#L1093)
``` go
func (g *Group) Append(c ...Code) *Group
```
Append inserts the append built in function




### <a name="Group.Block">func</a> (\*Group) [Block](/src/target/generated.go?s=2371:2410#L101)
``` go
func (g *Group) Block(c ...Code) *Group
```
Block inserts curly braces containing a statements list




### <a name="Group.Bool">func</a> (\*Group) [Bool](/src/target/generated.go?s=4086:4115#L181)
``` go
func (g *Group) Bool() *Group
```
Bool inserts the bool identifier




### <a name="Group.Break">func</a> (\*Group) [Break](/src/target/generated.go?s=13254:13284#L637)
``` go
func (g *Group) Break() *Group
```
Break inserts the break keyword




### <a name="Group.Byte">func</a> (\*Group) [Byte](/src/target/generated.go?s=4456:4485#L200)
``` go
func (g *Group) Byte() *Group
```
Byte inserts the byte identifier




### <a name="Group.Call">func</a> (\*Group) [Call](/src/target/generated.go?s=2816:2854#L121)
``` go
func (g *Group) Call(c ...Code) *Group
```
Call inserts parenthesis containing a comma separated list




### <a name="Group.Cap">func</a> (\*Group) [Cap](/src/target/generated.go?s=22525:22562#L1106)
``` go
func (g *Group) Cap(c ...Code) *Group
```
Cap inserts the cap built in function




### <a name="Group.Case">func</a> (\*Group) [Case](/src/target/generated.go?s=15152:15181#L732)
``` go
func (g *Group) Case() *Group
```
Case inserts the case keyword




### <a name="Group.Chan">func</a> (\*Group) [Chan](/src/target/generated.go?s=16957:16986#L827)
``` go
func (g *Group) Chan() *Group
```
Chan inserts the chan keyword




### <a name="Group.Clause">func</a> (\*Group) [Clause](/src/target/generated.go?s=1012:1052#L41)
``` go
func (g *Group) Clause(c ...Code) *Group
```
Clause inserts a semicolon separated list




### <a name="Group.Close">func</a> (\*Group) [Close](/src/target/generated.go?s=22853:22892#L1119)
``` go
func (g *Group) Close(c ...Code) *Group
```
Close inserts the close built in function




### <a name="Group.Comment">func</a> (\*Group) [Comment](/src/target/comments.go?s=138:188#L3)
``` go
func (g *Group) Comment(comments ...string) *Group
```



### <a name="Group.Commentf">func</a> (\*Group) [Commentf](/src/target/comments.go?s=501:565#L21)
``` go
func (g *Group) Commentf(format string, a ...interface{}) *Group
```



### <a name="Group.Complex">func</a> (\*Group) [Complex](/src/target/generated.go?s=23199:23240#L1132)
``` go
func (g *Group) Complex(c ...Code) *Group
```
Complex inserts the complex built in function




### <a name="Group.Complex128">func</a> (\*Group) [Complex128](/src/target/generated.go?s=5277:5312#L238)
``` go
func (g *Group) Complex128() *Group
```
Complex128 inserts the complex128 identifier




### <a name="Group.Complex64">func</a> (\*Group) [Complex64](/src/target/generated.go?s=4856:4890#L219)
``` go
func (g *Group) Complex64() *Group
```
Complex64 inserts the complex64 identifier




### <a name="Group.Const">func</a> (\*Group) [Const](/src/target/generated.go?s=18813:18843#L922)
``` go
func (g *Group) Const() *Group
```
Const inserts the const keyword




### <a name="Group.Continue">func</a> (\*Group) [Continue](/src/target/generated.go?s=20699:20732#L1017)
``` go
func (g *Group) Continue() *Group
```
Continue inserts the continue keyword




### <a name="Group.Copy">func</a> (\*Group) [Copy](/src/target/generated.go?s=23533:23571#L1145)
``` go
func (g *Group) Copy(c ...Code) *Group
```
Copy inserts the copy built in function




### <a name="Group.Decls">func</a> (\*Group) [Decls](/src/target/generated.go?s=3706:3745#L161)
``` go
func (g *Group) Decls(c ...Code) *Group
```
Decls inserts parenthesis containing a statement list




### <a name="Group.Default">func</a> (\*Group) [Default](/src/target/generated.go?s=13636:13668#L656)
``` go
func (g *Group) Default() *Group
```
Default inserts the default keyword




### <a name="Group.Defer">func</a> (\*Group) [Defer](/src/target/generated.go?s=15519:15549#L751)
``` go
func (g *Group) Defer() *Group
```
Defer inserts the defer keyword




### <a name="Group.Delete">func</a> (\*Group) [Delete](/src/target/generated.go?s=23870:23910#L1158)
``` go
func (g *Group) Delete(c ...Code) *Group
```
Delete inserts the delete built in function




### <a name="Group.Do">func</a> (\*Group) [Do](/src/target/group.go?s=720:761#L28)
``` go
func (g *Group) Do(f func(*Group)) *Group
```
Do calls the provided function with the group as a parameter




### <a name="Group.Else">func</a> (\*Group) [Else](/src/target/generated.go?s=17318:17347#L846)
``` go
func (g *Group) Else() *Group
```
Else inserts the else keyword




### <a name="Group.Empty">func</a> (\*Group) [Empty](/src/target/tokens.go?s=1947:1977#L87)
``` go
func (g *Group) Empty() *Group
```
Empty token produces no output but is followed by a
separator in a list.




### <a name="Group.Error">func</a> (\*Group) [Error](/src/target/generated.go?s=5671:5701#L257)
``` go
func (g *Group) Error() *Group
```
Error inserts the error identifier




### <a name="Group.Fallthrough">func</a> (\*Group) [Fallthrough](/src/target/generated.go?s=19219:19255#L941)
``` go
func (g *Group) Fallthrough() *Group
```
Fallthrough inserts the fallthrough keyword




### <a name="Group.False">func</a> (\*Group) [False](/src/target/generated.go?s=12150:12180#L580)
``` go
func (g *Group) False() *Group
```
False inserts the false identifier




### <a name="Group.Float32">func</a> (\*Group) [Float32](/src/target/generated.go?s=6062:6094#L276)
``` go
func (g *Group) Float32() *Group
```
Float32 inserts the float32 identifier




### <a name="Group.Float64">func</a> (\*Group) [Float64](/src/target/generated.go?s=6459:6491#L295)
``` go
func (g *Group) Float64() *Group
```
Float64 inserts the float64 identifier




### <a name="Group.For">func</a> (\*Group) [For](/src/target/generated.go?s=21066:21094#L1036)
``` go
func (g *Group) For() *Group
```
For inserts the for keyword




### <a name="Group.Func">func</a> (\*Group) [Func](/src/target/generated.go?s=14006:14035#L675)
``` go
func (g *Group) Func() *Group
```
Func inserts the func keyword




### <a name="Group.Go">func</a> (\*Group) [Go](/src/target/generated.go?s=15871:15898#L770)
``` go
func (g *Group) Go() *Group
```
Go inserts the go keyword




### <a name="Group.GoString">func</a> (\*Group) [GoString](/src/target/group.go?s=2621:2654#L142)
``` go
func (g *Group) GoString() string
```



### <a name="Group.Goto">func</a> (\*Group) [Goto](/src/target/generated.go?s=17679:17708#L865)
``` go
func (g *Group) Goto() *Group
```
Goto inserts the goto keyword




### <a name="Group.Id">func</a> (\*Group) [Id](/src/target/tokens.go?s=2568:2610#L125)
``` go
func (g *Group) Id(names ...string) *Group
```



### <a name="Group.If">func</a> (\*Group) [If](/src/target/generated.go?s=19589:19616#L960)
``` go
func (g *Group) If() *Group
```
If inserts the if keyword




### <a name="Group.Imag">func</a> (\*Group) [Imag](/src/target/generated.go?s=24201:24239#L1171)
``` go
func (g *Group) Imag(c ...Code) *Group
```
Imag inserts the imag built in function




### <a name="Group.Import">func</a> (\*Group) [Import](/src/target/generated.go?s=21436:21467#L1055)
``` go
func (g *Group) Import() *Group
```
Import inserts the import keyword




### <a name="Group.Index">func</a> (\*Group) [Index](/src/target/generated.go?s=1930:1969#L81)
``` go
func (g *Group) Index(c ...Code) *Group
```
Index inserts square brackets containing a colon separated list




### <a name="Group.Int">func</a> (\*Group) [Int](/src/target/generated.go?s=6832:6860#L314)
``` go
func (g *Group) Int() *Group
```
Int inserts the int identifier




### <a name="Group.Int16">func</a> (\*Group) [Int16](/src/target/generated.go?s=7575:7605#L352)
``` go
func (g *Group) Int16() *Group
```
Int16 inserts the int16 identifier




### <a name="Group.Int32">func</a> (\*Group) [Int32](/src/target/generated.go?s=7954:7984#L371)
``` go
func (g *Group) Int32() *Group
```
Int32 inserts the int32 identifier




### <a name="Group.Int64">func</a> (\*Group) [Int64](/src/target/generated.go?s=8333:8363#L390)
``` go
func (g *Group) Int64() *Group
```
Int64 inserts the int64 identifier




### <a name="Group.Int8">func</a> (\*Group) [Int8](/src/target/generated.go?s=7199:7228#L333)
``` go
func (g *Group) Int8() *Group
```
Int8 inserts the int8 identifier




### <a name="Group.Interface">func</a> (\*Group) [Interface](/src/target/generated.go?s=14397:14431#L694)
``` go
func (g *Group) Interface() *Group
```
Interface inserts the interface keyword




### <a name="Group.Iota">func</a> (\*Group) [Iota](/src/target/generated.go?s=12523:12552#L599)
``` go
func (g *Group) Iota() *Group
```
Iota inserts the iota identifier




### <a name="Group.Len">func</a> (\*Group) [Len](/src/target/generated.go?s=24520:24557#L1184)
``` go
func (g *Group) Len(c ...Code) *Group
```
Len inserts the len built in function




### <a name="Group.List">func</a> (\*Group) [List](/src/target/generated.go?s=600:638#L21)
``` go
func (g *Group) List(c ...Code) *Group
```
List inserts a comma separated list




### <a name="Group.Lit">func</a> (\*Group) [Lit](/src/target/lit.go?s=108:149#L2)
``` go
func (g *Group) Lit(v interface{}) *Group
```



### <a name="Group.Make">func</a> (\*Group) [Make](/src/target/generated.go?s=24842:24880#L1197)
``` go
func (g *Group) Make(c ...Code) *Group
```
Make inserts the make built in function




### <a name="Group.Map">func</a> (\*Group) [Map](/src/target/generated.go?s=16220:16248#L789)
``` go
func (g *Group) Map() *Group
```
Map inserts the map keyword




### <a name="Group.New">func</a> (\*Group) [New](/src/target/generated.go?s=25161:25198#L1210)
``` go
func (g *Group) New(c ...Code) *Group
```
New inserts the new built in function




### <a name="Group.Nil">func</a> (\*Group) [Nil](/src/target/generated.go?s=12887:12915#L618)
``` go
func (g *Group) Nil() *Group
```
Nil inserts the nil identifier




### <a name="Group.Null">func</a> (\*Group) [Null](/src/target/tokens.go?s=1517:1546#L65)
``` go
func (g *Group) Null() *Group
```
Null token produces no output but also no separator
in a list.




### <a name="Group.Op">func</a> (\*Group) [Op](/src/target/tokens.go?s=2249:2285#L106)
``` go
func (g *Group) Op(op string) *Group
```



### <a name="Group.Package">func</a> (\*Group) [Package](/src/target/generated.go?s=18058:18090#L884)
``` go
func (g *Group) Package() *Group
```
Package inserts the package keyword




### <a name="Group.Panic">func</a> (\*Group) [Panic](/src/target/generated.go?s=25489:25528#L1223)
``` go
func (g *Group) Panic(c ...Code) *Group
```
Panic inserts the panic built in function




### <a name="Group.Params">func</a> (\*Group) [Params](/src/target/generated.go?s=3266:3306#L141)
``` go
func (g *Group) Params(c ...Code) *Group
```
Params inserts parenthesis containing a comma separated list




### <a name="Group.Parens">func</a> (\*Group) [Parens](/src/target/generated.go?s=198:238#L1)
``` go
func (g *Group) Parens(c ...Code) *Group
```
Parens inserts parenthesis




### <a name="Group.Print">func</a> (\*Group) [Print](/src/target/generated.go?s=25823:25862#L1236)
``` go
func (g *Group) Print(c ...Code) *Group
```
Print inserts the print built in function




### <a name="Group.Println">func</a> (\*Group) [Println](/src/target/generated.go?s=26169:26210#L1249)
``` go
func (g *Group) Println(c ...Code) *Group
```
Println inserts the println built in function




### <a name="Group.Range">func</a> (\*Group) [Range](/src/target/generated.go?s=19950:19980#L979)
``` go
func (g *Group) Range() *Group
```
Range inserts the range keyword




### <a name="Group.Real">func</a> (\*Group) [Real](/src/target/generated.go?s=26503:26541#L1262)
``` go
func (g *Group) Real(c ...Code) *Group
```
Real inserts the real built in function




### <a name="Group.Recover">func</a> (\*Group) [Recover](/src/target/generated.go?s=26846:26887#L1275)
``` go
func (g *Group) Recover(c ...Code) *Group
```
Recover inserts the recover built in function




### <a name="Group.Return">func</a> (\*Group) [Return](/src/target/generated.go?s=27172:27212#L1288)
``` go
func (g *Group) Return(c ...Code) *Group
```
Return inserts the return keyword




### <a name="Group.Rune">func</a> (\*Group) [Rune](/src/target/generated.go?s=8706:8735#L409)
``` go
func (g *Group) Rune() *Group
```
Rune inserts the rune identifier




### <a name="Group.Select">func</a> (\*Group) [Select](/src/target/generated.go?s=14785:14816#L713)
``` go
func (g *Group) Select() *Group
```
Select inserts the select keyword




### <a name="Group.String">func</a> (\*Group) [String](/src/target/generated.go?s=9088:9119#L428)
``` go
func (g *Group) String() *Group
```
String inserts the string identifier




### <a name="Group.Struct">func</a> (\*Group) [Struct](/src/target/generated.go?s=16590:16621#L808)
``` go
func (g *Group) Struct() *Group
```
Struct inserts the struct keyword




### <a name="Group.Switch">func</a> (\*Group) [Switch](/src/target/generated.go?s=18440:18471#L903)
``` go
func (g *Group) Switch() *Group
```
Switch inserts the switch keyword




### <a name="Group.True">func</a> (\*Group) [True](/src/target/generated.go?s=11774:11803#L561)
``` go
func (g *Group) True() *Group
```
True inserts the true identifier




### <a name="Group.Type">func</a> (\*Group) [Type](/src/target/generated.go?s=20314:20343#L998)
``` go
func (g *Group) Type() *Group
```
Type inserts the type keyword




### <a name="Group.Uint">func</a> (\*Group) [Uint](/src/target/generated.go?s=9464:9493#L447)
``` go
func (g *Group) Uint() *Group
```
Uint inserts the uint identifier




### <a name="Group.Uint16">func</a> (\*Group) [Uint16](/src/target/generated.go?s=10225:10256#L485)
``` go
func (g *Group) Uint16() *Group
```
Uint16 inserts the uint16 identifier




### <a name="Group.Uint32">func</a> (\*Group) [Uint32](/src/target/generated.go?s=10613:10644#L504)
``` go
func (g *Group) Uint32() *Group
```
Uint32 inserts the uint32 identifier




### <a name="Group.Uint64">func</a> (\*Group) [Uint64](/src/target/generated.go?s=11001:11032#L523)
``` go
func (g *Group) Uint64() *Group
```
Uint64 inserts the uint64 identifier




### <a name="Group.Uint8">func</a> (\*Group) [Uint8](/src/target/generated.go?s=9840:9870#L466)
``` go
func (g *Group) Uint8() *Group
```
Uint8 inserts the uint8 identifier




### <a name="Group.Uintptr">func</a> (\*Group) [Uintptr](/src/target/generated.go?s=11395:11427#L542)
``` go
func (g *Group) Uintptr() *Group
```
Uintptr inserts the uintptr identifier




### <a name="Group.Values">func</a> (\*Group) [Values](/src/target/generated.go?s=1470:1510#L61)
``` go
func (g *Group) Values(c ...Code) *Group
```
Values inserts curly braces containing a comma separated list




### <a name="Group.Var">func</a> (\*Group) [Var](/src/target/generated.go?s=21797:21825#L1074)
``` go
func (g *Group) Var() *Group
```
Var inserts the var keyword




## <a name="Token">type</a> [Token](/src/target/tokens.go?s=331:400#L11)
``` go
type Token struct {
    *Group
    // contains filtered or unexported fields
}
```













- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
