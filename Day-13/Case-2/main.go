package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

const (
	up = 0
	down = 1
)

type layer struct {
	position int
	direction int
	scope int
}

var layers map[int]*layer

func main() {

	layers = make(map[int]*layer)
	maxPosition := 0

	file, _ := os.Open("input.txt")
	input := bufio.NewScanner(file)

	for input.Scan() {
		
		line := strings.Split(input.Text(), ": ")
		depth, _ := strconv.Atoi(line[0])
		scope, _ := strconv.Atoi(line[1])
		layers[depth] = &layer{ position: 1, direction: down, scope: scope }
		if depth > maxPosition {
			maxPosition = depth
		}
		
	}

	delay := 0

	for {

		resetLayers()
		for d := 0; d < delay; d++ {
			moveLayers()
		}
		
		detected := false

		for p := 0; p <= maxPosition; p++ {

			if l, ok := layers[p]; ok && l.position == 1 {
				detected = true
			}
			moveLayers()

		}

		if !detected {
			break
		}
		delay++

	}

	fmt.Println("We can pass through without being detected with a delay of", delay)

}


func resetLayers() {
	for _, l := range layers {
		l.position = 1
		l.direction = down
	}
}


func moveLayers() {
	for _, l := range layers {
		if l.position == 1 && l.direction == up {
			l.direction = down
			l.position++
		} else if l.position == l.scope && l.direction == down {
			l.direction = up
			l.position--
		} else if l.direction == down {
			l.position++
		} else if l.direction == up {
			l.position--
		}
	}
}