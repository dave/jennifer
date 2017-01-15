package jen

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
)

type Statement struct {
	items []Code
}

func newStatement() *Statement {
	return &Statement{}
}

func (s Statement) isNull() bool {
	for _, c := range s.items {
		if !c.isNull() {
			return false
		}
	}
	return true
}

func (s Statement) render(f *File, w io.Writer) error {
	first := true
	for _, code := range s.items {
		if code.isNull() {
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
