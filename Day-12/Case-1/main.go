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
	marked bool
	connections []*program
}

func main() {

	programs := make(map[int]*program)
	
	file, _ := os.Open("input.txt")
	input := bufio.NewScanner(file)

	for input.Scan() {

		line := strings.Split(input.Text(), " <-> ")
		lineRight := strings.Split(line[1], ", ")

		id, _ := strconv.Atoi(line[0])

        parent, ok := programs[id]
        if !ok {
            parent = &program{id: id}
            programs[id] = parent
        }
		
		connections := []*program{}
		for _, v := range lineRight {
			connectionID, _ := strconv.Atoi(v)
			if _, ok := programs[connectionID]; ok {
				connections = append(connections, programs[connectionID])
			} else {
				connection := &program{ id: connectionID }
				connections = append(connections, connection)
				programs[connectionID] = connection
			}
		}

		parent.connections = connections

	}

	markPrograms(programs[0])

	counter := 0
	for i := 0; i < len(programs); i++ {
		if programs[i].marked {
			counter++
		}
		fmt.Println(programs[i].id, programs[i].marked, programs[i].connections)
	}

	fmt.Println(counter)

}


func markPrograms(parent *program) {
	parent.marked = true
	for i := 0; i < len(parent.connections); i++ {
		if parent.connections[i].marked == false {
			markPrograms(parent.connections[i])
		}
	}

}