package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	inputFile = flag.String("input", "input", "The input file")
)

func SweepDepthIncrease(inputs []int) int {
	lastInput := 0
	ret := 0

	for n, d := range inputs {
		if n == 0 {
			lastInput = d
			continue
		}
		if d > lastInput {
			ret++
			lastInput = d
		}
		lastInput = d
	}
	return ret
}

func main() {
	flag.Parse()
	input, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("error opening input file %q: %v", *inputFile, err)
	}
	inputFields := strings.Fields(string(input))
	inputInts := make([]int, len(inputFields))
	for n, field := range inputFields {
		theInt, err := strconv.Atoi(field)
		if err != nil {
			log.Fatalf("error converting field %d (%q) to integer: %v", n, field, err)
		}
		inputInts[n] = theInt
	}

	out := SweepDepthIncrease(inputInts)
	log.Printf("After %d measurements, we found %d measurements that are larger than the previous measurements", len(inputInts), out)
}
