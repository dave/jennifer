package jen

import (
	"fmt"
	"io"
	"strings"
)

type tokenType string

const (
	identifierToken tokenType = "identifier"
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

func (t token) isNull() bool {
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
	case identifierToken:
		id := t.content.(string)
		var name, path, alias, full string
		if sep := strings.LastIndex(id, "."); sep > -1 {
			name = id[sep+1:]
			path = id[:sep]
			alias = f.register(path)
		} else {
			name = id
		}
		if alias != "" {
			full = fmt.Sprintf("%s.%s", alias, name)
		} else {
			full = fmt.Sprintf("%s", name)
		}
		if _, err := w.Write([]byte(full)); err != nil {
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

func Id(items ...interface{}) *Statement {
	return newStatement().Id(items...)
}

func (g *Group) Id(items ...interface{}) *Statement {
	s := Id(items...)
	g.items = append(g.items, s)
	return s
}

func (s *Statement) Id(items ...interface{}) *Statement {
	first := true
	for _, item := range items {
		if !first {
			s.Op(".")
		}
		switch item := item.(type) {
		case string:
			first = false
			t := token{
				typ:     identifierToken,
				content: item,
			}
			*s = append(*s, t)
		case Code:
			if item.isNull() {
				break
			}
			first = false
			*s = append(*s, item)
		}
	}
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

func Dot() *Statement {
	return newStatement().Dot()
}

func (g *Group) Dot() *Statement {
	s := Dot()
	g.items = append(g.items, s)
	return s
}

func (s *Statement) Dot() *Statement {
	t := token{
		typ:     delimiterToken,
		content: ".",
	}
	*s = append(*s, t)
	return s
}

func Comma() *Statement {
	return newStatement().Comma()
}

func (g *Group) Comma() *Statement {
	s := Comma()
	g.items = append(g.items, s)
	return s
}

func (s *Statement) Comma() *Statement {
	t := token{
		typ:     delimiterToken,
		content: ",",
	}
	*s = append(*s, t)
	return s
}
