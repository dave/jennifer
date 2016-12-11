package jen

import (
	"bytes"
	"io"
	"sort"
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
	t := token{
		Group:   g,
		typ:     literalToken,
		content: v,
	}
	g.items = append(g.items, t)
	return g
}

func LitFunc(f func() interface{}) *Group {
	return newStatement().LitFunc(f)
}

func (g *Group) LitFunc(f func() interface{}) *Group {
	if startNewStatement(g.syntax) {
		s := LitFunc(f)
		g.items = append(g.items, s)
		return s
	}
	t := token{
		Group:   g,
		typ:     literalToken,
		content: f(),
	}
	g.items = append(g.items, t)
	return g
}

// Dict inserts a map literal
func Dict(m map[Code]Code) *Group {
	return newStatement().Dict(m)
}

// Dict inserts a map literal
func (g *Group) Dict(m map[Code]Code) *Group {
	if startNewStatement(g.syntax) {
		s := Dict(m)
		g.items = append(g.items, s)
		return s
	}
	ml := dict{
		Group: g,
		m:     m,
	}
	g.items = append(g.items, ml)
	return g
}

func DictFunc(f func(map[Code]Code)) *Group {
	return newStatement().DictFunc(f)
}

func (g *Group) DictFunc(f func(map[Code]Code)) *Group {
	if startNewStatement(g.syntax) {
		s := DictFunc(f)
		g.items = append(g.items, s)
		return s
	}
	m := map[Code]Code{}
	f(m)
	ml := dict{
		Group: g,
		m:     m,
	}
	g.items = append(g.items, ml)
	return g
}

type dict struct {
	*Group
	m map[Code]Code
}

func (l dict) isNull() bool {
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
		if v.isNull() {
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
