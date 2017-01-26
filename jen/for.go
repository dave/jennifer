package jen

// For inserts the for keyword
func For(c ...Code) *Statement {
	return newStatement().For(c...)
}

// For inserts the for keyword
func (g *Group) For(c ...Code) *Statement {
	s := For(c...)
	g.items = append(g.items, s)
	return s
}

// For inserts the for keyword
func (s *Statement) For(c ...Code) *Statement {
	s.Id("for")
	g := &Group{
		syntax: clauseSyntax,
		items:  c,
	}
	s.items = append(s.items, g)
	return s
}
