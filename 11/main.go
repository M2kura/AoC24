package main

import (
	"fmt"
	"AoC24/utils"
	"math"
)

func blink25times(stonesRow []int) int {
	stonesLen := len(stonesRow)
	for i:=0; i<25; i++ {
		for j:=0; j<stonesLen; j++ {
			if stonesRow[j] == 0 {
				stonesRow[j] = 1
			} else if digits:=utils.CountDigits(stonesRow[j]);digits%2 == 0 {
				rightNumber := 0
				digitsHalf := digits/2
				for k:=0;k<digitsHalf;k++{
					rightNumber = int(math.Pow(10.0, float64(k))) * (stonesRow[j]%10) + rightNumber
					stonesRow[j] /= 10
				}
				stonesRow = append(stonesRow[:j+1], append([]int{rightNumber}, stonesRow[j+1:]...)...)
				stonesLen++
				j++
			} else {
				stonesRow[j] *= 2024
			}
		}
	}
	return stonesLen
}

func main() {
	delim := []rune{' ', '\n'}
	content := utils.FileToStr("input")
	stonesRow := utils.StrToSpliceInt(content, delim)
	stonesLen := blink25times(stonesRow)
	fmt.Println(stonesLen)
}
