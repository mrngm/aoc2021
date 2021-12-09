package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	RegexpLineSegment = regexp.MustCompile(`^(\d+),(\d+)\s+->\s+(\d+),(\d+)$`)
)

type Point struct {
	X int
	Y int
}

type LineSegment struct {
	Begin *Point
	End   *Point
}

func (ls *LineSegment) IsHorizontal() bool {
	return ls.Begin.X == ls.End.X
}

func (ls *LineSegment) IsVertical() bool {
	return ls.Begin.Y == ls.End.Y
}

func (ls *LineSegment) PointCoverage() []*Point {
	var ret []*Point
	if !ls.IsHorizontal() && !ls.IsVertical() {
		log.Printf("XXX: non-horizontal/vertical line segments are currently not supported: %v,%v", ls.Begin, ls.End)
		return ret
	}
	ret = append(ret, ls.Begin)
	ret = append(ret, ls.End)
	if ls.IsHorizontal() {
		// Only Y coordinates differ. Range over the difference between the Y coordinates
		for y := ls.Begin.Y; y != ls.End.Y; {
			switch {
			case y > ls.End.Y:
				y--
				if y == ls.End.Y {
					// Don't add the point
					continue
				}
			case y < ls.End.Y:
				y++
				if y == ls.End.Y {
					// Don't add the point
					continue
				}
			default:
				// The line segment is actually a dot.
				continue
			}
			ret = append(ret, &Point{ls.Begin.X, y})
		}
	}

	if ls.IsVertical() {
		// Only X coordinates differ. Range over the difference between the X coordinates
		for x := ls.Begin.X; x != ls.End.X; {
			switch {
			case x > ls.End.X:
				x--
				if x == ls.End.X {
					// Don't add the point
					continue
				}
			case x < ls.End.X:
				x++
				if x == ls.End.X {
					// Don't add the point
					continue
				}
			default:
				// The line segment is actually a dot.
				continue
			}
			ret = append(ret, &Point{x, ls.Begin.Y})
		}
	}

	return ret
}

func ReadVentLines(in string) []string {
	var ret []string
	lines := strings.Split(in, "\n")
	for n, line := range lines {
		if !RegexpLineSegment.MatchString(line) {
			log.Printf("line %d (%q) did not match expected regexp, ignoring", n, line)
			continue
		}
		ret = append(ret, strings.TrimSpace(line))
	}

	return ret
}

func ExtractLineSegments(in []string) []*LineSegment {
	var ret []*LineSegment

	for n, line := range in {
		matches := RegexpLineSegment.FindAllStringSubmatch(line, -1)
		if len(matches) != 1 {
			log.Printf("line %d (%q) did not return 1 match, but %d, ignoring", n, line, len(matches))
			continue
		}
		if len(matches[0]) != 5 {
			log.Printf("line %d (%q) did not yield 5 submatches, but %d, ignoring", n, line, len(matches[0]))
			continue
		}
		x1, err := strconv.Atoi(matches[0][1])
		if err != nil {
			log.Printf("could not parse x1 (%q): %v, ignoring line", matches[0][1], err)
			continue
		}
		y1, err := strconv.Atoi(matches[0][2])
		if err != nil {
			log.Printf("could not parse y1 (%q): %v, ignoring line", matches[0][2], err)
			continue
		}
		x2, err := strconv.Atoi(matches[0][3])
		if err != nil {
			log.Printf("could not parse x2 (%q): %v, ignoring line", matches[0][3], err)
			continue
		}
		y2, err := strconv.Atoi(matches[0][4])
		if err != nil {
			log.Printf("could not parse y2 (%q): %v, ignoring line", matches[0][4], err)
			continue
		}
		segment := &LineSegment{
			&Point{x1, y1},
			&Point{x2, y2},
		}
		ret = append(ret, segment)
	}

	return ret
}

func GenerateCoverageDiagram(in []*LineSegment) [][]int {
	topLeftX, topLeftY := 0, 0
	bottomRightX, bottomRightY := 0, 0
	// Determine maximum range
	for _, segment := range in {
		if segment.Begin.X > bottomRightX {
			bottomRightX = segment.Begin.X
		}
		if segment.End.X > bottomRightX {
			bottomRightX = segment.End.X
		}
		if segment.Begin.Y > bottomRightY {
			bottomRightY = segment.Begin.Y
		}
		if segment.End.Y > bottomRightY {
			bottomRightY = segment.End.Y
		}
		if segment.Begin.X < topLeftX {
			topLeftX = segment.Begin.X
		}
		if segment.End.X < topLeftX {
			topLeftX = segment.End.X
		}
		if segment.Begin.Y < topLeftY {
			topLeftY = segment.Begin.Y
		}
		if segment.End.Y < topLeftY {
			topLeftY = segment.End.Y
		}
	}
	log.Printf("Coverage diagram ranges from (%d, %d) to (%d, %d)", topLeftX, topLeftY, bottomRightX, bottomRightY)

	if topLeftX < 0 || topLeftY < 0 || bottomRightX < 0 || bottomRightY < 0 {
		log.Printf("negative coordinates are currently not supported")
		return nil
	}

	diagram := make([][]int, bottomRightY+1)
	for y := range diagram {
		diagram[y] = make([]int, bottomRightX+1)
	}

	for _, segment := range in {
		points := segment.PointCoverage()
		for _, p := range points {
			diagram[p.Y][p.X]++
		}
	}

	return diagram
}

func DiagramToString(diagram [][]int) string {
	log.Printf("diagram: %+v", diagram)

	var ret string
	for y := range diagram {
		for x := range diagram[y] {
			switch diagram[y][x] {
			case 0:
				ret += "."
			default:
				ret += strconv.Itoa(diagram[y][x])
			}
		}
		ret += "\n"
	}

	// See cloud_avoidance_test.go for example output. In this function, we only consider overlapping horizontal/vertical line segments.

	return ret
}

func NumberOfOverlappingPoints(diagram [][]int, threshold int) int {
	// Count the number of overlapping points equal or larger to threshold
	ret := 0
	for y := range diagram {
		for x := range diagram[y] {
			p := diagram[y][x]
			switch {
			case p >= threshold:
				ret++
			}
		}
	}

	return ret
}
