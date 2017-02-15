package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"text/template"

	"log"

	"github.com/davelondon/gopackages"
	"github.com/davelondon/jennifer/rebecca"
)

func main() {

	defaultPackage, _ := gopackages.GetPackageFromDir(os.Getenv("GOPATH"), ".")
	// ignore error - means we can't find a package at the current dir

	pkgFlag := flag.String("package", defaultPackage, "Package to scan")
	flag.Parse()
	pkg := *pkgFlag

	if pkg == "" {
		wd, _ := os.Getwd()
		log.Fatalf("Can't find package at current dir (%s) and no package specified with 'package' flag.", wd)
	}

	dir, err := gopackages.GetDirFromPackage(os.Environ(), os.Getenv("GOPATH"), pkg)
	if err != nil {
		log.Fatal(err)
	}

	m, err := rebecca.NewCodeMap(pkg, dir)
	if err != nil {
		log.Fatal(err)
	}

	funcMap := template.FuncMap{
		"example": m.ExampleFunc(false),
		"code":    m.ExampleFunc(true),
		"output":  m.OutputFunc,
		"link":    m.LinkFunc,
		"doc":     m.DocFunc,
	}

	tpl := template.Must(template.New("main").Funcs(funcMap).ParseGlob("*.md.tpl"))

	buf := &bytes.Buffer{}
	if err := tpl.ExecuteTemplate(buf, "README.md.tpl", nil); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("README.md", buf.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
}
