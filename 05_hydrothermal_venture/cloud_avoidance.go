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

type EnumDirection int

// Although the diagrams have placed (0, 0) in the top left, EnumDirection constants assumes (0, 0) in the bottom left, just like regular XY diagrams
const (
	DirectionPoint EnumDirection = iota
	DirectionHorizontalRight
	DirectionHorizontalLeft
	DirectionVerticalUp
	DirectionVerticalDown
	DirectionDiagonalIncrease
	DirectionDiagonalDecrease
	DirectionDiagonalReverseIncrease // from left to right, bottom to top
	DirectionDiagonalReverseDecrease // from left to right, top to bottom
)

type Point struct {
	X int
	Y int
}

type LineSegment struct {
	Begin     *Point
	End       *Point
	Direction EnumDirection
}

func (ls *LineSegment) PointCoverage() []*Point {
	var ret []*Point

	switch ls.Direction {
	case DirectionHorizontalRight:
		for x := ls.Begin.X; x <= ls.End.X; x++ {
			ret = append(ret, &Point{x, ls.Begin.Y})
		}
	case DirectionHorizontalLeft:
		for x := ls.Begin.X; x >= ls.End.X; x-- {
			ret = append(ret, &Point{x, ls.Begin.Y})
		}
	case DirectionVerticalUp:
		for y := ls.Begin.Y; y <= ls.End.Y; y++ {
			ret = append(ret, &Point{ls.Begin.X, y})
		}
	case DirectionVerticalDown:
		for y := ls.Begin.Y; y >= ls.End.Y; y-- {
			ret = append(ret, &Point{ls.Begin.X, y})
		}
	case DirectionDiagonalIncrease:
		x := ls.Begin.X
		y := ls.Begin.Y
		for {
			if x > ls.End.X || y > ls.End.Y {
				break
			}
			ret = append(ret, &Point{x, y})
			x++
			y++
		}
	case DirectionDiagonalDecrease:
		x := ls.Begin.X
		y := ls.Begin.Y
		for {
			if x > ls.End.X || y < ls.End.Y {
				break
			}
			ret = append(ret, &Point{x, y})
			x++
			y--
		}
	case DirectionDiagonalReverseIncrease:
		x := ls.Begin.X
		y := ls.Begin.Y
		for {
			if x < ls.End.X || y > ls.End.Y {
				break
			}
			ret = append(ret, &Point{x, y})
			x--
			y++
		}
	case DirectionDiagonalReverseDecrease:
		x := ls.Begin.X
		y := ls.Begin.Y
		for {
			if x < ls.End.X || y < ls.End.Y {
				break
			}
			ret = append(ret, &Point{x, y})
			x--
			y--
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
		var direction EnumDirection

		// Although the diagrams have placed (0, 0) in the top left, this part assumes (0, 0) in the bottom left, just like regular XY diagrams
		if x1 == x2 {
			if y1 > y2 {
				direction = DirectionVerticalDown
			} else {
				direction = DirectionVerticalUp
			}
		} else if y1 == y2 {
			if x1 > x2 {
				direction = DirectionHorizontalLeft
			} else {
				direction = DirectionHorizontalRight
			}
		} else {
			if x1 > x2 {
				if y1 > y2 {
					direction = DirectionDiagonalReverseDecrease
				} else {
					direction = DirectionDiagonalReverseIncrease
				}
			} else {
				if y1 > y2 {
					direction = DirectionDiagonalDecrease
				} else {
					direction = DirectionDiagonalIncrease
				}
			}

		}
		segment := &LineSegment{
			&Point{x1, y1},
			&Point{x2, y2},
			direction,
		}
		ret = append(ret, segment)
	}

	return ret
}

func GenerateCoverageDiagram(in []*LineSegment, skipDiagonal bool) [][]int {
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
		if skipDiagonal {
			switch segment.Direction {
			case DirectionHorizontalRight, DirectionHorizontalLeft, DirectionVerticalUp, DirectionVerticalDown:
			default:
				continue
			}
		}
		points := segment.PointCoverage()
		log.Printf("len(points): %d, segment: %+v", len(points), segment)
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
