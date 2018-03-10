package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
)

func main() {
	// notest

	var out = flag.String("output", "./package-names.go", "Output filename to write")
	var pkg = flag.String("package", "main", "Package name in generated file")
	var name = flag.String("name", "PackageNames", "Name of the variable to define")
	var filter = flag.String("filter", ".*", "Regex to filter paths (operates on full path including vendor directory)")
	var standard = flag.Bool("standard", false, "Use standard library packages")
	var novendor = flag.Bool("novendor", false, "Exclude packages in vendor directories")
	var goListPath = flag.String("path", "all", "Path to pass to go list command")
	flag.Parse()

	buf := &bytes.Buffer{}
	if err := hints(buf, *pkg, *name, *goListPath, *filter, *standard, *novendor); err != nil {
		log.Fatal(err.Error())
	}
	if err := ioutil.WriteFile(*out, buf.Bytes(), 0644); err != nil {
		log.Fatal(err.Error())
	}
}
