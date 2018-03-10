package main

import (
	"fmt"
	"go/build"
	"io"
	"os/exec"
	"strings"

	"regexp"

	"path/filepath"

	. "github.com/dave/jennifer/jen"
)

func hints(w io.Writer, pkg, name, goListPath, filter string, standard, novendor bool) error {

	// notest

	file := NewFile(pkg)

	file.HeaderComment("This file is generated - do not edit.")
	file.Line()

	packages, err := getPackages(goListPath, filter, standard, novendor)
	if err != nil {
		return err
	}
	/*
		// <name> contains package name hints
		var <name> = map[string]string{
			...
		}
	*/
	file.Commentf("%s contains package name hints", name)
	file.Var().Id(name).Op("=").Map(String()).String().Values(DictFunc(func(d Dict) {
		for path, name := range packages {
			d[Lit(path)] = Lit(name)
		}
	}))

	return file.Render(w)
}

func getPackages(goListPath, filter string, standard, novendor bool) (map[string]string, error) {

	// notest

	r, err := regexp.Compile(filter)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("go", "list", "-e", "-f", "{{ .Standard }} {{ .ImportPath }} {{ .Name }}", goListPath)
	cmd.Env = []string{
		fmt.Sprintf("GOPATH=%s", build.Default.GOPATH),
		fmt.Sprintf("GOROOT=%s", build.Default.GOROOT),
	}
	if standard {
		cmd.Dir = filepath.Join(build.Default.GOROOT, "src")
	} else {
		cmd.Dir = filepath.Join(build.Default.GOPATH, "src")
	}
	b, err := cmd.Output()
	if err != nil {
		if x, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("go list command returned an error - %s: %s", err.Error(), string(x.Stderr))
		}
		return nil, fmt.Errorf("go list command returned an error: %s", err.Error())
	}
	all := strings.Split(strings.TrimSpace(string(b)), "\n")

	packages := map[string]string{}
	for _, j := range all {

		parts := strings.Split(j, " ")

		isStandard := parts[0] == "true"
		if isStandard != standard {
			continue
		}

		path := parts[1]
		name := parts[2]

		if novendor && hasVendor(path) {
			continue
		}

		if name == "main" {
			continue
		}

		if !r.MatchString(path) {
			continue
		}

		path = unvendorPath(path)

		if packages[path] != "" {
			continue
		}

		packages[path] = name
	}
	return packages, nil
}

func unvendorPath(path string) string {
	// notest
	i, ok := findVendor(path)
	if !ok {
		return path
	}
	return path[i+len("vendor/"):]
}

// FindVendor looks for the last non-terminating "vendor" path element in the given import path.
// If there isn't one, FindVendor returns ok=false.
// Otherwise, FindVendor returns ok=true and the index of the "vendor".
// Copied from cmd/go/internal/load
func findVendor(path string) (index int, ok bool) {
	// notest
	// Two cases, depending on internal at start of string or not.
	// The order matters: we must return the index of the final element,
	// because the final one is where the effective import path starts.
	switch {
	case strings.Contains(path, "/vendor/"):
		return strings.LastIndex(path, "/vendor/") + 1, true
	case strings.HasPrefix(path, "vendor/"):
		return 0, true
	}
	return 0, false
}

func hasVendor(path string) bool {
	// notest
	_, v := findVendor(path)
	return v
}
