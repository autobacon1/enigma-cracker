package main

import (
	"testing"
)

func TestGeneratePairs(t *testing.T) {

	pairs := generate()

	for i := 0; i < len(pairs); i++ {

		for j := 0; j < len(pairs); j++ {

			if i == j {
				break
			} else {

				if pairs[i] == pairs[j] {
					t.Error("Error, found duplicated value in pair array.")
				}
			}
		}
	}
}

func TestGeneratePairsSize(t *testing.T) {
	pairs := generate()
	if len(pairs) != 20 {
		t.Error("Pairs length not 20")
	}
}

func BenchmarkGeneratePairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generate()
	}
}

func BenchmarkGeneratePairs2(b *testing.B) {
	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		generate2(letters)
	}
}

func BenchmarkSwapLetter(b *testing.B) {
	pairs := generate()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		swapLetter(pairs, pairs[0])
	}
}
