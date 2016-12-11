package jen

import (
	"bytes"
	"go/format"
	"io"
)

type Group struct {
	syntax syntaxType
	items  []Code
}

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

func (g Group) isNull() bool {
	i := syntaxInfo[g.syntax]
	if i.open != "" || i.close != "" {
		return false
	}
	for _, c := range g.items {
		if !c.isNull() {
			return false
		}
	}
	return true
}

func (g Group) render(f *File, w io.Writer) error {
	i := syntaxInfo[g.syntax]
	if i.open != "" {
		if _, err := w.Write([]byte(i.open)); err != nil {
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
		if !first && i.seperator != "" {
			if _, err := w.Write([]byte(i.seperator)); err != nil {
				return err
			}
		}
		if err := code.render(f, w); err != nil {
			return err
		}
		first = false
	}
	if i.close != "" {
		if _, err := w.Write([]byte(i.close)); err != nil {
			return err
		}
	}
	return nil
}

func (g *Group) GoString() string {
	f := NewFile("")
	buf := &bytes.Buffer{}
	if err := g.render(f, buf); err != nil {
		panic(err)
	}
	b, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	return string(b)
}
