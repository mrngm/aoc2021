package main

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	testStringInput        = "16,1,2,0,4,2,7,1,2,14"
	testInput              = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	testMedian             = 2
	testFuelConsumption    = 37
	testFuelConsumptionTo5 = 168
	testFuelConsumptionTo2 = 206
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

func TestFuelConsumptionTo5(t *testing.T) {
	in := ParseInput(testStringInput)
	if consumption := FuelConsumptionUsingMean(in, -1); consumption != testFuelConsumptionTo5 {
		t.Fatalf("found consumption value %d, expected %d", consumption, testFuelConsumptionTo5)
	}
}

func TestFuelConsumptionTo2(t *testing.T) {
	in := ParseInput(testStringInput)
	if consumption := FuelConsumptionUsingMean(in, testMedian); consumption != testFuelConsumptionTo2 {
		t.Fatalf("found consumption value %d, expected %d", consumption, testFuelConsumptionTo2)
	}
}

func TestDistanceFuelUsage(t *testing.T) {
	tests := []struct {
		Distance int
		Usage    int
	}{
		{1, 1},
		{2, 3},
		{3, 6},
		{4, 10},
		{5, 15},
		{9, 45},
		{11, 66},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Distance%d", test.Distance), func(t *testing.T) {
			if tryUsage := DistanceFuelUsage(test.Distance); tryUsage != test.Usage {
				t.Fatalf("distance %d: expected fuel usage %d, got %d", test.Distance, test.Usage, tryUsage)
			}
		})
	}
}
