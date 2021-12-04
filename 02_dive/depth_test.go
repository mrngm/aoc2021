package main

import (
	"reflect"
	"testing"
)

const (
	testHPos     = 15
	testDepth    = 10
	testAimHPos  = 15
	testAimDepth = 60
)

var (
	testCourse = []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	testInput  = []CourseOp{
		CourseOp{DirectionForward, 5},
		CourseOp{DirectionDown, 5},
		CourseOp{DirectionForward, 8},
		CourseOp{DirectionUp, 3},
		CourseOp{DirectionDown, 8},
		CourseOp{DirectionForward, 2},
	}
)

func TestDive(t *testing.T) {
	hPos, depth := Dive(testInput)
	if hPos != testHPos && depth != testDepth {
		t.Fatalf("%s: hPos: %d, depth: %d, expected: %d, %d", t.Name(), hPos, depth, testHPos, testDepth)
	}
	t.Logf("%s: hPos: %d, depth: %d, expected: %d, %d", t.Name(), hPos, depth, testHPos, testDepth)
}

func TestParseCourse(t *testing.T) {
	course := ParseDiveCourse(testCourse)
	if !reflect.DeepEqual(course, testInput) {
		t.Fatalf("%s: parsing dive course (%+v) failed, expected course: %+v", t.Name(), course, testInput)
	}
	t.Logf("%s: parsing dive course (%+v) succeeded (expected course: %+v)", t.Name(), course, testInput)
}

func TestDiveAim(t *testing.T) {
	hPos, depth := DiveAim(testInput)
	if hPos != testAimHPos && depth != testAimDepth {
		t.Fatalf("%s: hPos: %d, depth: %d, expected: %d, %d", t.Name(), hPos, depth, testAimHPos, testAimDepth)
	}
	t.Logf("%s: hPos: %d, depth: %d, expected: %d, %d", t.Name(), hPos, depth, testAimHPos, testAimDepth)
}
