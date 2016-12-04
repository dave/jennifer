//go:generate genjen
package jen

import (
	"bytes"
	"context"
	"fmt"
	"go/format"
	"io"
	"os"
	"strconv"
)

type Code interface {
	Render(ctx context.Context, w io.Writer) error
}

func NewFile() *StatementList {
	return new(StatementList)
}

type StatementList []Code

type Statement []Code

func (g *Statement) Insert(code ...Code) *Statement {
	*g = append(*g, code...)
	return g
}

func (f *StatementList) Insert(code ...Code) *Statement {
	g := new(Statement)
	g.Insert(code...)
	*f = append(*f, g)
	return g
}

func WriteFile(ctx context.Context, l *StatementList, filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := Render(ctx, l, f); err != nil {
		return err
	}
	return nil
}

func Render(ctx context.Context, l *StatementList, w io.Writer) error {
	body := &bytes.Buffer{}
	if err := l.Render(ctx, body); err != nil {
		return err
	}
	global := FromContext(ctx)
	source := &bytes.Buffer{}
	if _, err := fmt.Fprintf(source, "package %s\n\n", global.Name); err != nil {
		return err
	}
	if len(global.Imports) == 1 {
		for path, alias := range global.Imports {
			if _, err := fmt.Fprintf(source, "import %s %s\n\n", alias, strconv.Quote(path)); err != nil {
				return err
			}
		}
	} else if len(global.Imports) > 1 {
		if _, err := fmt.Fprint(source, "import (\n"); err != nil {
			return err
		}
		for path, alias := range global.Imports {
			if _, err := fmt.Fprintf(source, "%s %s\n", alias, strconv.Quote(path)); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprint(source, ")\n\n"); err != nil {
			return err
		}
	}
	if _, err := source.Write(body.Bytes()); err != nil {
		return err
	}
	formatted, err := format.Source(source.Bytes())
	if err != nil {
		return err
	}
	_, err = w.Write(formatted)
	if err != nil {
		return err
	}
	return nil
}

func (f StatementList) Render(ctx context.Context, w io.Writer) error {
	for i, code := range f {
		if i > 0 {
			if _, err := w.Write([]byte(";")); err != nil {
				return err
			}
		}
		if err := code.Render(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (g Statement) Render(ctx context.Context, w io.Writer) error {
	for _, code := range g {
		if err := code.Render(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
