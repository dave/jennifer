package main

import (
	"bytes"
	"io/ioutil"
)

func main() {
	// notest
	buf := &bytes.Buffer{}
	if err := render(buf); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("./jen/generated.go", buf.Bytes(), 0644); err != nil {
		panic(err)
	}
}
