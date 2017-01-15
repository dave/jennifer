package jen

import (
	"fmt"
	"io"
	"strings"
)

func Comment(str string) *Statement {
	return newStatement().Comment(str)
}

func (g *Group) Comment(str string) *Statement {
	s := Comment(str)
	g.items = append(g.items, s)
	return s
}

func (s *Statement) Comment(str string) *Statement {
	c := comment{
		comment: str,
	}
	s.items = append(s.items, c)
	return s
}

func Commentf(format string, a ...interface{}) *Statement {
	return newStatement().Commentf(format, a...)
}

func (g *Group) Commentf(format string, a ...interface{}) *Statement {
	s := Commentf(format, a...)
	g.items = append(g.items, s)
	return s
}

func (s *Statement) Commentf(format string, a ...interface{}) *Statement {
	c := comment{
		comment: fmt.Sprintf(format, a...),
	}
	s.items = append(s.items, c)
	return s
}

type comment struct {
	comment string
}

func (c comment) isNull() bool {
	return false
}

func (c comment) render(f *File, w io.Writer) error {
	if strings.Contains(c.comment, "\n") {
		if _, err := w.Write([]byte("/*\n")); err != nil {
			return err
		}
	} else {
		if _, err := w.Write([]byte("// ")); err != nil {
			return err
		}
	}

	if _, err := w.Write([]byte(c.comment)); err != nil {
		return err
	}
	if strings.Contains(c.comment, "\n") {
		if !strings.HasSuffix(c.comment, "\n") {
			if _, err := w.Write([]byte("\n")); err != nil {
				return err
			}
		}
		if _, err := w.Write([]byte("*/")); err != nil {
			return err
		}
	}
	return nil
}
