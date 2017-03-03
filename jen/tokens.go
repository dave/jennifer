package jen

import (
	"fmt"
	"io"
)

type tokenType string

const (
	packageToken    tokenType = "package"
	identifierToken tokenType = "identifier"
	qualifiedToken  tokenType = "qualified"
	keywordToken    tokenType = "keyword"
	operatorToken   tokenType = "operator"
	delimiterToken  tokenType = "delimiter"
	literalToken    tokenType = "literal"
	nullToken       tokenType = "null"
	layoutToken     tokenType = "layout"
)

type token struct {
	typ     tokenType
	content interface{}
}

func (t token) isNull(f *File) bool {
	if t.typ == packageToken {
		// package token is null if the path is the local package path
		return f.isLocal(t.content.(string))
	}
	return t.typ == nullToken
}

func (t token) render(f *File, w io.Writer, s *Statement) error {
	switch t.typ {
	case literalToken:
		// TODO: Does this work in all cases?
		if _, err := w.Write([]byte(fmt.Sprintf("%#v", t.content))); err != nil {
			return err
		}
	case keywordToken, operatorToken, layoutToken:
		if _, err := w.Write([]byte(fmt.Sprintf("%s", t.content))); err != nil {
			return err
		}
	case packageToken:
		path := t.content.(string)
		alias := f.register(path)
		if _, err := w.Write([]byte(alias)); err != nil {
			return err
		}
	case identifierToken:
		if _, err := w.Write([]byte(t.content.(string))); err != nil {
			return err
		}
	case nullToken:
		// do nothing
	}
	return nil
}

// Null adds a null item. Null items render nothing and are not followed by a
// separator in lists.
func Null() *Statement {
	return newStatement().Null()
}

// Null adds a null item. Null items render nothing and are not followed by a
// separator in lists.
func (g *Group) Null() *Statement {
	s := Null()
	g.items = append(g.items, s)
	return s
}

// Null adds a null item. Null items render nothing and are not followed by a
// separator in lists.
func (s *Statement) Null() *Statement {
	t := token{
		typ: nullToken,
	}
	*s = append(*s, t)
	return s
}

// Empty adds an empty item. Empty items render nothing but are followed by a
// separator in lists.
func Empty() *Statement {
	return newStatement().Empty()
}

// Empty adds an empty item. Empty items render nothing but are followed by a
// separator in lists.
func (g *Group) Empty() *Statement {
	s := Empty()
	g.items = append(g.items, s)
	return s
}

// Empty adds an empty item. Empty items render nothing but are followed by a
// separator in lists.
func (s *Statement) Empty() *Statement {
	t := token{
		typ:     operatorToken,
		content: "",
	}
	*s = append(*s, t)
	return s
}

// Op renders the provided operator / token.
func Op(op string) *Statement {
	return newStatement().Op(op)
}

// Op renders the provided operator / token.
func (g *Group) Op(op string) *Statement {
	s := Op(op)
	g.items = append(g.items, s)
	return s
}

// Op renders the provided operator / token.
func (s *Statement) Op(op string) *Statement {
	t := token{
		typ:     operatorToken,
		content: op,
	}
	*s = append(*s, t)
	return s
}

// Id renders an identifier.
func Id(name string) *Statement {
	return newStatement().Id(name)
}

// Id renders an identifier.
func (g *Group) Id(name string) *Statement {
	s := Id(name)
	g.items = append(g.items, s)
	return s
}

// Id renders an identifier.
func (s *Statement) Id(name string) *Statement {
	t := token{
		typ:     identifierToken,
		content: name,
	}
	*s = append(*s, t)
	return s
}

// Qual renders a qualified identifier. Imports are automatically added when
// used with a File. If the path matches the local path, the package name is
// omitted. If package names conflict they are automatically renamed.
func Qual(path, name string) *Statement {
	return newStatement().Qual(path, name)
}

// Qual renders a qualified identifier. Imports are automatically added when
// used with a File. If the path matches the local path, the package name is
// omitted. If package names conflict they are automatically renamed.
func (g *Group) Qual(path, name string) *Statement {
	s := Qual(path, name)
	g.items = append(g.items, s)
	return s
}

// Qual renders a qualified identifier. Imports are automatically added when
// used with a File. If the path matches the local path, the package name is
// omitted. If package names conflict they are automatically renamed.
func (s *Statement) Qual(path, name string) *Statement {
	g := Sel(
		token{
			typ:     packageToken,
			content: path,
		},
		token{
			typ:     identifierToken,
			content: name,
		},
	)
	*s = append(*s, g)
	return s
}

// Line inserts a blank line.
func Line() *Statement {
	return newStatement().Line()
}

// Line inserts a blank line.
func (g *Group) Line() *Statement {
	s := Line()
	g.items = append(g.items, s)
	return s
}

// Line inserts a blank line.
func (s *Statement) Line() *Statement {
	t := token{
		typ:     layoutToken,
		content: "\n",
	}
	*s = append(*s, t)
	return s
}
