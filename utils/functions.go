package utils

import (
	"os"
	"io/ioutil"
	"log"
	"strconv"
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

func Contain(slice []int, num int) bool {
	for _, v := range slice {
		if v == num { return true }
	}
	return false
}

func WriteToFile(name string, content string) {
	file, _ := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	file.WriteString(content)
}
