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
		open:      ".(",
		close:     ")",
		seperator: " ",
	},
}
