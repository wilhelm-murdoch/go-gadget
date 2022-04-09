package gadget_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilhelm-murdoch/go-gadget"
)

func TestFields(t *testing.T) {
	var err error
	var find string

	file, err := gadget.NewFile("sink/sink.go")
	assert.Nil(t, err, "Expected to open existing file, but got: %s", err)

	find = "NormalRandomType"
	t1 := file.Types.Find(func(i int, item *gadget.Type) bool {
		return item.Name == find
	})

	assert.NotNil(t, t1, "Expected to return `%s`, but got nothing instead.", find)
	assert.Equal(t, t1.Fields.Length(), 0, "Expected to return zero fields, but got %d instead.", t1.Fields.Length())

	find = "GenericRandomType"
	t2 := file.Types.Find(func(i int, item *gadget.Type) bool {
		return item.Name == find
	})

	assert.NotNil(t, t2, "Expected to return `%s`, but got nothing instead.", find)
	assert.Equal(t, t2.Fields.Length(), 0, "Expected to return zero fields, but got %d instead.", t2.Fields.Length())

	find = "InterfaceTest"
	t3 := file.Types.Find(func(i int, item *gadget.Type) bool {
		return item.Name == find
	})

	first, ok := t3.Fields.AtFirst()
	assert.True(t, ok, "Expected to return a single field, but got nothing instead.")
	assert.NotNil(t, t3, "Expected to return `%s`, but got nothing instead.", find)
	assert.Equal(t, t3.Fields.Length(), 1, "Expected to return one field, but got %d instead.", t3.Fields.Length())
	assert.Equal(t, first.String(), "ImplementMe", "Expected to return a specific field, but got nothing instead.")

	find = "NormalStructTest"
	t4 := file.Types.Find(func(i int, item *gadget.Type) bool {
		return item.Name == find
	})

	assert.NotNil(t, t4, "Expected to return `%s`, but got nothing instead.", find)
	assert.Equal(t, t4.Fields.Length(), 5, "Expected to return 5 fields, but got %d instead.", t4.Fields.Length())

	f1 := t4.Fields.CountBy(func(f *gadget.Field) bool {
		return f.IsEmbedded
	})
	assert.Equal(t, f1, 1, "Expected to return one embedded field, but got %d instead.", f1)

	f2 := t4.Fields.CountBy(func(f *gadget.Field) bool {
		return f.IsExported == false
	})
	assert.Equal(t, f2, 2, "Expected to return two unexported fields, but got %d instead.", f2)

	f3 := t4.Fields.Find(func(i int, item *gadget.Field) bool {
		return item.Name == "private"
	})
	assert.NotNil(t, f3, "Expected to return one field, but got %d instead.")
	assert.False(t, f3.IsExported, "Expected a private field, but got an exported on instead.")
	assert.Equal(t, f3.Signature, "private string", "Expected `private string`, but got `%s` instead.", f3.Signature)
}
