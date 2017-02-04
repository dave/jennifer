package jen

import (
	"fmt"
	"io"
	"sort"
	"strconv"
)

// Tag adds a struct tag
func Tag(items map[string]string) *Statement {
	return newStatement().Tag(items)
}

// Tag adds a struct tag
func (g *Group) Tag(items map[string]string) *Statement {
	s := Tag(items)
	g.items = append(g.items, s)
	return s
}

// Tag adds a struct tag
func (s *Statement) Tag(items map[string]string) *Statement {
	c := tag{
		items: items,
	}
	*s = append(*s, c)
	return s
}

type tag struct {
	items map[string]string
}

func (t tag) isNull(f *File) bool {
	return len(t.items) == 0
}

func (t tag) render(f *File, w io.Writer) error {

	if t.isNull(f) {
		return nil
	}

	var s string

	var sorted []string
	for k := range t.items {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)

	for _, k := range sorted {
		v := t.items[k]
		if len(s) > 0 {
			s += " "
		}
		s += fmt.Sprintf(`%s:"%s"`, k, v)
	}

	if strconv.CanBackquote(s) {
		s = "`" + s + "`"
	} else {
		s = strconv.Quote(s)
	}

	if _, err := w.Write([]byte(s)); err != nil {
		return err
	}

	return nil
}
