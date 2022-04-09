package sink_test

import (
	"fmt"

	"github.com/wilhelm-murdoch/go-gadget/sink"
)

func ExampleNewNormalStructTest() {
	s := sink.NewNormalStructTest("Wilhelm", "Murdoch", 40)

	fmt.Println(s.First, s.Last)

	// Output:
	// Wilhelm Murdoch
}
