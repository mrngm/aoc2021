package main

import (
	"testing"
)

var (
	testBoards          = 3
	testNumbers         = 27
	testWinningNumber   = 24
	testWinningBoardSum = 188
	testInput           = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

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
		t.Fatalf("expected %d numbers, got %d", testNumbers, len(numbers))
	}
	t.Logf("expected %d numbers, got %d", testNumbers, len(numbers))
	if len(boards) != testBoards {
		t.Fatalf("expected %d boards, got %d", testBoards, len(boards))
	}
	t.Logf("expected %d boards, got %d", testBoards, len(boards))
}

func TestPlayGame(t *testing.T) {
	numbers, boards := ParseInput(testInput)
	winner, boardSum := PlayGame(numbers, boards)
	if winner != testWinningNumber {
		t.Fatalf("got %d as winning number, expected: %d", winner, testWinningNumber)
	}
	t.Logf("got %d as winning number, expected: %d", winner, testWinningNumber)
	if boardSum != testWinningBoardSum {
		t.Fatalf("got %d as winning board sum, expected: %d", boardSum, testWinningBoardSum)
	}
	t.Logf("got %d as winning board sum, expected: %d", boardSum, testWinningBoardSum)
}
