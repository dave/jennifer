package jen

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
)

type Group struct {
	syntax syntaxType
	items  []Code
}

func (g *Group) isNull() bool {
	if g == nil {
		return true
	}
	i := syntaxInfo[g.syntax]
	if i.open != "" || i.close != "" {
		return false
	}
	for _, c := range g.items {
		if !c.isNull() {
			return false
		}
	}
	return true
}

func (g *Group) render(f *File, w io.Writer) error {
	i := syntaxInfo[g.syntax]
	if i.open != "" {
		if _, err := w.Write([]byte(i.open)); err != nil {
			return err
		}
	}
	first := true
	for _, code := range g.items {
		if code == nil || code.isNull() {
			// Null() token produces no output but also
			// no separator. Empty() token products no
			// output but adds a separator.
			continue
		}
		if !first && i.seperator != "" {
			if _, err := w.Write([]byte(i.seperator)); err != nil {
				return err
			}
		}
		if err := code.render(f, w); err != nil {
			return err
		}
		first = false
	}
	if i.close != "" {
		if _, err := w.Write([]byte(i.close)); err != nil {
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
