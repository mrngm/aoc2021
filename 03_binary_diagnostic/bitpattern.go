package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

var (
	inputFile   = flag.String("input", "input", "The input file")
	lifeSupport = flag.Bool("lifesupport", false, "Do life support calculation")
)

func BitstringToUint(in string) uint {
	ret := uint(0)
	for _, c := range in {
		ret <<= 1
		if c == '1' {
			ret |= 1
		}
	}
	return ret
}

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

func ConsiderPosition(in []string, pos int, isOxygen bool) []string {
	ret := make([]string, 0)

	columnOnes, columnZeroes := CountOnesZeroes(in)
	for _, line := range in {
		if line == "" {
			continue
		}
		if pos >= len(columnOnes) || pos >= len(columnZeroes) {
			log.Printf("pos %d larger than columns, aborting", pos)
			return []string{}
		}
		if columnOnes[pos] >= columnZeroes[pos] {
			if isOxygen && line[pos] == '1' {
				ret = append(ret, line)
			} else if !isOxygen && line[pos] == '0' {
				ret = append(ret, line)
			}
		} else {
			if isOxygen && line[pos] == '0' {
				ret = append(ret, line)
			} else if !isOxygen && line[pos] == '1' {
				ret = append(ret, line)
			}
		}
	}

	if len(ret) > 0 && len(ret) != 1 {
		pos++
		ret = ConsiderPosition(ret, pos, isOxygen)
	}

	return ret
}

func CountOnesZeroes(in []string) ([]int, []int) {
	columnOnes := make([]int, len(in[0]))
	columnZeroes := make([]int, len(in[0]))

	for _, line := range in {
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

	return columnOnes, columnZeroes
}

func LifeSupportRating(input []string) (int, int) {
	oxygen := ConsiderPosition(input, 0, true)
	co2 := ConsiderPosition(input, 0, false)

	return int(BitstringToUint(oxygen[0])), int(BitstringToUint(co2[0]))
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

	if *lifeSupport {
		oxygen, co2 := LifeSupportRating(lines)
		log.Printf("After running the diagnostic report, we found the oxygen generator rating to be %d and CO2 scrubber rating to be %d: %d\n", oxygen, co2, oxygen*co2)
	} else {
		gamma, epsilon := PowerConsumption(lines)
		log.Printf("After running the diagnostic report, we found gamma to be %d and epsilon to be %d: %d\n", gamma, epsilon, gamma*epsilon)
	}
}
