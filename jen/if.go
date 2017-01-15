package jen

// For inserts the if keyword
func If(c ...Code) *Statement {
	return newStatement().If(c...)
}

func (g *Group) If(c ...Code) *Statement {
	s := If(c...)
	g.items = append(g.items, s)
	return s
}

// If inserts the if keyword
func (s *Statement) If(c ...Code) *Statement {
	s.Id("if")
	g := Group{
		syntax: clauseSyntax,
		items:  c,
	}
	s.items = append(s.items, g)
	return s
}
