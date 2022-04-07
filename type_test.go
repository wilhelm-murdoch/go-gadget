package gadget_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilhelm-murdoch/go-gadget"
)

type TestArray []string
type TestFunc func(string) error
type TestChan chan int
type TestMap map[string]string
type TestInterface interface {
	Thing()
	Boop()
	Stuff()
	Flarp()
}

func TestNewType(t *testing.T) {
	var err error
	var file *gadget.File

	file, err = gadget.NewFile("type_test.go")
	assert.Nil(t, err, "Expected to open existing file, but got: %s", err)

	var expectedTypes, foundTypes []string
	expectedTypes = []string{"TestArray", "TestFunc", "TestChan", "TestMap", "TestInterface"}
	file.Types.Each(func(i int, t *gadget.Type) bool {
		foundTypes = append(foundTypes, t.String())
		return false
	})
	assert.Equal(t, expectedTypes, foundTypes, "Was expecting `%s`, but got `%s` instead.", expectedTypes, foundTypes)
}
