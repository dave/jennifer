package main

var Keywords = []string{"break", "default", "func", "select", "defer", "go", "struct", "chan", "else", "goto", "const", "fallthrough", "range", "type", "continue", "var"}

// "return", "map", "switch", "for", "interface", "case" and "if" are special cases
// "import" and "package" are handled automatically, so not needed.

var Types = []string{"bool", "byte", "complex64", "complex128", "error", "float32", "float64", "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"}

var Constants = []string{"true", "false", "iota"}

var Zero = []string{"nil"}

var Errs = []string{"err"}

var Identifiers = append(append(append(append([]string{}, Types...), Constants...), Zero...), Errs...)

var Functions = []string{"append", "cap", "close", "complex", "copy", "delete", "imag", "len", "make", "new", "panic", "print", "println", "real", "recover"}

var Groups = []struct {
	Name      string
	Desc      string
	List      bool
	Open      string
	Close     string
	Separator string
}{
	{
		Name:      "Parens",
		Desc:      "renders a single item in parenthesis. Use for type conversion or to specify evaluation order.",
		List:      false,
		Open:      "(",
		Close:     ")",
		Separator: "",
	},
	{
		Name:      "List",
		Desc:      "renders a comma separated list with no open or closing tokens. Use for multiple return functions.",
		List:      true,
		Open:      "",
		Close:     "",
		Separator: ",",
	},
	{
		Name:      "Values",
		Desc:      "renders a comma separated list enclosed by curly braces. Use for slice literals.",
		List:      true,
		Open:      "{",
		Close:     "}",
		Separator: ",",
	},
	{
		Name:      "Index",
		Desc:      "renders a colon separated list enclosed by square brackets. Use for array / slice indexes and definitions.",
		List:      true,
		Open:      "[",
		Close:     "]",
		Separator: ":",
	},
	{
		Name:      "Block",
		Desc:      "renders a statement list enclosed by curly braces. Use for all code blocks.",
		List:      true,
		Open:      "{",
		Close:     "}",
		Separator: "\n",
	},
	{
		Name:      "Defs",
		Desc:      "renders a list of statements enclosed in parenthesis. Use for definition lists.",
		List:      true,
		Open:      "(",
		Close:     ")",
		Separator: "\n",
	},
	{
		Name:      "Call",
		Desc:      "renders a comma separated list enclosed by parenthesis. Use for function calls.",
		List:      true,
		Open:      "(",
		Close:     ")",
		Separator: ",",
	},
	{
		Name:      "Params",
		Desc:      "renders a comma separated list enclosed by parenthesis. Use for function parameters and method receivers.",
		List:      true,
		Open:      "(",
		Close:     ")",
		Separator: ",",
	},
	{
		Name:      "CaseBlock",
		Desc:      "renders a statement list preceded by a colon. Use to build switch / select statements.",
		List:      true,
		Open:      ":",
		Close:     "",
		Separator: "\n",
	},
	{
		Name:      "Assert",
		Desc:      "renders a period followed by a single item enclosed by parenthesis. Use for type assertions.",
		List:      false,
		Open:      ".(",
		Close:     ")",
		Separator: "",
	},
	{
		Name:      "Map",
		Desc:      "renders the map keyword followed by a single item enclosed by square brackets. Use for map definitions.",
		List:      false,
		Open:      "map[",
		Close:     "]",
		Separator: "",
	},
	{
		Name:      "If",
		Desc:      "renders the if keyword followed by a semicolon separated list.",
		List:      true,
		Open:      "if ",
		Close:     "",
		Separator: ";",
	},
	{
		Name:      "Return",
		Desc:      "renders the return keyword, followed by a comma separated list.",
		List:      true,
		Open:      "return ",
		Close:     "",
		Separator: ",",
	},
	{
		Name:      "For",
		Desc:      "renders the for keyword, followed by a semicolon separated list.",
		List:      true,
		Open:      "for ",
		Close:     "",
		Separator: ";",
	},
	{
		Name:      "Switch",
		Desc:      "renders the switch keyword, followed by a semicolon separated list.",
		List:      true,
		Open:      "switch ",
		Close:     "",
		Separator: ";",
	},
	{
		Name:      "Interface",
		Desc:      "renders the interface keyword, followed by curly braces containing a statement list.",
		List:      true,
		Open:      "interface{",
		Close:     "}",
		Separator: "\n",
	},
	{
		Name:      "Case",
		Desc:      "renders the case keyword, followed by a comma separated list.",
		List:      true,
		Open:      "case ",
		Close:     "",
		Separator: ",",
	},
	{
		Name:      "Sel",
		Desc:      "renders a chain of selectors separated by periods.",
		List:      true,
		Open:      "",
		Close:     "",
		Separator: ".",
	},
}
