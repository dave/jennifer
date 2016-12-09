package jen

import (
	"context"
	"io"
)

func MapLit(lit map[Code]Code) *Statement {
	s := new(Statement)
	return s.MapLit(lit)
}

func (l *StatementList) MapLit(lit map[Code]Code) *Statement {
	s := MapLit(lit)
	*l = append(*l, s)
	return s
}

func (s *Statement) MapLit(lit map[Code]Code) *Statement {
	l := mapLit{
		Statement: s,
		m:         lit,
	}
	*s = append(*s, l)
	return s
}

func MapLitFunc(f func(map[Code]Code)) *Statement {
	s := new(Statement)
	return s.MapLitFunc(f)
}

func (l *StatementList) MapLitFunc(f func(map[Code]Code)) *Statement {
	s := MapLitFunc(f)
	*l = append(*l, s)
	return s
}

func (s *Statement) MapLitFunc(f func(map[Code]Code)) *Statement {
	l := mapLit{
		Statement: s,
		f:         f,
	}
	*s = append(*s, l)
	return s
}

type mapLit struct {
	*Statement
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
	if _, err := w.Write([]byte("} ")); err != nil {
		return err
	}
	return nil
}
