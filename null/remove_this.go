package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
)

var src = []byte(`//lkjhklgjflk
package somepackage //fdfsfsd
import (
	"time" // comments

	// fasfasdasdas

	"fmt"
)

//fsdfsdsd

func SomeFunc() {
	if time.Now() == time.Now() {
		panic(time.Now())
	}
	fmt.Println("check")
}
`)

func main() {
	set := token.NewFileSet()
	f, err := parser.ParseFile(set, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	resDecls := make([]ast.Decl, 0, len(f.Decls))
	f.End()
	for i := range f.Decls {
		d, ok := f.Decls[i].(*ast.GenDecl)
		if !ok {
			resDecls = append(resDecls, f.Decls[i])
			continue
		}
		if d.Tok == token.IMPORT {
			continue
		}
		resDecls = append(resDecls, f.Decls[i])
	}
	f.Decls = resDecls
	var b bytes.Buffer
	err = printer.Fprint(&b, set, f)
	if err != nil {
		panic(err)
	}
	raw := []byte(b.String())

	y := f.Name.End()

	raw = append(raw[:f.Package-1], raw[y:]...)
	fmt.Println(b.String())
	fmt.Println()
	fmt.Println(string(raw))
}
