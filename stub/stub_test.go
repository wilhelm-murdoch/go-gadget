package stub_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wilhelm-murdoch/go-gadget/stub"
)

func ExampleStub_HelloWorld() {
	fmt.Println(stub.HelloWorld())

	// Output:
	// Hello, world!
}

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, stub.HelloWorld(), "Hello, world!")
}
