package main

import (
	"fmt"

	"github.com/kalmduck/algorithms/backtrack"
)

var n int

// Board stores the positions of queens.
// Each cell represents one row, the value b[r] specifies which column
// the queen in that row occupies.
type Board []int

func posToSquare(pos backtrack.Position) Square {
	s, ok := pos.(Square)
	if !ok {
		panic("Position couldn't be converted to Square.")
	}
	return s
}

func (b Board) String() string {
	var s string
	s += "\n"
	for i := 0; i < n; i++ {
		s += "|"
		for j := 0; j < n; j++ {
			if b[i]-1 == j {
				s += "Q"
			} else {
				s += "+"
			}
			s += "|"
		}
		s += "\n"
	}
	return s
}

func (b Board) Valid(pos backtrack.Position) bool {
	s := posToSquare(pos)
	for i := 0; i < s.r; i++ { // for every row that has a queen so far
		if b[i]-1 == s.c { // this column is already covered
			return false
		}
		if checkDiag(i, b[i]-1, s.r, s.c) {
			// should cover the diags
			return false
		}
	}
	return true
}

func checkDiag(qr, qc, r, c int) bool {
	var rowDiff, colDiff int
	if rowDiff = r - qr; rowDiff < 0 {
		rowDiff = -rowDiff
	}
	if colDiff = c - qc; colDiff < 0 {
		colDiff = -colDiff
	}
	return rowDiff == colDiff
}

func (b Board) Record(pos backtrack.Position) {
	s := posToSquare(pos)
	b[s.r] = s.c + 1
}

func (b Board) Done(pos backtrack.Position) bool {
	s := posToSquare(pos)
	return s.r == n-1
}

func (b Board) Undo(pos backtrack.Position) {
	s := posToSquare(pos)
	b[s.r] = 0
}

type Square struct {
	r, c int
}

func (s Square) NextVal() backtrack.Position {
	var p backtrack.Position
	p = Square{s.r, s.c + 1}
	return p
}

func (s Square) End() bool {
	return !(s.c < n)
}

func (s Square) NextPos() backtrack.Position {
	var p backtrack.Position
	p = Square{s.r + 1, 0}
	return p
}

func main() {
	n = 16
	b := make(Board, n)
	tracker := backtrack.New(b)
	var p backtrack.Position
	p = Square{0, 0}
	if tracker.Solve(p) {
		fmt.Println("Found solution: ", b)
	} else {
		fmt.Println("no solution: ", b)
	}
}
