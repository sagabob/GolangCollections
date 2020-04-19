package testing

import (
	"fmt"
	"testing"
)

func BenchmarkHello(b *testing.B) { // Benchmark definition
	for i := 0; i < b.N; i++ {
		fmt.Sprintln("Hello High Performance Go")
	}
}
