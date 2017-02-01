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

func (t token) render(f *File, w io.Writer) error {
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

// Null token produces no output but also no separator
// in a list.
func Null() *Statement {
	return newStatement().Null()
}

// Null token produces no output but also no separator
// in a list.
func (g *Group) Null() *Statement {
	s := Null()
	g.items = append(g.items, s)
	return s
}

// Null token produces no output but also no separator
// in a list.
func (s *Statement) Null() *Statement {
	t := token{
		typ: nullToken,
	}
	*s = append(*s, t)
	return s
}

// Empty token produces no output but is followed by a
// separator in a list.
func Empty() *Statement {
	return newStatement().Empty()
}

// Empty token produces no output but is followed by a
// separator in a list.
func (g *Group) Empty() *Statement {
	s := Empty()
	g.items = append(g.items, s)
	return s
}

// Empty token produces no output but is followed by a
// separator in a list.
func (s *Statement) Empty() *Statement {
	t := token{
		typ:     operatorToken,
		content: "",
	}
	*s = append(*s, t)
	return s
}

func Op(op string) *Statement {
	return newStatement().Op(op)
}

func (g *Group) Op(op string) *Statement {
	s := Op(op)
	g.items = append(g.items, s)
	return s
}

func (s *Statement) Op(op string) *Statement {
	t := token{
		typ:     operatorToken,
		content: op,
	}
	*s = append(*s, t)
	return s
}

func Id(name string) *Statement {
	return newStatement().Id(name)
}

func (g *Group) Id(name string) *Statement {
	s := Id(name)
	g.items = append(g.items, s)
	return s
}

func (s *Statement) Id(name string) *Statement {
	t := token{
		typ:     identifierToken,
		content: name,
	}
	*s = append(*s, t)
	return s
}

func Qual(path, name string) *Statement {
	return newStatement().Qual(path, name)
}

func (g *Group) Qual(path, name string) *Statement {
	s := Qual(path, name)
	g.items = append(g.items, s)
	return s
}

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

func Line() *Statement {
	return newStatement().Line()
}

func (g *Group) Line() *Statement {
	s := Line()
	g.items = append(g.items, s)
	return s
}

func (s *Statement) Line() *Statement {
	t := token{
		typ:     layoutToken,
		content: "\n",
	}
	*s = append(*s, t)
	return s
}
