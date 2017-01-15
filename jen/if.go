package jen

// For inserts the for keyword
func If(c ...Code) *Group { return newStatement().If(c...) }

// If inserts the if keyword
func (g *Group) If(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := If(c...)
		g.items = append(g.items, s)
		return s
	}
	g.Id("if")
	s := Group{
		syntax: clauseSyntax,
		items:  c,
	}
	g.items = append(g.items, s)
	return g
}
