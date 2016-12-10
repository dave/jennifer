package jen

// Return inserts the return keyword
func Return(c ...Code) *Group { return newStatement().Return(c...) }

// Return inserts the return keyword
func (g *Group) Return(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Return(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("return").List(c...)
}
