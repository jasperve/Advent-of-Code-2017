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

	counter := 0
	repeatAfter := 0
	iterCounter := 0
	iterToGo := 0

	OUTER:
	for {

		for _, move := range strings.Split(string(input), ",") {

			if repeatAfter != 0 {
				iterCounter++
			}

			counter++

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

			if iterCounter == iterToGo && iterToGo != 0 {
				break OUTER
			}

			if string(line) == "abcdefghijklmnop" {
				repeatAfter = counter
				iterToGo = 1000000000%repeatAfter
			}

			
			
		}

	}

	fmt.Println(string(line))

}