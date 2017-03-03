// Package jen is a code generator for Go
package jen

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"sort"
	"strconv"
)

// Code represents an item of code that can be rendered.
type Code interface {
	render(f *File, w io.Writer, s *Statement) error
	isNull(f *File) bool
}

// Save renders the file and saves to the filename provided.
func (f *File) Save(filename string) error {
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filename, buf.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

// Render renders the file to the provided writer.
func (f *File) Render(w io.Writer) error {
	body := &bytes.Buffer{}
	if err := f.render(f, body, nil); err != nil {
		return err
	}
	source := &bytes.Buffer{}
	if f.comments != nil {
		for _, c := range f.comments {
			if err := Comment(c).render(f, source, nil); err != nil {
				return err
			}
			if _, err := fmt.Fprint(source, "\n"); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fprintf(source, "package %s\n\n", f.name); err != nil {
		return err
	}
	if err := f.renderImports(source); err != nil {
		return err
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

func (f *File) renderImports(source io.Writer) error {
	if len(f.imports) == 1 {
		for path, alias := range f.imports {
			if _, err := fmt.Fprintf(source, "import %s %s\n\n", alias, strconv.Quote(path)); err != nil {
				return err
			}
		}
	} else if len(f.imports) > 1 {
		if _, err := fmt.Fprint(source, "import (\n"); err != nil {
			return err
		}
		// We must sort the imports to ensure repeatable
		// source.
		paths := []string{}
		for path := range f.imports {
			paths = append(paths, path)
		}
		sort.Strings(paths)
		for _, path := range paths {
			if _, err := fmt.Fprintf(source, "%s %s\n", f.imports[path], strconv.Quote(path)); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprint(source, ")\n\n"); err != nil {
			return err
		}
	}
	return nil
}
