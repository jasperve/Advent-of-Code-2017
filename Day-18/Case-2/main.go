package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type instruction struct {
	class string
	x     string
	y     string
}

var instructions []instruction

func main() {

	loadInstructions("input.txt")

	var wg = sync.WaitGroup{}

	channelA := make(chan int, 100000000)
	channelB := make(chan int, 100000000)

	var program0LockSince int64
	var program1LockSince int64
	program0Result := 0
	program1Result := 0

	wg.Add(2)
	go startProgram(&wg, 0, &program0LockSince, &program1LockSince, channelA, channelB, &program0Result)
	go startProgram(&wg, 1, &program1LockSince, &program0LockSince, channelB, channelA, &program1Result)
	wg.Wait()

	fmt.Println(program1Result)

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

func startProgram(wg *sync.WaitGroup, programID int, selfLockSince *int64, otherLockSince *int64, sendChannel chan int, receiveChannel chan int, result *int) {

	register := map[string]int{"p": programID}
	position := 0
	amountSent := 0

INSTRUCTIONSLOOP:
	for position >= 0 && position < len(instructions) {

		fmt.Println(programID, amountSent)

		x := instructions[position].x
		y := instructions[position].y
		oldPosition := position

		switch instructions[position].class {
		case "rcv":
			if *selfLockSince == 0 {
				*selfLockSince = time.Now().UnixNano()
			} 
			select {
			case register[x] = <-receiveChannel:
				*selfLockSince = 0
			default:
				if *otherLockSince < time.Now().UnixNano() - 5000000000 {
					break INSTRUCTIONSLOOP
				} else {
					continue INSTRUCTIONSLOOP
				}
			}
		case "snd":
			sendChannel <- register[x]
			amountSent++
		case "jgz":
			if vx, err := strconv.Atoi(x); err == nil {
				if vx > 0 {
					if vy, err := strconv.Atoi(y); err == nil {
						position += vy
					} else {
						position += register[y]
					}
				}
			} else {
				if register[x] > 0 {
					if vy, err := strconv.Atoi(y); err == nil {
						position += vy
					} else {
						position += register[y]
					}
				} 
			}
		case "set":
			if i, err := strconv.Atoi(y); err != nil {
				register[x] = register[y]
			} else {
				register[x] = i
			}
		case "add":
			if i, err := strconv.Atoi(y); err != nil {
				register[x] += register[y]
			} else {
				register[x] += i
			}
		case "mul":
			if i, err := strconv.Atoi(y); err != nil {
				register[x] *= register[y]
			} else {
				register[x] *= i
			}
		case "mod":
			if i, err := strconv.Atoi(y); err != nil {
				register[x] %= register[y]
			} else {
				register[x] %= i
			}
		}

		if position == oldPosition {
			position++
		}

	}

	*result = amountSent
	wg.Done()

}
