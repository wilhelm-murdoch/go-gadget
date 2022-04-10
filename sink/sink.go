// Package sink covers most, if not all, patterns to adequately-test gadget's
// capabilities.
package sink

import "fmt"

// a block comment describing const assignments:
const (
	ONE   = 1 // represents the number 1
	TWO   = 2 // represents the number 2
	THREE = 3 // represents the number 3
)

// a block comment describing var assignments:
var (
	one   = "one"   // represents the english spelling of 1
	two   = "two"   // represents the english spelling of 2
	three = "three" // represents the english spelling of 3
)

type NormalFuncType func() bool
type NormalChannelType chan bool
type NormalRandomType map[string]int
type GenericRandomType[T any] []string

// a block comment describing var collections.
var collection map[string]map[string]string // this should be picked up as an inline comment.

// PrintVars prints out a value on each line.
func PrintVars() {
	fmt.Println(one)
	fmt.Println(two) // this will be the string "two".
	fmt.Println(three)
}

// AssignCollection assigns values to var collection.
func AssignCollection() {
	collection = make(map[string]map[string]string)

	// Here we assign values.
	collection["one"] = map[string]string{"foo": "bar"}
	collection["two"] = map[string]string{"merp": "flakes"}
}

// PrintConst prints out a value on each line.
func PrintConst() {
	fmt.Println(ONE)
	fmt.Println(TWO)
	fmt.Println(THREE) // this will be number 3.
}

// InterfaceTest is an example of an interface definition.
type InterfaceTest interface {
	ImplementMe() // this should be added to your data structures to implement this interface.
}

// EmbeddedStructTest represents an example of an embedded struct.
type EmbeddedStructTest struct {
	Occupation string // a standard job title.
}

// NormalStructTest represents an example of a top-level struct.
type NormalStructTest struct {
	First               string // first name
	Last                string // last name
	Age                 int    // age
	private             string // a dark secret
	*EmbeddedStructTest        // an embedded struct
}

// NewNormalStructTest returns a new instance of NormalStructTest.
func NewNormalStructTest(first, last string, age int) *NormalStructTest {
	return &NormalStructTest{
		First:              first,
		Last:               last,
		Age:                age,
		EmbeddedStructTest: &EmbeddedStructTest{"SWE"}, // will always be "SWE"
	}
}

// GetPrivate is an accessor method that returns a dark secret.
func (nst *NormalStructTest) GetPrivate() string {
	return nst.private
}

// GetOccupation is an accessor method that returns an occupation.
func (nst *NormalStructTest) GetOccupation() string {
	return nst.Occupation
}

// GetFullName is an function that attempts to return a full name.
func (nst NormalStructTest) GetFullName() string {
	return fmt.Sprint(nst.First, nst.Last)
}

// notExported is an example of a function that will not be exported.
//lint:ignore U1000 This struct is only used to test Field.IsExported.
func (nst NormalStructTest) notExported() string {
	return "I should not be exported!"
}

// GenericStructTest represents an example of a generic struct.
type GenericStructTest[T any] struct {
	First   string
	Last    string
	Age     int
	private string
}

// NewGenericStructTest returns a new instance of GenericStructTest.
func NewGenericStructTest[T any](first, last string, age int) *GenericStructTest[T] {
	return &GenericStructTest[T]{first, last, age, "hidden"}
}

// GetPrivate is an accessor method that returns a dark secret.
func (nst *GenericStructTest[T]) GetPrivate() string {
	return nst.private
}

// GetFullName is an function that attempts to return a full name.
func (nst GenericStructTest[T]) GetFullName() string {
	return fmt.Sprint(nst.First, nst.Last)
}

// IsBlank is an function that does not have a body.
func (nst GenericStructTest[T]) IsBlank() {}
