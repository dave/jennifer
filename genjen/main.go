package main

import (
	"bytes"
	"io/ioutil"
)

func main() {
	buf := &bytes.Buffer{}
	if err := Render(buf); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("./generated.go", buf.Bytes(), 0644); err != nil {
		panic(err)
	}
}
