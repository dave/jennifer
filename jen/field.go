package jen

func Field(name string) *Statement {
	s := new(Statement)
	return s.Field(name)
}

func (l *StatementList) Field(name string) *Statement {
	s := Field(name)
	*l = append(*l, s)
	return s
}

func (s *Statement) Field(name string) *Statement {
	*s = append(*s, Sel(), Id(name))
	return s
}
