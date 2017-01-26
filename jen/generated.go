package jen

/*
	This file is generated by genjen - do not edit!
*/

// Parens inserts parenthesis
func Parens(c Code) *Statement {
	return newStatement().Parens(c)
}

// Parens inserts parenthesis
func (g *Group) Parens(c Code) *Statement {
	s := Parens(c)
	g.items = append(g.items, s)
	return s
}

// Parens inserts parenthesis
func (s *Statement) Parens(c Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{c},
		open:      "(",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// Parens inserts parenthesis
func ParensFunc(f func(*Group)) *Statement {
	return newStatement().ParensFunc(f)
}

// Parens inserts parenthesis
func (g *Group) ParensFunc(f func(*Group)) *Statement {
	s := ParensFunc(f)
	g.items = append(g.items, s)
	return s
}

// Parens inserts parenthesis
func (s *Statement) ParensFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		open:      "(",
		separator: "",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// List inserts a comma separated list
func List(c ...Code) *Statement {
	return newStatement().List(c...)
}

// List inserts a comma separated list
func (g *Group) List(c ...Code) *Statement {
	s := List(c...)
	g.items = append(g.items, s)
	return s
}

// List inserts a comma separated list
func (s *Statement) List(c ...Code) *Statement {
	g := &Group{
		close:     "",
		items:     c,
		open:      "",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// List inserts a comma separated list
func ListFunc(f func(*Group)) *Statement {
	return newStatement().ListFunc(f)
}

// List inserts a comma separated list
func (g *Group) ListFunc(f func(*Group)) *Statement {
	s := ListFunc(f)
	g.items = append(g.items, s)
	return s
}

// List inserts a comma separated list
func (s *Statement) ListFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "",
		open:      "",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Values inserts curly braces containing a comma separated list
func Values(c ...Code) *Statement {
	return newStatement().Values(c...)
}

// Values inserts curly braces containing a comma separated list
func (g *Group) Values(c ...Code) *Statement {
	s := Values(c...)
	g.items = append(g.items, s)
	return s
}

// Values inserts curly braces containing a comma separated list
func (s *Statement) Values(c ...Code) *Statement {
	g := &Group{
		close:     "}",
		items:     c,
		open:      "{",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Values inserts curly braces containing a comma separated list
func ValuesFunc(f func(*Group)) *Statement {
	return newStatement().ValuesFunc(f)
}

// Values inserts curly braces containing a comma separated list
func (g *Group) ValuesFunc(f func(*Group)) *Statement {
	s := ValuesFunc(f)
	g.items = append(g.items, s)
	return s
}

// Values inserts curly braces containing a comma separated list
func (s *Statement) ValuesFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "}",
		open:      "{",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Slice inserts curly braces containing a comma separated list
func Slice(c ...Code) *Statement {
	return newStatement().Slice(c...)
}

// Slice inserts curly braces containing a comma separated list
func (g *Group) Slice(c ...Code) *Statement {
	s := Slice(c...)
	g.items = append(g.items, s)
	return s
}

// Slice inserts curly braces containing a comma separated list
func (s *Statement) Slice(c ...Code) *Statement {
	g := &Group{
		close:     "}",
		items:     c,
		open:      "{",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Slice inserts curly braces containing a comma separated list
func SliceFunc(f func(*Group)) *Statement {
	return newStatement().SliceFunc(f)
}

// Slice inserts curly braces containing a comma separated list
func (g *Group) SliceFunc(f func(*Group)) *Statement {
	s := SliceFunc(f)
	g.items = append(g.items, s)
	return s
}

// Slice inserts curly braces containing a comma separated list
func (s *Statement) SliceFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "}",
		open:      "{",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Index inserts square brackets containing a colon separated list
func Index(c ...Code) *Statement {
	return newStatement().Index(c...)
}

// Index inserts square brackets containing a colon separated list
func (g *Group) Index(c ...Code) *Statement {
	s := Index(c...)
	g.items = append(g.items, s)
	return s
}

// Index inserts square brackets containing a colon separated list
func (s *Statement) Index(c ...Code) *Statement {
	g := &Group{
		close:     "]",
		items:     c,
		open:      "[",
		separator: ":",
	}
	*s = append(*s, g)
	return s
}

// Index inserts square brackets containing a colon separated list
func IndexFunc(f func(*Group)) *Statement {
	return newStatement().IndexFunc(f)
}

// Index inserts square brackets containing a colon separated list
func (g *Group) IndexFunc(f func(*Group)) *Statement {
	s := IndexFunc(f)
	g.items = append(g.items, s)
	return s
}

// Index inserts square brackets containing a colon separated list
func (s *Statement) IndexFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "]",
		open:      "[",
		separator: ":",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Block inserts curly braces containing a statements list
func Block(c ...Code) *Statement {
	return newStatement().Block(c...)
}

// Block inserts curly braces containing a statements list
func (g *Group) Block(c ...Code) *Statement {
	s := Block(c...)
	g.items = append(g.items, s)
	return s
}

// Block inserts curly braces containing a statements list
func (s *Statement) Block(c ...Code) *Statement {
	g := &Group{
		close:     "}",
		items:     c,
		open:      "{\n",
		separator: "\n",
	}
	*s = append(*s, g)
	return s
}

// Block inserts curly braces containing a statements list
func BlockFunc(f func(*Group)) *Statement {
	return newStatement().BlockFunc(f)
}

// Block inserts curly braces containing a statements list
func (g *Group) BlockFunc(f func(*Group)) *Statement {
	s := BlockFunc(f)
	g.items = append(g.items, s)
	return s
}

// Block inserts curly braces containing a statements list
func (s *Statement) BlockFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "}",
		open:      "{\n",
		separator: "\n",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Call inserts parenthesis containing a comma separated list
func Call(c ...Code) *Statement {
	return newStatement().Call(c...)
}

// Call inserts parenthesis containing a comma separated list
func (g *Group) Call(c ...Code) *Statement {
	s := Call(c...)
	g.items = append(g.items, s)
	return s
}

// Call inserts parenthesis containing a comma separated list
func (s *Statement) Call(c ...Code) *Statement {
	g := &Group{
		close:     ")",
		items:     c,
		open:      "(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Call inserts parenthesis containing a comma separated list
func CallFunc(f func(*Group)) *Statement {
	return newStatement().CallFunc(f)
}

// Call inserts parenthesis containing a comma separated list
func (g *Group) CallFunc(f func(*Group)) *Statement {
	s := CallFunc(f)
	g.items = append(g.items, s)
	return s
}

// Call inserts parenthesis containing a comma separated list
func (s *Statement) CallFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		open:      "(",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Params inserts parenthesis containing a comma separated list
func Params(c ...Code) *Statement {
	return newStatement().Params(c...)
}

// Params inserts parenthesis containing a comma separated list
func (g *Group) Params(c ...Code) *Statement {
	s := Params(c...)
	g.items = append(g.items, s)
	return s
}

// Params inserts parenthesis containing a comma separated list
func (s *Statement) Params(c ...Code) *Statement {
	g := &Group{
		close:     ")",
		items:     c,
		open:      "(",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Params inserts parenthesis containing a comma separated list
func ParamsFunc(f func(*Group)) *Statement {
	return newStatement().ParamsFunc(f)
}

// Params inserts parenthesis containing a comma separated list
func (g *Group) ParamsFunc(f func(*Group)) *Statement {
	s := ParamsFunc(f)
	g.items = append(g.items, s)
	return s
}

// Params inserts parenthesis containing a comma separated list
func (s *Statement) ParamsFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		open:      "(",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Decls inserts parenthesis containing a statement list
func Decls(c ...Code) *Statement {
	return newStatement().Decls(c...)
}

// Decls inserts parenthesis containing a statement list
func (g *Group) Decls(c ...Code) *Statement {
	s := Decls(c...)
	g.items = append(g.items, s)
	return s
}

// Decls inserts parenthesis containing a statement list
func (s *Statement) Decls(c ...Code) *Statement {
	g := &Group{
		close:     ")",
		items:     c,
		open:      "(",
		separator: ";",
	}
	*s = append(*s, g)
	return s
}

// Decls inserts parenthesis containing a statement list
func DeclsFunc(f func(*Group)) *Statement {
	return newStatement().DeclsFunc(f)
}

// Decls inserts parenthesis containing a statement list
func (g *Group) DeclsFunc(f func(*Group)) *Statement {
	s := DeclsFunc(f)
	g.items = append(g.items, s)
	return s
}

// Decls inserts parenthesis containing a statement list
func (s *Statement) DeclsFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		open:      "(",
		separator: ";",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// CaseBlock inserts a statement list preceeded by a colon
func CaseBlock(c ...Code) *Statement {
	return newStatement().CaseBlock(c...)
}

// CaseBlock inserts a statement list preceeded by a colon
func (g *Group) CaseBlock(c ...Code) *Statement {
	s := CaseBlock(c...)
	g.items = append(g.items, s)
	return s
}

// CaseBlock inserts a statement list preceeded by a colon
func (s *Statement) CaseBlock(c ...Code) *Statement {
	g := &Group{
		close:     "",
		items:     c,
		open:      ":\n",
		separator: "\n",
	}
	*s = append(*s, g)
	return s
}

// CaseBlock inserts a statement list preceeded by a colon
func CaseBlockFunc(f func(*Group)) *Statement {
	return newStatement().CaseBlockFunc(f)
}

// CaseBlock inserts a statement list preceeded by a colon
func (g *Group) CaseBlockFunc(f func(*Group)) *Statement {
	s := CaseBlockFunc(f)
	g.items = append(g.items, s)
	return s
}

// CaseBlock inserts a statement list preceeded by a colon
func (s *Statement) CaseBlockFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "",
		open:      ":\n",
		separator: "\n",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Assert inserts a type assertion
func Assert(c Code) *Statement {
	return newStatement().Assert(c)
}

// Assert inserts a type assertion
func (g *Group) Assert(c Code) *Statement {
	s := Assert(c)
	g.items = append(g.items, s)
	return s
}

// Assert inserts a type assertion
func (s *Statement) Assert(c Code) *Statement {
	g := &Group{
		close:     ")",
		items:     []Code{c},
		open:      ".(",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// Assert inserts a type assertion
func AssertFunc(f func(*Group)) *Statement {
	return newStatement().AssertFunc(f)
}

// Assert inserts a type assertion
func (g *Group) AssertFunc(f func(*Group)) *Statement {
	s := AssertFunc(f)
	g.items = append(g.items, s)
	return s
}

// Assert inserts a type assertion
func (s *Statement) AssertFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     ")",
		open:      ".(",
		separator: "",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Map inserts the map keyword, followed by square brackets
func Map(c Code) *Statement {
	return newStatement().Map(c)
}

// Map inserts the map keyword, followed by square brackets
func (g *Group) Map(c Code) *Statement {
	s := Map(c)
	g.items = append(g.items, s)
	return s
}

// Map inserts the map keyword, followed by square brackets
func (s *Statement) Map(c Code) *Statement {
	g := &Group{
		close:     "]",
		items:     []Code{c},
		open:      "map[",
		separator: "",
	}
	*s = append(*s, g)
	return s
}

// Map inserts the map keyword, followed by square brackets
func MapFunc(f func(*Group)) *Statement {
	return newStatement().MapFunc(f)
}

// Map inserts the map keyword, followed by square brackets
func (g *Group) MapFunc(f func(*Group)) *Statement {
	s := MapFunc(f)
	g.items = append(g.items, s)
	return s
}

// Map inserts the map keyword, followed by square brackets
func (s *Statement) MapFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "]",
		open:      "map[",
		separator: "",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// If inserts the if keyword, followed by a semicolon separated list
func If(c ...Code) *Statement {
	return newStatement().If(c...)
}

// If inserts the if keyword, followed by a semicolon separated list
func (g *Group) If(c ...Code) *Statement {
	s := If(c...)
	g.items = append(g.items, s)
	return s
}

// If inserts the if keyword, followed by a semicolon separated list
func (s *Statement) If(c ...Code) *Statement {
	g := &Group{
		close:     "",
		items:     c,
		open:      "if ",
		separator: ";",
	}
	*s = append(*s, g)
	return s
}

// If inserts the if keyword, followed by a semicolon separated list
func IfFunc(f func(*Group)) *Statement {
	return newStatement().IfFunc(f)
}

// If inserts the if keyword, followed by a semicolon separated list
func (g *Group) IfFunc(f func(*Group)) *Statement {
	s := IfFunc(f)
	g.items = append(g.items, s)
	return s
}

// If inserts the if keyword, followed by a semicolon separated list
func (s *Statement) IfFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "",
		open:      "if ",
		separator: ";",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Return inserts the return keyword, followed by a comma separated list
func Return(c ...Code) *Statement {
	return newStatement().Return(c...)
}

// Return inserts the return keyword, followed by a comma separated list
func (g *Group) Return(c ...Code) *Statement {
	s := Return(c...)
	g.items = append(g.items, s)
	return s
}

// Return inserts the return keyword, followed by a comma separated list
func (s *Statement) Return(c ...Code) *Statement {
	g := &Group{
		close:     "",
		items:     c,
		open:      "return ",
		separator: ",",
	}
	*s = append(*s, g)
	return s
}

// Return inserts the return keyword, followed by a comma separated list
func ReturnFunc(f func(*Group)) *Statement {
	return newStatement().ReturnFunc(f)
}

// Return inserts the return keyword, followed by a comma separated list
func (g *Group) ReturnFunc(f func(*Group)) *Statement {
	s := ReturnFunc(f)
	g.items = append(g.items, s)
	return s
}

// Return inserts the return keyword, followed by a comma separated list
func (s *Statement) ReturnFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "",
		open:      "return ",
		separator: ",",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// For inserts the for keyword, followed by a semicolon separated list
func For(c ...Code) *Statement {
	return newStatement().For(c...)
}

// For inserts the for keyword, followed by a semicolon separated list
func (g *Group) For(c ...Code) *Statement {
	s := For(c...)
	g.items = append(g.items, s)
	return s
}

// For inserts the for keyword, followed by a semicolon separated list
func (s *Statement) For(c ...Code) *Statement {
	g := &Group{
		close:     "",
		items:     c,
		open:      "for ",
		separator: ";",
	}
	*s = append(*s, g)
	return s
}

// For inserts the for keyword, followed by a semicolon separated list
func ForFunc(f func(*Group)) *Statement {
	return newStatement().ForFunc(f)
}

// For inserts the for keyword, followed by a semicolon separated list
func (g *Group) ForFunc(f func(*Group)) *Statement {
	s := ForFunc(f)
	g.items = append(g.items, s)
	return s
}

// For inserts the for keyword, followed by a semicolon separated list
func (s *Statement) ForFunc(f func(*Group)) *Statement {
	g := &Group{
		close:     "",
		open:      "for ",
		separator: ";",
	}
	f(g)
	*s = append(*s, g)
	return s
}

// Bool inserts the bool identifier
func Bool() *Statement {
	return newStatement().Bool()
}

// Bool inserts the bool identifier
func (g *Group) Bool() *Statement {
	s := Bool()
	g.items = append(g.items, s)
	return s
}

// Bool inserts the bool identifier
func (s *Statement) Bool() *Statement {
	t := token{
		content: "bool",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Byte inserts the byte identifier
func Byte() *Statement {
	return newStatement().Byte()
}

// Byte inserts the byte identifier
func (g *Group) Byte() *Statement {
	s := Byte()
	g.items = append(g.items, s)
	return s
}

// Byte inserts the byte identifier
func (s *Statement) Byte() *Statement {
	t := token{
		content: "byte",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Complex64 inserts the complex64 identifier
func Complex64() *Statement {
	return newStatement().Complex64()
}

// Complex64 inserts the complex64 identifier
func (g *Group) Complex64() *Statement {
	s := Complex64()
	g.items = append(g.items, s)
	return s
}

// Complex64 inserts the complex64 identifier
func (s *Statement) Complex64() *Statement {
	t := token{
		content: "complex64",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Complex128 inserts the complex128 identifier
func Complex128() *Statement {
	return newStatement().Complex128()
}

// Complex128 inserts the complex128 identifier
func (g *Group) Complex128() *Statement {
	s := Complex128()
	g.items = append(g.items, s)
	return s
}

// Complex128 inserts the complex128 identifier
func (s *Statement) Complex128() *Statement {
	t := token{
		content: "complex128",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Error inserts the error identifier
func Error() *Statement {
	return newStatement().Error()
}

// Error inserts the error identifier
func (g *Group) Error() *Statement {
	s := Error()
	g.items = append(g.items, s)
	return s
}

// Error inserts the error identifier
func (s *Statement) Error() *Statement {
	t := token{
		content: "error",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Float32 inserts the float32 identifier
func Float32() *Statement {
	return newStatement().Float32()
}

// Float32 inserts the float32 identifier
func (g *Group) Float32() *Statement {
	s := Float32()
	g.items = append(g.items, s)
	return s
}

// Float32 inserts the float32 identifier
func (s *Statement) Float32() *Statement {
	t := token{
		content: "float32",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Float64 inserts the float64 identifier
func Float64() *Statement {
	return newStatement().Float64()
}

// Float64 inserts the float64 identifier
func (g *Group) Float64() *Statement {
	s := Float64()
	g.items = append(g.items, s)
	return s
}

// Float64 inserts the float64 identifier
func (s *Statement) Float64() *Statement {
	t := token{
		content: "float64",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Int inserts the int identifier
func Int() *Statement {
	return newStatement().Int()
}

// Int inserts the int identifier
func (g *Group) Int() *Statement {
	s := Int()
	g.items = append(g.items, s)
	return s
}

// Int inserts the int identifier
func (s *Statement) Int() *Statement {
	t := token{
		content: "int",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Int8 inserts the int8 identifier
func Int8() *Statement {
	return newStatement().Int8()
}

// Int8 inserts the int8 identifier
func (g *Group) Int8() *Statement {
	s := Int8()
	g.items = append(g.items, s)
	return s
}

// Int8 inserts the int8 identifier
func (s *Statement) Int8() *Statement {
	t := token{
		content: "int8",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Int16 inserts the int16 identifier
func Int16() *Statement {
	return newStatement().Int16()
}

// Int16 inserts the int16 identifier
func (g *Group) Int16() *Statement {
	s := Int16()
	g.items = append(g.items, s)
	return s
}

// Int16 inserts the int16 identifier
func (s *Statement) Int16() *Statement {
	t := token{
		content: "int16",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Int32 inserts the int32 identifier
func Int32() *Statement {
	return newStatement().Int32()
}

// Int32 inserts the int32 identifier
func (g *Group) Int32() *Statement {
	s := Int32()
	g.items = append(g.items, s)
	return s
}

// Int32 inserts the int32 identifier
func (s *Statement) Int32() *Statement {
	t := token{
		content: "int32",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Int64 inserts the int64 identifier
func Int64() *Statement {
	return newStatement().Int64()
}

// Int64 inserts the int64 identifier
func (g *Group) Int64() *Statement {
	s := Int64()
	g.items = append(g.items, s)
	return s
}

// Int64 inserts the int64 identifier
func (s *Statement) Int64() *Statement {
	t := token{
		content: "int64",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Rune inserts the rune identifier
func Rune() *Statement {
	return newStatement().Rune()
}

// Rune inserts the rune identifier
func (g *Group) Rune() *Statement {
	s := Rune()
	g.items = append(g.items, s)
	return s
}

// Rune inserts the rune identifier
func (s *Statement) Rune() *Statement {
	t := token{
		content: "rune",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// String inserts the string identifier
func String() *Statement {
	return newStatement().String()
}

// String inserts the string identifier
func (g *Group) String() *Statement {
	s := String()
	g.items = append(g.items, s)
	return s
}

// String inserts the string identifier
func (s *Statement) String() *Statement {
	t := token{
		content: "string",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Uint inserts the uint identifier
func Uint() *Statement {
	return newStatement().Uint()
}

// Uint inserts the uint identifier
func (g *Group) Uint() *Statement {
	s := Uint()
	g.items = append(g.items, s)
	return s
}

// Uint inserts the uint identifier
func (s *Statement) Uint() *Statement {
	t := token{
		content: "uint",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Uint8 inserts the uint8 identifier
func Uint8() *Statement {
	return newStatement().Uint8()
}

// Uint8 inserts the uint8 identifier
func (g *Group) Uint8() *Statement {
	s := Uint8()
	g.items = append(g.items, s)
	return s
}

// Uint8 inserts the uint8 identifier
func (s *Statement) Uint8() *Statement {
	t := token{
		content: "uint8",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Uint16 inserts the uint16 identifier
func Uint16() *Statement {
	return newStatement().Uint16()
}

// Uint16 inserts the uint16 identifier
func (g *Group) Uint16() *Statement {
	s := Uint16()
	g.items = append(g.items, s)
	return s
}

// Uint16 inserts the uint16 identifier
func (s *Statement) Uint16() *Statement {
	t := token{
		content: "uint16",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Uint32 inserts the uint32 identifier
func Uint32() *Statement {
	return newStatement().Uint32()
}

// Uint32 inserts the uint32 identifier
func (g *Group) Uint32() *Statement {
	s := Uint32()
	g.items = append(g.items, s)
	return s
}

// Uint32 inserts the uint32 identifier
func (s *Statement) Uint32() *Statement {
	t := token{
		content: "uint32",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Uint64 inserts the uint64 identifier
func Uint64() *Statement {
	return newStatement().Uint64()
}

// Uint64 inserts the uint64 identifier
func (g *Group) Uint64() *Statement {
	s := Uint64()
	g.items = append(g.items, s)
	return s
}

// Uint64 inserts the uint64 identifier
func (s *Statement) Uint64() *Statement {
	t := token{
		content: "uint64",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Uintptr inserts the uintptr identifier
func Uintptr() *Statement {
	return newStatement().Uintptr()
}

// Uintptr inserts the uintptr identifier
func (g *Group) Uintptr() *Statement {
	s := Uintptr()
	g.items = append(g.items, s)
	return s
}

// Uintptr inserts the uintptr identifier
func (s *Statement) Uintptr() *Statement {
	t := token{
		content: "uintptr",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// True inserts the true identifier
func True() *Statement {
	return newStatement().True()
}

// True inserts the true identifier
func (g *Group) True() *Statement {
	s := True()
	g.items = append(g.items, s)
	return s
}

// True inserts the true identifier
func (s *Statement) True() *Statement {
	t := token{
		content: "true",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// False inserts the false identifier
func False() *Statement {
	return newStatement().False()
}

// False inserts the false identifier
func (g *Group) False() *Statement {
	s := False()
	g.items = append(g.items, s)
	return s
}

// False inserts the false identifier
func (s *Statement) False() *Statement {
	t := token{
		content: "false",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Iota inserts the iota identifier
func Iota() *Statement {
	return newStatement().Iota()
}

// Iota inserts the iota identifier
func (g *Group) Iota() *Statement {
	s := Iota()
	g.items = append(g.items, s)
	return s
}

// Iota inserts the iota identifier
func (s *Statement) Iota() *Statement {
	t := token{
		content: "iota",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Nil inserts the nil identifier
func Nil() *Statement {
	return newStatement().Nil()
}

// Nil inserts the nil identifier
func (g *Group) Nil() *Statement {
	s := Nil()
	g.items = append(g.items, s)
	return s
}

// Nil inserts the nil identifier
func (s *Statement) Nil() *Statement {
	t := token{
		content: "nil",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Err inserts the err identifier
func Err() *Statement {
	return newStatement().Err()
}

// Err inserts the err identifier
func (g *Group) Err() *Statement {
	s := Err()
	g.items = append(g.items, s)
	return s
}

// Err inserts the err identifier
func (s *Statement) Err() *Statement {
	t := token{
		content: "err",
		typ:     identifierToken,
	}
	*s = append(*s, t)
	return s
}

// Break inserts the break keyword
func Break() *Statement {
	return newStatement().Break()
}

// Break inserts the break keyword
func (g *Group) Break() *Statement {
	s := Break()
	g.items = append(g.items, s)
	return s
}

// Break inserts the break keyword
func (s *Statement) Break() *Statement {
	t := token{
		content: "break",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Default inserts the default keyword
func Default() *Statement {
	return newStatement().Default()
}

// Default inserts the default keyword
func (g *Group) Default() *Statement {
	s := Default()
	g.items = append(g.items, s)
	return s
}

// Default inserts the default keyword
func (s *Statement) Default() *Statement {
	t := token{
		content: "default",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Func inserts the func keyword
func Func() *Statement {
	return newStatement().Func()
}

// Func inserts the func keyword
func (g *Group) Func() *Statement {
	s := Func()
	g.items = append(g.items, s)
	return s
}

// Func inserts the func keyword
func (s *Statement) Func() *Statement {
	t := token{
		content: "func",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Interface inserts the interface keyword
func Interface() *Statement {
	return newStatement().Interface()
}

// Interface inserts the interface keyword
func (g *Group) Interface() *Statement {
	s := Interface()
	g.items = append(g.items, s)
	return s
}

// Interface inserts the interface keyword
func (s *Statement) Interface() *Statement {
	t := token{
		content: "interface",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Select inserts the select keyword
func Select() *Statement {
	return newStatement().Select()
}

// Select inserts the select keyword
func (g *Group) Select() *Statement {
	s := Select()
	g.items = append(g.items, s)
	return s
}

// Select inserts the select keyword
func (s *Statement) Select() *Statement {
	t := token{
		content: "select",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Case inserts the case keyword
func Case() *Statement {
	return newStatement().Case()
}

// Case inserts the case keyword
func (g *Group) Case() *Statement {
	s := Case()
	g.items = append(g.items, s)
	return s
}

// Case inserts the case keyword
func (s *Statement) Case() *Statement {
	t := token{
		content: "case",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Defer inserts the defer keyword
func Defer() *Statement {
	return newStatement().Defer()
}

// Defer inserts the defer keyword
func (g *Group) Defer() *Statement {
	s := Defer()
	g.items = append(g.items, s)
	return s
}

// Defer inserts the defer keyword
func (s *Statement) Defer() *Statement {
	t := token{
		content: "defer",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Go inserts the go keyword
func Go() *Statement {
	return newStatement().Go()
}

// Go inserts the go keyword
func (g *Group) Go() *Statement {
	s := Go()
	g.items = append(g.items, s)
	return s
}

// Go inserts the go keyword
func (s *Statement) Go() *Statement {
	t := token{
		content: "go",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Struct inserts the struct keyword
func Struct() *Statement {
	return newStatement().Struct()
}

// Struct inserts the struct keyword
func (g *Group) Struct() *Statement {
	s := Struct()
	g.items = append(g.items, s)
	return s
}

// Struct inserts the struct keyword
func (s *Statement) Struct() *Statement {
	t := token{
		content: "struct",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Chan inserts the chan keyword
func Chan() *Statement {
	return newStatement().Chan()
}

// Chan inserts the chan keyword
func (g *Group) Chan() *Statement {
	s := Chan()
	g.items = append(g.items, s)
	return s
}

// Chan inserts the chan keyword
func (s *Statement) Chan() *Statement {
	t := token{
		content: "chan",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Else inserts the else keyword
func Else() *Statement {
	return newStatement().Else()
}

// Else inserts the else keyword
func (g *Group) Else() *Statement {
	s := Else()
	g.items = append(g.items, s)
	return s
}

// Else inserts the else keyword
func (s *Statement) Else() *Statement {
	t := token{
		content: "else",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Goto inserts the goto keyword
func Goto() *Statement {
	return newStatement().Goto()
}

// Goto inserts the goto keyword
func (g *Group) Goto() *Statement {
	s := Goto()
	g.items = append(g.items, s)
	return s
}

// Goto inserts the goto keyword
func (s *Statement) Goto() *Statement {
	t := token{
		content: "goto",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Switch inserts the switch keyword
func Switch() *Statement {
	return newStatement().Switch()
}

// Switch inserts the switch keyword
func (g *Group) Switch() *Statement {
	s := Switch()
	g.items = append(g.items, s)
	return s
}

// Switch inserts the switch keyword
func (s *Statement) Switch() *Statement {
	t := token{
		content: "switch",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Const inserts the const keyword
func Const() *Statement {
	return newStatement().Const()
}

// Const inserts the const keyword
func (g *Group) Const() *Statement {
	s := Const()
	g.items = append(g.items, s)
	return s
}

// Const inserts the const keyword
func (s *Statement) Const() *Statement {
	t := token{
		content: "const",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Fallthrough inserts the fallthrough keyword
func Fallthrough() *Statement {
	return newStatement().Fallthrough()
}

// Fallthrough inserts the fallthrough keyword
func (g *Group) Fallthrough() *Statement {
	s := Fallthrough()
	g.items = append(g.items, s)
	return s
}

// Fallthrough inserts the fallthrough keyword
func (s *Statement) Fallthrough() *Statement {
	t := token{
		content: "fallthrough",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Range inserts the range keyword
func Range() *Statement {
	return newStatement().Range()
}

// Range inserts the range keyword
func (g *Group) Range() *Statement {
	s := Range()
	g.items = append(g.items, s)
	return s
}

// Range inserts the range keyword
func (s *Statement) Range() *Statement {
	t := token{
		content: "range",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Type inserts the type keyword
func Type() *Statement {
	return newStatement().Type()
}

// Type inserts the type keyword
func (g *Group) Type() *Statement {
	s := Type()
	g.items = append(g.items, s)
	return s
}

// Type inserts the type keyword
func (s *Statement) Type() *Statement {
	t := token{
		content: "type",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Continue inserts the continue keyword
func Continue() *Statement {
	return newStatement().Continue()
}

// Continue inserts the continue keyword
func (g *Group) Continue() *Statement {
	s := Continue()
	g.items = append(g.items, s)
	return s
}

// Continue inserts the continue keyword
func (s *Statement) Continue() *Statement {
	t := token{
		content: "continue",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Var inserts the var keyword
func Var() *Statement {
	return newStatement().Var()
}

// Var inserts the var keyword
func (g *Group) Var() *Statement {
	s := Var()
	g.items = append(g.items, s)
	return s
}

// Var inserts the var keyword
func (s *Statement) Var() *Statement {
	t := token{
		content: "var",
		typ:     keywordToken,
	}
	*s = append(*s, t)
	return s
}

// Append inserts the built in function append
func Append(c ...Code) *Statement {
	return newStatement().Append(c...)
}

// Append inserts the built in function append
func (g *Group) Append(c ...Code) *Statement {
	s := Append(c...)
	g.items = append(g.items, s)
	return s
}

// Append inserts the built in function append
func (s *Statement) Append(c ...Code) *Statement {
	return s.Id("append").Call(c...)
}

// Cap inserts the built in function cap
func Cap(c ...Code) *Statement {
	return newStatement().Cap(c...)
}

// Cap inserts the built in function cap
func (g *Group) Cap(c ...Code) *Statement {
	s := Cap(c...)
	g.items = append(g.items, s)
	return s
}

// Cap inserts the built in function cap
func (s *Statement) Cap(c ...Code) *Statement {
	return s.Id("cap").Call(c...)
}

// Close inserts the built in function close
func Close(c ...Code) *Statement {
	return newStatement().Close(c...)
}

// Close inserts the built in function close
func (g *Group) Close(c ...Code) *Statement {
	s := Close(c...)
	g.items = append(g.items, s)
	return s
}

// Close inserts the built in function close
func (s *Statement) Close(c ...Code) *Statement {
	return s.Id("close").Call(c...)
}

// Complex inserts the built in function complex
func Complex(c ...Code) *Statement {
	return newStatement().Complex(c...)
}

// Complex inserts the built in function complex
func (g *Group) Complex(c ...Code) *Statement {
	s := Complex(c...)
	g.items = append(g.items, s)
	return s
}

// Complex inserts the built in function complex
func (s *Statement) Complex(c ...Code) *Statement {
	return s.Id("complex").Call(c...)
}

// Copy inserts the built in function copy
func Copy(c ...Code) *Statement {
	return newStatement().Copy(c...)
}

// Copy inserts the built in function copy
func (g *Group) Copy(c ...Code) *Statement {
	s := Copy(c...)
	g.items = append(g.items, s)
	return s
}

// Copy inserts the built in function copy
func (s *Statement) Copy(c ...Code) *Statement {
	return s.Id("copy").Call(c...)
}

// Delete inserts the built in function delete
func Delete(c ...Code) *Statement {
	return newStatement().Delete(c...)
}

// Delete inserts the built in function delete
func (g *Group) Delete(c ...Code) *Statement {
	s := Delete(c...)
	g.items = append(g.items, s)
	return s
}

// Delete inserts the built in function delete
func (s *Statement) Delete(c ...Code) *Statement {
	return s.Id("delete").Call(c...)
}

// Imag inserts the built in function imag
func Imag(c ...Code) *Statement {
	return newStatement().Imag(c...)
}

// Imag inserts the built in function imag
func (g *Group) Imag(c ...Code) *Statement {
	s := Imag(c...)
	g.items = append(g.items, s)
	return s
}

// Imag inserts the built in function imag
func (s *Statement) Imag(c ...Code) *Statement {
	return s.Id("imag").Call(c...)
}

// Len inserts the built in function len
func Len(c ...Code) *Statement {
	return newStatement().Len(c...)
}

// Len inserts the built in function len
func (g *Group) Len(c ...Code) *Statement {
	s := Len(c...)
	g.items = append(g.items, s)
	return s
}

// Len inserts the built in function len
func (s *Statement) Len(c ...Code) *Statement {
	return s.Id("len").Call(c...)
}

// Make inserts the built in function make
func Make(c ...Code) *Statement {
	return newStatement().Make(c...)
}

// Make inserts the built in function make
func (g *Group) Make(c ...Code) *Statement {
	s := Make(c...)
	g.items = append(g.items, s)
	return s
}

// Make inserts the built in function make
func (s *Statement) Make(c ...Code) *Statement {
	return s.Id("make").Call(c...)
}

// New inserts the built in function new
func New(c ...Code) *Statement {
	return newStatement().New(c...)
}

// New inserts the built in function new
func (g *Group) New(c ...Code) *Statement {
	s := New(c...)
	g.items = append(g.items, s)
	return s
}

// New inserts the built in function new
func (s *Statement) New(c ...Code) *Statement {
	return s.Id("new").Call(c...)
}

// Panic inserts the built in function panic
func Panic(c ...Code) *Statement {
	return newStatement().Panic(c...)
}

// Panic inserts the built in function panic
func (g *Group) Panic(c ...Code) *Statement {
	s := Panic(c...)
	g.items = append(g.items, s)
	return s
}

// Panic inserts the built in function panic
func (s *Statement) Panic(c ...Code) *Statement {
	return s.Id("panic").Call(c...)
}

// Print inserts the built in function print
func Print(c ...Code) *Statement {
	return newStatement().Print(c...)
}

// Print inserts the built in function print
func (g *Group) Print(c ...Code) *Statement {
	s := Print(c...)
	g.items = append(g.items, s)
	return s
}

// Print inserts the built in function print
func (s *Statement) Print(c ...Code) *Statement {
	return s.Id("print").Call(c...)
}

// Println inserts the built in function println
func Println(c ...Code) *Statement {
	return newStatement().Println(c...)
}

// Println inserts the built in function println
func (g *Group) Println(c ...Code) *Statement {
	s := Println(c...)
	g.items = append(g.items, s)
	return s
}

// Println inserts the built in function println
func (s *Statement) Println(c ...Code) *Statement {
	return s.Id("println").Call(c...)
}

// Real inserts the built in function real
func Real(c ...Code) *Statement {
	return newStatement().Real(c...)
}

// Real inserts the built in function real
func (g *Group) Real(c ...Code) *Statement {
	s := Real(c...)
	g.items = append(g.items, s)
	return s
}

// Real inserts the built in function real
func (s *Statement) Real(c ...Code) *Statement {
	return s.Id("real").Call(c...)
}

// Recover inserts the built in function recover
func Recover(c ...Code) *Statement {
	return newStatement().Recover(c...)
}

// Recover inserts the built in function recover
func (g *Group) Recover(c ...Code) *Statement {
	s := Recover(c...)
	g.items = append(g.items, s)
	return s
}

// Recover inserts the built in function recover
func (s *Statement) Recover(c ...Code) *Statement {
	return s.Id("recover").Call(c...)
}
