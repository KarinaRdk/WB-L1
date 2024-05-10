package main

import (
	"testing"
)

func BenchmarkMu (b *testing.B) {
	withMu()
}

func BenchmarkAtomic(b *testing.B) {
	withAtomic()
}

func BenchmarkChanel(b *testing.B) {
	withChan()
}
