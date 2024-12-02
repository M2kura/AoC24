package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"strings"
	"bufio"
	"strconv"
)

func getNumsSlice(str string) []int {
	nums_str := strings.Split(str, " ")
	var nums []int 
	for _, num_str := range nums_str {
		num, err := strconv.Atoi(num_str)
		if err != nil { log.Fatal(err) }
		nums = append(nums, num)
	}
	return nums
}

func safecheck(nums []int, problemDampener bool) bool {
	for i := 0; i < len(nums) - 1; i++ {
		difference := nums[i] - nums[i+1]
		if nums[0] < nums[1] {
			difference = -difference
		}

		if difference < 1 || difference > 3 {
			if problemDampener {
				for j, _ := range nums {
					sub_nums := append(append([]int(nil), nums[:j]...), nums[j+1:]...)
					if (safecheck(sub_nums, false)) { return true }
				}
			}
			return false
		}
	}
	return true
}

func fileToStr(path string) string {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func main() {
	reader := strings.NewReader(fileToStr("input"))
	scanner := bufio.NewScanner(reader)
	res := 0

	for scanner.Scan() {
		nums := getNumsSlice(scanner.Text())
		if safecheck(nums, false) { res += 1 }
	}
	
	fmt.Println(res)
	// Part 2

	res = 0
	reader.Reset(fileToStr("input"))
	scanner = bufio.NewScanner(reader)

	for scanner.Scan() {
		nums := getNumsSlice(scanner.Text())
		if safecheck(nums, true) { res += 1 }
	}

	fmt.Println(res)
}
