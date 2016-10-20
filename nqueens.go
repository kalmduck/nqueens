package main

import (
	"fmt"
	"math"

	"github.com/kalmduck/algorithms/backtrack"
)

var n int

// Board stores the positions of queens.
// Each cell represents one row, the value b[r] specifies which column
// the queen in that row occupies.
type Board []int

var _ backtrack.Problem = (Board)(nil)
var _ backtrack.Position = (*Square)(nil)

func posToSquare(pos backtrack.Position) Square {
	s, ok := pos.(Square)
	if !ok {
		panic("Position couldn't be converted to Square.")
	}
	return s
}

func (b Board) Valid(pos backtrack.Position) bool {
	s := posToSquare(pos)
	for i := 0; i < s.r; i++ { // for every row that has a queen so far
		if b[i] == s.c+1 { // this column is already covered
			fmt.Println("covered column", s)
			fmt.Println(b)
			return false
		}
		if math.Abs(float64(b[i]-b[s.r])) == math.Abs(float64((i)-(s.r))) {
			// should cover the diags
			fmt.Println("covered diag", s)
			fmt.Println(b)
			return false
		}
	}
	return true
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
	n = 4
	b := make(Board, n)
	tracker := backtrack.New(b)
	var p backtrack.Position
	p = Square{0, 0}
	if tracker.Solve(p) {
		fmt.Println("Found solution: ", b)
	} else {
		fmt.Println("no solution: ", b)
	}
	fmt.Println(p)
	fmt.Println(b)
}
