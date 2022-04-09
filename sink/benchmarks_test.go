package sink_test

import (
	"fmt"
	"testing"
)

func BenchmarkKitchenSink(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("intentionally not implemented")
	}
}
