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

type forbiddenScope struct {
	base int
	addition int
}

func main() {

	layers := make(map[int]*layer)
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

	forbiddenScopes := []forbiddenScope{}

	for i, l := range layers {
		forbiddenScopes = append(forbiddenScopes, forbiddenScope{ base: (l.scope-1)*2-i, addition: (l.scope-1)*2 })
	}

	delay := 0

	OUTER:
	for  {

		for s := 0; s < len(forbiddenScopes); s++ {
			if delay == forbiddenScopes[s].base || (delay-forbiddenScopes[s].base)%forbiddenScopes[s].addition == 0 { 
				delay++
				continue OUTER 
			}
		}
			
		fmt.Println("result:", delay)
		break

	}
}