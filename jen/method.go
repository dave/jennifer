package jen

func Method(name string, code ...Code) *Statement {
	s := new(Statement)
	return s.Method(name, code...)
}

func (l *StatementList) Method(name string, code ...Code) *Statement {
	s := Method(name, code...)
	*l = append(*l, s)
	return s
}

func (s *Statement) Method(name string, code ...Code) *Statement {
	g := group{
		Statement: s,
		code:      code,
		open:      "(",
		close:     ")",
		seperator: ",",
	}
	*s = append(*s, Sel(), Id(name), g)
	return s
}
