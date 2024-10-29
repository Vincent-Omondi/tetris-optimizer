package solver

import (
	"github/Vincent-Omondi/tetris-optimizer/internal/board"
	"github/Vincent-Omondi/tetris-optimizer/internal/tetromino"
)

// Solve attempts to find a solution for placing all tetrominos on the board
func Solve(b board.Board, tetrominos []tetromino.Tetromino) board.Board {
	if len(tetrominos) == 0 {
		return b
	}

	current := tetrominos[0]
	for y := range b {
		for x := range b[y] {
			if b.CanPlace(current.Shape, x, y) {
				b.Place(current.Shape, x, y)
				if result := Solve(b, tetrominos[1:]); result != nil {
					return result
				}
				b.Remove(current.Shape, x, y)
			}
		}
	}
	return nil
}