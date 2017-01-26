package jen

// Add creates a new statement and appends the provided code.
func Add(code ...Code) *Statement {
	return newStatement().Add(code...)
}

// Add creates a new statement in the group and appends the provided code.
func (g *Group) Add(code ...Code) *Statement {
	s := Add(code...)
	g.items = append(g.items, s)
	return s
}

// Add appends the provided code to the statement.
func (s *Statement) Add(code ...Code) *Statement {
	*s = append(*s, code...)
	return s
}
