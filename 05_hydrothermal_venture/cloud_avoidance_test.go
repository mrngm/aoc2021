package main

import (
	"testing"
)

var (
	testInput = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
	testInputLines                 = 10
	testInputSegments              = 10
	testNumberOverlappingPoints    = 5
	testOverlappingPointsThreshold = 2

	testCoverageDiagram = `.......1..
..1....1..
..1....1..
.......1..
.112111211
..........
..........
..........
..........
222111....
`
)

func TestReadVentLines(t *testing.T) {
	lines := ReadVentLines(testInput)
	if len(lines) != testInputLines {
		t.Fatalf("read %d lines, but expected %d", len(lines), testInputLines)
	}
	t.Logf("read %d lines, expected %d", len(lines), testInputLines)
}

func TestExtractLineSegments(t *testing.T) {
	lines := ReadVentLines(testInput)
	segments := ExtractLineSegments(lines)
	if len(segments) != testInputSegments {
		t.Fatalf("extracted %d segments, but expected %d", len(segments), testInputSegments)
	}
	t.Logf("extracted %d segments, expected %d", len(segments), testInputSegments)
}

func TestGenerateCoverageDiagramString(t *testing.T) {
	lines := ReadVentLines(testInput)
	segments := ExtractLineSegments(lines)
	diagram := GenerateCoverageDiagram(segments)
	diagramString := DiagramToString(diagram)
	if diagramString != testCoverageDiagram {
		t.Fatalf("coverage diagram\n%q\n did not match expected one:\n%q\n", diagramString, testCoverageDiagram)
	}
	t.Logf("successfully generated coverage diagram")
}

func TestNumberOfOverlappingPoints(t *testing.T) {
	lines := ReadVentLines(testInput)
	segments := ExtractLineSegments(lines)
	diagram := GenerateCoverageDiagram(segments)
	n := NumberOfOverlappingPoints(diagram, testOverlappingPointsThreshold)
	if n != testNumberOverlappingPoints {
		t.Fatalf("found %d overlapping points (threshold %d), but expected %d", n, testOverlappingPointsThreshold, testNumberOverlappingPoints)
	}
	t.Logf("found %d overlapping points (threshold %d), expected %d", n, testOverlappingPointsThreshold, testNumberOverlappingPoints)
}
