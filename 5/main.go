package main

import (
	"fmt"
	"AoC24/utils"
	"regexp"
	"strings"
)

func mapRules(str string) map[int][]int {
	rules := make(map[int][]int)
	re := regexp.MustCompile(`\d{2}\|\d{2}`)
	matches := re.FindAllString(str, -1)
	for _, rule := range matches {
		num1 := utils.Atoi(rule[0:2])
		num2 := utils.Atoi(rule[3:])
		_, exists := rules[num1]
		if exists {
			rules[num1] = append(rules[num1], num2)
		} else {
			rules[num1] = []int{num2}
		}
	}
	return rules
}

func getRawUpdates(str string, rules map[int][]int) []string {
	for i := 0; i < len(str); i++ {
		if rune(str[i]) == '\n' && rune(str[i+1]) == '\n' {
			str = str[i+2:]
			break
		}
	}
	updates := strings.Split(str, "\n")
	updates = updates[:len(updates)-1]
	return updates
}

func separateUpdates(updates []string, rules map[int][]int) ([][]int, [][]int) {
	goodRes := [][]int{}
	badRes := [][]int{}
	MainLoop:
	for _, update := range updates {
		nums_str := strings.Split(update, ",")
		nums := []int{}
		for _, num := range nums_str {
			nums = append(nums, utils.Atoi(num))
		}

		for i := 0; i < len(nums)-1; i++ {
			for j := 0; j < i; j++ {
				if utils.Contain(rules[nums[i]], nums[j]) {
					badRes = append(badRes , nums)
					continue MainLoop
				}
			}
		}
		goodRes = append(goodRes , nums)
	}
	return goodRes, badRes
}

func findSumOfNonValid(updates [][]int, rules map[int][]int) int {
	sum := 0
	MainLoop:
	for _, update := range updates {
		for _, num1 := range update {
			numsAfter := 0
			for _, num2 := range update {
				if utils.Contain(rules[num1], num2) {
					numsAfter++
				}
			}
			if numsAfter == len(update)/2 {
				sum += num1
				continue MainLoop
			}
		}
	}
	return sum
}

func countSum(updates [][]int) int {
	sum := 0
	for _, update := range updates {
		sum += update[len(update)/2]
	}
	return sum
}

func main() {
	content := utils.FileToStr("input")
	rules := mapRules(content)
	updatesRaw := getRawUpdates(content, rules)
	updatesValid, updatesNonValid:= separateUpdates(updatesRaw, rules)
	sum := countSum(updatesValid)
	nonValidSum := findSumOfNonValid(updatesNonValid, rules)
	fmt.Println(sum)
	fmt.Println(nonValidSum)
}
