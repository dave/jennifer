package jen

import (
	"bytes"
	"fmt"
	"go/parser"
	asttoken "go/token"
	"regexp"
	"strings"
)

// NewFile Creates a new file, with the specified package name.
func NewFile(packageName string) *File {
	return &File{
		Group: &Group{
			multi: true,
		},
		name:    packageName,
		imports: map[string]importdef{},
		hints:   map[string]importdef{},
	}
}

// NewFilePath creates a new file while specifying the package path - the
// package name is inferred from the path.
func NewFilePath(packagePath string) *File {
	return &File{
		Group: &Group{
			multi: true,
		},
		name:    guessAlias(packagePath),
		path:    packagePath,
		imports: map[string]importdef{},
		hints:   map[string]importdef{},
	}
}

// NewFilePathName creates a new file with the specified package path and name.
func NewFilePathName(packagePath, packageName string) *File {
	return &File{
		Group: &Group{
			multi: true,
		},
		name:    packageName,
		path:    packagePath,
		imports: map[string]importdef{},
		hints:   map[string]importdef{},
	}
}

func NewFileFromSource(src []byte) (*File, error) {
	set := asttoken.NewFileSet()
	file, err := parser.ParseFile(set, "", src, 0)
	if err != nil {
		return nil, fmt.Errorf("parse source error: %v", err)
	}
	return NewFileFromAst(file, set)
}

// File represents a single source file. Package imports are managed
// automaticaly by File.
type File struct {
	*Group
	name        string
	path        string
	imports     map[string]importdef
	hints       map[string]importdef
	comments    []string
	headers     []string
	cgoPreamble []string
	// If you're worried about generated package aliases conflicting with local variable names, you
	// can set a prefix here. Package foo becomes {prefix}_foo.
	PackagePrefix string
	// CanonicalPath adds a canonical import path annotation to the package clause.
	CanonicalPath string
}

// importdef is used to differentiate packages where we know the package name from packages where the
// import is aliased. If alias == false, then name is the actual package name, and the import will be
// rendered without an alias. If used == false, the import has not been used in code yet and should be
// excluded from the import block.
type importdef struct {
	name  string
	alias bool
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

// Anon adds an anonymous import.
func (f *File) Anon(paths ...string) {
	for _, p := range paths {
		f.imports[p] = importdef{name: "_", alias: true}
	}
}

// ImportName provides the package name for a path. If specified, the alias will be omitted from the
// import block. This is optional. If not specified, a sensible package name is used based on the path
// and this is added as an alias in the import block.
func (f *File) ImportName(path, name string) {
	f.hints[path] = importdef{name: name, alias: false}
}

// ImportNames allows multiple names to be imported as a map. Use the [gennames](gennames) command to
// automatically generate a go file containing a map of a selection of package names.
func (f *File) ImportNames(names map[string]string) {
	for path, name := range names {
		f.hints[path] = importdef{name: name, alias: false}
	}
}

// ImportAlias provides the alias for a package path that should be used in the import block.
func (f *File) ImportAlias(path, alias string) {
	f.hints[path] = importdef{name: alias, alias: true}
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
		if alias == v.name {
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

	// if the path has been registered previously, simply return the name
	def := f.imports[path]
	if def.name != "" && def.name != "_" {
		return def.name
	}

	// special case for "C" pseudo-package
	if path == "C" {
		f.imports["C"] = importdef{name: "C", alias: false}
		return "C"
	}

	if ImportAliasFromSources {
		if alias, ok := importsCache[path]; ok && alias != "" {
			f.hints[path] = importdef{name: alias, alias: false}
		}
	}

	var name string
	var alias bool

	if hint := f.hints[path]; hint.name != "" {
		// look up the path in the list of provided package names and aliases by ImportName / ImportAlias
		name = hint.name
		alias = hint.alias
	} else if standardLibraryHints[path] != "" {
		// look up the path in the list of standard library packages
		name = standardLibraryHints[path]
		alias = false
	} else {
		// if a hint is not found for the package, guess the alias from the package path
		name = guessAlias(path)
		alias = true
	}

	// If the name is invalid or has been registered already, make it unique by appending a number
	unique := name
	i := 0
	for !f.isValidAlias(unique) {
		i++
		unique = fmt.Sprintf("%s%d", name, i)
	}

	// If we've changed the name to make it unique, it should definitely be an alias
	if unique != name {
		alias = true
	}

	// Only add a prefix if the name is an alias
	if f.PackagePrefix != "" && alias {
		unique = f.PackagePrefix + "_" + unique
	}

	// Register the eventual name
	f.imports[path] = importdef{name: unique, alias: alias}

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
