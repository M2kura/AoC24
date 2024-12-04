package main

import (
	"fmt"
	"AoC24/utils"
	"unicode"
	"strconv"
)

func getMulSum(str string) int {
	sum := 0
	enabled := true
	MainLoop:
	for i := 0; i < len(str); i++ {
		if str[i] == 'd' {
			if substr := str[i:i+7]; substr == "don't()" {
				enabled = false
				i += 6
			} else if substr := str[i:i+4]; substr == "do()" {
				enabled = true
				i += 3
			}
		}
		if enabled && str[i] == 'm' {
			var str1, str2 string
			i++
			chars1 := "ul("
			for j := 0; j < 3; j++ {
				if str[i] != chars1[j] { continue MainLoop }
				i++
			}
			for k := 0; k < 4; k++ {
				if k == 0 {
					if !unicode.IsDigit(rune(str[i])) || str[i] == '0' { continue MainLoop }
				} else {
					if !unicode.IsDigit(rune(str[i])) { continue MainLoop }
				}
				str1 += string(str[i])
				i++
				if str[i] == ',' {
					break
				} else if k == 3 && str[i] != ',' { continue MainLoop }
			}
			i++
			for l := 0; l < 4; l++ {
				if l == 0 {
					if !unicode.IsDigit(rune(str[i])) || str[i] == '0' { continue MainLoop }
				} else {
					if !unicode.IsDigit(rune(str[i])) { continue MainLoop }
				}
				str2 += string(str[i])
				i++
				if str[i] == ')' {
					break
				} else if l == 3 && str[i] != ')' { continue MainLoop }
			}
			num1, _ := strconv.Atoi(str1)
			num2, _ := strconv.Atoi(str2)
			sum += num1 * num2
		}
	}
	return sum
}

func main() {
	content := utils.FileToStr("input")
	res := getMulSum(content)
	fmt.Println(res)
}
