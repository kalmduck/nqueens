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

// posToSquare sweeps away some of the mess caused by
// using the interface.
func posToSquare(pos backtrack.Position) Square {
	s, ok := pos.(Square)
	if !ok {
		panic("Position couldn't be converted to Square.")
	}
	return s
}

// String converts the current state of the the Board to a string.
// The position of a queen is marked with a 'Q', and empty squares are
// marked with a '+'.
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

// Valid checks if the given position is a valid place for a queen given
// the current board state. Returns false if the square is already covered by
// a queen on the board.  Assumes that queens have only been placed in
// previous rows.
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

// checkDiag compares the position of the a queen (qr, qc) with the position
// we're thinking about placing a queen (r, c).  If abs(r-qr) == abs(c-qc),
// the position we're considering is on the diagonal with a queen and can't
// be used.
func checkDiag(qr, qc, r, c int) bool {
	var rowDiff, colDiff int
	if rowDiff = r - qr; rowDiff < 0 { // abs
		rowDiff = -rowDiff
	}
	if colDiff = c - qc; colDiff < 0 { // abs
		colDiff = -colDiff
	}
	return rowDiff == colDiff
}

// Record marks this position with a queen.
// The column choices range from 1 - n, not the typical
// zero offset.
func (b Board) Record(pos backtrack.Position) {
	s := posToSquare(pos)
	b[s.r] = s.c + 1
}

// Done returns true if we have just placed the n-th queen, i.e.
// all the queens are placed, and we have found a solution.
func (b Board) Done(pos backtrack.Position) bool {
	s := posToSquare(pos)
	return s.r == n-1
}

// Undo removes a queen from the board at the passed position.
// This is accomplished by marking that position as 0 again.
func (b Board) Undo(pos backtrack.Position) {
	s := posToSquare(pos)
	b[s.r] = 0
}

// Square represents a particular location on the chess board grid
type Square struct {
	r, c int
}

// NextVal returns a Position of the next column value in the
// current row.
func (s Square) NextVal() backtrack.Position {
	var p backtrack.Position
	p = Square{s.r, s.c + 1}
	return p
}

// End returns true if the Square is at the end of all
// Possible columns.
func (s Square) End() bool {
	return !(s.c < n)
}

// NextPos advances to the next row in the chess board. In the algorithm,
// this is called after placing a queen in the current row to
// move on to the beginning of the next row
func (s Square) NextPos() backtrack.Position {
	var p backtrack.Position
	p = Square{s.r + 1, 0}
	return p
}

func main() {
	n = 16
	b := make(Board, n)
	tracker := backtrack.New(b)
	p := backtrack.Position(Square{0, 0})
	if tracker.Solve(p) {
		fmt.Println("Found solution: ", b)
	} else {
		fmt.Println("no solution: ", b)
	}
}
