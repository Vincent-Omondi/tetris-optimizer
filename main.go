package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strings"
)

type Tetromino struct {
	shape  []string
	letter rune
}

const (
	maxFileSize   = 1024 * 1024 // 1MB max file size
	maxBoardSize  = 20          // Reasonable max board size
	tetrominoSize = 4           // Size of each tetromino
	maxTetrominos = 26          // Max number of tetrominos (A-Z)
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<tetrominos_file.txt>")
		os.Exit(1)
	}

	tetrominos, err := readTetrominoes()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	size := int(math.Ceil(math.Sqrt(float64(len(tetrominos) * 4))))

	for size <= maxBoardSize {
		if board := solve(createBoard(size), tetrominos); board != nil {
			printBoard(board)
			return
		}
		size++
	}
	fmt.Println("ERROR: no solution found")
}

func readTetrominoes() ([]Tetromino, error) {
	// Check arguments first
	if len(os.Args) < 2 {
		return nil, fmt.Errorf("usage: %s <tetrominos_file.txt>", os.Args[0])
	}

	filename := os.Args[1]

	// File type check
	if !strings.HasSuffix(strings.ToLower(filename), ".txt") {
		return nil, fmt.Errorf("invalid file type: must be .txt")
	}

	// File stats check
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	if fileInfo.Size() > maxFileSize {
		return nil, fmt.Errorf("file too large (max %d bytes)", maxFileSize)
	}
	if !fileInfo.Mode().IsRegular() {
		return nil, fmt.Errorf("not a regular file")
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Check for non-ASCII characters
	for _, b := range content {
		if b > 127 {
			return nil, fmt.Errorf("file contains invalid characters (non-ASCII)")
		}
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

	var tetrominos []Tetromino
	for i, block := range blocks {
		lines := strings.Split(strings.TrimSpace(block), "\n")
		if !isValidTetromino(lines) {
			return nil, fmt.Errorf("invalid tetromino at position %d", i+1)
		}

		shape := make([]string, len(lines))
		letter := rune('A' + i)
		for j, line := range lines {
			shape[j] = strings.ReplaceAll(line, "#", string(letter))
		}
		tetrominos = append(tetrominos, Tetromino{shape: trimTetromino(shape), letter: letter})
	}

	return tetrominos, nil
}

func isValidTetromino(lines []string) bool {
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

	connectedBlocks := floodFill(lines, visited, startX, startY)

	// All blocks must be connected and count must match
	return connectedBlocks == 4
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

func trimTetromino(shape []string) []string {
	// Find bounds
	minX, maxX, minY, maxY := 4, -1, 4, -1
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

func createBoard(size int) [][]string {
	board := make([][]string, size)
	for i := range board {
		board[i] = make([]string, size)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	return board
}

func solve(board [][]string, tetrominos []Tetromino) [][]string {
	if len(tetrominos) == 0 {
		return board
	}

	current := tetrominos[0]
	for y := range board {
		for x := range board[y] {
			if canPlace(board, current.shape, x, y) {
				place(board, current.shape, x, y)
				if result := solve(board, tetrominos[1:]); result != nil {
					return result
				}
				remove(board, current.shape, x, y)
			}
		}
	}
	return nil
}

func canPlace(board [][]string, shape []string, x, y int) bool {
	for dy := range shape {
		for dx, char := range shape[dy] {
			if char != '.' {
				if y+dy >= len(board) || x+dx >= len(board[0]) || board[y+dy][x+dx] != "." {
					return false
				}
			}
		}
	}
	return true
}

func place(board [][]string, shape []string, x, y int) {
	for dy := range shape {
		for dx, char := range shape[dy] {
			if char != '.' {
				board[y+dy][x+dx] = string(char)
			}
		}
	}
}

func remove(board [][]string, shape []string, x, y int) {
	for dy := range shape {
		for dx, char := range shape[dy] {
			if char != '.' {
				board[y+dy][x+dx] = "."
			}
		}
	}
}

func printBoard(board [][]string) {
	if board == nil {
		return
	}

	for _, row := range board {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
	fmt.Println()
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
