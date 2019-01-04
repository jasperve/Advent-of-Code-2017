package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

const hashLength = 256

func main() {

	sparseHash := []int{}
	for i := 0; i < hashLength; i++ {
		sparseHash = append(sparseHash, i)
	}

	input, _ := ioutil.ReadFile("input.txt")
	//input = append(input, 3, 4, 1, 5, 17, 31, 73, 47, 23)

	position := 0
	skip := 0

	for _, v := range strings.Split(string(input), ",") {

		vInt, _ := strconv.Atoi(v)

		tempList := []int{}
		for i := 0; i < vInt; i++ {
			tempList = append(tempList, sparseHash[(position+i)%len(sparseHash)])
		}

		for i := 0; i < vInt; i++ {
			sparseHash[(position+i)%len(sparseHash)] = tempList[(len(tempList)-1-i)%len(sparseHash)]
		}

		position += vInt + skip
		skip++
	}

	fmt.Println(sparseHash)

}