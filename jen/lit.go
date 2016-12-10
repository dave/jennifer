package jen

import (
	"context"
	"io"
)

func Lit(v interface{}) *Group {
	return newStatement().Lit(v)
}

func (g *Group) Lit(v interface{}) *Group {
	if startNewStatement(g.syntax) {
		s := Lit(v)
		g.items = append(g.items, s)
		return s
	}
	switch v := v.(type) {
	case map[Code]Code:
		ml := mapLit{
			Group: g,
			m:     v,
		}
		g.items = append(g.items, ml)
		return g
	case func(map[Code]Code):
		m := map[Code]Code{}
		v(m)
		ml := mapLit{
			Group: g,
			m:     m,
		}
		g.items = append(g.items, ml)
		return g
	case func() interface{}:
		i := v()
		t := Token{
			Group:   g,
			typ:     literalToken,
			content: i,
		}
		g.items = append(g.items, t)
		return g
	default:
		t := Token{
			Group:   g,
			typ:     literalToken,
			content: v,
		}
		g.items = append(g.items, t)
		return g
	}

}

type mapLit struct {
	*Group
	m map[Code]Code
}

func (l mapLit) IsNull() bool {
	return false
}

func (l mapLit) Render(ctx context.Context, w io.Writer) error {
	if _, err := w.Write([]byte("{")); err != nil {
		return err
	}
	first := true
	for k, v := range l.m {
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
