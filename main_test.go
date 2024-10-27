package main

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// Helper function to create temporary test files
func createTempFile(content string) (string, error) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "tetris-test-*")
	if err != nil {
		return "", err
	}

	// Create a file inside the temporary directory
	tmpFilePath := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(tmpFilePath, []byte(content), 0644); err != nil {
		os.RemoveAll(tmpDir) // Clean up directory on error
		return "", err
	}

	return tmpFilePath, nil
}

// Clean up helper function
func cleanupTempFile(path string) {
	if path != "" {
		os.RemoveAll(filepath.Dir(path)) // Remove the entire temp directory
	}
}

// Test cases for isValidTetromino
func TestIsValidTetromino(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected bool
	}{
		{
			name: "Valid I tetromino",
			input: []string{
				"#...",
				"#...",
				"#...",
				"#...",
			},
			expected: true,
		},
		{
			name: "Valid square tetromino",
			input: []string{
				"....",
				".##.",
				".##.",
				"....",
			},
			expected: true,
		},
		{
			name: "Valid L tetromino",
			input: []string{
				"#...",
				"#...",
				"##..",
				"....",
			},
			expected: true,
		},
		{
			name: "Invalid size",
			input: []string{
				"#...",
				"#...",
				"##..",
			},
			expected: false,
		},
		{
			name: "Invalid character",
			input: []string{
				"#...",
				"#...",
				"#X..",
				"....",
			},
			expected: false,
		},
		{
			name: "Invalid block count",
			input: []string{
				"#...",
				"#...",
				"#...",
				"....",
			},
			expected: false,
		},
		{
			name: "Invalid connections",
			input: []string{
				"#...",
				"....",
				"..#.",
				"..##",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidTetromino(tt.input)
			if result != tt.expected {
				t.Errorf("isValidTetromino() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Test cases for trimTetromino
func TestTrimTetromino(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name: "Already trimmed",
			input: []string{
				"AA",
				"AA",
			},
			expected: []string{
				"AA",
				"AA",
			},
		},
		{
			name: "Need trimming",
			input: []string{
				"....",
				".AA.",
				".AA.",
				"....",
			},
			expected: []string{
				"AA",
				"AA",
			},
		},
		{
			name: "Single column",
			input: []string{
				".A..",
				".A..",
				".A..",
				".A..",
			},
			expected: []string{
				"A",
				"A",
				"A",
				"A",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := trimTetromino(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("trimTetromino() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Test cases for readTetrominoes
func TestReadTetrominoes(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{
			name: "Valid single tetromino",
			input: `....
....
..##
..##`,
			expectError: false,
		},
		{
			name: "Valid multiple tetrominoes",
			input: `....
....
..##
..##

.##.
.##.
....
....`,
			expectError: false,
		},
		{
			name: "Invalid character",
			input: `....
..X.
..##
..##`,
			expectError: true,
		},
		{
			name: "Invalid size",
			input: `...
...
..##`,
			expectError: true,
		},
		{
			name: "Too many tetrominoes",
			input: strings.Repeat(`....
....
..##
..##

`, 27),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original args and restore after test
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			// Create temporary file
			tmpfile, err := createTempFile(tt.input)
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer cleanupTempFile(tmpfile) // Clean up after test

			// Set up args for test
			os.Args = []string{"cmd", tmpfile}

			_, err = readTetrominoes()
			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

// Test cases for board operations
func TestBoardOperations(t *testing.T) {
	// Test createBoard
	t.Run("CreateBoard", func(t *testing.T) {
		size := 3
		board := createBoard(size)
		if len(board) != size {
			t.Errorf("Board height = %d, want %d", len(board), size)
		}
		for i := range board {
			if len(board[i]) != size {
				t.Errorf("Board width at row %d = %d, want %d", i, len(board[i]), size)
			}
			for j := range board[i] {
				if board[i][j] != "." {
					t.Errorf("Board cell at (%d,%d) = %s, want .", i, j, board[i][j])
				}
			}
		}
	})

	// Test canPlace, place, and remove
	t.Run("PlaceAndRemove", func(t *testing.T) {
		board := createBoard(4)
		shape := []string{
			"AA",
			"AA",
		}

		// Test placement
		if !canPlace(board, shape, 0, 0) {
			t.Error("canPlace() returned false, want true")
		}

		place(board, shape, 0, 0)
		expected := [][]string{
			{"A", "A", ".", "."},
			{"A", "A", ".", "."},
			{".", ".", ".", "."},
			{".", ".", ".", "."},
		}
		if !reflect.DeepEqual(board, expected) {
			t.Errorf("After place(), board = %v, want %v", board, expected)
		}

		// Test removal
		remove(board, shape, 0, 0)
		emptyBoard := createBoard(4)
		if !reflect.DeepEqual(board, emptyBoard) {
			t.Errorf("After remove(), board = %v, want %v", board, emptyBoard)
		}
	})
}

// Test full solution
// Test full solution
func TestSolve(t *testing.T) {
	t.Run("Simple Solution", func(t *testing.T) {
		tetrominos := []Tetromino{
			{
				shape: []string{
					"AA",
					"AA",
				},
				letter: 'A',
			},
			{
				shape: []string{
					"BB",
					"BB",
				},
				letter: 'B',
			},
		}

		// Create a 4x4 board which is big enough for both pieces
		board := createBoard(4)
		result := solve(board, tetrominos)

		if result == nil {
			t.Error("solve() returned nil, expected solution")
		} else {
			// Verify that both pieces are present
			foundA := false
			foundB := false
			for i := range result {
				for j := range result[i] {
					if result[i][j] == "A" {
						foundA = true
					}
					if result[i][j] == "B" {
						foundB = true
					}
				}
			}
			if !foundA || !foundB {
				t.Errorf("Solution missing pieces. Found A: %v, Found B: %v", foundA, foundB)
			}
		}
	})

	t.Run("L Shape Solution", func(t *testing.T) {
		tetrominos := []Tetromino{
			{
				shape: []string{
					"A..",
					"AAA",
				},
				letter: 'A',
			},
		}

		board := createBoard(3)
		result := solve(board, tetrominos)

		if result == nil {
			t.Error("solve() returned nil, expected solution")
		}
	})

	t.Run("Multiple Shapes", func(t *testing.T) {
		tetrominos := []Tetromino{
			{
				shape: []string{
					"AA",
					"A.",
					"A.",
				},
				letter: 'A',
			},
			{
				shape: []string{
					"BB",
					"BB",
				},
				letter: 'B',
			},
		}

		board := createBoard(4)
		result := solve(board, tetrominos)

		if result == nil {
			t.Error("solve() returned nil, expected solution")
		}
	})

	t.Run("Impossible Solution", func(t *testing.T) {
		tetrominos := []Tetromino{
			{
				shape: []string{
					"AAA",
					"AAA",
				},
				letter: 'A',
			},
		}

		board := createBoard(2) // Too small for the piece
		result := solve(board, tetrominos)

		if result != nil {
			t.Error("solve() returned solution for impossible case")
		}
	})

	t.Run("Empty Tetrominos", func(t *testing.T) {
		board := createBoard(2)
		result := solve(board, []Tetromino{})

		if result == nil {
			t.Error("solve() returned nil for empty tetrominos")
		}
	})

	t.Run("Single Square", func(t *testing.T) {
		tetrominos := []Tetromino{
			{
				shape: []string{
					"AA",
					"AA",
				},
				letter: 'A',
			},
		}

		board := createBoard(2)
		result := solve(board, tetrominos)

		expected := [][]string{
			{"A", "A"},
			{"A", "A"},
		}

		if result == nil {
			t.Error("solve() returned nil, expected solution")
		} else if !reflect.DeepEqual(result, expected) {
			t.Errorf("solve() returned %v, want %v", result, expected)
		}
	})
}

// Add a test for edge case with maximum board size
func TestSolveMaxBoardSize(t *testing.T) {
	// Create multiple tetrominos that require a larger board
	var tetrominos []Tetromino
	for i := 0; i < 4; i++ {
		tetrominos = append(tetrominos, Tetromino{
			shape: []string{
				string(rune('A' + i)),
				string(rune('A' + i)),
				string(rune('A' + i)),
				string(rune('A' + i)),
			},
			letter: rune('A' + i),
		})
	}

	board := createBoard(8) // Should be large enough for 4 I-shaped pieces
	result := solve(board, tetrominos)

	if result == nil {
		t.Error("solve() returned nil for valid large board")
	}
}

// Add a test for invalid piece placement
func TestInvalidPiecePlacement(t *testing.T) {
	tetrominos := []Tetromino{
		{
			shape: []string{
				"AA",
				"AA",
			},
			letter: 'A',
		},
		{
			shape: []string{
				"BB",
				"BB",
			},
			letter: 'B',
		},
	}

	// Create a board that's too narrow
	board := createBoard(1)
	result := solve(board, tetrominos)

	if result != nil {
		t.Error("solve() returned solution for invalid board size")
	}
}

// Test invalid file path
func TestInvalidFilePath(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "nonexistent-file.txt"}
	_, err := readTetrominoes()
	if err == nil {
		t.Error("Expected error for nonexistent file, got nil")
	}
}

// Test invalid argument count
func TestInvalidArgumentCount(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd"} // Missing filename argument
	_, err := readTetrominoes()
	if err == nil {
		t.Error("Expected error for missing argument, got nil")
	}
}
