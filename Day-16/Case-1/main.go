package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

func main() {

	input, _ := ioutil.ReadFile("input.txt")
	line := []byte("abcdefghijklmnop")

	for _, move := range strings.Split(string(input), ",") {

		switch move[0] {
		case 's':
			position, _ := strconv.Atoi(move[1:])
			line = append(line[len(line)-position:], line[:len(line)-position]...)
		case 'x':
			positions := strings.Split(move[1:], "/")
			positionA, _ := strconv.Atoi(positions[0])
			positionB, _ := strconv.Atoi(positions[1])
			line[positionA], line[positionB] = line[positionB], line[positionA]
		case 'p':
			partners := strings.Split(move[1:], "/")
			var positionA, positionB int
			for l := 0; l < len(line); l++ {
				if string(line[l]) == partners[0] {
					positionA = l
				} else if string(line[l]) == partners[1] {
					positionB = l
				}
			}
			line[positionA], line[positionB] = line[positionB], line[positionA]
		}	
		
	}

	fmt.Println(string(line))

}