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

	totalChars := 0
	for i := 0; i < len(matches); i++ {
		garbage := input[matches[i][0]:matches[i][1]]
		for j := 0; j < len(garbage); j++ {
			if (string(garbage[j]) == "!") {
				garbage = garbage[:j] + garbage[j+2:]
				j--
			}
		}
		totalChars += len(garbage) - 2
	}
	fmt.Println(totalChars)

}