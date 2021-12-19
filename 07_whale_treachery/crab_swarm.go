package main

import (
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	inputFile = flag.String("input", "input", "The input file")
)

func ParseInput(in string) []int {
	in = strings.TrimSpace(in)
	splitted := strings.Split(in, ",")
	ret := make([]int, len(splitted))
	for n, digit := range splitted {
		d, err := strconv.Atoi(digit)
		if err != nil {
			ret[n] = -1
			continue
		}
		ret[n] = d
	}
	return ret
}

func Median(in []int) int {
	sort.Ints(in)
	middle := len(in) / 2
	middle1 := (len(in) / 2) - 1
	if len(in)%2 == 1 {
		return in[middle]
	}
	return (in[middle] + in[middle1]) / 2
}

func FuelConsumption(in []int) int {
	median := Median(in)
	ret := 0
	for _, crab := range in {
		if crab >= median {
			ret += crab - median
		} else {
			ret += median - crab
		}
	}
	return ret
}

func main() {
	flag.Parse()
	input, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("error opening input file %q: %v", *inputFile, err)
	}

	crabs := ParseInput(string(input))
	consumption := FuelConsumption(crabs)
	log.Printf("For %d crabs, we need %d fuel to align them all", len(crabs), consumption)
}
