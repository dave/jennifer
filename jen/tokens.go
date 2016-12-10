package jen

import (
	"context"
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
)

type Token struct {
	*Group
	typ     tokenType
	content interface{}
}

func (t Token) isNull() bool {
	return t.typ == nullToken
}

func (t Token) render(ctx context.Context, w io.Writer) error {
	switch t.typ {
	case literalToken:
		// TODO: this does not work in all cases
		if _, err := w.Write([]byte(fmt.Sprintf("%#v", t.content))); err != nil {
			return err
		}
	case keywordToken, operatorToken:
		if _, err := w.Write([]byte(fmt.Sprintf("%s", t.content))); err != nil {
			return err
		}
	case identifierToken:
		id := t.content.(string)
		var name, path, alias, full string
		if sep := strings.LastIndex(id, "."); sep > -1 {
			name = id[sep+1:]
			path = id[:sep]
			file := FromContext(ctx)
			alias = file.register(path)
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
func Null() *Group {
	return newStatement().Null()
}

// Null token produces no output but also no separator
// in a list.
func (g *Group) Null() *Group {
	if startNewStatement(g.syntax) {
		s := Null()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group: g,
		typ:   nullToken,
	}
	g.items = append(g.items, t)
	return g
}

// Empty token produces no output but is followed by a
// separator in a list.
func Empty() *Group {
	return newStatement().Empty()
}

// Empty token produces no output but is followed by a
// separator in a list.
func (g *Group) Empty() *Group {
	if startNewStatement(g.syntax) {
		s := Empty()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     operatorToken,
		content: "",
	}
	g.items = append(g.items, t)
	return g
}

func Op(op string) *Group {
	return newStatement().Op(op)
}

func (g *Group) Op(op string) *Group {
	if startNewStatement(g.syntax) {
		s := Op(op)
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     operatorToken,
		content: op,
	}
	g.items = append(g.items, t)
	return g
}

func Id(names ...string) *Group {
	return newStatement().Id(names...)
}

func (g *Group) Id(names ...string) *Group {
	if startNewStatement(g.syntax) {
		s := Id(names...)
		g.items = append(g.items, s)
		return s
	}
	for i, n := range names {
		if i > 0 {
			g.Op(".")
		}
		t := Token{
			Group:   g,
			typ:     identifierToken,
			content: n,
		}
		g.items = append(g.items, t)
	}
	return g
}
