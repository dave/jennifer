package jen

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

// NewFile Creates a new file, with the specified package name.
func NewFile(packageName string) *File {
	return &File{
		Group: &Group{
			separator: "\n",
		},
		name:    packageName,
		imports: map[string]string{},
	}
}

// NewFilePath creates a new file while specifying the package path - the
// package name is inferred from the path.
func NewFilePath(packagePath string) *File {
	return &File{
		Group: &Group{
			separator: "\n",
		},
		name:    guessAlias(packagePath),
		path:    packagePath,
		imports: map[string]string{},
	}
}

// NewFilePathName creates a new file with the specified package path and name.
func NewFilePathName(packagePath, packageName string) *File {
	return &File{
		Group: &Group{
			separator: "\n",
		},
		name:    packageName,
		path:    packagePath,
		imports: map[string]string{},
	}
}

// File represents a single source file. Package imports are managed
// automaticaly by File.
type File struct {
	*Group
	name        string
	path        string
	imports     map[string]string
	comments    []string
	headers     []string
	cgoPreamble []string
	// If you're worried about package aliases conflicting with local variable
	// names, you can set a prefix here. Package foo becomes {prefix}_foo.
	PackagePrefix string
}

// HeaderComment adds a comment to the top of the file, above any package
// comments. A blank line is rendered below the header comments, ensuring
// header comments are not included in the package doc.
func (f *File) HeaderComment(comment string) {
	f.headers = append(f.headers, comment)
}

// PackageComment adds a comment to the top of the file, above the package
// keyword.
func (f *File) PackageComment(comment string) {
	f.comments = append(f.comments, comment)
}

// CgoPreamble adds a cgo preamble comment that is rendered directly before the "C" pseudo-package
// import.
func (f *File) CgoPreamble(comment string) {
	f.cgoPreamble = append(f.cgoPreamble, comment)
}

// Anon adds an anonymous import:
func (f *File) Anon(paths ...string) {
	for _, p := range paths {
		f.imports[p] = "_"
	}
}

func (f *File) isLocal(path string) bool {
	return f.path == path
}

var reserved = []string{
	/* keywords */
	"break", "default", "func", "interface", "select", "case", "defer", "go", "map", "struct", "chan", "else", "goto", "package", "switch", "const", "fallthrough", "if", "range", "type", "continue", "for", "import", "return", "var",
	/* predeclared */
	"bool", "byte", "complex64", "complex128", "error", "float32", "float64", "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "true", "false", "iota", "nil", "append", "cap", "close", "complex", "copy", "delete", "imag", "len", "make", "new", "panic", "print", "println", "real", "recover",
	/* common variables */
	"err",
}

func isReservedWord(alias string) bool {
	for _, name := range reserved {
		if alias == name {
			return true
		}
	}
	return false
}

func (f *File) isValidAlias(alias string) bool {
	// the import alias is invalid if it's a reserved word
	if isReservedWord(alias) {
		return false
	}
	// the import alias is invalid if it's already been registered
	for _, v := range f.imports {
		if alias == v {
			return false
		}
	}
	return true
}

func (f *File) register(path string) string {
	if f.isLocal(path) {
		// notest
		// should never get here becasue in Qual the packageToken will be null,
		// so render will never be called.
		return ""
	}
	if f.imports[path] != "" && f.imports[path] != "_" {
		return f.imports[path]
	}
	if path == "C" {
		// special case for "C" pseudo-package
		f.imports[path] = "C"
		return "C"
	}
	alias := guessAlias(path)
	unique := alias
	i := 0
	for !f.isValidAlias(unique) {
		i++
		unique = fmt.Sprintf("%s%d", alias, i)
	}
	if f.PackagePrefix != "" {
		unique = f.PackagePrefix + "_" + unique
	}
	f.imports[path] = unique
	return unique
}

// GoString renders the File for testing. Any error will cause a panic.
func (f *File) GoString() string {
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		panic(err)
	}
	return buf.String()
}

func guessAlias(path string) string {
	alias := path

	if strings.HasSuffix(alias, "/") {
		// training slashes are usually tolerated, so we can get rid of one if
		// it exists
		alias = alias[:len(alias)-1]
	}

	if strings.Contains(alias, "/") {
		// if the path contains a "/", use the last part
		alias = alias[strings.LastIndex(alias, "/")+1:]
	}

	// alias should be lower case
	alias = strings.ToLower(alias)

	// alias should now only contain alphanumerics
	importsRegex := regexp.MustCompile(`[^a-z0-9]`)
	alias = importsRegex.ReplaceAllString(alias, "")

	return alias
}
