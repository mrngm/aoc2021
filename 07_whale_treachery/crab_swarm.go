package main

import (
	"flag"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	inputFile  = flag.String("input", "input", "The input file")
	meanMethod = flag.Bool("useMean", false, "Use mean instead of median")
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

func Mean(in []int) float64 {
	intermediate := 0
	for _, i := range in {
		intermediate += i
	}
	return float64(intermediate) / float64(len(in))
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

func DistanceFuelUsage(distance int) int {
	if distance <= 0 {
		return 0
	}
	if distance == 1 {
		return 1
	}
	return distance + DistanceFuelUsage(distance-1)
}

func FuelConsumptionUsingMean(in []int, mean int) int {
	if mean == -1 {
		// Calculate the mean from our input
		mean = int(math.Floor(Mean(in)))
	}
	ret := 0
	for _, crab := range in {
		if crab >= mean {
			ret += DistanceFuelUsage(crab - mean)
		} else {
			ret += DistanceFuelUsage(mean - crab)
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
	var consumption int
	if *meanMethod {
		consumption = FuelConsumptionUsingMean(crabs, -1)
	} else {
		consumption = FuelConsumption(crabs)
	}
	log.Printf("For %d crabs, we need %d fuel to align them all", len(crabs), consumption)
}
