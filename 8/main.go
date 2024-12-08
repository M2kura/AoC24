package main

import (
	"AoC24/utils"
	"fmt"
)

func inMatrix(matrix [][]rune, x int, y int) bool {
	if x >= len(matrix) || y >= len(matrix[0]) || x < 0 || y < 0 {
		return false
	}
	return true
}

func mapSignals(matrix [][]rune) map[rune][][]int {
	signals := make(map[rune][][]int)
	for i, row := range matrix {
		for j := range row {
			if matrix[i][j] != '.' {
				if _, exists := signals[matrix[i][j]]; !exists {
					signals[matrix[i][j]] = [][]int{}
				}
				signals[matrix[i][j]] = append(signals[matrix[i][j]], []int{i, j})
			}
		}
	}
	return signals
}

func isUnique(antinodes [][]int, x int, y int) bool {
	for _, antinode := range antinodes {
		if antinode[0] == x && antinode[1] == y {
			return false
		}
	}
	return true
}

func countAntinodes(matrix [][]rune, harmonics bool) int {
	signals := mapSignals(matrix)
	uniqueAntinodes := [][]int{}
	count := 0
	for _, cords := range signals {
		for _, xx := range cords {
			for _, yy := range cords {
				if xx[0] == yy[0] && xx[1] == yy[1] { continue }
				distance := []int{xx[0]-yy[0],xx[1]-yy[1]}
				if !harmonics {
					x := xx[0] + distance[0]
					y := xx[1] + distance[1]
					if inMatrix(matrix, x, y) && isUnique(uniqueAntinodes, x, y) {
						count++
						uniqueAntinodes = append(uniqueAntinodes, []int{x, y})
					}
				} else {
					antinode := []int{xx[0], xx[1]}
					for inMatrix(matrix, antinode[0], antinode[1]) {
						if isUnique(uniqueAntinodes, antinode[0], antinode[1]) {
							count++
							uniqueAntinodes = append(uniqueAntinodes, []int{antinode[0], antinode[1]})
						}
						antinode = []int{antinode[0]+distance[0],antinode[1]+distance[1]}
					}
				}
			}
		}
	}
	return count
}

func main() {
	content := utils.FileToStr("input")
	matrix := utils.StrToMatrix(content)
	sum := countAntinodes(matrix, false)
	fmt.Println(sum)
	sumAll := countAntinodes(matrix, true)
	fmt.Println(sumAll)
}
