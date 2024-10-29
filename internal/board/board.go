package board

import "fmt"

// Board represents the game board
type Board [][]string

// New creates a new board with the specified size
func New(size int) Board {
	board := make(Board, size)
	for i := range board {
		board[i] = make([]string, size)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	return board
}

// CanPlace checks if a shape can be placed at the given position
func (b Board) CanPlace(shape []string, x, y int) bool {
	for dy := range shape {
		for dx, char := range shape[dy] {
			if char != '.' {
				if y+dy >= len(b) || x+dx >= len(b[0]) || b[y+dy][x+dx] != "." {
					return false
				}
			}
		}
	}
	return true
}

// Place puts a shape on the board at the given position
func (b Board) Place(shape []string, x, y int) {
	for dy := range shape {
		for dx, char := range shape[dy] {
			if char != '.' {
				b[y+dy][x+dx] = string(char)
			}
		}
	}
}

// Remove takes a shape off the board at the given position
func (b Board) Remove(shape []string, x, y int) {
	for dy := range shape {
		for dx, char := range shape[dy] {
			if char != '.' {
				b[y+dy][x+dx] = "."
			}
		}
	}
}

// Print displays the board
func Print(board Board) {
	if board == nil {
		return
	}

	for _, row := range board {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}