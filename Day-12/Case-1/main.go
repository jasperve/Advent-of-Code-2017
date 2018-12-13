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
	pipes []program
}

func main() {

	programs := make(map[int]program)
	
	file, _ := os.Open("input-test.txt")
	input := bufio.NewScanner(file)

	for input.Scan() {

		line := strings.Split(input.Text(), " <-> ")
		lineRight := strings.Split(line[1], ", ")

		id, _ := strconv.Atoi(line[0])

		parent := program{}
		if _, ok := programs[id]; ok {
			parent = programs[id]
		} else {
			parent.id = id
			programs[id] = parent
		}
		
		pipes := []program{}
		for _, v := range lineRight {
			pipe_id, _ := strconv.Atoi(v)
			if _, ok := programs[pipe_id]; ok {
				pipes = append(pipes, programs[pipe_id])
			} else {
				pipe := program{ id: pipe_id }
				pipes = append(pipes, pipe)
				programs[pipe_id] = pipe
			}
		}

		parent.pipes = pipes
		programs[id] = parent

	}

	allPipes := listPipes(programs[0])

}


func listPipes(parent program) []*program {

	allPipes := []program{}

	for _, pipe := range parent.pipes {
		allPipes = listPipes(pipe)

		fmt.Println(" - ", pipe.id)
	}

}