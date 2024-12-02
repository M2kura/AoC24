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

func safecheck(nums []int, firstTime bool) bool {
	for i := 0; i < len(nums) - 1; i++ {
		difference := nums[i] - nums[i+1]
		if nums[0] < nums[1] {
			difference = -difference
		}

		if difference < 1 || difference > 3 {
			if firstTime  {
				nums1 := append(append([]int(nil), nums[:i]...), nums[i+1:]...)
				nums2 := append(append([]int(nil), nums[:i+1]...), nums[i+2:]...)
				if safecheck(nums1, false) || safecheck(nums2, false) {
					return true
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
	file, _ := os.OpenFile("test", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	var wrong_nums [][]int

	for scanner.Scan() {
		nums := getNumsSlice(scanner.Text())
		if safecheck(nums, true) {
			res += 1
		} else {
			file.WriteString(fmt.Sprint(nums) + "\n")
			wrong_nums = append(wrong_nums, nums)
		}
	}

	fmt.Println(res)

	// I could just do this in the safecheck function and live a happy life
	var valid_nums [][]int
	for _, wrong_num := range wrong_nums {
		for i, _ := range wrong_num {
			slice := append(append([]int(nil), wrong_num[:i]...), wrong_num[i+1:]...)
			if safecheck(slice, false) {
				valid_nums = append(valid_nums, wrong_num)
			}
		}
	}

	file2, _ := os.OpenFile("test_new", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file2.Close()
	for _, nums := range valid_nums {
		file2.WriteString(fmt.Sprint(nums) + "\n")
	}
}
