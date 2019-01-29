package jen_test

import (
	"bytes"
	"testing"

	. "github.com/dave/jennifer/jen"
)

func TestStatement_Render(t *testing.T) {
	file := NewFile("main")
	file.ImportAlias("fmt", "fmtalias")

	statement := file.Func().Id("main").Params().Block(
		Qual("fmt", "Println").Call(Lit("something")),
	)

	expect := `func main() {
	fmtalias.Println("something")
}`

	var got bytes.Buffer

	err := statement.RenderWithFile(&got, file)
	if err != nil {
		t.Fatal(err)
	}

	if got.String() != expect {
		t.Fatalf("Got: %v, expect: %v", got.String(), expect)
	}
}
