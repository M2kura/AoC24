package main

import (
	"fmt"
	"AoC24/utils"
)

func makeMemory(str string) []int {
	res := []int{}
	for i, char := range str {
		if i%2 == 0 {
			for j:=0;j<utils.Atoi(string(char));j++ {
				res = append(res, i/2)
			}
		} else {
			if char == '\n' { break }
			for j:=0;j<utils.Atoi(string(char));j++ {
				res = append(res, -1)
			}
		}
	}
	return res
}

func moveBlocks(memory []int) {
	curBlockId := -2
	MainLoop:
	for i:=len(memory)-1;i>0;i-- {
		if memory[i] != -1 {
			if memory[i] != curBlockId {
				curBlockId = memory[i]
			}
			for j:=0;j<len(memory)-1;j++ {
				if j == i {
					break MainLoop
				} else if memory[j] == -1 {
					memory[j] = curBlockId
					break
				}
			}
			memory[i] = -1
		}
	}
}

func moveFiles(memory []int) {
	curFileId := -2
	fileLength := 1
	for i:=len(memory)-1;i>0;i-- {
		if memory[i] != -1 {
			curFileId = memory[i]
			for memory[i-1] == curFileId {
				fileLength++
				i--
				if i == 0 { break }
			}
			InnerLoop:
			for j:=0;j<len(memory)-1;j++ {
				if j == i {
					break
				} else if memory[j] == -1 {
					for k:=1;k<fileLength;k++ {
						if memory[j+k] != -1 {
							continue InnerLoop
						}
					}
					for k:=0;k<fileLength;k++ {
						memory[j+k] = curFileId
						memory[i+k] = -1
					}
					break
				}
			}
			fileLength = 1
		}
	}
}

func checkSum(memory []int) int {
	sum := 0
	for i, num := range memory {
		if num != -1 {
			sum += i * num
		}
	}
	return sum
}

func main() {
	content := utils.FileToStr("input")
	memory := makeMemory(content)
	moveBlocks(memory)
	sum := checkSum(memory)
	fmt.Println(sum)
	memory = makeMemory(content)
	moveFiles(memory)
	sum = checkSum(memory)
	fmt.Println(sum)
}
