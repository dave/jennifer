package jen

type syntax struct {
	typ  syntaxType
	name string
	path string
}

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
