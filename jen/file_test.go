package jen

import (
	"fmt"
	"testing"
)

func TestGuessAlias(t *testing.T) {

	data := map[string]string{
		"A":        "a",
		"a":        "a",
		"a$":       "a",
		"a/b":      "b",
		"a/b/c":    "c",
		"a/b/c-d":  "d",
		"a/b/c-d/": "d",
		"a.b":      "a",
		"a/b.c":    "b",
		"a/b-c.d":  "c",
	}
	for path, expected := range data {
		if guessAlias(path) != expected {
			fmt.Printf("guessAlias test failed %s should return %s but got %s", path, expected, guessAlias(path))
			t.Fail()
		}
	}
}

func TestValidAlias(t *testing.T) {
	data := map[string]bool{
		"a":  true,
		"b":  false,
		"go": false,
	}
	f := NewFile("test")
	f.register("b")
	for alias, expected := range data {
		if f.isValidAlias(alias) != expected {
			fmt.Printf("isValidAlias test failed %s should return %t but got %t", alias, expected, f.isValidAlias(alias))
			t.Fail()
		}
	}
}
