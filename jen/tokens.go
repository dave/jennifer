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
	*Statement
	typ     tokenType
	content interface{}
}

func (t Token) IsNull() bool {
	return t.typ == nullToken
}

func (t Token) Render(ctx context.Context, w io.Writer) error {
	switch t.typ {
	case literalToken:
		// TODO: this does not work in all cases
		if _, err := w.Write([]byte(fmt.Sprintf("%#v ", t.content))); err != nil {
			return err
		}
	case keywordToken, operatorToken:
		if _, err := w.Write([]byte(fmt.Sprintf("%s ", t.content))); err != nil {
			return err
		}
	case identifierToken:
		id := t.content.(string)
		var name, path, alias, full string
		if sep := strings.LastIndex(id, "."); sep > -1 {
			name = id[sep+1:]
			path = id[:sep]
			file := FromContext(ctx)
			if file.Path != path {
				alias = file.register(path)
			}
		} else {
			name = id
		}
		if alias != "" {
			full = fmt.Sprintf("%s.%s ", alias, name)
		} else {
			full = fmt.Sprintf("%s ", name)
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
// in a block.
func Null() *Statement {
	s := new(Statement)
	return s.Null()
}

// Null token produces no output but also no separator
// in a block.
func (l *StatementList) Null() *Statement {
	s := Null()
	*l = append(*l, s)
	return s
}

// Null token produces no output but also no separator
// in a block.
func (s *Statement) Null() *Statement {
	t := Token{
		Statement: s,
		typ:       nullToken,
	}
	*s = append(*s, t)
	return s
}

// Empty token produces no output but is followed by a
// separator in a block.
func Empty() *Statement {
	s := new(Statement)
	return s.Empty()
}

// Empty token produces no output but is followed by a
// separator in a block.
func (l *StatementList) Empty() *Statement {
	s := Empty()
	*l = append(*l, s)
	return s
}

// Empty token produces no output but is followed by a
// separator in a block.
func (s *Statement) Empty() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "",
	}
	*s = append(*s, t)
	return s
}

func Op(op string) *Statement {
	s := new(Statement)
	return s.Op(op)
}

func (l *StatementList) Op(op string) *Statement {
	s := Op(op)
	*l = append(*l, s)
	return s
}

func (s *Statement) Op(op string) *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   op,
	}
	*s = append(*s, t)
	return s
}

func Id(name string) *Statement {
	s := new(Statement)
	return s.Id(name)
}

func (l *StatementList) Id(name string) *Statement {
	s := Id(name)
	*l = append(*l, s)
	return s
}

func (s *Statement) Id(name string) *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   name,
	}
	*s = append(*s, t)
	return s
}

func Lit(v interface{}) *Statement {
	s := new(Statement)
	return s.Lit(v)
}

func (l *StatementList) Lit(v interface{}) *Statement {
	s := Lit(v)
	*l = append(*l, s)
	return s
}

func (s *Statement) Lit(v interface{}) *Statement {
	t := Token{
		Statement: s,
		typ:       literalToken,
		content:   v,
	}
	*s = append(*s, t)
	return s
}
