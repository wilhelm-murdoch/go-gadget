package gadget

import "github.com/wilhelm-murdoch/go-collection"

// Package represents a golang package as well as all associated files,
// functions and other declarations.
type Package struct {
	Name  string                        `json:"name"`  // The name of the current package.
	Files *collection.Collection[*File] `json:"files"` // A collection of golang files associated with this package.
}

// NewPackage returns a Package instance with an initialised collection used for
// assigning and iterating through files.
func NewPackage(name string) *Package {
	return &Package{
		Name:  name,
		Files: collection.New[*File](),
	}
}

// String implements the Stringer inteface and returns the current package's
// name.
func (p *Package) String() string {
	return p.Name
}
