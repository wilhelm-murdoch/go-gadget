package substub_test

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/wilhelm-murdoch/go-gadget/stub/substub"
)

func ExampleSubstub_HelloName() {
	fmt.Println(substub.HelloName("wilhelm"))

	// Output:
	// Hello, wilhelm!
}

func TestHelloName(t *testing.T) {
	assert.Equal(t, substub.HelloName("wilhelm"), "Hello, wilhelm!")
}
