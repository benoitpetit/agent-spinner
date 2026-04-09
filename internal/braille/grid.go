// Package braille implements braille character grid operations.
// This is an internal package - not part of the public API.
package braille

// DotMap maps grid positions to braille dot bit values.
// Row 0: dot1 (0x01), dot4 (0x08)
// Row 1: dot2 (0x02), dot5 (0x10)
// Row 2: dot3 (0x04), dot6 (0x20)
// Row 3: dot7 (0x40), dot8 (0x80)
var DotMap = [4][2]int{
	{0x01, 0x08},
	{0x02, 0x10},
	{0x04, 0x20},
	{0x40, 0x80},
}

// GridToBraille converts a 2D boolean grid into a braille string.
// grid[row][col] = true means dot is raised.
// Width must be even (2 dot-columns per braille char).
func GridToBraille(grid [][]bool) string {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return ""
	}
	rows, cols := len(grid), len(grid[0])
	charCount := (cols + 1) / 2
	result := make([]rune, 0, charCount)

	for c := 0; c < charCount; c++ {
		code := 0x2800
		for r := 0; r < 4 && r < rows; r++ {
			for d := 0; d < 2; d++ {
				col := c*2 + d
				if col < cols && r < len(grid) && col < len(grid[r]) && grid[r][col] {
					code |= DotMap[r][d]
				}
			}
		}
		result = append(result, rune(code))
	}
	return string(result)
}

// MakeGrid creates an empty grid of given dimensions.
// Returns nil for invalid dimensions.
func MakeGrid(rows, cols int) [][]bool {
	if rows <= 0 || cols <= 0 {
		return nil
	}
	grid := make([][]bool, rows)
	for i := range grid {
		grid[i] = make([]bool, cols)
	}
	return grid
}

// CloneGrid creates a deep copy of a grid.
func CloneGrid(grid [][]bool) [][]bool {
	if len(grid) == 0 {
		return nil
	}
	newGrid := make([][]bool, len(grid))
	for i, row := range grid {
		newGrid[i] = make([]bool, len(row))
		copy(newGrid[i], row)
	}
	return newGrid
}
