package main

import "testing"

func soma(a, b int) int {
	return a + b
}

func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		soma(2, 3)
	}
}
