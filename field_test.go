package gadget_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilhelm-murdoch/go-gadget"
)

type SubThing struct{}

type Thing struct {
	// a doc
	Name string // a comment
	//lint:ignore U1000 This struct is only used to test Field.IsExported.
	private string
	SubThing
}

func AnEmptyFunc() {}

func TestNewField(t *testing.T) {
	var err error

	file, err := gadget.NewFile("field_test.go")
	assert.Nil(t, err, "Expected to open existing file, but got: %s", err)

	sub := file.Types.Find(func(i int, item *gadget.Type) bool {
		return item.Name == "SubThing"
	})
	assert.NotNil(t, sub, "Expected to return *gadget.Type, but got nothing instead.")
	assert.Equal(t, sub.Fields.Length(), 0, "Expected to return *gadget.Type with zero fields, but got %d instead.", sub.Fields.Length())

	thing := file.Types.Find(func(i int, item *gadget.Type) bool {
		return item.Name == "Thing"
	})
	assert.NotNil(t, thing, "Expected to return *gadget.Type, but got nothing instead.")
	assert.Equal(t, thing.Fields.Length(), 3, "Expected to return *gadget.Type with zero fields, but got %d instead.", sub.Fields.Length())

	f1 := thing.Fields.Find(func(i int, item *gadget.Field) bool {
		return item.Name == "Name"
	})
	assert.Equal(t, f1.String(), f1.Name, "Expected the field name and stringer implementation to be equal.")
	assert.NotNil(t, f1, "Expected to return *gadget.Field, but got nothing instead.")
	assert.Equal(t, f1.Name, "Name", "Expected to return *gadget.Field named \"Name\", but got %s nothing instead.", f1.Name)
	assert.NotNil(t, f1.Comment, "Expected to have a comment associated with field, but got nothing instead.")
	assert.NotNil(t, f1.Doc, "Expected to have a document block associated with field, but got nothing instead.")
	assert.Equal(t, f1.Signature, "Name string", "Expected a properly-formatted signature for field named \"Name\", but got `%s` instead.", f1.Signature)
	assert.True(t, f1.IsExported, "Expected field to be marked as exported")
	assert.False(t, f1.IsEmbedded, "Expected field to not be an embedded type.")

	f2 := thing.Fields.Find(func(i int, item *gadget.Field) bool {
		return item.Name == "private"
	})
	assert.Equal(t, f2.String(), f2.Name, "Expected the field name and stringer implementation to be equal.")
	assert.NotNil(t, f2, "Expected to return *gadget.Field, but got nothing instead.")
	assert.Equal(t, f2.Name, "private", "Expected to return *gadget.Field named \"private\", but got %s nothing instead.", f1.Name)
	assert.Equal(t, f2.Comment, "", "Expected to not have a comment associated with field.")
	assert.Equal(t, f2.Doc, "", "Expected not to have a document block associated with field.")
	assert.Equal(t, f2.Signature, "private string", "Expected a properly-formatted signature for field named \"private\", but got `%s` instead.", f1.Signature)
	assert.False(t, f2.IsExported, "Expected field to be marked as private ( non-exported )")
	assert.False(t, f2.IsEmbedded, "Expected field to not be an embedded type.")

	f3 := thing.Fields.Find(func(i int, item *gadget.Field) bool {
		return item.Name == "SubThing"
	})
	assert.NotNil(t, f3, "Expected to return *gadget.Field, but got nothing instead.")
	assert.Equal(t, f3.String(), f3.Name, "Expected the field name and stringer implementation to be equal.")
	assert.Equal(t, f3.Name, "SubThing", "Expected to return *gadget.Field named \"private\", but got %s nothing instead.", f1.Name)
	assert.Equal(t, f3.Comment, "", "Expected to not have a comment associated with field.")
	assert.Equal(t, f3.Doc, "", "Expected not to have a document block associated with field.")
	assert.Equal(t, f3.Signature, "SubThing", "Expected a properly-formatted signature for field named \"SubThing\", but got `%s` instead.", f1.Signature)
	assert.True(t, f3.IsEmbedded, "Expected field to not be an embedded type.")
}
