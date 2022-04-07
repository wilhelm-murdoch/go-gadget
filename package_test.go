package gadget_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilhelm-murdoch/go-collection"
	"github.com/wilhelm-murdoch/go-gadget"
)

func TestNewPackage(t *testing.T) {
	packages := collection.New[*gadget.Package]()
	p := gadget.NewPackage("flooopsie")
	packages.Push(p)
	assert.Equal(t, packages.Length(), 1, "Was expecting a single package, but got %d instead", packages.Length())

	found, ok := packages.At(0)
	assert.True(t, ok, "Was expecting value at index 0, but got nothing instead.")
	assert.Equal(t, found.String(), "flooopsie", "Was expecting value `flooopsie`, but got %s instead", found)
}
