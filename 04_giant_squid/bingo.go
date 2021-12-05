package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	inputFile   = flag.String("input", "input", "The input file")
	lastWinning = flag.Bool("lastWin", false, "Get the scores of the last winning board")
)

var (
	RegexpNumberRow = regexp.MustCompile(`^(\d+,)+\d+$`)
	RegexpEmptyRow  = regexp.MustCompile(`^$`)
	RegexpBoardRow  = regexp.MustCompile(`^\s*(\d+)\s*(\d+)\s*(\d+)\s*(\d+)\s*(\d+)\s*$`)
)

func ParseInput(in string) ([]int, []*Board) {
	numbers := make([]int, 0)
	numbersSet := false
	boards := make([]*Board, 0)
	currentBoard := -1
	currentRow := -1

	lines := strings.Split(in, "\n")
	for n, line := range lines {
		if RegexpNumberRow.MatchString(line) && !numbersSet {
			nrs := strings.Split(line, ",")
			for _, nr := range nrs {
				// The regex ensures our input only contains digits, no need to check err here
				num, _ := strconv.Atoi(nr)
				numbers = append(numbers, num)
			}
			numbersSet = true
			continue
		}
		if RegexpEmptyRow.MatchString(line) {
			currentBoard++
			currentRow = 0
			continue
		}
		if RegexpBoardRow.MatchString(line) {
			if currentRow == 0 {
				// Create the board
				b := NewBoard(5, 5) // The boards are guaranteed to be 5x5
				boards = append(boards, b)
			}

			rowVals := RegexpBoardRow.FindStringSubmatch(line)
			vals := make([]int, len(rowVals)-1)
			for i, val := range rowVals {
				if i == 0 {
					continue
				}
				// The regex ensures our input only contains digits, no need to check err here
				num, _ := strconv.Atoi(val)
				vals[i-1] = num
			}
			boards[currentBoard].SetRow(currentRow, vals)
			currentRow++
			continue
		}
		log.Printf("unrecognized line (%d), ignoring: %v", n, line)
	}

	return numbers, boards
}

func PlayGame(numbers []int, boards []*Board) (int, int) {
	for _, num := range numbers {
		log.Printf("pulled number %d", num)
	boardLoop:
		for bi, b := range boards {
			err := b.MarkNumber(num)
			switch {
			case errors.Is(err, ErrOutOfBounds):
				log.Printf("error marking %d on board# %d, ignoring: bounds error", num, bi)
			case errors.Is(err, ErrMarkedValueNotFound):
				// This board did not contain the given number
			case err == nil:
				// Successfully marked, check for bingo before continuing to the next board
				hasBingo, score := b.CheckBingo()
				if hasBingo {
					log.Printf("got bingo for board# %d: %v", bi, score)
					return num, score
				}
				continue boardLoop
			}
		}
	}

	return 0, 0
}

func PlayGameLastWinning(numbers []int, boards []*Board) (int, int) {
	winningBoards := make(map[int]struct{})
	lastNumber, lastScore := 0, 0
	for _, num := range numbers {
		log.Printf("pulled number %d", num)
	boardLoop:
		for bi, b := range boards {
			err := b.MarkNumber(num)
			switch {
			case errors.Is(err, ErrOutOfBounds):
				log.Printf("error marking %d on board# %d, ignoring: bounds error", num, bi)
			case errors.Is(err, ErrMarkedValueNotFound):
				// This board did not contain the given number
			case err == nil:
				// Successfully marked, check for bingo before continuing to the next board
				hasBingo, score := b.CheckBingo()
				if hasBingo {
					if _, ok := winningBoards[bi]; ok {
						continue boardLoop
					}
					log.Printf("got bingo for board# %d: %v", bi, score)
					lastNumber = num
					lastScore = score
					if len(winningBoards) == len(boards)-1 {
						// We are the last board, let the giant squid win
						return lastNumber, lastScore
					}
					winningBoards[bi] = struct{}{}
				}
			}
		}
	}

	return lastNumber, lastScore
}

func main() {
	flag.Parse()
	input, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("error opening input file %q: %v", *inputFile, err)
	}
	numbers, boards := ParseInput(string(input))

	if *lastWinning {
		winningNumber, score := PlayGameLastWinning(numbers, boards)
		log.Printf("The winning number (%d) yielded score %d on the last winning board: %d", winningNumber, score, winningNumber*score)
	} else {
		winningNumber, score := PlayGame(numbers, boards)
		log.Printf("The winning number (%d) yielded score %d: %d", winningNumber, score, winningNumber*score)
	}
}
