package jen

import (
	"context"
	"io"
)

func MapLit(lit map[Code]Code) *Group {
	return newStatement().MapLit(lit)
}

func (g *Group) MapLit(code map[Code]Code) *Group {
	if startNewStatement(g.syntax) {
		s := MapLit(code)
		g.items = append(g.items, s)
		return s
	}
	m := mapLit{
		Group: g,
		m:     code,
	}
	g.items = append(g.items, m)
	return g
}

func MapLitFunc(f func(map[Code]Code)) *Group {
	return newStatement().MapLitFunc(f)
}

func (g *Group) MapLitFunc(f func(map[Code]Code)) *Group {
	if startNewStatement(g.syntax) {
		s := MapLitFunc(f)
		g.items = append(g.items, s)
		return s
	}
	m := mapLit{
		Group: g,
		f:     f,
	}
	g.items = append(g.items, m)
	return g
}

type mapLit struct {
	*Group
	m map[Code]Code
	f func(map[Code]Code)
}

func (l mapLit) IsNull() bool {
	return false
}

func (l mapLit) Render(ctx context.Context, w io.Writer) error {
	if _, err := w.Write([]byte("{")); err != nil {
		return err
	}
	m := l.m
	if m == nil {
		m = map[Code]Code{}
	}
	if l.f != nil {
		l.f(m)
	}
	first := true
	for k, v := range m {
		if v.IsNull() {
			// Null() token produces no output but also
			// no separator. Empty() token products no
			// output but adds a separator.
			continue
		}
		if first {
			if _, err := w.Write([]byte("\n")); err != nil {
				return err
			}
			first = false
		}
		if err := k.Render(ctx, w); err != nil {
			return err
		}
		if _, err := w.Write([]byte(":")); err != nil {
			return err
		}
		if err := v.Render(ctx, w); err != nil {
			return err
		}
		if _, err := w.Write([]byte(",\n")); err != nil {
			return err
		}
	}
	if _, err := w.Write([]byte("}")); err != nil {
		return err
	}
	return nil
}
