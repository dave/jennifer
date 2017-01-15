package jen

// Do creates a new statement and calls the provided function with it as a
// parameter
func Do(f func(*Statement)) *Statement {
	return newStatement().Do(f)
}

// Do creates a new statement in the group and calls the provided function with
// it as a parameter
func (g *Group) Do(f func(*Statement)) *Statement {
	s := Do(f)
	g.items = append(g.items, s)
	return s
}

// Do calls the provided function with the statement as a parameter
func (s *Statement) Do(f func(*Statement)) *Statement {
	f(s)
	return s
}
