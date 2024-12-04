package jen

import "io"

// OrderedDictFunc executes a func(*OrderedDict) to generate the value. Use OrderedDict.Add to append key/value pairs.
// Use with Values for map or composite literals.
func OrderedDictFunc(f func(*OrderedDict)) *OrderedDict {
	o := &OrderedDict{}
	f(o)
	return o
}

// OrderedDict renders as key/value pairs. Use with Values for map or composite
// literals.
type OrderedDict struct {
	items []kv
}

type kv struct {
	k Code
	v Code
}

// Add appends a key/value pairs.
func (o *OrderedDict) Add(key Code, value Code) *OrderedDict {
	o.items = append(o.items, kv{k: key, v: value})
	return o
}

func (o *OrderedDict) render(f *File, w io.Writer, _ *Statement) error {
	first := true
	for _, item := range o.items {
		if first && len(o.items) > 1 {
			if _, err := w.Write([]byte("\n")); err != nil {
				return err
			}
			first = false
		}
		if err := item.k.render(f, w, nil); err != nil {
			return err
		}
		if _, err := w.Write([]byte(":")); err != nil {
			return err
		}
		if err := item.v.render(f, w, nil); err != nil {
			return err
		}
		if len(o.items) > 1 {
			if _, err := w.Write([]byte(",\n")); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *OrderedDict) isNull(f *File) bool {
	if len(o.items) == 0 {
		return true
	}
	for _, item := range o.items {
		if !item.k.isNull(f) && !item.v.isNull(f) {
			// if any of the key/value pairs are both not null, the OrderedDict is not
			// null
			return false
		}
	}
	return true
}
