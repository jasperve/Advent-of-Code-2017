package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const listLength = 256

func main() {

	list := []int{}
	for i := 0; i < listLength; i++ {
		list = append(list, i)
	}

	input := splitToIntSlice("input.txt", ",")

	position := 0
	skip := 0

	for _, v := range input {

		tempList := []int{}
		for ii := 0; ii < v; ii++ {
			tempList = append(tempList, list[(position+ii)%len(list)])
		}

		for ii := 0; ii < v; ii++ {
			list[(position+ii)%len(list)] = tempList[(len(tempList)-1-ii)%len(list)]
		}

		position += v+skip
		skip++

	}

	fmt.Println(list[0]*list[1])


}


func splitToIntSlice(location string, sep string) (out []int) {
	input, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatalln("FATAL: Unable to open file at location: %v", location)
	}
	for _, token := range strings.Split(string(input), sep) {
		value, err := strconv.Atoi(token)
		if err != nil {
			log.Fatalln("FATAL: Unable to convert %v", token)
		}
		out = append(out, value)
	}
	return out
}
