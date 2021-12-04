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
	sliding   = flag.Bool("sliding", false, "Use sliding window method")
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
		}
		lastInput = d
	}
	return ret
}

func SweepDepthIncreaseWithSlidingWindow(inputs []int) (int, int) {
	ret := 0
	nWindows := 0

	for n := 0; n+3 < len(inputs); n++ {
		nWindows++
	}

	for w := 0; w < nWindows; w++ {
		thisWindow := inputs[w+2] + inputs[w+1] + inputs[w]
		nextWindow := inputs[w+3] + inputs[w+2] + inputs[w+1]
		if nextWindow > thisWindow {
			ret++
		}
	}

	return ret, nWindows + 1
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

	if *sliding {
		out, nWindows := SweepDepthIncreaseWithSlidingWindow(inputInts)
		log.Printf("After %d measurements with %d sliding windows, we found %d measurements that are larger than the previous measurements", len(inputInts), nWindows, out)
	} else {
		out := SweepDepthIncrease(inputInts)
		log.Printf("After %d measurements, we found %d measurements that are larger than the previous measurements", len(inputInts), out)
	}
}
