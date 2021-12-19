package main

import (
	"testing"
	"reflect"
)

var (
	testStringInput = "3,4,3,1,2"
	testInput = []int{3,4,3,1,2}
	testDays18 = 18
	testAfter18Days = "6,0,6,4,5,6,0,1,1,2,6,0,1,1,1,2,2,3,3,4,6,7,8,8,8,8"
	testAfter18DaysFishCount = 26
	testDays80 = 80
	testAfter80DaysFishCount = 5934
)

func TestParseInput(t *testing.T) {
	in := ParseInput(testStringInput)
	if !reflect.DeepEqual(in, testInput) {
		t.Fatalf("string input and parsed input are not equal")
	}
}

func TestInitialize(t *testing.T) {
	init := Initialize(testInput)
	for n, fish := range init {
		if fish.TimeLeft != testInput[n] {
			t.Fatalf("fish %d did not initialize properly, timeleft %d, expected %d", n ,fish.TimeLeft, testInput[n])
		}
	}
}

func TestIterate18WithFormatting(t *testing.T) {
	init := Initialize(testInput)
	for d := 0; d <= testDays18; d++ {
		init = Iterate(init, d, testDays18, true)
	}
	if FormatFish(init) != testAfter18Days {
		t.Fatalf("after %d days, output did not match expected output", testDays18)
	}
	if len(init) != testAfter18DaysFishCount {
		t.Fatalf("after %d days, number of fish %d did not match expected %d", testDays18, len(init), testAfter18DaysFishCount)
	}
}

func TestIterate80(t *testing.T) {
	init := Initialize(testInput)
	for d := 0; d <= testDays80; d++ {
		init = Iterate(init, d, testDays80, false)
	}
	if len(init) != testAfter80DaysFishCount {
		t.Fatalf("after %d days, number of fish %d did not match expected %d", testDays80, len(init), testAfter80DaysFishCount)
	}
}
