package main

import (
	"testing"
)

const (
	testPowerConsumptionGammaRate   = 22
	testPowerConsumptionEpsilonRate = 9
	testPowerConsumption            = 198
)

var (
	testInput = []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
)

func TestPowerConsumption(t *testing.T) {
	gamma, epsilon := PowerConsumption(testInput)
	if gamma != testPowerConsumptionGammaRate || epsilon != testPowerConsumptionEpsilonRate {
		t.Fatalf("%s: got gamma/epsilon %d/%d, expected %d/%d", t.Name(), gamma, epsilon, testPowerConsumptionGammaRate, testPowerConsumptionEpsilonRate)
	}
	t.Logf("%s: got gamma/epsilon %d/%d, expected %d/%d", t.Name(), gamma, epsilon, testPowerConsumptionGammaRate, testPowerConsumptionEpsilonRate)
}
