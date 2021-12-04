package main

import (
	"testing"
)

const (
	testOutput = 7
)

var (
	testInput = []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
)

func TestSweepDepthIncreases(t *testing.T) {
	out := SweepDepthIncrease(testInput)
	if out != testOutput {
		t.Fatalf("out: %v, expected: %v", out, testOutput)
	}
	t.Logf("out: %v, expected: %v", out, testOutput)
}
