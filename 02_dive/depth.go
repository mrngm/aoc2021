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

type Direction int

const (
	DirectionUnknown Direction = iota
	DirectionForward
	DirectionDown
	DirectionUp
)

type CourseOp struct {
	Course Direction
	Steps  int
}

func Dive(inputs []CourseOp) (int, int) {
	hPos, depth := 0, 0

	for _, op := range inputs {
		switch op.Course {
		case DirectionForward:
			hPos += op.Steps
		case DirectionDown:
			depth += op.Steps
		case DirectionUp:
			depth -= op.Steps
		}
	}

	return hPos, depth
}

func ParseDiveCourse(course []string) []CourseOp {
	ret := make([]CourseOp, 0)

	for n, c := range course {
		fields := strings.Fields(c)
		if len(fields) != 2 {
			log.Printf("course[%d]: unexpected number of fields (%d), expecting 2, ignoring line %q", n, len(fields), c)
			continue
		}

		direction, steps := fields[0], fields[1]
		nSteps, err := strconv.Atoi(steps)
		if err != nil {
			log.Printf("course[%d]: could not parse steps (%q), ignoring line %q", n, steps, c)
			continue
		}

		var dir Direction
		switch direction {
		case "forward":
			dir = DirectionForward
		case "up":
			dir = DirectionUp
		case "down":
			dir = DirectionDown
		}

		ret = append(ret, CourseOp{
			Course: dir,
			Steps:  nSteps,
		})
	}

	return ret
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
	diveCourse := ParseDiveCourse(lines)
	hPos, depth := Dive(diveCourse)
	log.Printf("After %d course operations, we ended up at horizontal position %d, and depth %d: %d", len(diveCourse), hPos, depth, hPos*depth)
}
