package jen

// For inserts the for keyword
func For(c ...Code) *Group { return newStatement().For(c...) }

// For inserts the for keyword
func (g *Group) For(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := For(c...)
		g.items = append(g.items, s)
		return s
	}
	g.Id("for")
	s := Group{
		syntax: clauseSyntax,
		items:  c,
	}
	g.items = append(g.items, s)
	return g
}
