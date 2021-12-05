package main

import (
	"errors"
	"fmt"
)

var (
	ErrOutOfBounds         = errors.New("out of bounds")
	ErrMarkUnequalValue    = errors.New("marked value unequal to board position")
	ErrMarkedValueNotFound = errors.New("marked value could not be found")
)

type Board struct {
	contents [][]int
	marked   [][]bool
	rows     int
	cols     int
}

func NewBoard(rows, cols int) *Board {
	contents := make([][]int, rows)
	marks := make([][]bool, rows)
	for n, _ := range contents {
		contents[n] = make([]int, cols)
	}
	for n, _ := range marks {
		marks[n] = make([]bool, cols)
	}

	return &Board{
		contents: contents,
		marked:   marks,
		rows:     rows,
		cols:     cols,
	}
}

func (b *Board) SetRow(n int, vals []int) error {
	if len(vals) != b.cols || n > b.rows-1 {
		return fmt.Errorf("len(vals) (%d) or n (%d) out of bounds: %d/%d", len(vals), n, b.cols, b.rows-1)
	}
	for i, val := range vals {
		b.contents[n][i] = val
	}
	return nil
}

func (b *Board) Position(x, y int) int {
	if x < 0 || y < 0 || x > b.cols-1 || y > b.rows-1 {
		return -1
	}
	return b.contents[y][x]
}

func (b *Board) IsMarked(x, y int) bool {
	if x < 0 || y < 0 || x > b.cols-1 || y > b.rows-1 {
		return false
	}
	return b.marked[y][x]
}

func (b *Board) MarkNumber(value int) error {
	for x := 0; x < b.cols; x++ {
		for y := 0; y < b.rows; y++ {
			err := b.Mark(x, y, value)
			switch {
			case errors.Is(err, ErrOutOfBounds):
				return err
			case errors.Is(err, ErrMarkUnequalValue):
				// Let's try another position
			case err == nil:
				// Successfully marked!
				return nil
			}
		}
	}
	return ErrMarkedValueNotFound
}
func (b *Board) Mark(x, y, value int) error {
	if x < 0 || y < 0 || x > b.cols-1 || y > b.rows-1 {
		return ErrOutOfBounds
	}
	if b.Position(x, y) != value {
		return ErrMarkUnequalValue
	}
	b.marked[y][x] = true
	return nil
}

func (b *Board) CheckBingo() (bool, int) {
	hasBingo := false
	score := 0

	// Check rows for bingo
	for y := 0; y < b.rows; y++ {
		nMarks := 0
		for x := 0; x < b.cols; x++ {
			if b.IsMarked(x, y) {
				nMarks++
			}
		}
		if nMarks == b.cols {
			// This row won!
			hasBingo = true
			break
		}
	}

	if !hasBingo {
		// Check columns for bingo
		for x := 0; x < b.cols; x++ {
			nMarks := 0
			for y := 0; y < b.rows; y++ {
				if b.IsMarked(x, y) {
					nMarks++
				}
			}
			if nMarks == b.rows {
				// This column won!
				hasBingo = true
				break
			}
		}
	}

	if hasBingo {
		// Sum all unmarked positions
		for x := 0; x < b.cols; x++ {
			for y := 0; y < b.rows; y++ {
				if !b.IsMarked(x, y) {
					score += b.Position(x, y)
				}
			}
		}
	}

	return hasBingo, score
}
