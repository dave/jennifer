package jen

import (
	"fmt"
	"io"
	"strings"
)

func Comment(str string) *Group {
	return newStatement().Comment(str)
}

func (g *Group) Comment(str string) *Group {
	if startNewStatement(g.syntax) {
		s := Comment(str)
		g.items = append(g.items, s)
		return s
	}
	c := comment{
		Group:   g,
		comment: str,
	}
	g.items = append(g.items, c)
	return g
}

func Commentf(format string, a ...interface{}) *Group {
	return newStatement().Commentf(format, a...)
}

func (g *Group) Commentf(format string, a ...interface{}) *Group {
	if startNewStatement(g.syntax) {
		s := Commentf(format, a...)
		g.items = append(g.items, s)
		return s
	}
	c := comment{
		Group:   g,
		comment: fmt.Sprintf(format, a...),
	}
	g.items = append(g.items, c)
	return g
}

type comment struct {
	*Group
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
