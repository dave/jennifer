//go:generate genjen2
package jen

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"go/format"
	"io"
	"os"
	"sort"
	"strconv"
)

type Code interface {
	render(ctx context.Context, w io.Writer) error
	isNull() bool
}

func NewFile(name string) *Group {
	return &Group{
		syntax: syntax{
			typ:  fileSyntax,
			name: name,
		},
	}
}

func NewFilePath(name, path string) *Group {
	return &Group{
		syntax: syntax{
			typ:  fileSyntax,
			name: name,
			path: path,
		},
	}
}

func newStatement(c ...Code) *Group {
	return &Group{
		syntax: syntax{
			typ: statementSyntax,
		},
		items: c,
	}
}

func startNewStatement(s syntax) bool {
	switch s.typ {
	case fileSyntax, blockSyntax:
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
	if g.syntax.typ != fileSyntax {
		return errors.New("Group passed to RenderFile must be a File.")
	}
	global := FromContext(ctx)
	global.Name = g.syntax.name
	global.Path = g.syntax.path
	body := &bytes.Buffer{}
	if err := g.render(ctx, body); err != nil {
		return err
	}
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
		// We must sort the imports to ensure repeatable
		// source.
		paths := []string{}
		for path := range global.Imports {
			paths = append(paths, path)
		}
		sort.Strings(paths)
		for _, path := range paths {
			if _, err := fmt.Fprintf(source, "%s %s\n", global.Imports[path], strconv.Quote(path)); err != nil {
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
