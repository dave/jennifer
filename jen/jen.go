//go:generate genjen2
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
	IsNull() bool
}

func NewFile() *Group {
	return &Group{
		syntax: FileSyntax,
	}
}

func newStatement(c ...Code) *Group {
	return &Group{
		syntax: StatementSyntax,
		items:  c,
	}
}

func startNewStatement(s syntaxType) bool {
	switch s {
	case FileSyntax, BlockSyntax:
		return true
	}
	return false
}

func WriteFile(ctx context.Context, g *Group, filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := RenderFile(ctx, g, f); err != nil {
		return err
	}
	return nil
}

func RenderFile(ctx context.Context, g *Group, w io.Writer) error {
	body := &bytes.Buffer{}
	if err := g.Render(ctx, body); err != nil {
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
		return fmt.Errorf("Error %s while formatting source:\n%s", err, source.String())
	}
	_, err = w.Write(formatted)
	if err != nil {
		return err
	}
	return nil
}
