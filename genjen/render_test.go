package main

import (
	"io/ioutil"
	"regexp"
	"testing"

	"bytes"
)

func TestRender(t *testing.T) {

	buf := &bytes.Buffer{}
	if err := render(buf); err != nil {
		t.Fatal(err.Error())
	}
	generatedString := buf.String()

	existingFilePath := "../jen/generated.go"
	existingBytes, err := ioutil.ReadFile(existingFilePath)
	if err != nil {
		t.Fatal(err.Error())
	}
	existingString := string(existingBytes)

	// The "goimports" tool will often re-order the imports, so this is a
	// kludge to remove it before comparing. This is not ideal!
	importsRegex := regexp.MustCompile(`(?ms:\nimport \(\n.*\n\)\n)`)
	generatedString = importsRegex.ReplaceAllString(generatedString, "-")
	existingString = importsRegex.ReplaceAllString(existingString, "-")

	if generatedString != existingString {
		t.Fatalf("Generated code is not what is present:\n%s", generatedString)
	}
}
