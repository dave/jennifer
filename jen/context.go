package jen

import (
	"context"
	"fmt"
	"strings"
)

type syntaxType string

const syntaxKey syntaxType = ""

const (
	FileSyntax      syntaxType = "file"
	ParamsSyntax    syntaxType = "params"
	StatementSyntax syntaxType = "statement"
	BlockSyntax     syntaxType = "block"
	ParensSyntax    syntaxType = "parens"
	ValuesSyntax    syntaxType = "values"
	IndexSyntax     syntaxType = "index"
	CallSyntax      syntaxType = "call"
	DeclsSyntax     syntaxType = "decls"
	BracesSyntax    syntaxType = "braces"
	ListSyntax      syntaxType = "list"
)

type contextType int

var contextKey contextType

func Context(ctx context.Context, name string) context.Context {
	f := &global{
		Name:    name,
		Imports: make(map[string]string),
	}
	return context.WithValue(ctx, contextKey, f)
}

func ContextPath(ctx context.Context, name, path string) context.Context {
	f := &global{
		Name:    name,
		Path:    path,
		Imports: make(map[string]string),
	}
	return context.WithValue(ctx, contextKey, f)
}

func FromContext(ctx context.Context) *global {
	val := ctx.Value(contextKey)
	if val == nil {
		panic("jen context not found")
	}
	return val.(*global)
}

type global struct {
	Name    string
	Path    string
	Imports map[string]string
}

func (f *global) register(path string) string {
	if f.Imports[path] != "" && f.Imports[path] != "_" {
		return f.Imports[path]
	}
	alias := ""
	if sep := strings.LastIndex(path, "/"); sep > -1 {
		alias = path[sep+1:]
	} else {
		alias = path
	}
	unique := alias
	find := func(a string) bool {
		for _, v := range f.Imports {
			if a == v {
				return true
			}
		}
		return false
	}
	i := 0
	for find(unique) {
		i++
		unique = fmt.Sprintf("%s%d", alias, i)
	}
	f.Imports[path] = unique
	return unique
}
