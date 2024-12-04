package jen

import (
	"bytes"
	"io"
	"sort"
)

// Dict renders as key/value pairs. Use with Values for map or composite
// literals.
type Dict map[Code]Code

// DictFunc executes a func(Dict) to generate the value. Use with Values for
// map or composite literals.
func DictFunc(f func(Dict)) Dict {
	d := Dict{}
	f(d)
	return d
}

func (d Dict) render(f *File, w io.Writer, s *Statement) error {
	lookup := make(map[string]kv, len(d))
	keys := make([]string, 0, len(d))

	buf := &bytes.Buffer{}
	for k, v := range d {
		if k.isNull(f) || v.isNull(f) {
			continue
		}
		if err := k.render(f, buf, nil); err != nil {
			return err
		}
		keys = append(keys, buf.String())
		lookup[buf.String()] = kv{k: k, v: v}
		buf.Reset()
	}

	// must order keys to ensure repeatable source
	sort.Strings(keys)

	ordered := make([]kv, 0, len(keys))
	for _, key := range keys {
		ordered = append(ordered, kv{k: lookup[key].k, v: lookup[key].v})
	}

	dict := &OrderedDict{items: ordered}
	return dict.render(f, w, s)
}

func (d Dict) isNull(f *File) bool {
	if d == nil || len(d) == 0 {
		return true
	}
	for k, v := range d {
		if !k.isNull(f) && !v.isNull(f) {
			// if any of the key/value pairs are both not null, the Dict is not
			// null
			return false
		}
	}
	return true
}
