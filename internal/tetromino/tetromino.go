package tetromino

import (
	"bytes"
	"fmt"
	"github/Vincent-Omondi/tetris-optimizer/pkg/validator"
	"os"
	"strings"
)

const (
	maxFileSize   = 1024 * 1024 // 1MB max file size
	tetrominoSize = 4           // Size of each tetromino
	maxTetrominos = 26          // Max number of tetrominos (A-Z)
)

// Tetromino represents a single tetromino piece
type Tetromino struct {
	Shape  []string
	Letter rune
}

// ReadFromFile reads and validates tetrominos from a file
func ReadFromFile(filename string) ([]Tetromino, error) {
	// File validation
	if err := validator.ValidateFile(filename, maxFileSize); err != nil {
		return nil, err
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Check for non-ASCII characters
	if !validator.IsASCII(content) {
		return nil, fmt.Errorf("file contains invalid characters (non-ASCII)")
	}

	// Normalize line endings
	content = bytes.ReplaceAll(content, []byte("\r\n"), []byte("\n"))

	blocks := strings.Split(strings.TrimSpace(string(content)), "\n\n")
	if len(blocks) == 0 {
		return nil, fmt.Errorf("no tetrominos found")
	}
	if len(blocks) > maxTetrominos {
		return nil, fmt.Errorf("too many tetrominos (max %d)", maxTetrominos)
	}

	return parseBlocks(blocks)
}

func parseBlocks(blocks []string) ([]Tetromino, error) {
	var tetrominos []Tetromino
	for i, block := range blocks {
		lines := strings.Split(strings.TrimSpace(block), "\n")
		if !validator.IsValidTetromino(lines) {
			return nil, fmt.Errorf("invalid tetromino at position %d", i+1)
		}

		shape := make([]string, len(lines))
		letter := rune('A' + i)
		for j, line := range lines {
			shape[j] = strings.ReplaceAll(line, "#", string(letter))
		}
		tetrominos = append(tetrominos, Tetromino{
			Shape:  trimTetromino(shape),
			Letter: letter,
		})
	}
	return tetrominos, nil
}

func trimTetromino(shape []string) []string {
	// Find bounds
	minX, maxX, minY, maxY := tetrominoSize, -1, tetrominoSize, -1
	for y, line := range shape {
		for x, char := range line {
			if char != '.' {
				minX = min(minX, x)
				maxX = max(maxX, x)
				minY = min(minY, y)
				maxY = max(maxY, y)
			}
		}
	}

	// Trim to bounds
	result := make([]string, maxY-minY+1)
	for i := range result {
		result[i] = shape[i+minY][minX : maxX+1]
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}