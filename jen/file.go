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
	alias := ""
	if sep := strings.LastIndex(path, "/"); sep > -1 {
		alias = path[sep+1:]
	} else {
		alias = path
	}
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
