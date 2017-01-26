package jen

import (
	"bytes"
	"io"
	"sort"
)

func Lit(v interface{}) *Statement {
	return newStatement().Lit(v)
}

func (g *Group) Lit(v interface{}) *Statement {
	s := Lit(v)
	g.items = append(g.items, s)
	return s
}

func (s *Statement) Lit(v interface{}) *Statement {
	t := token{
		typ:     literalToken,
		content: v,
	}
	*s = append(*s, t)
	return s
}

func LitFunc(f func() interface{}) *Statement {
	return newStatement().LitFunc(f)
}

func (g *Group) LitFunc(f func() interface{}) *Statement {
	s := LitFunc(f)
	g.items = append(g.items, s)
	return s
}

func (s *Statement) LitFunc(f func() interface{}) *Statement {
	t := token{
		typ:     literalToken,
		content: f(),
	}
	*s = append(*s, t)
	return s
}

// Dict inserts a map literal
func Dict(m map[Code]Code) *Statement {
	return newStatement().Dict(m)
}

// Dict inserts a map literal
func (g *Group) Dict(m map[Code]Code) *Statement {
	s := Dict(m)
	g.items = append(g.items, s)
	return s
}

// Dict inserts a map literal
func (s *Statement) Dict(m map[Code]Code) *Statement {
	d := dict{
		m: m,
	}
	*s = append(*s, d)
	return s
}

func DictFunc(f func(map[Code]Code)) *Statement {
	return newStatement().DictFunc(f)
}

func (g *Group) DictFunc(f func(map[Code]Code)) *Statement {
	s := DictFunc(f)
	g.items = append(g.items, s)
	return s
}

func (s *Statement) DictFunc(f func(map[Code]Code)) *Statement {
	m := map[Code]Code{}
	f(m)
	ml := dict{
		m: m,
	}
	*s = append(*s, ml)
	return s
}

type dict struct {
	m map[Code]Code
}

func (l dict) isNull(f *File) bool {
	return false
}

func (l dict) render(f *File, w io.Writer) error {
	if _, err := w.Write([]byte("{")); err != nil {
		return err
	}
	first := true
	// must order keys to ensure repeatable source
	type kv struct {
		k Code
		v Code
	}
	lookup := map[string]kv{}
	keys := []string{}
	for k, v := range l.m {
		buf := &bytes.Buffer{}
		if err := k.render(f, buf); err != nil {
			return err
		}
		keys = append(keys, buf.String())
		lookup[buf.String()] = kv{k: k, v: v}
	}
	sort.Strings(keys)
	for _, key := range keys {
		k := lookup[key].k
		v := lookup[key].v
		if v.isNull(f) {
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
		if err := k.render(f, w); err != nil {
			return err
		}
		if _, err := w.Write([]byte(":")); err != nil {
			return err
		}
		if err := v.render(f, w); err != nil {
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
