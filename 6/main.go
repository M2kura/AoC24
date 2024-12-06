package main

import (
	"fmt"
	"AoC24/utils"
)

func findCursor(matrix [][]rune) []int {
	cursor := []rune{'^','>','v','<'}
	for i, row := range matrix {
		for j := range row {
			if utils.ContainRune(cursor, matrix[i][j]) {
				return []int{i, j}
			}
		}
	}
	return nil
}

func makeMove(matrix [][]rune, cursorCord []int, movesCount *int) {
	cursorRune := matrix[cursorCord[0]][cursorCord[1]]
	if cursorRune == '^' {
		makeMoveUp(matrix, cursorCord[0], cursorCord[1], movesCount)
	} else if cursorRune == '>' {
		makeMoveRight(matrix, cursorCord[0], cursorCord[1], movesCount)
	} else if cursorRune == 'v' {
		makeMoveDown(matrix, cursorCord[0], cursorCord[1], movesCount)
	} else if cursorRune == '<' {
		makeMoveLeft(matrix, cursorCord[0], cursorCord[1], movesCount)
	}
}

func makeMoveUp(matrix [][]rune, x int, y int, movesCount *int) {
	if x != 0 && matrix[x-1][y] == '#' {
		matrix[x][y] = '>'
	} else {
		*movesCount++
		if x != 0 {
			matrix[x-1][y] = '^'
		}
		matrix[x][y] = 'X'
	}
}

func makeMoveRight(matrix [][]rune, x int, y int, movesCount *int) {
	if y != len(matrix[x])-1 && matrix[x][y+1] == '#' {
		matrix[x][y] = 'v'
	} else {
		*movesCount++
		if y != len(matrix[x])-1 {
			matrix[x][y+1] = '>'
		}
		matrix[x][y] = 'X'
	}
}

func makeMoveDown(matrix [][]rune, x int, y int, movesCount *int) {
	if x != len(matrix)-1 && matrix[x+1][y] == '#' {
		matrix[x][y] = '<'
	} else {
		*movesCount++
		if x != len(matrix)-1 {
			matrix[x+1][y] = 'v'
		}
		matrix[x][y] = 'X'
	}
}

func makeMoveLeft(matrix [][]rune, x int, y int, movesCount *int) {
	if y != 0 && matrix[x][y-1] == '#' {
		matrix[x][y] = '^'
	} else {
		*movesCount++
		if y != 0 {
			matrix[x][y-1] = '<'
		}
		matrix[x][y] = 'X'
	}
}

func markThePath(matrix [][]rune) int {
	cursor := findCursor(matrix)
	movesCount := 0
	for cursor != nil {
		makeMove(matrix, cursor, &movesCount)
		if movesCount > 10000 {
			return -1
		}
		cursor = findCursor(matrix)
	}
	return movesCount
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

func tryObstr(matrix [][]rune) int {
	count := 0
	for i, row := range matrix {
		for j := range row {
			if !(i == 79 && j == 87) && matrix[i][j] == 'X' {
				matrixCopy := utils.StrToMatrix(utils.FileToStr("input"))
				matrixCopy[i][j] = '#'
				steps := markThePath(matrixCopy)
				if steps == -1 {
					count++
				}
				fmt.Printf("For block on row %d, col %d - %d steps\n", i+1, j+1, steps)
			}
		}
	}
	return count
}

func main() {
	content := utils.FileToStr("input")
	matrix := utils.StrToMatrix(content)
	movesCount := markThePath(matrix)
	sum := sumOfMarks(matrix)
	fmt.Println(sum)
	fmt.Println("Move coutn deffault:", movesCount, "\n-----------")
	fmt.Println(tryObstr(matrix))
}
