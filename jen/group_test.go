package jen_test

import (
	"bytes"
	"testing"

	. "github.com/dave/jennifer/jen"
)

func TestGroup_Render(t *testing.T) {
	file := NewFile("main")
	file.ImportAlias("fmt", "fmtalias")

	var g *Group
	BlockFunc(func(group *Group) {
		g = group
	})

	g.Qual("fmt", "Errorf").Call(Lit("error"))

	expect := `{
	fmtalias.Errorf("error")
}`

	var got bytes.Buffer

	err := g.RenderWithFile(&got, file)
	if err != nil {
		t.Fatal(err)
	}

	if got.String() != expect {
		t.Fatalf("Got: %v, expect: %v", got.String(), expect)
	}
}

func TestGroup_GoStringUnsafe(t *testing.T) {
	tt := []struct {
		statement *Statement
		expect    string
	}{
		{Func(), `func`},
		{Map(String()).Int(), `map[string] int`},
		{Interface(), `interface{}`},
	}

	for _, tc := range tt {
		g := &Group{}
		g.Add(tc.statement)

		got, err := g.GoStringUnsafe()
		if err != nil {
			t.Fatal(err)
		}

		if got != tc.expect {
			t.Fatalf("Got: %v, expect: %v", got, tc.expect)
		}
	}
}
