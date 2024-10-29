package main

import (
	"fmt"
	"github/Vincent-Omondi/tetris-optimizer/internal/board"
	"github/Vincent-Omondi/tetris-optimizer/internal/solver"
	"github/Vincent-Omondi/tetris-optimizer/internal/tetromino"
	"math"
	"os"
)

const (
	maxBoardSize = 20 // Reasonable max board size
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<tetrominos_file.txt>")
		os.Exit(1)
	}

	tetrominos, err := tetromino.ReadFromFile(os.Args[1])
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	size := int(math.Ceil(math.Sqrt(float64(len(tetrominos) * 4))))

	for size <= maxBoardSize {
		b := board.New(size)
		if solution := solver.Solve(b, tetrominos); solution != nil {
			board.Print(solution)
			return
		}
		size++
	}
	fmt.Println("ERROR: no solution found")
}