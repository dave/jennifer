package jen

import (
	"bytes"
	"io"
	"sort"
)

// Lit renders a literal, using the format provided by the fmt package %#v
// verb.
func Lit(v interface{}) *Statement {
	return newStatement().Lit(v)
}

// Lit renders a literal, using the format provided by the fmt package %#v
// verb.
func (g *Group) Lit(v interface{}) *Statement {
	s := Lit(v)
	g.items = append(g.items, s)
	return s
}

// Lit renders a literal, using the format provided by the fmt package %#v
// verb.
func (s *Statement) Lit(v interface{}) *Statement {
	t := token{
		typ:     literalToken,
		content: v,
	}
	*s = append(*s, t)
	return s
}

// LitFunc renders a literal, using the format provided by the fmt package %#v
// verb. LitFunc generates the value to render by executing the provided
// function.
func LitFunc(f func() interface{}) *Statement {
	return newStatement().LitFunc(f)
}

// LitFunc renders a literal, using the format provided by the fmt package %#v
// verb. LitFunc generates the value to render by executing the provided
// function.
func (g *Group) LitFunc(f func() interface{}) *Statement {
	s := LitFunc(f)
	g.items = append(g.items, s)
	return s
}

// LitFunc renders a literal, using the format provided by the fmt package %#v
// verb. LitFunc generates the value to render by executing the provided
// function.
func (s *Statement) LitFunc(f func() interface{}) *Statement {
	t := token{
		typ:     literalToken,
		content: f(),
	}
	*s = append(*s, t)
	return s
}

// Dict takes a map[Code]Code and renders a list of colon separated key value
// pairs, enclosed in curly braces. Use for map or composite literals.
func Dict(m map[Code]Code) *Statement {
	return newStatement().Dict(m)
}

// Dict takes a map[Code]Code and renders a list of colon separated key value
// pairs, enclosed in curly braces. Use for map or composite literals.
func (g *Group) Dict(m map[Code]Code) *Statement {
	s := Dict(m)
	g.items = append(g.items, s)
	return s
}

// Dict takes a map[Code]Code and renders a list of colon separated key value
// pairs, enclosed in curly braces. Use for map or composite literals.
func (s *Statement) Dict(m map[Code]Code) *Statement {
	d := dict{
		m: m,
	}
	*s = append(*s, d)
	return s
}

// DictFunc executes a func(map[Code]Code) to generate the value. The value is
// rendered as a list of colon separated key value pairs, enclosed in curly
// braces. Use for map or composite literals.
func DictFunc(f func(map[Code]Code)) *Statement {
	return newStatement().DictFunc(f)
}

// DictFunc executes a func(map[Code]Code) to generate the value. The value is
// rendered as a list of colon separated key value pairs, enclosed in curly
// braces. Use for map or composite literals.
func (g *Group) DictFunc(f func(map[Code]Code)) *Statement {
	s := DictFunc(f)
	g.items = append(g.items, s)
	return s
}

// DictFunc executes a func(map[Code]Code) to generate the value. The value is
// rendered as a list of colon separated key value pairs, enclosed in curly
// braces. Use for map or composite literals.
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

func (l dict) render(f *File, w io.Writer, s *Statement) error {
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
		if err := k.render(f, buf, nil); err != nil {
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
			// Null() token renders no output but also no separator. Empty()
			// token renders no output but adds a separator.
			continue
		}
		if first {
			if _, err := w.Write([]byte("\n")); err != nil {
				return err
			}
			first = false
		}
		if err := k.render(f, w, nil); err != nil {
			return err
		}
		if _, err := w.Write([]byte(":")); err != nil {
			return err
		}
		if err := v.render(f, w, nil); err != nil {
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
