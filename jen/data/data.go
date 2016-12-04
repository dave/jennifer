package data

var Keywords = []string{"break", "default", "func", "interface", "select", "case", "defer", "go", "map", "struct", "chan", "else", "goto", "package", "switch", "const", "fallthrough", "if", "range", "type", "continue", "for", "import", "return", "var"}

var Types = []string{"bool", "byte", "complex64", "complex128", "error", "float32", "float64", "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"}

var Constants = []string{"true", "false", "iota"}

var Zero = []string{"nil"}

var Functions = []string{"append", "cap", "close", "complex", "copy", "delete", "imag", "len", "make", "new", "panic", "print", "println", "real", "recover"}

var Blocks = []struct {
	Name      string
	Desc      string
	Open      string
	Close     string
	Seperator string
}{
	{
		Name:      "List",
		Desc:      "a comma seperated list",
		Seperator: ",",
	},
	{
		Name:  "Parens",
		Desc:  "parenthesis",
		Open:  "(",
		Close: ")",
	},
	{
		Name:  "Braces",
		Desc:  "curly braces",
		Open:  "{",
		Close: "}",
	},
	{
		Name:      "Values",
		Desc:      "curly braces containing a comma separated list",
		Open:      "{",
		Close:     "}",
		Seperator: ",",
	},
	{
		Name:      "Index",
		Desc:      "square brackets containing a colon separated list",
		Open:      "[",
		Close:     "]",
		Seperator: ":",
	},
	{
		Name:      "Block",
		Desc:      "curly braces containing a statements list",
		Open:      "{",
		Close:     "}",
		Seperator: ";",
	},
	{
		Name:      "Call",
		Desc:      "parenthesis containing a comma separated list",
		Open:      "(",
		Close:     ")",
		Seperator: ",",
	},
	{
		Name:      "Params",
		Desc:      "parenthesis containing a comma separated list",
		Open:      "(",
		Close:     ")",
		Seperator: ",",
	},
	{
		Name:      "Decls",
		Desc:      "parenthesis containing a statement list",
		Open:      "(",
		Close:     ")",
		Seperator: ";",
	},
}

/*
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=
*/
var Operators = []struct {
	Name string
	Desc string
	Op   string
}{
	/*
		+    sum
		-    difference
		*    product
		/    quotient
		%    remainder

		&    bitwise AND
		|    bitwise OR
		^    bitwise XOR
		&^   bit clear (AND NOT)

		<<   left shift
		>>   right shift
	*/
	{Name: "Sum", Desc: "sum", Op: "+"},
	{Name: "Diff", Desc: "difference", Op: "-"},
	{Name: "Product", Desc: "product", Op: "*"},
	{Name: "Quotient", Desc: "quotient", Op: "/"},
	{Name: "Remainder", Desc: "remainder", Op: "%"},
	/*
		=    assign
		:=   short assign
		,    list separator?
		;    statement separator?
		:    key separator?
		.    selector
	*/
	{Name: "As", Desc: "assignment", Op: "="},
	{Name: "Sas", Desc: "short assignment", Op: ":="},
	{Name: "Sel", Desc: "selector", Op: "."},
	/*
		&=   bitwise AND assign
		|=   bitwise OR assign
		^=   bitwise XOR assign
		&^=  bit clear (AND NOT) assign
		<<=  left shift assign
		>>=  right shift assign
	*/
	/*
		==    equal
		!=    not equal
		<     less
		<=    less or equal
		>     greater
		>=    greater or equal
	*/
	{Name: "Eq", Desc: "equal", Op: "=="},
	{Name: "Neq", Desc: "not equal", Op: "!="},
	{Name: "Lt", Desc: "less than", Op: "<"},
	{Name: "Lte", Desc: "less than or equal", Op: "<="},
	{Name: "Gt", Desc: "greater than", Op: ">"},
	{Name: "Gte", Desc: "greater than or equal", Op: ">="},
	/*
		&&    conditional AND
		||    conditional OR
		!     conditional NOT
	*/
	{Name: "And", Desc: "conditional and", Op: "&&"},
	{Name: "Or", Desc: "conditional or", Op: "||"},
	{Name: "Not", Desc: "conditional not", Op: "!"},
	/*
		<-    receive
	*/
	{Name: "Rcv", Desc: "channel receive", Op: "<-"},
	/*
		...   variadic
	*/
	{Name: "Vari", Desc: "variadic", Op: "..."},
	/*
		++ 	  increment
		--    decrement
		+=    increment assign
		-=    decrement assign
		*=    product assign
		/=    quotient assign
		%=    remainder assign
	*/
	{Name: "Inc", Desc: "increment", Op: "++"},
	{Name: "Dec", Desc: "decrement", Op: "--"},
	/*
		*     pointer
		&     address
	*/
	{Name: "Ptr", Desc: "pointer", Op: "*"},
	{Name: "Adr", Desc: "address", Op: "&"},
}

var Identifiers = append(append(append([]string{}, Types...), Constants...), Zero...)
