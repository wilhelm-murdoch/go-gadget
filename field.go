package gadget

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"
)

// Field represents a field used in either a struct or an interface.
type Field struct {
	Name       string `json:"name"`              // The name of the field.
	IsExported bool   `json:"is_exported"`       // Determines whether the field is exported.
	Line       int    `json:"line"`              // The line number this field appears on in the associated source file.
	Signature  string `json:"body"`              // The full definition of the field including name, arguments and return values.
	Comment    string `json:"comment,omitempty"` // Any inline comments associated with the field.
	Doc        string `json:"doc,omitempty"`     // The comment block directly above this field's definition.
	astField   *ast.Field
	parent     *File
}

// NewField returns a field instance and attempts to populate all associated
// fields with meaningful values.
func NewField(f *ast.Field, parent *File) *Field {
	return (&Field{
		Doc:      strings.TrimSpace(f.Doc.Text()),
		Comment:  strings.TrimSpace(f.Comment.Text()),
		Line:     parent.tokenSet.File(f.Pos()).Line(f.Pos()),
		astField: f,
		parent:   parent,
	}).Parse()
}

// Parse is responsible for browsing through f.astField, f.parent to populate
// the current fields's fields. ( Chainable )
func (f *Field) Parse() *Field {
	// A field can be a nested struct. So, in the absense of a field name, instead
	// assume we're dealing with a struct and use the field's type to populate the
	// name field.
	if f.astField.Names != nil {
		f.Name = f.astField.Names[len(f.astField.Names)-1].Name
	} else {
		f.Name = fmt.Sprintf("<nested struct>: %v", f.astField.Type)
	}

	pattern := regexp.MustCompile(`(?m)^[A-Z]{1}`)
	if pattern.MatchString(f.Name) {
		f.IsExported = true
	}

	f.parseSignature()

	return f
}

// parseSignature determines the position of the current field within the
// associated source file and extracts the relevant line of code. We only want
// the content before any inline comments. This will also replace consecutive
// spaces with a single space.
func (f *Field) parseSignature() {
	line := strings.TrimSpace(string(GetLinesFromFile(f.parent.Path, f.Line, f.Line)))
	if n := strings.IndexByte(line, '/'); n >= 0 {
		line = line[:n]
	}

	pattern := regexp.MustCompile(`\s+`)
	f.Signature = pattern.ReplaceAllString(line, " ")
}

// String implements the Stringer struct and returns the current package's name.
func (f *Field) String() string {
	return f.Name
}
