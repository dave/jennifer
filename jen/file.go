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
	name     string
	path     string
	imports  map[string]string
	comments []string
	// If you're worried about package aliases conflicting with local variable
	// names, you can set a prefix here. Package foo becomes {prefix}_foo.
	PackagePrefix string
}

// PackageComment adds a comment to the top of the file, above the package
// keyword.
func (f *File) PackageComment(comment string) {
	f.comments = append(f.comments, comment)
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

func (f *File) register(path string) string {
	if f.isLocal(path) {
		return ""
	}
	if f.imports[path] != "" && f.imports[path] != "_" {
		return f.imports[path]
	}
	alias := guessAlias(path)
	unique := alias
	find := func(a string) bool {
		for _, v := range f.imports {
			if a == v {
				return true
			}
		}
		return false
	}
	i := 0
	for find(unique) {
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
	if strings.Contains(alias, "-") {
		// the name usually follows a hyphen - e.g. github.com/foo/go-bar if
		// the package name contains a "-", use the last part
		alias = alias[strings.LastIndex(alias, "-")+1:]
	}
	if strings.Contains(alias, ".") {
		// dot is commonly usually used as a version - e.g. github.com/foo/bar.v1
		// if the package name contains a ".", use the first part
		alias = alias[:strings.Index(alias, ".")]
	}
	// alias should be lower case
	alias = strings.ToLower(alias)

	// alias should now only contain alphanumerics
	importsRegex := regexp.MustCompile(`[^a-z0-9]`)
	alias = importsRegex.ReplaceAllString(alias, "")

	return alias
}
