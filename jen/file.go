package jen

import (
	"bytes"
	"fmt"
	"strings"
)

func NewFile(name string) *File {
	return &File{
		Group: &Group{
			syntax: fileSyntax,
		},
		name:    name,
		imports: map[string]string{},
	}
}

func NewFilePath(name, path string) *File {
	return &File{
		Group: &Group{
			syntax: fileSyntax,
		},
		name:    name,
		path:    path,
		imports: map[string]string{},
	}
}

type File struct {
	*Group
	name    string
	path    string
	imports map[string]string
}

func (f *File) register(path string) string {
	if f.path == path {
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
	f.imports[path] = unique
	return unique
}

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
	return alias
}
