package jen

type syntaxType string

const (
	fileSyntax      syntaxType = "file"
	paramsSyntax    syntaxType = "params"
	statementSyntax syntaxType = "statement"
	blockSyntax     syntaxType = "block"
	parensSyntax    syntaxType = "parens"
	valuesSyntax    syntaxType = "values"
	indexSyntax     syntaxType = "index"
	callSyntax      syntaxType = "call"
	declsSyntax     syntaxType = "decls"
	listSyntax      syntaxType = "list"
	clauseSyntax    syntaxType = "clause"
)

var syntaxInfo = map[syntaxType]struct {
	open      string
	close     string
	seperator string
}{
	fileSyntax: {
		seperator: "\n",
	},
	statementSyntax: {
		seperator: " ",
	},
	parensSyntax: {
		open:      "(",
		close:     ")",
		seperator: " ",
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
		open:      "{",
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
}
