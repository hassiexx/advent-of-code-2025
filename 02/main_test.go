package main

import (
	"testing"

	"github.com/hassiexx/advent-of-code-2025/common"
)

func TestPart1(t *testing.T) {
	const expected = 1227775554

	input := common.ReadInput(2, true)
	res := solve(input, false)

	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestPart2(t *testing.T) {
	const expected = 4174379265

	input := common.ReadInput(2, true)
	res := solve(input, true)

	if res != expected {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func BenchmarkPart1(b *testing.B) {
	input := common.ReadInput(2, false)

	for b.Loop() {
		solve(input, false)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := common.ReadInput(2, false)

	for b.Loop() {
		solve(input, true)
	}
}
