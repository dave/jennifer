package jen

import (
	"context"
	"fmt"
	"io"
)

func Comment(comments ...string) *Group {
	return newStatement().Comment(comments...)
}

func (g *Group) Comment(comments ...string) *Group {
	if startNewStatement(g.syntax) {
		s := Comment(comments...)
		g.items = append(g.items, s)
		return s
	}
	c := comment{
		Group:    g,
		comments: comments,
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
		Group:    g,
		comments: []string{fmt.Sprintf(format, a...)},
	}
	g.items = append(g.items, c)
	return g
}

type comment struct {
	*Group
	comments []string
}

func (c comment) IsNull() bool {
	return false
}

func (c comment) Render(ctx context.Context, w io.Writer) error {
	if len(c.comments) > 1 {
		if _, err := w.Write([]byte("/*\n")); err != nil {
			return err
		}
	} else {
		if _, err := w.Write([]byte("// ")); err != nil {
			return err
		}
	}
	for _, str := range c.comments {
		if _, err := w.Write([]byte(str)); err != nil {
			return err
		}
	}
	if len(c.comments) > 1 {
		if _, err := w.Write([]byte("\n*/")); err != nil {
			return err
		}
	}
	return nil
}
