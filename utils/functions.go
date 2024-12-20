package utils

import (
	"os"
	"io/ioutil"
	"log"
	"strconv"
	"regexp"
)

func FileToStr(path string) string {
	file, err := os.Open(path)
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

func Atoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil { log.Fatal(err) }
	return num
}

func ContainInt(slice []int, num int) bool {
	for _, v := range slice {
		if v == num { return true }
	}
	return false
}

func ContainRune(slice []rune, char rune) bool {
	for _, v := range slice {
		if v == char { return true }
	}
	return false
}

func WriteToFile(name string, content string) {
	file, _ := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	file.WriteString(content)
}

func StrToMatrix(str string) [][]rune {
	matrix := [][]rune{{},}
	row, col := 0, 0
	for i, char := range str {
		if char == '\n' {
			if i != len(str) - 1 {
				matrix = append(matrix, []rune{})
			}
			row++
			col = 0
		} else {
			matrix[row] = append(matrix[row], char)
			col++
		}
	}
	return matrix
}

func StrToMatrixInt(str string) [][]int {
	matrix := [][]int{{},}
	row, col := 0, 0
	for i, num := range str {
		if num == '\n' {
			if i != len(str) - 1 {
				matrix = append(matrix, []int{})
			}
			row++
			col = 0
		} else {
			matrix[row] = append(matrix[row], Atoi(string(num)))
			col++
		}
	}
	return matrix
}

func DeepCopy(matrix [][]rune) [][]rune {
	copyMatrix := make([][]rune, len(matrix))
	for i := range matrix {
		copyMatrix[i] = make([]rune, len(matrix[i]))
		copy(copyMatrix[i], matrix[i])
	}
	return copyMatrix
}

func RuneMatrixToFile(name string, matrix [][]rune) {
	file, _ := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	for i, row := range matrix {
		str := ""
		for j := range row {
			str += string(matrix[i][j])
		}
		str += "\n"
		file.WriteString(str)
	}
}

func CountDigits(num int) int {
	count := 0 
	for num != 0 {
		num = num/10
		count++
	}
	return count
}

func StrToSpliceInt(str string, delimiters []rune) []int {
	numSlice := []int{}
	delList := "["
	for _, d := range delimiters {
		delList += string(d)
	}
	delList += "]"
	re := regexp.MustCompile(delList)
	strSlice := re.Split(str, -1)
	strSlice = strSlice[:len(strSlice)-1]
	for _, num := range strSlice {
		numSlice = append(numSlice, Atoi(num))
	}
	return numSlice
}
