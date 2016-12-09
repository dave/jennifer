package jen

import (
	"context"
	"fmt"
	"io"
)

func Comment(comments ...string) *Statement {
	s := new(Statement)
	return s.Comment(comments...)
}

func (l *StatementList) Comment(comments ...string) *Statement {
	s := Comment(comments...)
	*l = append(*l, s)
	return s
}

func (s *Statement) Comment(comments ...string) *Statement {
	c := comment{
		Statement: s,
		comments:  comments,
	}
	*s = append(*s, c)
	return s
}

func Commentf(format string, a ...interface{}) *Statement {
	s := new(Statement)
	return s.Comment(fmt.Sprintf(format, a...))
}

func (l *StatementList) Commentf(format string, a ...interface{}) *Statement {
	s := Comment(fmt.Sprintf(format, a...))
	*l = append(*l, s)
	return s
}

func (s *Statement) Commentf(format string, a ...interface{}) *Statement {
	c := comment{
		Statement: s,
		comments:  []string{fmt.Sprintf(format, a...)},
	}
	*s = append(*s, c)
	return s
}

type comment struct {
	*Statement
	comments []string
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
