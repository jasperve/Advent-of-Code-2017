package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {

	i, _ := ioutil.ReadFile("input.txt")
	input := string(i)
	matches := regexp.MustCompile("<>|<.*?([^!]|[^!](!!)*)>").FindAllStringSubmatchIndex(input, -1)

	for i := len(matches)-1; i >= 0; i-- {

		fmt.Println(matches[i][0], matches[i][1])
		input = input[:matches[i][0]]+input[matches[i][1]:]
	}

	fmt.Println(input)

}