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
