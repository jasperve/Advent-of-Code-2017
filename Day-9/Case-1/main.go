package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {

	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	matches := regexp.MustCompile("<>|<(!!)*>|<.*?([^!]|[^!](!!)*)>").FindAllStringSubmatchIndex(input, -1)

	for i := len(matches) - 1; i >= 0; i-- {
		input = input[:matches[i][0]] + input[matches[i][1]:]
	}

	total, level := 0, 0
	for _, v := range input {
		if string(v) == "{" {
			level++
		} else if string(v) == "}" {
			total += level
			level--
		}
	}

	fmt.Println(total)

}