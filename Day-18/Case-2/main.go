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

func main() {

	register := make(map[string]int)
	var lastPlayedFreq, position int

	functions := make(map[string]func(string, string))

	functions["snd"] = func(x string, y string) { lastPlayedFreq = register[x] }

	functions["set"] = func(x string, y string) {
		if i, err := strconv.Atoi(y); err != nil {
			register[x] = register[y]
		} else {
			register[x] = i
		}
	}

	functions["add"] = func(x string, y string) {
		if i, err := strconv.Atoi(y); err != nil {
			register[x] += register[y]
		} else {
			register[x] += i
		}
	}

	functions["mul"] = func(x string, y string) {
		if i, err := strconv.Atoi(y); err != nil {
			register[x] *= register[y]
		} else {
			register[x] *= i
		}
	}

	functions["mod"] = func(x string, y string) {
		if i, err := strconv.Atoi(y); err != nil {
			register[x] %= register[y]
		} else {
			register[x] %= i
		}
	}

	functions["rcv"] = func(x string, y string) {
		if register[x] != 0 {
			register[x] = lastPlayedFreq
		}
	}
	functions["jgz"] = func(x string, y string) {
		if register[x] > 0 {
			if i, err := strconv.Atoi(y); err == nil {
				position += i
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

	for position >= 0 && position < len(instructions) {

		if instructions[position].class == "rcv" && register[instructions[position].x] != 0 {
			fmt.Println("first time with recovered frequency", lastPlayedFreq)
			return
		}

		positionBefore := position
		functions[instructions[position].class](instructions[position].x, instructions[position].y)
		if position == positionBefore {
			position++
		}

	}

}
