package main

import (
	"AoC24/utils"
	"fmt"
	// "os"
)

func strToMatrix(str string) [][]rune {
	matrix := [][]rune{{},}
	row, col := 0, 0
	for i, char := range str {
		if char == '\n' {
			if i != len(str) - 1 {
				matrix = append(matrix, []rune{})
			}
			row++
			col = 0
		} else {
			matrix[row] = append(matrix[row], char)
			col++
		}
	}
	return matrix
}

func countWords(matrix [][]rune) int {
	count := 0
	xmas := []rune{'X', 'M', 'A', 'S'}
	for i, row := range matrix {
		for j, char := range row {
			if char == 'X' {
				// UP
				if i >= 3 {
					for k := 1; k < 4; k++ {
						if matrix[i-k][j] != xmas[k] { break }
						if k == 3 { count++ }
					}
				}
				// UP RIGHT
				if i >= 3 && j <= len(row) - 4 {
					for k := 1; k < 4; k++ {
						if matrix[i-k][j+k] != xmas[k] { break }
						if k == 3 { count++ }
					}
				}
				// RIGHT
				if j <= len(row) - 4 {
					for k := 1; k < 4; k++ {
						if matrix[i][j+k] != xmas[k] { break }
						if k == 3 { count++ }
					}
				}
				// DOWN RIGHT
				if i <= len(matrix) - 4 && j <= len(row) - 4 {
					for k := 1; k < 4; k++ {
						if matrix[i+k][j+k] != xmas[k] { break }
						if k == 3 { count++ }
					}
				}
				// DOWN
				if i <= len(matrix) - 4 {
					for k := 1; k < 4; k++ {
						if matrix[i+k][j] != xmas[k] { break }
						if k == 3 { count++ }
					}
				}
				// DOWN LEFT
				if i <= len(matrix) - 4 && j >= 3 {
					for k := 1; k < 4; k++ {
						if matrix[i+k][j-k] != xmas[k] { break }
						if k == 3 { count++ }
					}
				}
				// LEFT
				if j >= 3 {
					for k := 1; k < 4; k++ {
						if matrix[i][j-k] != xmas[k] { break }
						if k == 3 { count++ }
					}
				}
				// UP LEFT
				if i >= 3 && j >= 3 {
					for k := 1; k < 4; k++ {
						if matrix[i-k][j-k] != xmas[k] { break }
						if k == 3 { count++ }
					}
				}
			}
		}
	}
	return count
}

func countXmas(matrix [][]rune) int {
	count := 0
	for i := 0; i < len(matrix) - 2; i++ {
		for j := 0; j < len(matrix[i]) - 2; j++ {
			if (matrix[i][j] == 'S' || matrix[i][j] == 'M') && matrix[i+1][j+1] == 'A' {
				if ((matrix[i][j] == 'S' && matrix[i+2][j+2] == 'M') ||
				(matrix[i][j] == 'M' && matrix[i+2][j+2] == 'S')) && 
				((matrix[i][j+2] == 'S' && matrix[i+2][j] == 'M') ||
				(matrix[i][j+2] == 'M' && matrix[i+2][j] == 'S')) {
					count++ 
				}
			}
		}
	}
	return count
}

func main() {
	content := utils.FileToStr("input")
	matrix := strToMatrix(content)
/*
	file, _ := os.OpenFile("test", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	for _, row := range matrix {
		file.WriteString(fmt.Sprint(row) + "\n")
	}
*/
	fmt.Println(countWords(matrix))
	fmt.Println(countXmas(matrix))
}
