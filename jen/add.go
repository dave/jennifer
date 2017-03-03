package jen

// Add appends the provided items to the statement.
func Add(code ...Code) *Statement {
	return newStatement().Add(code...)
}

// Add appends the provided items to the statement.
func (g *Group) Add(code ...Code) *Statement {
	s := Add(code...)
	g.items = append(g.items, s)
	return s
}

// Add appends the provided items to the statement.
func (s *Statement) Add(code ...Code) *Statement {
	for _, item := range code {
		if st, ok := item.(*Statement); ok && st != nil {
			for _, inner := range *st {
				s.Add(inner)
			}
		} else {
			*s = append(*s, item)
		}
	}
	return s
}
