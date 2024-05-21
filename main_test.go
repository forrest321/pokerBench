package main

import "testing"

func BenchmarkUseFo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run()
	}
}
