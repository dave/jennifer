package jen

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
)

type Statement []Code

func newStatement() *Statement {
	return &Statement{}
}

func (s *Statement) Clone() *Statement {
	return &Statement{s}
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

func (s *Statement) render(f *File, w io.Writer) error {
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
		if err := code.render(f, w); err != nil {
			return err
		}
		first = false
	}
	return nil
}

func (s *Statement) GoString() string {
	f := NewFile("")
	buf := &bytes.Buffer{}
	if err := s.render(f, buf); err != nil {
		panic(err)
	}
	b, err := format.Source(buf.Bytes())
	if err != nil {
		panic(fmt.Errorf("Error while formatting source: %s\nSource: %s", err, buf.String()))
	}
	return string(b)
}
