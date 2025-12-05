package main

import (
	"testing"

	"github.com/hassiexx/advent-of-code-2025/common"
)

func TestPart1(t *testing.T) {
	const expected = 3

	input := common.ReadInput(1, true)
	res := solve(input)

	if expected != res {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func BenchmarkPart1(b *testing.B) {
	input := common.ReadInput(1, false)

	for b.Loop() {
		solve(input)
	}
}
