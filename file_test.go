package gadget_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilhelm-murdoch/go-gadget"
)

func TestNewFile(t *testing.T) {
	var err error
	var file *gadget.File

	_, err = gadget.NewFile("sink/sink.go")
	assert.Nil(t, err, "Expected to open existing file, but got: %s", err)

	_, err = gadget.NewFile("not_found.go")
	assert.NotNil(t, err, "Expected to exit with error when opening non-existant file.")

	find := "sink/sink.go"
	file, _ = gadget.NewFile(find)
	assert.Equal(t, file.String(), find, "Expected to return `%s`, but got `%s` instead.", file, find)

	file, _ = gadget.NewFile("sink/sink_test.go")
	assert.True(t, file.IsTest, "Expected to return `true` for a test file.")

	file, _ = gadget.NewFile("cmd/gadget/main.go")
	assert.True(t, file.IsMain, "Expected to return `true` for a main file.")

	file, _ = gadget.NewFile("sink/benchmarks_test.go")
	assert.True(t, file.HasBenchmarks, "Expected to return `true` for containing benchmark functions.")

	file, _ = gadget.NewFile("sink/examples_test.go")
	assert.True(t, file.HasExamples, "Expected to return `true` for containing example functions.")
}

func TestFileGetAstAttributes(t *testing.T) {
	find := "sink/sink.go"
	file, _ := gadget.NewFile(find)
	assert.Equal(t, file.String(), find, "Expected to return `%s`, but got `%s` instead.", find, file)

	astFile, tokenSet := file.GetAstAttributes()
	name := tokenSet.File(astFile.Pos()).Name()
	assert.Equal(t, name, find, "Expected to return `%s`, but got `%s` instead.", find, name)
}
