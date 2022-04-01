package gadget

import (
	"go/ast"
	"strings"

	"github.com/wilhelm-murdoch/go-collection"
)

// Struct represents a golang struct definition.
type Struct struct {
	Name      string                         `json:"name"`       // The name of the struct.
	LineStart int                            `json:"line_start"` // The line number in the associated source file where this struct is initially defined.
	LineEnd   int                            `json:"line_end"`   // The line number in the associated source file where the definition block ends.
	LineCount int                            `json:"line_count"` // The total number of lines, including body, the struct occupies.
	Comment   string                         `json:"comment"`    // Any inline comments associated with the struct.
	Doc       string                         `json:"doc"`        // The comment block directly above this struct's definition.
	Signature string                         `json:"signature"`  // The full definition of the struct itself.
	Body      string                         `json:"body"`       // The full body of the struct sourced directly from the associated file; comments included.
	Fields    *collection.Collection[*Field] `json:"fields"`     // A collection of fields and their associated metadata.
	astType   *ast.StructType
	astSpec   *ast.TypeSpec
	parent    *File
}

// NewStruct returns an struct instance and attempts to populate all
// associated fields with meaningful values.
func NewStruct(it *ast.StructType, ts *ast.TypeSpec, parent *File) *Struct {
	return (&Struct{
		Name:    ts.Name.Name,
		Fields:  collection.New[*Field](),
		Doc:     ts.Doc.Text(),
		astType: it,
		astSpec: ts,
		parent:  parent,
	}).Parse()
}

// Parse is responsible for browsing through f.astSpec, f.astType, f.parent to
// populate the current struct's fields. ( Chainable )
func (s *Struct) Parse() *Struct {
	s.parseLines()
	s.parseBody()
	s.parseSignature()
	s.parseFields()
	// s.parseComments()

	return s
}

// parseLines determines the current struct's opening and closing line
// positions.
func (s *Struct) parseLines() {
	s.LineStart = s.parent.tokenSet.File(s.astSpec.Pos()).Line(s.astSpec.Pos())
	s.LineEnd = s.parent.tokenSet.Position(s.astSpec.End()).Line
	s.LineCount = (s.LineEnd + 1) - s.LineStart
}

// parseBody attempts to make a few adjustments to the *ast.BlockStmt which
// represents the current struct's body. We remove the opening and closing
// braces as well as the first occurrent `\t` sequence on each line.
func (s *Struct) parseBody() {
	s.Body = AdjustSource(string(GetLinesFromFile(s.parent.Path, s.LineStart+1, s.LineEnd-1)), false)
}

// parseSignature attempts to determine the current structs's type and assigns
// it to the Signature field of struct Function.
func (s *Struct) parseSignature() {
	line := strings.TrimSpace(string(GetLinesFromFile(s.parent.Path, s.LineStart, s.LineStart)))
	s.Signature = line[:len(line)-1]
}

// parseFields iterates through the struct's list of defined methods to
// populate the Fields collection.
func (s *Struct) parseFields() {
	for _, field := range s.astType.Fields.List {
		s.Fields.Push(NewField(field, s.parent))
	}
}

// parseComments is responsible for collecting and organising comments
// representing a comment block immediately above the struct definition.
// @TODO: Not yet implemented!
func (s *Struct) parseComments() {
	panic("Not yet implemented.")
	// if s.parent.astFile.Comments != nil {
	// 	for _, cg := range s.parent.astFile.Comments {
	// 		for c := len(cg.List) - 1; c >= 0; c-- {
	// 			fmt.Println(cg.List[c].Pos(), s.astSpec.Pos())
	// 		}
	// 	}
	// }
}

// String implements the Stringer struct and returns the current package's
// name.
func (s *Struct) String() string {
	return s.Name
}
