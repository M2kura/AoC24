package main

import (
	"AoC24/utils"
	"fmt"
	"regexp"
	"strings"
)

func extractNumbers(str string) [][]int {
	nums := [][]int{}
	pattern := `\d+`
	re := regexp.MustCompile(pattern)
	lines := strings.Split(str, "\n")
	lines = lines[:len(lines)-1]
	for _, line := range lines {
		num_str := re.FindAllString(line, -1)
		num_slice := []int{}
		for _, num := range num_str {
			num_slice = append(num_slice, utils.Atoi(num))
		}
		nums = append(nums, num_slice)
	}
	return nums 
}

func calculate(nums []int, indexNext int, current int, results *[]int) {
	if indexNext == len(nums) {
		*results = append(*results, current)
		return
	}
	calculate(nums, indexNext+1, current+nums[indexNext], results)
	calculate(nums, indexNext+1, current*nums[indexNext], results)
	digits := utils.CountDigits(nums[indexNext])
	for i := 0; i < digits; i++ { current *= 10 }
	calculate(nums, indexNext+1, current+nums[indexNext], results)
}

func getSum(nums [][]int) int {
	sum := 0
	for _, line := range nums {
		results := []int{}
		calculate(line[1:], 1, line[1], &results)
		if utils.ContainInt(results, line[0]) { sum += line[0] }
	}
	return sum
}

func main() {
	content := utils.FileToStr("input")
	eqs := extractNumbers(content)
	sum := getSum(eqs)
	fmt.Println(sum)
}
