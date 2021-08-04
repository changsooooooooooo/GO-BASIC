package fibonacci

import "testing"

func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci1(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci2(20)
	}
}
