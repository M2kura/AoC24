package main

import (
	"io/ioutil"
	"os"
	"bufio"
	"strings"
	"strconv"
	"sort"
	"fmt"
	"log"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	reader := strings.NewReader(string(content))
	scanner := bufio.NewScanner(reader)

	var left, right []int

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "   ")
		leftNum, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err)
		}
		rightNum, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, leftNum)
		right = append(right, rightNum)
	}
	
	sort.Ints(left)
	sort.Ints(right)
	
	res := 0

	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]
		if distance < 0 {
			distance = -distance
		}
		res += distance
	}
	
	fmt.Println(res)
	
	// Part 2

	score := 0

	for _, leftNum := range left {
		count := 0
		for _, rightNum := range right {
			if leftNum == rightNum {
				count++
			}
		}
		score += count * leftNum
	}

	fmt.Println(score)
}
