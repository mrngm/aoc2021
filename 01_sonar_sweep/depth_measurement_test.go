package main

import (
	"testing"
)

const (
	testOutput              = 7
	testOutputSlidingWindow = 5
)

var (
	testInput = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
)

func TestSweepDepthIncreases(t *testing.T) {
	out := SweepDepthIncrease(testInput)
	if out != testOutput {
		t.Fatalf("%s: out: %v, expected: %v", t.Name(), out, testOutput)
	}
	t.Logf("%s: out: %v, expected: %v", t.Name(), out, testOutput)
}

func TestSweepDepthIncreaseWithSlidingWindow(t *testing.T) {
	out, _ := SweepDepthIncreaseWithSlidingWindow(testInput)
	if out != testOutputSlidingWindow {
		t.Fatalf("%s: out: %v, expected: %v", t.Name(), out, testOutputSlidingWindow)
	}
	t.Logf("%s: out: %v, expected: %v", t.Name(), out, testOutputSlidingWindow)
}
