package gadget

import (
	"go/ast"
	"strings"

	"github.com/wilhelm-murdoch/go-collection"
)

// Interface represents a golang interface definition.
type Interface struct {
	Name      string                         `json:"name"`       // The name of the interface.
	LineStart int                            `json:"line_start"` // The line number in the associated source file where this interface is initially defined.
	LineEnd   int                            `json:"line_end"`   // The line number in the associated source file where the definition block ends.
	LineCount int                            `json:"line_count"` // The total number of lines, including body, the interface occupies.
	Comment   string                         `json:"comment"`    // Any inline comments associated with the interface.
	Doc       string                         `json:"doc"`        // The comment block directly above this interface's definition.
	Signature string                         `json:"signature"`  // The full definition of the interface itself.
	Body      string                         `json:"body"`       // The full body of the interface sourced directly from the associated file; comments included.
	Fields    *collection.Collection[*Field] `json:"fields"`     // A collection of fields and their associated metadata.
	astType   *ast.InterfaceType
	astSpec   *ast.TypeSpec
	parent    *File
}

// NewInterface returns an interface instance and attempts to populate all
// associated fields with meaningful values.
func NewInterface(it *ast.InterfaceType, ts *ast.TypeSpec, parent *File) *Interface {
	return (&Interface{
		Name:    ts.Name.Name,
		Fields:  collection.New[*Field](),
		Doc:     ts.Doc.Text(),
		astType: it,
		astSpec: ts,
		parent:  parent,
	}).Parse()
}

// Parse is responsible for browsing through f.astSpec, f.astType, f.parent to
// populate the current interface's fields. ( Chainable )
func (i *Interface) Parse() *Interface {
	i.parseLines()
	i.parseBody()
	i.parseSignature()
	i.parseFields()
	// i.parseComments()

	return i
}

// parseLines determines the current interface's opening and closing line
// positions.
func (i *Interface) parseLines() {
	i.LineStart = i.parent.tokenSet.File(i.astSpec.Pos()).Line(i.astSpec.Pos())
	i.LineEnd = i.parent.tokenSet.Position(i.astSpec.End()).Line
	i.LineCount = (i.LineEnd + 1) - i.LineStart
}

// parseBody attempts to make a few adjustments to the *ast.BlockStmt which
// represents the current interface's body. We remove the opening and closing
// braces as well as the first occurrent `\t` sequence on each line.
func (i *Interface) parseBody() {
	i.Body = AdjustSource(string(GetLinesFromFile(i.parent.Path, i.LineStart+1, i.LineEnd-1)), false)
}

// parseSignature attempts to determine the current interfaces's type and assigns
// it to the Signature field of struct Function.
func (i *Interface) parseSignature() {
	line := strings.TrimSpace(string(GetLinesFromFile(i.parent.Path, i.LineStart, i.LineStart)))
	i.Signature = line[:len(line)-1]
}

// parseFields iterates through the interface's list of defined methods to
// populate the Fields collection.
func (i *Interface) parseFields() {
	for _, field := range i.astType.Methods.List {
		i.Fields.Push(NewField(field, i.parent))
	}
}

// parseComments is responsible for collecting and organising comments
// representing a comment block immediately above the interface definition.
// @TODO: Not yet implemented!
func (i *Interface) parseComments() {
	panic("Not yet implemented.")
}

// String implements the Stringer interface and returns the current package's
// name.
func (i *Interface) String() string {
	return i.Name
}
