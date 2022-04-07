package gadget_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilhelm-murdoch/go-gadget"
)

func TestNewFile(t *testing.T) {
	var err error
	var file *gadget.File

	_, err = gadget.NewFile("file.go")
	assert.Nil(t, err, "Expected to open existing file, but got: %s", err)

	_, err = gadget.NewFile("not_found.go")
	assert.NotNil(t, err, "Expected to exit with error when opening non-existant file.")

	file, _ = gadget.NewFile("file.go")
	assert.Equal(t, file.String(), "file.go", "Expected to return `file.go`, but got %s instead.", file)

	file, _ = gadget.NewFile("file_test.go")
	assert.True(t, file.IsTest, "Expected to return `true` for a test file.")

	file, _ = gadget.NewFile("cmd/gadget/main.go")
	assert.True(t, file.IsMain, "Expected to return `true` for a main file.")

	file, _ = gadget.NewFile("benchmarks_test.go")
	assert.True(t, file.HasBenchmarks, "Expected to return `true` for containing benchmark functions.")

	file, _ = gadget.NewFile("gadget_examples_test.go")
	assert.True(t, file.HasExamples, "Expected to return `true` for containing example functions.")
}
