package main

import (
	"fmt"
	"AoC24/utils"
)

func markThePath(matrix [][]rune) bool {
	isOut := false
	cursor := []rune{'^','>','v','<'}
	for !isOut {
		MainLoop:
		for i, row := range matrix {
			for j := range row {
				if utils.ContainRune(cursor, matrix[i][j]) {
					if matrix[i][j] == '^' {
						if i != 0 {
							if matrix[i-1][j] == '#' {
								matrix[i][j] = '>'
							} else {
								matrix[i-1][j] = '^'
								matrix[i][j] = 'X'
							}
						} else {
							matrix[i][j] = 'X'
							isOut = true
							break MainLoop
						}
					} else if matrix[i][j] == '>' {
						if j != len(row)-1 {
							if matrix[i][j+1] == '#' {
								matrix[i][j] = 'v'
							} else {
								matrix[i][j+1] = '>'
								matrix[i][j] = 'X'
							}
						} else {
							matrix[i][j] = 'X'
							isOut = true
							break MainLoop
						}
					} else if matrix[i][j] == 'v' {
						if i != len(matrix)-1 {
							if matrix[i+1][j] == '#' {
								matrix[i][j] = '<'
							} else {
								matrix[i+1][j] = 'v'
								matrix[i][j] = 'X'
							}
						} else {
							matrix[i][j] = 'X'
							isOut = true
							break MainLoop
						}
					} else if matrix[i][j] == '<' {
						if j != 0 {
							if matrix[i][j-1] == '#' {
								matrix[i][j] = '^'
							} else {
								matrix[i][j-1] = '<'
								matrix[i][j] = 'X'
							}
						} else {
							matrix[i][j] = 'X'
							isOut = true
							break MainLoop
						}
					}
				}
			}
		}
	}
	return isOut
}

func sumOfMarks(matrix [][]rune) int {
	sum := 0
	for i, row := range matrix {
		for j := range row {
			if matrix[i][j] == 'X' {
				sum++
			}
		}
	}
	return sum
}

func main() {
	content := utils.FileToStr("input")
	matrix := utils.StrToMatrix(content)
	markThePath(matrix)
	sum := sumOfMarks(matrix)
	fmt.Println(sum)
}
