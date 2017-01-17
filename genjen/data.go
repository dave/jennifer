package main

var Keywords = []string{"break", "default", "func", "interface", "select", "case", "defer", "go", "struct", "chan", "else", "goto", "switch", "const", "fallthrough", "range", "type", "continue", "var"}

// "return", "for" and "if" are special cases
// "import" and "package" are handled automatically, so not needed.
// "map" is treated as a block

var Types = []string{"bool", "byte", "complex64", "complex128", "error", "float32", "float64", "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"}

var Constants = []string{"true", "false", "iota"}

var Zero = []string{"nil"}

var Errs = []string{"err"}

var Identifiers = append(append(append(append([]string{}, Types...), Constants...), Zero...), Errs...)

var Functions = []string{"append", "cap", "close", "complex", "copy", "delete", "imag", "len", "make", "new", "panic", "print", "println", "real", "recover"}

var Blocks = []struct {
	Name   string
	Syntax string
	Desc   string
	List   bool
}{
	{
		Name:   "Parens",
		Syntax: "parensSyntax",
		Desc:   "parenthesis",
		List:   false,
	},
	{
		Name:   "List",
		Syntax: "listSyntax",
		Desc:   "a comma separated list",
		List:   true,
	},
	{
		Name:   "Values",
		Syntax: "valuesSyntax",
		Desc:   "curly braces containing a comma separated list",
		List:   true,
	},
	{
		Name:   "Slice",
		Syntax: "valuesSyntax",
		Desc:   "curly braces containing a comma separated list",
		List:   true,
	},
	{
		Name:   "Index",
		Syntax: "indexSyntax",
		Desc:   "square brackets containing a colon separated list",
		List:   true,
	},
	{
		Name:   "Block",
		Syntax: "blockSyntax",
		Desc:   "curly braces containing a statements list",
		List:   true,
	},
	{
		Name:   "Call",
		Syntax: "callSyntax",
		Desc:   "parenthesis containing a comma separated list",
		List:   true,
	},
	{
		Name:   "Params",
		Syntax: "paramsSyntax",
		Desc:   "parenthesis containing a comma separated list",
		List:   true,
	},
	{
		Name:   "Decls",
		Syntax: "declsSyntax",
		Desc:   "parenthesis containing a statement list",
		List:   true,
	},
	{
		Name:   "CaseBlock",
		Syntax: "caseSyntax",
		Desc:   "a statement list preceeded by a colon",
		List:   true,
	},
	{
		Name:   "Assert",
		Syntax: "assertSyntax",
		Desc:   "a type assertion",
		List:   false,
	},
	{
		Name:   "Map",
		Syntax: "mapSyntax",
		Desc:   "the map keyword, followed by square brackets",
		List:   false,
	},
}

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
