package jen

type syntaxType string

const (
	fileSyntax   syntaxType = "file"
	paramsSyntax syntaxType = "params"
	blockSyntax  syntaxType = "block"
	parensSyntax syntaxType = "parens"
	valuesSyntax syntaxType = "values"
	indexSyntax  syntaxType = "index"
	callSyntax   syntaxType = "call"
	declsSyntax  syntaxType = "decls"
	listSyntax   syntaxType = "list"
	clauseSyntax syntaxType = "clause"
	caseSyntax   syntaxType = "case"
	assertSyntax syntaxType = "assert"
	mapSyntax    syntaxType = "map"
	ifSyntax     syntaxType = "if"
	returnSyntax syntaxType = "return"
	forSyntax    syntaxType = "for"
)

var syntaxInfo = map[syntaxType]struct {
	open      string
	close     string
	seperator string
}{
	fileSyntax: {
		seperator: "\n",
	},
	parensSyntax: {
		open:  "(",
		close: ")",
	},
	listSyntax: {
		seperator: ",",
	},
	clauseSyntax: {
		seperator: ";",
	},
	valuesSyntax: {
		open:      "{",
		close:     "}",
		seperator: ",",
	},
	indexSyntax: {
		open:      "[",
		close:     "]",
		seperator: ":",
	},
	blockSyntax: {
		open:      "{\n",
		close:     "}",
		seperator: "\n",
	},
	callSyntax: {
		open:      "(",
		close:     ")",
		seperator: ",",
	},
	paramsSyntax: {
		open:      "(",
		close:     ")",
		seperator: ",",
	},
	declsSyntax: {
		open:      "(",
		close:     ")",
		seperator: ";",
	},
	caseSyntax: {
		open:      ":\n",
		seperator: "\n",
	},
	assertSyntax: {
		open:  ".(",
		close: ")",
	},
	mapSyntax: {
		open:  "map[",
		close: "]",
	},
	ifSyntax: {
		open:      "if ",
		close:     "",
		seperator: ";",
	},
	returnSyntax: {
		open:      "return ",
		close:     "",
		seperator: ",",
	},
	forSyntax: {
		open:      "for ",
		close:     "",
		seperator: ";",
	},
}