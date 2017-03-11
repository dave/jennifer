package jen

// Lit renders a literal, using the format provided by the fmt package %#v
// verb.
func Lit(v interface{}) *Statement {
	return newStatement().Lit(v)
}

// Lit renders a literal, using the format provided by the fmt package %#v
// verb.
func (g *Group) Lit(v interface{}) *Statement {
	s := Lit(v)
	g.items = append(g.items, s)
	return s
}

// Lit renders a literal, using the format provided by the fmt package %#v
// verb.
func (s *Statement) Lit(v interface{}) *Statement {
	t := token{
		typ:     literalToken,
		content: v,
	}
	*s = append(*s, t)
	return s
}

// LitFunc renders a literal, using the format provided by the fmt package %#v
// verb. LitFunc generates the value to render by executing the provided
// function.
func LitFunc(f func() interface{}) *Statement {
	return newStatement().LitFunc(f)
}

// LitFunc renders a literal, using the format provided by the fmt package %#v
// verb. LitFunc generates the value to render by executing the provided
// function.
func (g *Group) LitFunc(f func() interface{}) *Statement {
	s := LitFunc(f)
	g.items = append(g.items, s)
	return s
}

// LitFunc renders a literal, using the format provided by the fmt package %#v
// verb. LitFunc generates the value to render by executing the provided
// function.
func (s *Statement) LitFunc(f func() interface{}) *Statement {
	t := token{
		typ:     literalToken,
		content: f(),
	}
	*s = append(*s, t)
	return s
}
