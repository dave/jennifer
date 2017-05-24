package jen

import (
	"fmt"
	"testing"
)

func TestGuessAlias(t *testing.T) {

	data := map[string]string{
		"A":             "a",
		"a":             "a",
		"a$":            "a",
		"a/b":           "b",
		"a/b/c":         "c",
		"a/b/c-d":       "cd",
		"a/b/c-d/":      "cd",
		"a.b":           "ab",
		"a/b.c":         "bc",
		"a/b-c.d":       "bcd",
		"a/bb-ccc.dddd": "bbcccdddd",
		"a/foo-go":      "foogo",
	}
	for path, expected := range data {
		if guessAlias(path) != expected {
			fmt.Printf("guessAlias test failed %s should return %s but got %s\n", path, expected, guessAlias(path))
			t.Fail()
		}
	}
}

func TestValidAlias(t *testing.T) {
	data := map[string]bool{
		"a":   true,  // ok
		"b":   false, // already registered
		"go":  false, // keyword
		"int": false, // predeclared
		"err": false, // common name
	}
	f := NewFile("test")
	f.register("b")
	for alias, expected := range data {
		if f.isValidAlias(alias) != expected {
			fmt.Printf("isValidAlias test failed %s should return %t but got %t\n", alias, expected, f.isValidAlias(alias))
			t.Fail()
		}
	}
}
