package gadget_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilhelm-murdoch/go-gadget"
)

type GenericTest[T any] struct {
	Name string
}

func (gt *GenericTest[T]) Merp() {
}

func TestNewFunction(t *testing.T) {
	file, err := gadget.NewFile("function.go")
	assert.Nil(t, err, "Expected to open existing file, but got: %s", err)

	var functions []string
	file.Functions.Each(func(i int, f *gadget.Function) bool {
		functions = append(functions, f.String())
		return false
	})
	assert.True(t, len(functions) > 0, "Expected to receive a list of function names, but got nothing instead.")
}

func TestFunctionFields(t *testing.T) {
	file, err := gadget.NewFile("function.go")
	assert.Nil(t, err, "Expected to open existing file, but got: %s", err)

	find := "Parse"
	f := file.Functions.Find(func(i int, item *gadget.Function) bool {
		return item.Name == find
	})
	assert.NotNil(t, f, "Expected a function to reference, but got nothing instead.")
	assert.Equal(t, f.Name, find, "Expected function name %s, but got %s instead.", find, f.Name)
	assert.False(t, f.IsTest, "Expected a non-test function.")
	assert.False(t, f.IsBenchmark, "Expected a non-benchmark function.")
	assert.False(t, f.IsExample, "Expected a non-example function.")
	assert.True(t, f.IsExported, "Expected function to be exported.")
	assert.True(t, f.IsMethod, "Expected function to be a method of a struct.")
	assert.Equal(t, f.Receiver, "Function", "Expected function to have a receiver for type Function.")
	assert.True(t, len(f.Doc) > 0, "Expected function to have a documentation block above its definition.")
	assert.Equal(t, f.Output, "", "Expected function to have an empty output field.")
	assert.True(t, len(f.Body) > 0, "Expected function to contain a body.")
	assert.True(t, len(f.Signature) > 0, "Expected function to have a definition signature.")
	assert.Equal(t, f.Signature, "func (f *Function) Parse() *Function", "Function did not return expected signature.")
	assert.Equal(t, f.LineCount, f.LineEnd-f.LineStart, "Expected line count to equal the difference betwee the last line and the firest line.")
}

func TestFunctionExamples(t *testing.T) {
	file, err := gadget.NewFile("gadget_examples_test.go")
	assert.Nil(t, err, "Expected to open existing file, but got: %s", err)

	examples := file.Functions.Filter(func(f *gadget.Function) bool {
		return f.IsExample
	})
	assert.Greater(t, examples.Length(), 0, "Expected example functions, but got nothing instead.")

	examples.Each(func(i int, f *gadget.Function) bool {
		assert.Greater(t, len(f.Output), 0, "Expected example %s to have output, but got nothing instead.", f.Name)
		return false
	})
}

func TestFunctionGeneric(t *testing.T) {
	file, err := gadget.NewFile("function_test.go")
	assert.Nil(t, err, "Expected to open existing file, but got: %s", err)

	find := "Merp"
	f := file.Functions.Find(func(i int, item *gadget.Function) bool {
		return item.Name == find
	})
	assert.NotNil(t, f, "Expected a function to reference, but got nothing instead.")
	assert.Equal(t, f.Signature, "func (gt *GenericTest[T]) Merp()")
}
