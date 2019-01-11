package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"time"
)

type instruction struct {
	class string
	x     string
	y     string
}

func main() {

	register := make(map[string]int)
	var position, mulCounter int

	functions := make(map[string]func(string, string))

	functions["set"] = func(x string, y string) {
		if i, err := strconv.Atoi(y); err != nil {
			register[x] = register[y]
		} else {
			register[x] = i
		}
	}

	functions["sub"] = func(x string, y string) {
		if i, err := strconv.Atoi(y); err != nil {
			register[x] -= register[y]
		} else {
			register[x] -= i
		}
	}

	functions["mul"] = func(x string, y string) {
		if i, err := strconv.Atoi(y); err != nil {
			register[x] *= register[y]
		} else {
			register[x] *= i
		}
		mulCounter++
	}

	functions["jnz"] = func(x string, y string) {
		if vx, err := strconv.Atoi(x); err == nil {
			if vx > 0 {
				if vy, err := strconv.Atoi(y); err == nil {
					position += vy
				} else {
					position += register[y]
				}
	
			}
		} else {
			if register[x] != 0 {
				if vy, err := strconv.Atoi(y); err == nil {
					position += vy
				} else {
					position += register[y]
				}
			} else {
				position++
			}
		}
	}


	instructions := []instruction{}

	file, _ := os.Open("input.txt")
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

	register["a"] = 1

	theOne := false


	for position >= 0 && position < len(instructions) {

		positionBefore := position
		fmt.Println(register["e"], register["d"])

		if theOne {

			fmt.Println(instructions[position])
			return

		}

		if register["d"] == 3 && instructions[position].class == "jnz" {

			fmt.Println(register["e"], register["d"])
			fmt.Println("before: ", positionBefore)
			theOne = true
		}

		functions[instructions[position].class](instructions[position].x, instructions[position].y)
		if position == positionBefore {
			position++
		}

		if theOne {
			fmt.Println(register["e"], register["d"])
			fmt.Println("after:", position)
		}


	}

	fmt.Println(mulCounter)


}
