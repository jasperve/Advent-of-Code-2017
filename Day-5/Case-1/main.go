package main

import (
	"log"
	"strings"
	"io/ioutil"
	"strconv"
	"fmt"
)

func main() {

	input, _ := ioutil.ReadFile("input-test.txt")

	instructions := make(map[int]int)
	for i, instruction := range strings.Split(string(input), "\r\n"){
		instructionInt, err := strconv.Atoi(instruction)
		if err != nil {
			log.Fatalln(err)
		}
		instructions[i] = instructionInt
	}

	position := 0
	jumps := 0

	for {
		if _, ok := instructions[position]; !ok { break }
		instruction := instructions[position]
		instructions[position]++
		position += instruction
		jumps++
	}

	fmt.Println("Number of jumps needed: ", jumps)

}