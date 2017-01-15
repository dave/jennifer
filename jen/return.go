package jen

// Return inserts the return keyword
func Return(c ...Code) *Statement {
	return newStatement().Return(c...)
}

// Return inserts the return keyword
func (g *Group) Return(c ...Code) *Statement {
	s := Return(c...)
	g.items = append(g.items, s)
	return s
}

// Return inserts the return keyword
func (s *Statement) Return(c ...Code) *Statement {
	return s.Id("return").List(c...)
}
