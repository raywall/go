package main

import "testing"

func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		soma(2, 3)
	}
}
