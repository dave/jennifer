package jen

import (
	"context"
	"io"

	"github.com/davelondon/jennifer/jen/data"
)

type Group struct {
	syntax syntaxType
	items  []Code
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

func (g *Group) AddFunc(f func() Code) *Group {
	if startNewStatement(g.syntax) {
		s := newStatement().AddFunc(f)
		g.items = append(g.items, s)
		return s
	}
	code := f()
	g.items = append(g.items, code)
	return g
}

func (g Group) IsNull() bool {
	d := data.Blocks[string(g.syntax)]
	if d.Open != "" || d.Close != "" {
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
	d := data.Blocks[string(g.syntax)]
	if d.Open != "" {
		if _, err := w.Write([]byte(d.Open)); err != nil {
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
		if !first && d.Seperator != "" {
			if _, err := w.Write([]byte(d.Seperator)); err != nil {
				return err
			}
		}
		if err := code.Render(ctx, w); err != nil {
			return err
		}
		first = false
	}
	if d.Close != "" {
		if _, err := w.Write([]byte(d.Close)); err != nil {
			return err
		}
	}
	return nil
}
