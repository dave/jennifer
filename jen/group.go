package jen

import (
	"bytes"
	"context"
	"go/format"
	"io"
)

type Group struct {
	syntax syntax
	items  []Code
}

func Add(code ...Code) *Group {
	return newStatement().Add(code...)
}

func (g *Group) Add(code ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := newStatement(code...)
		g.items = append(g.items, s)
		return s
	}
	g.items = append(g.items, code...)
	return g
}

func Do(f func(*Group)) *Group {
	return newStatement().Do(f)
}

func (g *Group) Do(f func(*Group)) *Group {
	if startNewStatement(g.syntax) {
		s := newStatement().Do(f)
		g.items = append(g.items, s)
		return s
	}
	f(g)
	return g
}

var info = map[syntaxType]struct {
	Open      string
	Close     string
	Seperator string
}{
	fileSyntax: {
		Seperator: "\n",
	},
	statementSyntax: {
		Seperator: " ",
	},
	parensSyntax: {
		Open:  "(",
		Close: ")",
	},
	listSyntax: {
		Seperator: ",",
	},
	bracesSyntax: {
		Open:  "{",
		Close: "}",
	},
	valuesSyntax: {
		Open:      "{",
		Close:     "}",
		Seperator: ",",
	},
	indexSyntax: {
		Open:      "[",
		Close:     "]",
		Seperator: ":",
	},
	blockSyntax: {
		Open:      "{",
		Close:     "}",
		Seperator: "\n",
	},
	callSyntax: {
		Open:      "(",
		Close:     ")",
		Seperator: ",",
	},
	paramsSyntax: {
		Open:      "(",
		Close:     ")",
		Seperator: ",",
	},
	declsSyntax: {
		Open:      "(",
		Close:     ")",
		Seperator: ";",
	},
}

func (g Group) IsNull() bool {
	i := info[g.syntax.typ]
	if i.Open != "" || i.Close != "" {
		return false
	}
	for _, c := range g.items {
		if !c.IsNull() {
			return false
		}
	}
	return true
}

func (g Group) Render(ctx context.Context, w io.Writer) error {
	i := info[g.syntax.typ]
	if i.Open != "" {
		if _, err := w.Write([]byte(i.Open)); err != nil {
			return err
		}
	}
	first := true
	for _, code := range g.items {
		if code.IsNull() {
			// Null() token produces no output but also
			// no separator. Empty() token products no
			// output but adds a separator.
			continue
		}
		if !first && i.Seperator != "" {
			if _, err := w.Write([]byte(i.Seperator)); err != nil {
				return err
			}
		}
		if err := code.Render(ctx, w); err != nil {
			return err
		}
		first = false
	}
	if i.Close != "" {
		if _, err := w.Write([]byte(i.Close)); err != nil {
			return err
		}
	}
	return nil
}

func (g *Group) GoString() string {
	ctx := Context(context.Background())
	if g.syntax.typ == fileSyntax {
		buf := &bytes.Buffer{}
		if err := RenderFile(ctx, g, buf); err != nil {
			panic(err)
		}
		return buf.String()
	}
	buf := &bytes.Buffer{}
	if err := g.Render(ctx, buf); err != nil {
		panic(err)
	}
	b, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	return string(b)
}
