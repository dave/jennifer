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
