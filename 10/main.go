package main

import (
	"fmt"
	"AoC24/utils"
)

func checkList(list *[][]int, x, y int) bool {
	for _, cords := range *list {
		if cords[0] == x && cords[1] == y { return false }
	}
	return true
}

func checkTrailhead(matrix [][]int, listOf9 *[][]int, x, y, num int) int {
	sum := 0
	if x != 0 && matrix[x-1][y] == num+1 {
		if matrix[x-1][y] == 9 && checkList(listOf9, x-1, y) {
			sum++
			*listOf9 = append(*listOf9, []int{x-1,y})
		} else {
			sum += checkTrailhead(matrix, listOf9, x-1, y, num+1)
		}
	}
	if y != len(matrix[0])-1 && matrix[x][y+1] == num+1 {
		if matrix[x][y+1] == 9 && checkList(listOf9, x, y+1) {
			sum++
			*listOf9 = append(*listOf9, []int{x,y+1})
		} else {
			sum += checkTrailhead(matrix, listOf9, x, y+1, num+1)
		}
	}
	if x != len(matrix)-1 && matrix[x+1][y] == num+1{
		if matrix[x+1][y] == 9 && checkList(listOf9, x+1, y) {
			sum++
			*listOf9 = append(*listOf9, []int{x+1,y})
		} else {
			sum += checkTrailhead(matrix, listOf9, x+1, y, num+1)
		}
	}
	if y != 0 && matrix[x][y-1] == num+1 {
		if matrix[x][y-1] == 9 && checkList(listOf9, x, y-1) {
			sum++
			*listOf9 = append(*listOf9, []int{x,y-1})
		} else {
			sum += checkTrailhead(matrix, listOf9, x, y-1, num+1)
		}
	}
	return sum
}

func checkTrailheadAll(matrix [][]int, x, y, num int) int {
	sum := 0
	if x != 0 && matrix[x-1][y] == num+1 {
		if matrix[x-1][y] == 9 {
			sum++
		} else {
			sum += checkTrailheadAll(matrix, x-1, y, num+1)
		}
	}
	if y != len(matrix[0])-1 && matrix[x][y+1] == num+1 {
		if matrix[x][y+1] == 9 {
			sum++
		} else {
			sum += checkTrailheadAll(matrix, x, y+1, num+1)
		}
	}
	if x != len(matrix)-1 && matrix[x+1][y] == num+1{
		if matrix[x+1][y] == 9 {
			sum++
		} else {
			sum += checkTrailheadAll(matrix, x+1, y, num+1)
		}
	}
	if y != 0 && matrix[x][y-1] == num+1 {
		if matrix[x][y-1] == 9 {
			sum++
		} else {
			sum += checkTrailheadAll(matrix, x, y-1, num+1)
		}
	}
	return sum
}

func findTrailheads(matrix [][]int, all bool) int {
	sum := 0
	for i, row := range matrix {
		for j := range row {
			if matrix[i][j] == 0 {
				if all {
					sum += checkTrailheadAll(matrix, i, j, 0)
				} else {
					listOf9 := [][]int{}
					sum += checkTrailhead(matrix, &listOf9, i, j, 0)
				}
			}
		}
	}
	return sum
}

func main() {
	content := utils.FileToStr("input")
	matrix := utils.StrToMatrixInt(content)
	trailheads := findTrailheads(matrix, false)
	trailheadsAll := findTrailheads(matrix, true)
	fmt.Println(trailheads)
	fmt.Println(trailheadsAll)
}
