package jen

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
)

type Group struct {
	items     []Code
	open      string
	close     string
	separator string
}

func (g *Group) isNull(f *File) bool {
	if g == nil {
		return true
	}
	if g.open != "" || g.close != "" {
		return false
	}
	for _, c := range g.items {
		if !c.isNull(f) {
			return false
		}
	}
	return true
}

func (g *Group) render(f *File, w io.Writer) error {
	if g.open != "" {
		if _, err := w.Write([]byte(g.open)); err != nil {
			return err
		}
	}
	first := true
	for _, code := range g.items {
		if code == nil || code.isNull(f) {
			// Null() token produces no output but also
			// no separator. Empty() token products no
			// output but adds a separator.
			continue
		}
		if first && g.separator == "\n" {
			// For blocks seperated with new lines, we always insert a new line
			// before the first item (but only if there is an item).
			if _, err := w.Write([]byte("\n")); err != nil {
				return err
			}
		}
		if !first && g.separator != "" {
			if _, err := w.Write([]byte(g.separator)); err != nil {
				return err
			}
		}
		if err := code.render(f, w); err != nil {
			return err
		}
		first = false
	}
	if !first && g.separator == "\n" && g.close != "" {
		// For blocks seperated with new lines and with a closing token, we
		// always insert a new line after the last item (but only if there is
		// an item). This is to ensure that if the statement finishes with a
		// comment, the closing token is not commented out.
		// TODO: This seems really brittle.
		if _, err := w.Write([]byte("\n")); err != nil {
			return err
		}
	}
	if g.close != "" {
		if _, err := w.Write([]byte(g.close)); err != nil {
			return err
		}
	}
	return nil
}

func (g *Group) GoString() string {
	f := NewFile("")
	buf := &bytes.Buffer{}
	if err := g.render(f, buf); err != nil {
		panic(err)
	}
	b, err := format.Source(buf.Bytes())
	if err != nil {
		panic(fmt.Errorf("Error while formatting source: %s\nSource: %s", err, buf.String()))
	}
	return string(b)
}
