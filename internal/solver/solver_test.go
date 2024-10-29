package solver

import (
	"github/Vincent-Omondi/tetris-optimizer/internal/board"
	"github/Vincent-Omondi/tetris-optimizer/internal/tetromino"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name       string
		boardSize  int
		tetrominos []tetromino.Tetromino
		wantSolved bool
	}{
		{
			name:      "empty tetrominos",
			boardSize: 3,
			tetrominos: []tetromino.Tetromino{},
			wantSolved: true,
		},
		{
			name:      "single tetromino",
			boardSize: 2,
			tetrominos: []tetromino.Tetromino{
				{
					Shape:  []string{"AA", "AA"},
					Letter: 'A',
				},
			},
			wantSolved: true,
		},
		{
			name:      "impossible fit",
			boardSize: 2,
			tetrominos: []tetromino.Tetromino{
				{
					Shape:  []string{"AAA"},
					Letter: 'A',
				},
			},
			wantSolved: false,
		},
		{
			name:      "multiple tetrominos",
			boardSize: 4,
			tetrominos: []tetromino.Tetromino{
				{
					Shape:  []string{"AA", "A."},
					Letter: 'A',
				},
				{
					Shape:  []string{"BB", "BB"},
					Letter: 'B',
				},
			},
			wantSolved: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := board.New(tt.boardSize)
			got := Solve(b, tt.tetrominos)
			if (got != nil) != tt.wantSolved {
				t.Errorf("Solve() = %v, wantSolved %v", got != nil, tt.wantSolved)
			}
		})
	}
}