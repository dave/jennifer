package jen

import (
	"context"
	"io"
)

type group struct {
	*Statement
	code      []Code
	open      string
	close     string
	seperator string
}

func (g group) IsNull() bool {
	for _, c := range g.code {
		if !c.IsNull() {
			return false
		}
	}
	return true
}

func (g group) Render(ctx context.Context, w io.Writer) error {
	if g.open != "" {
		if _, err := w.Write([]byte(g.open)); err != nil {
			return err
		}
	}
	first := true
	for _, code := range g.code {
		if code.IsNull() {
			// Null() token produces no output but also
			// no separator. Empty() token products no
			// output but adds a separator.
			continue
		}
		if !first && g.seperator != "" {
			if _, err := w.Write([]byte(g.seperator)); err != nil {
				return err
			}
		}
		if err := code.Render(ctx, w); err != nil {
			return err
		}
		first = false
	}
	if g.close != "" {
		if _, err := w.Write([]byte(g.close + " ")); err != nil {
			return err
		}
	}
	return nil
}
