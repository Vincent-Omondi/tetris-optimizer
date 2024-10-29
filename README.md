
# Tetris Optimizer

## Overview

Tetris Optimizer is a command-line program written in Go that arranges tetrominoes in the smallest square possible. Given a text file containing a list of tetrominoes, it uses backtracking to find an optimal arrangement, ensuring each tetromino is uniquely identifiable and follows proper validation checks. This project demonstrates the application of algorithms, file handling, and problem-solving within Go's standard library constraints.

## Objectives

1. Parse a file containing tetromino shapes, specified as `#` for filled cells and `.` for empty cells.
2. Arrange tetrominoes within the smallest square possible.
3. Display each tetromino using uppercase letters (e.g., A for the first tetromino, B for the second).
4. Ensure the program exits with an error message if:
   - The file format or tetromino format is invalid.
   - It cannot form a complete square and leaves spaces between tetrominoes.

## Features

- **Optimal Arrangement**: Using backtracking, the program places tetrominoes in the smallest square possible.
- **Tetromino Identification**: Each tetromino is assigned an uppercase Latin letter (A-Z) for clarity.
- **Error Handling**: Displays an error message for invalid input or when placement is impossible.
- **File Validation**: Ensures correct format, ASCII-only characters, and appropriate size limits.

## Requirements

- **Go 1.22+**
- **Standard Library**: The project only uses standard Go packages.

## Usage

### Running the Program

To execute the program, run the following command in the terminal:

```sh
go run . <path/to/tetrominoes_file.txt>
```

### Example Usage

```sh
$ go run . sample.txt
```

#### Sample Input File (`sample.txt`)

Below is an example format for the tetrominoes file:

```
...#
...#
...#
...#

....
....
....
####

.###
...#
....
....

....
..##
.##.
....

....
.##.
.##.
....

....
....
##..
.##.

##..
.#..
.#..
....

....
###.
.#..
....
```

### Expected Output

If the tetrominoes can be arranged, the output will display each tetromino with an assigned uppercase letter within the smallest square possible:

```
ABBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```

In case of errors in formatting or file issues, the output will display:

```sh
ERROR
```

## Project Structure

- **`main.go`**: Entry point of the program. It reads the input file and invokes the solver to arrange tetrominoes.
- **`internal/board`**: Defines the `Board` structure, responsible for creating, placing, and removing tetrominoes.
- **`internal/solver`**: Implements the recursive solver for arranging tetrominoes.
- **`internal/tetromino`**: Handles tetromino parsing and format validation.
- **`pkg/validator`**: Provides validation for file format, file size, ASCII-only content, and tetromino shape connectivity.
- **`sample.txt`**: Example tetromino file for testing purposes.

## How It Works

1. **File Validation**: The program checks the file for size, format, and ASCII-only content.
2. **Parsing Tetrominoes**: Reads tetromino shapes, validates them, and assigns a unique letter.
3. **Square Calculation**: Determines the minimum square size based on the number of tetrominoes.
4. **Backtracking Algorithm**: Attempts to place each tetromino in the smallest possible square. If it fails, it increases the square size until a solution is found or the maximum board size is exceeded.

## Development and Testing

### Running Tests

The project includes unit tests located within the `internal/board`, `internal/solver`, and `internal/tetromino` directories. To run all tests, execute:

```sh
go test ./...
```

### Adding Test Cases

Test cases can be added to validate edge cases and additional tetromino configurations in the respective `*_test.go` files.

## Limitations

- Only accepts files up to 1MB.
- Supports up to 26 tetrominoes (A-Z).
- Spaces may remain if a complete square formation is impossible.

## Contributions

Contributions are welcome! Please fork the repository, make changes, and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

To enhance the project's capabilities, you may consider adding advanced error handling, additional testing scenarios, or expanding support for different board sizes.
