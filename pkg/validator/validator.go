package validator

import (
	"fmt"
	"os"
	"strings"
)

const tetrominoSize = 4 // Size of each tetromino

// ValidateFile performs basic file validation
func ValidateFile(filename string, maxSize int64) error {
	// File type check
	if !strings.HasSuffix(strings.ToLower(filename), ".txt") {
		return fmt.Errorf("invalid file type: must be .txt")
	}

	// File stats check
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return err
	}
	if fileInfo.Size() > maxSize {
		return fmt.Errorf("file too large (max %d bytes)", maxSize)
	}
	if !fileInfo.Mode().IsRegular() {
		return fmt.Errorf("not a regular file")
	}

	return nil
}

// IsASCII checks if all bytes are ASCII characters
func IsASCII(content []byte) bool {
	for _, b := range content {
		if b > 127 {
			return false
		}
	}
	return true
}

// IsValidTetromino validates the shape and connectivity of a tetromino
func IsValidTetromino(lines []string) bool {
	// Check basic dimensions
	if len(lines) != tetrominoSize {
		return false
	}

	// Count blocks and find start position
	blocks := 0
	startX, startY := -1, -1

	for i, line := range lines {
		// Check line length
		if len(line) != tetrominoSize {
			return false
		}

		// Check each character and count blocks
		for j, char := range line {
			if char != '.' && char != '#' {
				return false
			}
			if char == '#' {
				blocks++
				if startX == -1 {
					startX, startY = j, i
				}
			}
		}
	}

	// Must have exactly 4 blocks
	if blocks != 4 || startX == -1 {
		return false
	}

	// Verify connectivity using flood fill
	visited := make([][]bool, tetrominoSize)
	for i := range visited {
		visited[i] = make([]bool, tetrominoSize)
	}

	return floodFill(lines, visited, startX, startY) == 4
}

func floodFill(lines []string, visited [][]bool, x, y int) int {
	// Check bounds and validity
	if x < 0 || y < 0 || x >= tetrominoSize || y >= tetrominoSize ||
		visited[y][x] || lines[y][x] != '#' {
		return 0
	}

	// Mark as visited
	visited[y][x] = true
	count := 1

	// Check all four adjacent positions
	directions := [][2]int{
		{0, 1},  // down
		{0, -1}, // up
		{1, 0},  // right
		{-1, 0}, // left
	}

	for _, dir := range directions {
		count += floodFill(lines, visited, x+dir[0], y+dir[1])
	}

	return count
}