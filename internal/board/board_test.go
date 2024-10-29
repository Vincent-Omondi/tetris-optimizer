package board

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		wantSize int
	}{
		{"size 3", 3, 3},
		{"size 5", 5, 5},
		{"size 0", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := New(tt.size)
			if len(board) != tt.wantSize {
				t.Errorf("New(%d) got size %d, want %d", tt.size, len(board), tt.wantSize)
			}

			// Check initialization
			for i := range board {
				if len(board[i]) != tt.wantSize {
					t.Errorf("New(%d) row %d got size %d, want %d", tt.size, i, len(board[i]), tt.wantSize)
				}
				for j := range board[i] {
					if board[i][j] != "." {
						t.Errorf("New(%d) position [%d][%d] = %s, want .", tt.size, i, j, board[i][j])
					}
				}
			}
		})
	}
}

func TestCanPlace(t *testing.T) {
	board := New(4)
	tests := []struct {
		name  string
		shape []string
		x, y  int
		want  bool
		setup func() // Optional setup function to prepare board state
	}{
		{
			name:  "valid placement empty board",
			shape: []string{"AA", "A."},
			x:     0,
			y:     0,
			want:  true,
		},
		{
			name:  "out of bounds right",
			shape: []string{"AA", "A."},
			x:     3,
			y:     0,
			want:  false,
		},
		{
			name:  "out of bounds bottom",
			shape: []string{"AA", "A."},
			x:     0,
			y:     3,
			want:  false,
		},
		{
			name:  "overlapping piece",
			shape: []string{"AA", "A."},
			x:     0,
			y:     0,
			want:  true,
			setup: func() {
				board[0][0] = "B"
			},
		},
		{
			name:  "valid placement with nearby pieces",
			shape: []string{"AA", "A."},
			x:     1,
			y:     1,
			want:  true,
			setup: func() {
				board[0][0] = "B"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := New(4) // Fresh board for each test
			if tt.setup != nil {
				tt.setup()
			}

			got := board.CanPlace(tt.shape, tt.x, tt.y)
			if got != tt.want {
				t.Errorf("CanPlace(%v, %d, %d) = %v, want %v", tt.shape, tt.x, tt.y, got, tt.want)
			}
		})
	}
}

func TestPlace(t *testing.T) {
	tests := []struct {
		name      string
		boardSize int
		shape     []string
		x, y      int
		want      string
	}{
		{
			name:      "simple placement",
			boardSize: 4,
			shape:     []string{"AA", "A."},
			x:         0,
			y:         0,
			want:      "AA..\nA...\n....\n....\n",
		},
		{
			name:      "corner placement",
			boardSize: 3,
			shape:     []string{"B"},
			x:         2,
			y:         2,
			want:      "...\n...\n..B\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := New(tt.boardSize)
			board.Place(tt.shape, tt.x, tt.y)

			var buf bytes.Buffer
			for i := range board {
				for j := range board[i] {
					buf.WriteString(board[i][j])
				}
				buf.WriteString("\n")
			}

			if got := buf.String(); got != tt.want {
				t.Errorf("Place(%v, %d, %d) produced board:\n%s\nwant:\n%s", tt.shape, tt.x, tt.y, got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name      string
		boardSize int
		shape     []string
		x, y      int
	}{
		{
			name:      "remove from corner",
			boardSize: 4,
			shape:     []string{"AA", "A."},
			x:         0,
			y:         0,
		},
		{
			name:      "remove single piece",
			boardSize: 3,
			shape:     []string{"B"},
			x:         1,
			y:         1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := New(tt.boardSize)
			// Place and then remove
			board.Place(tt.shape, tt.x, tt.y)
			board.Remove(tt.shape, tt.x, tt.y)

			// Verify board is empty
			for i := range board {
				for j := range board[i] {
					if board[i][j] != "." {
						t.Errorf("After Remove(%v, %d, %d), position [%d][%d] = %s, want .",
							tt.shape, tt.x, tt.y, i, j, board[i][j])
					}
				}
			}
		})
	}
}

