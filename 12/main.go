package main

import (
	"fmt"
	"AoC24/utils"
)

func getPlants(matrix [][]rune) []rune {
	plants := []rune{}
	for i, row := range matrix {
		for j, _ := range row {
			if !utils.ContainRune(plants, matrix[i][j]) {
				plants = append(plants, matrix[i][j])
			}
		}
	}
	return plants
}

func checkRegion(matrix [][]rune, x, y int, area, per *int) {
	plant := matrix[x][y]
	plantChecked := plant + ' '
	matrix[x][y] = plantChecked
	*area++
	if x > 0 {
		if matrix[x-1][y] == plant {
			checkRegion(matrix, x-1, y, area, per)
		} else if matrix[x-1][y] != plantChecked {
			*per++
		}
	} else {
		*per++
	}
	if y < len(matrix[0])-1 {
		if matrix[x][y+1] == plant {
			checkRegion(matrix, x, y+1, area, per)
		} else if matrix[x][y+1] != plantChecked {
			*per++
		}
	} else {
		*per++
	}
	if x < len(matrix[0])-1 {
		if matrix[x+1][y] == plant {
			checkRegion(matrix, x+1, y, area, per)
		} else if matrix[x+1][y] != plantChecked {
			*per++
		}
	} else {
		*per++
	}
	if y > 0 {
		if matrix[x][y-1] == plant {
			checkRegion(matrix, x, y-1, area, per)
		} else if matrix[x][y-1] != plantChecked {
			*per++
		}
	} else {
		*per++
	}
}

func countPrice(matrix [][]rune) int {
	plants := getPlants(matrix)
	price := 0
	for i, row := range matrix {
		for j, _ := range row {
			if utils.ContainRune(plants, matrix[i][j]) {
				area := 0
				per := 0
				checkRegion(matrix, i, j, &area, &per)
				price += area * per
			}
		}
	}
	return price
}

func main() {
	content := utils.FileToStr("input")
	matrix := utils.StrToMatrix(content)
	price := countPrice(matrix)
	fmt.Println(price)
}
