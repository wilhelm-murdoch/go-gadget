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

	find := "sink/sink.go"
	file, err = gadget.NewFile(find)
	assert.Nil(t, err, "Expected to open existing file, but got: %s", err)

	file.Values.Each(func(i int, v *gadget.Value) bool {
		vars = append(vars, v.String())
		return false
	})
	assert.Equal(t, len(vars), 7, "Expected 7 vars declared in `%s`, but got %d instead.", find, len(vars))
}
