package tetromino

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		wantErr  bool
		errMsg   string
		expected int // expected number of tetrominos
	}{
		{
			name: "valid single tetromino",
			content: `....
##..
.#..
.#..
`,
			wantErr:  false,
			expected: 1,
		},
		{
			name: "valid multiple tetrominos",
			content: `....
##..
.#..
.#..

..##
..##
....
....
`,
			wantErr:  false,
			expected: 2,
		},
		{
			name: "invalid shape",
			content: `....
##..
....
....
`,
			wantErr: true,
			errMsg:  "invalid tetromino at position 1",
		},
		{
			name: "too many tetrominos",
			content: strings.Repeat(`....
##..
.#..
.#..

`, 27),
			wantErr: true,
			errMsg:  "too many tetrominos",
		},
		{
			name: "non-ASCII characters",
			content: `....
##..
.#£.
.#..
`,
			wantErr: true,
			errMsg:  "invalid characters",
		},
		{
			name: "invalid size",
			content: `...
##.
.#.
`,
			wantErr: true,
			errMsg:  "invalid tetromino",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary file
			tmpDir := t.TempDir()
			tmpFile := filepath.Join(tmpDir, "test.txt")
			if err := os.WriteFile(tmpFile, []byte(tt.content), 0644); err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}

			tetrominos, err := ReadFromFile(tmpFile)
			if tt.wantErr {
				if err == nil {
					t.Errorf("ReadFromFile() expected error containing %q, got nil", tt.errMsg)
				} else if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("ReadFromFile() error = %v, want error containing %q", err, tt.errMsg)
				}
				return
			}

			if err != nil {
				t.Fatalf("ReadFromFile() unexpected error: %v", err)
			}

			if len(tetrominos) != tt.expected {
				t.Errorf("ReadFromFile() got %d tetrominos, want %d", len(tetrominos), tt.expected)
			}

			// Verify each tetromino has proper letter assignment
			for i, tetromino := range tetrominos {
				expectedLetter := rune('A' + i)
				if tetromino.Letter != expectedLetter {
					t.Errorf("Tetromino %d has letter %c, want %c", i, tetromino.Letter, expectedLetter)
				}
			}
		})
	}
}

func TestTrimTetromino(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name: "no trimming needed",
			input: []string{
				"AA",
				"A.",
			},
			want: []string{
				"AA",
				"A.",
			},
		},
		{
			name: "trim all sides",
			input: []string{
				"....",
				".AA.",
				".A..",
				"....",
			},
			want: []string{
				"AA",
				"A.",
			},
		},
		{
			name: "trim right and bottom",
			input: []string{
				"AA..",
				"A...",
				"....",
				"....",
			},
			want: []string{
				"AA",
				"A.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trimTetromino(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimTetromino() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinMax(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		wantMin int
		wantMax int
	}{
		{"positive numbers", 5, 3, 3, 5},
		{"negative numbers", -2, -5, -5, -2},
		{"equal numbers", 4, 4, 4, 4},
		{"zero and positive", 0, 7, 0, 7},
		{"negative and positive", -3, 2, -3, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin := min(tt.a, tt.b)
			if gotMin != tt.wantMin {
				t.Errorf("min(%d, %d) = %d, want %d", tt.a, tt.b, gotMin, tt.wantMin)
			}

			gotMax := max(tt.a, tt.b)
			if gotMax != tt.wantMax {
				t.Errorf("max(%d, %d) = %d, want %d", tt.a, tt.b, gotMax, tt.wantMax)
			}
		})
	}
}