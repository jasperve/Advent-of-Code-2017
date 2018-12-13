package main

import (

	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"

)

type program struct {
	id int
	pipes []*program
}

func main() {

	programs := make(map[int]*program)
	
	file, _ := os.Open("input-test.txt")
	input := bufio.NewScanner(file)

	for input.Scan() {

		line := strings.Split(input.Text(), " <-> ")
		lineRight := strings.Split(line[1], ", ")

		id, _ := strconv.Atoi(line[0])

		var parent *program
		if _, ok := programs[id]; ok {
			parent = programs[id]
			fmt.Println(parent.id)
		} else {
			parent = &program{}
			parent.id = id
			programs[id] = parent
		}
		
		fmt.Println("hier:", parent.id)

		pipes := []*program{}
		for _, v := range lineRight {
			pipeID, _ := strconv.Atoi(v)
			if _, ok := programs[pipeID]; ok {
				pipes = append(pipes, programs[pipeID])
			} else {
				pipe := &program{ id: pipeID }
				pipes = append(pipes, pipe)
				
				programs[pipeID] = pipe
			}
		}

		parent.pipes = pipes
		programs[id] = parent

	}

	allPipes := listPipes(programs[2])

	fmt.Println(allPipes)

}


func listPipes(parent *program) []*program {

	fmt.Println("Function called for ", parent.id)

	allPipes := []*program{}
	for _, pipe := range parent.pipes {
		fmt.Println(pipe)
		allPipes = listPipes(pipe)
	}
	return allPipes

}