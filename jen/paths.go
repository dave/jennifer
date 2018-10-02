package jen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	asttoken "go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	ImportAliasFromSources = false

	gopathCache  = ""
	importsCache = make(map[string]string)
)

func preloadAliasFromSource(path string, decl string) (string, error) {
	if std, ok := standardLibraryHints[path]; ok {
		importsCache[path] = std
		return std, nil
	}
	alias, ok := importsCache[path]
	if ok {
		return alias, nil
	}
	path, err := getRelatedFilePath(path)
	if err != nil {
		return "", err
	}
	pkgs, err := parser.ParseDir(asttoken.NewFileSet(), path, nonTestFilter, parser.PackageClauseOnly)
	if err != nil {
		return "", err
	}
	for k, pkg := range pkgs {
		// Name of type was not provided: take any import
		if decl == "" {
			alias = k
			break
		}
		// Remove all unexported declarations
		if !ast.PackageExports(pkg) {
			continue
		}
		// Try to find our decl in package
		if ast.FilterPackage(pkg, func(name string) bool { return name == decl }) {
			// filter returns true if package has declaration
			// make it to be sure, that we choose right alias
			// because golang spec allow to have multiple packages in one directory
			alias = k
			break
		}
	}
	importsCache[path] = alias
	return alias, nil
}

func getRelatedFilePath(pkg string) (string, error) {
	if gopathCache == "" {
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			var err error
			gopath, err = getDefaultGoPath()
			if err != nil {
				return "", fmt.Errorf("cannot determine GOPATH: %s", err)
			}
		}
		gopathCache = gopath
	}
	paths := allPaths(filepath.SplitList(gopathCache))
	for _, p := range paths {
		checkingPath := filepath.Join(p, pkg)
		if info, err := os.Stat(checkingPath); err == nil && info.IsDir() {
			return checkingPath, nil
		}
	}
	return "", fmt.Errorf("file '%v' is not in GOROOT or GOPATH. Checked paths:\n%s", pkg, strings.Join(paths, "\n"))
}

func allPaths(gopaths []string) []string {
	const _2 = 2
	res := make([]string, len(gopaths)+_2)
	res[0] = filepath.Join(build.Default.GOROOT, "src")
	res[1] = "vendor"
	for i := range res[_2:] {
		res[i+_2] = filepath.Join(gopaths[i], "src")
	}
	return res
}

func getDefaultGoPath() (string, error) {
	if build.Default.GOPATH != "" {
		return build.Default.GOPATH, nil
	}
	output, err := exec.Command("go", "env", "GOPATH").Output()
	return string(bytes.TrimSpace(output)), err
}

// filters all files with tests
func nonTestFilter(info os.FileInfo) bool {
	return !strings.HasSuffix(info.Name(), "_test.go")
}
