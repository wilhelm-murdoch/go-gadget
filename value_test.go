package gadget_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilhelm-murdoch/go-gadget"
)

func TestNewValue(t *testing.T) {
	var err error
	var file *gadget.File
	var vars []string

	file, err = gadget.NewFile("value_test.go")
	assert.Nil(t, err, "Expected to open existing file, but got: %s", err)

	file.Values.Each(func(i int, v *gadget.Value) bool {
		vars = append(vars, v.String())
		return false
	})
	assert.Equal(t, len(vars), 3, "Expected 3 vars declared in `value_test.go`, but got %d instead.", len(vars))
}
