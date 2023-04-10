// Package jennifer is a code generator for Go
package jennifer

//go:generate go install github.com/dave/jennifer/genjen@latest
//go:generate genjen
//go:generate go install github.com/dave/jennifer/gennames@latest
//go:generate gennames -output "jen/hints.go" -package "jen" -name "standardLibraryHints" -standard -novendor -path "./..."
//go:generate go install github.com/dave/rebecca/cmd/becca@latest
//go:generate becca -package=github.com/dave/jennifer/jen
