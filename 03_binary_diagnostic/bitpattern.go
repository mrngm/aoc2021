package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

var (
	inputFile = flag.String("input", "input", "The input file")
)

func PowerConsumption(input []string) (int, int) {
	columnOnes := make([]int, len(input[0]))
	columnZeroes := make([]int, len(input[0]))

	for _, line := range input {
		if line == "" {
			continue
		}
		for col, c := range line {
			switch c {
			case '1':
				columnOnes[col]++
			case '0':
				columnZeroes[col]++
			}
		}
	}

	gamma := uint(0)
	epsilon := uint(0)

	for n, _ := range input[0] {
		gamma <<= 1
		epsilon <<= 1
		if columnOnes[n] > columnZeroes[n] {
			gamma |= 1
		} else {
			epsilon |= 1
		}
	}

	return int(gamma), int(epsilon)
}

func main() {
	flag.Parse()
	input, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("error opening input file %q: %v", *inputFile, err)
	}
	lines := strings.Split(string(input), "\n")
	for n, line := range lines {
		lines[n] = strings.TrimSpace(line)
	}

	gamma, epsilon := PowerConsumption(lines)
	log.Printf("After running the diagnostic report, we found gamma to be %d and epsilon to be %d: %d\n", gamma, epsilon, gamma*epsilon)
}
