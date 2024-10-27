# Tetris Fit Solver

A Go program that solves the tetromino fitting puzzle. Given a set of tetromino pieces, it finds the smallest square board where all pieces can fit together without overlap.

## Description

This program takes a text file containing tetromino pieces and:
1. Validates the input pieces
2. Finds the smallest possible square board that can fit all pieces
3. Outputs the solution as a grid where each piece is represented by a letter (A-Z)

### What is a Tetromino?
A tetromino is a geometric shape composed of four squares, connected along their edges. These are the same pieces used in the classic game Tetris.

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/tetris-fit
cd tetris-fit

# Build the program
go build
```

## Usage

```bash
# Run the program with an input file
./tetris-fit input.txt
```

### Input File Format
- The input file must be a .txt file
- Each tetromino is represented by a 4x4 grid
- Use '#' for blocks and '.' for empty spaces
- Separate multiple tetrominoes with a blank line
- Maximum 26 tetrominoes (A-Z)

Example input.txt:
```
....
.##.
.##.
....

####
....
....
....
```

### Output Format
The program outputs a square grid where:
- Each tetromino is represented by a letter (A-Z)
- '.' represents empty spaces
- The grid is the smallest possible square that fits all pieces

Example output:
```
AABB
AABB
```

## Constraints and Validation

- Input file must be:
  - A valid .txt file
  - Less than 1MB in size
  - Contains only ASCII characters
- Each tetromino must:
  - Be 4x4 in size
  - Contain exactly 4 blocks
  - Be fully connected (no floating blocks)
  - Use only '#' and '.' characters
- Maximum 26 tetrominoes per input file
- Maximum board size is 20x20

## Error Messages

The program will display "ERROR" followed by a description when it encounters:
- Invalid file type or format
- Invalid tetromino shapes
- Too many tetrominoes
- No solution possible
- Missing or invalid arguments

## Running Tests

```bash
go test
```

The test suite includes:
- Input validation
- Tetromino validation
- Board operations
- Solution finding
- Edge cases
- File handling

## Examples

### Valid Tetromino Shapes:
```
# I-piece
####
....
....
....

# L-piece
#...
#...
##..
....

# Square
##..
##..
....
....
```

### Invalid Examples:
```
# Disconnected pieces (invalid)
##..
....
..##
....

# Wrong number of blocks (invalid)
###.
....
....
....
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Based on the classic tetromino puzzle problem
- Inspired by the game of Tetris