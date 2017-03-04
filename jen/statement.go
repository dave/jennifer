package jen

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
)

// Statement represents a simple list of code items. When rendered the items
// are separated by spaces.
type Statement []Code

func newStatement() *Statement {
	return &Statement{}
}

// Clone makes a copy of the Statement, so further tokens can be appended
// without affecting the original.
func (s *Statement) Clone() *Statement {
	return Add(s)
}

func (s *Statement) previous(c Code) Code {
	index := -1
	for i, item := range *s {
		if item == c {
			index = i
			break
		}
	}
	if index > 0 {
		return (*s)[index-1]
	}
	return nil
}

func (s *Statement) isNull(f *File) bool {
	if s == nil {
		return true
	}
	for _, c := range *s {
		if !c.isNull(f) {
			return false
		}
	}
	return true
}

func (s *Statement) render(f *File, w io.Writer, st *Statement, container *Group) error {
	first := true
	for _, code := range *s {
		if code == nil || code.isNull(f) {
			// Null() token produces no output but also
			// no separator. Empty() token products no
			// output but adds a separator.
			continue
		}
		if !first {
			if _, err := w.Write([]byte(" ")); err != nil {
				return err
			}
		}
		if err := code.render(f, w, s, container); err != nil {
			return err
		}
		first = false
	}
	return nil
}

// GoString renders the Statement for testing. Any error will cause a panic.
func (s *Statement) GoString() string {
	f := NewFile("")
	buf := &bytes.Buffer{}
	if err := s.render(f, buf, nil, nil); err != nil {
		panic(err)
	}
	b, err := format.Source(buf.Bytes())
	if err != nil {
		panic(fmt.Errorf("Error while formatting source: %s\nSource: %s", err, buf.String()))
	}
	return string(b)
}
