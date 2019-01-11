package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	class string
	x     string
	y     string
}

var instructions []instruction

func main() {

	loadInstructions("input.txt")
	fmt.Println(runInstructions())
	
}

func loadInstructions(location string) {

	file, _ := os.Open(location)
	input := bufio.NewScanner(file)

	for input.Scan() {

		inputSplit := strings.Split(input.Text(), " ")

		y := ""
		if len(inputSplit) == 3 {
			y = inputSplit[2]
		}

		newInstruction := instruction{
			class: inputSplit[0],
			x:     inputSplit[1],
			y:     y,
		}

		instructions = append(instructions, newInstruction)

	}

}

func runInstructions() int {

	registers := []map[string]int{ map[string]int{"p": 0}, map[string]int{"p": 1} }
	positions := []int{0, 0}
	amountSent := []int{0, 0}
	locked := []bool{false, false}
	lists := [][]int{ []int{}, []int{}}

	OUTER:
	for {

		for i := 0; i < 2; i++ {

			//fmt.Println(i, instructions[positions[i]], amountSent[i])
			
			x := instructions[positions[i]].x
			y := instructions[positions[i]].y

			switch instructions[positions[i]].class {
			case "rcv":
				if len(lists[i]) == 0 {
					locked[i] = true
				} else {
					registers[i][x], lists[i] = lists[i][0], lists[i][1:]
					locked[i] = false
					positions[i]++
				}
			case "snd":
				lists[(i+1)%2] = append(lists[(i+1)%2], registers[i][x])
				amountSent[i]++
				positions[i]++
			case "jgz":
				if vx, err := strconv.Atoi(x); err == nil {
					if vx > 0 {
						if vy, err := strconv.Atoi(y); err == nil {
							positions[i] += vy
						} else {
							positions[i] += registers[i][y]
						}
	
					}
				} else {
					if registers[i][x] > 0 {
						if vy, err := strconv.Atoi(y); err == nil {
							positions[i] += vy
						} else {
							positions[i] += registers[i][y]
						}
					} else {
						positions[i]++
					}
				}
			case "set":
				if v, err := strconv.Atoi(y); err != nil {
					registers[i][x] = registers[i][y]
				} else {
					registers[i][x] = v
				}
				positions[i]++
			case "add":
				if v, err := strconv.Atoi(y); err != nil {
					registers[i][x] += registers[i][y]
				} else {
					registers[i][x] += v
				}
				positions[i]++
			case "mul":
				if v, err := strconv.Atoi(y); err != nil {
					registers[i][x] *= registers[i][y]
				} else {
					registers[i][x] *= v
				}
				positions[i]++
			case "mod":
				if v, err := strconv.Atoi(y); err != nil {
					registers[i][x] %= registers[i][y]
				} else {
					registers[i][x] %= v
				}
				positions[i]++
			}

			if (positions[i] < 0 || positions[i] >= len(instructions)) || locked[0] && locked[1] {
				break OUTER
			}

		}

	}
	
	return amountSent[1]

}