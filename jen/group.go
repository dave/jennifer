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

// Add creates a new statement and appends the provided code.
//func Add(code ...Code) *Group {
//	return newStatement().Add(code...)
//}

// Add appends the provided code to the group.
func (g *Group) Add(code ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := newStatement(code...)
		g.items = append(g.items, s)
		return s
	}
	g.items = append(g.items, code...)
	return g
}

// Do creates a new statement and calls the provided function with it as a
// parameter
func Do(f func(*Group)) *Group {
	return newStatement().Do(f)
}

// Do calls the provided function with the group as a parameter
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
		Open:      "(",
		Close:     ")",
		Seperator: " ",
	},
	listSyntax: {
		Seperator: ",",
	},
	clauseSyntax: {
		Seperator: ";",
	},
	//bracesSyntax: {
	//	Open:  "{",
	//	Close: "}",
	//},
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

func (g Group) isNull() bool {
	i := info[g.syntax.typ]
	if i.Open != "" || i.Close != "" {
		return false
	}
	for _, c := range g.items {
		if !c.isNull() {
			return false
		}
	}
	return true
}

func (g Group) render(ctx context.Context, w io.Writer) error {
	i := info[g.syntax.typ]
	if i.Open != "" {
		if _, err := w.Write([]byte(i.Open)); err != nil {
			return err
		}
	}
	first := true
	for _, code := range g.items {
		if code.isNull() {
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
		if err := code.render(ctx, w); err != nil {
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
	if err := g.render(ctx, buf); err != nil {
		panic(err)
	}
	b, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	return string(b)
}
