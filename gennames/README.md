# gennames
For large projects, it may be useful to generate an index of package names for commonly used packages. 
The index of names can be added to each generated file using `File.ImportNames`. The `gennames` command 
is used internally to generate the list of standard library package names.

### Usage

```
Usage of gennames:
  -filter string
    	Regex to filter paths (operates on full path including vendor directory) (default ".*")
  -name string
    	Name of the variable to define (default "PackageNames")
  -novendor
    	Exclude packages in vendor directories
  -output string
    	Output filename to write (default "./package-names.go")
  -package string
    	Package name in generated file (default "main")
  -path string
    	Path to pass to go list command (default "all")
  -standard
    	Use standard library packages
```

### Path
Supply a `path` to pass to the `go list` command. You may use the wildcard `/...` to recursively return 
packages, but it's worth remembering that vendored packages are not returned by this method unless the 
path itself is a vendored path. Use `all` to return all packages in your `GOPATH` (including vendored 
packages), however remember this may take some time for a large `GOPATH`.

### Filter
Supply a regex `filter` to limit the packages that are returned by the `go list` command. The filter 
operates on the full vendored package path (e.g. `github.com/foo/bar/vendor/github.com/baz/qux`), however 
the package path added to the index is unvendored (e.g. `github.com/baz/qux`).

### Examples

```
gennames -filter "foo|bar"
```

Create a file named `package-names.go` with `package main` listing the names of all packages with paths 
containing `foo` or `bar`.

```
gennames -output "foo/names.go" -package "foo" -path "github.com/foo/bar/vendor/..."
```

Create a file named `foo/names.go` with `package foo` listing the names of all packages that are vendored 
inside `github.com/foo/bar`.

