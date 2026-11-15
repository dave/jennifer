package jen

import (
	"bytes"
	"fmt"
	"testing"
)

func TestFileExternalTestPackage(t *testing.T) {
	for _, tt := range []struct {
		name string
		path string
	}{
		{"pkg", "example.com/pkg"},
		{"client", "example.com/client/v2"},
		{"pkg_test", "example.com/pkg_test"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFilePathName(tt.path, tt.name+"_test")
			f.ImportName(tt.path, tt.name)
			f.Var().Id("value").Qual(tt.path, "Type")

			var buf bytes.Buffer
			if err := f.Render(&buf); err != nil {
				t.Fatal(err)
			}
			want := fmt.Sprintf("package %s_test\n\nimport %q\n\nvar value %s.Type\n", tt.name, tt.path, tt.name)
			if got := buf.String(); got != want {
				t.Errorf("got:\n%s\nwant:\n%s", got, want)
			}
		})
	}
}

func TestFileLocalPackageWithImportHint(t *testing.T) {
	for _, tt := range []struct {
		name  string
		hint  string
		alias bool
	}{
		{"pkg", "pkg", false},
		{"main", "main", false},
		{"pkg_test", "pkg_test", false},
		{"pkg_test", "pkg", true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFilePathName("example.com/pkg", tt.name)
			if tt.alias {
				f.ImportAlias("example.com/pkg", tt.hint)
			} else {
				f.ImportName("example.com/pkg", tt.hint)
			}
			f.Var().Id("value").Qual("example.com/pkg", "Type")

			var buf bytes.Buffer
			if err := f.Render(&buf); err != nil {
				t.Fatal(err)
			}
			want := fmt.Sprintf("package %s\n\nvar value Type\n", tt.name)
			if got := buf.String(); got != want {
				t.Errorf("got:\n%s\nwant:\n%s", got, want)
			}
		})
	}
}

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
		"123a":          "a",
		"a/321a.b":      "ab",
		"a/123":         "pkg",
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
