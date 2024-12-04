package utils

import (
	"os"
	"io/ioutil"
	"log"
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
