package main

import (
	"flag"
	"log"
	"os"
)

var (
	inputFile    = flag.String("input", "input", "The input file")
	threshold    = flag.Int("threshold", 2, "Minimum number of points that overlap")
	useDiagonals = flag.Bool("diagonals", false, "Include diagonals in calculation")
)

func main() {
	flag.Parse()
	input, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("error opening input file %q: %v", *inputFile, err)
	}

	lines := ReadVentLines(string(input))
	segments := ExtractLineSegments(lines)
	diagram := GenerateCoverageDiagram(segments, !*useDiagonals)
	n := NumberOfOverlappingPoints(diagram, *threshold)

	diagonals := ""
	if *useDiagonals {
		diagonals = " (with diagonals)"
	}
	log.Printf("Scanning for hydrothermal vents%s, using %d segments, we found %d points where at least %d points overlap", diagonals, len(segments), n, *threshold)
}
