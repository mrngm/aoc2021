package main

import (
	"flag"
	"log"
	"os"
)

var (
	inputFile = flag.String("input", "input", "The input file")
	threshold = flag.Int("threshold", 2, "Minimum number of points that overlap")
)

func main() {
	flag.Parse()
	input, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("error opening input file %q: %v", *inputFile, err)
	}

	lines := ReadVentLines(string(input))
	segments := ExtractLineSegments(lines)
	diagram := GenerateCoverageDiagram(segments)
	n := NumberOfOverlappingPoints(diagram, *threshold)

	log.Printf("Scanning for hydrothermal vents, using %d segments, we found %d points where at least %d points overlap", len(segments), n, *threshold)
}
