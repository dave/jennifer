package jen

import (
	"context"
	"io"
)

type block struct {
	*Statement
	code      []Code
	open      string
	close     string
	seperator string
}

func (d block) Render(ctx context.Context, w io.Writer) error {
	if d.open != "" {
		if _, err := w.Write([]byte(d.open)); err != nil {
			return err
		}
	}
	for i, code := range d.code {
		if i > 0 && d.seperator != "" {
			if _, err := w.Write([]byte(d.seperator)); err != nil {
				return err
			}
		}
		if err := code.Render(ctx, w); err != nil {
			return err
		}
	}
	if d.close != "" {
		if _, err := w.Write([]byte(d.close + " ")); err != nil {
			return err
		}
	}
	return nil
}
