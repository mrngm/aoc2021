package main

import (
	"testing"
)

var (
	testBoards              = 3
	testNumbers             = 27
	testWinningNumber       = 24
	testWinningBoardSum     = 188
	testLastWinningNumber   = 13
	testLastWinningBoardSum = 148
	testInput               = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`
)

func TestParseInput(t *testing.T) {
	numbers, boards := ParseInput(testInput)
	if len(numbers) != testNumbers {
		t.Fatalf("%s: expected %d numbers, got %d", t.Name(), testNumbers, len(numbers))
	}
	t.Logf("%s: expected %d numbers, got %d", t.Name(), testNumbers, len(numbers))
	if len(boards) != testBoards {
		t.Fatalf("%s: expected %d boards, got %d", t.Name(), testBoards, len(boards))
	}
	t.Logf("%s: expected %d boards, got %d", t.Name(), testBoards, len(boards))
}

func TestPlayGame(t *testing.T) {
	numbers, boards := ParseInput(testInput)
	winner, boardSum := PlayGame(numbers, boards)
	if winner != testWinningNumber {
		t.Fatalf("%s: got %d as winning number, expected: %d", t.Name(), winner, testWinningNumber)
	}
	t.Logf("%s: got %d as winning number, expected: %d", t.Name(), winner, testWinningNumber)
	if boardSum != testWinningBoardSum {
		t.Fatalf("%s: got %d as winning board sum, expected: %d", t.Name(), boardSum, testWinningBoardSum)
	}
	t.Logf("%s: got %d as winning board sum, expected: %d", t.Name(), boardSum, testWinningBoardSum)
}

func TestPlayGameLastWinning(t *testing.T) {
	numbers, boards := ParseInput(testInput)
	winner, boardSum := PlayGameLastWinning(numbers, boards)
	if winner != testLastWinningNumber {
		t.Fatalf("%s: got %d as winning number, expected: %d", t.Name(), winner, testLastWinningNumber)
	}
	t.Logf("%s: got %d as winning number, expected: %d", t.Name(), winner, testLastWinningNumber)
	if boardSum != testLastWinningBoardSum {
		t.Fatalf("%s: got %d as winning board sum, expected: %d", t.Name(), boardSum, testLastWinningBoardSum)
	}
	t.Logf("%s: got %d as winning board sum, expected: %d", t.Name(), boardSum, testLastWinningBoardSum)
}
