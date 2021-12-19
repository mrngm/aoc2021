package main

import (
	"reflect"
	"testing"
)

var (
	testStringInput     = "16,1,2,0,4,2,7,1,2,14"
	testInput           = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	testMedian          = 2
	testFuelConsumption = 37
)

func TestParseInput(t *testing.T) {
	in := ParseInput(testStringInput)
	if !reflect.DeepEqual(in, testInput) {
		t.Fatalf("string input and parsed input are not equal")
	}
}

func TestMedian(t *testing.T) {
	in := ParseInput(testStringInput)
	if median := Median(in); median != testMedian {
		t.Fatalf("found median value %d, expected %d", median, testMedian)
	}
}

func TestFuelConsumption(t *testing.T) {
	in := ParseInput(testStringInput)
	if consumption := FuelConsumption(in); consumption != testFuelConsumption {
		t.Fatalf("found consumption value %d, expected %d", consumption, testFuelConsumption)
	}
}
