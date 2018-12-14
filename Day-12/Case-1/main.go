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
	children []*program
}

func main() {

	programs := make(map[int]*program)
	
	file, _ := os.Open("input-test.txt")
	input := bufio.NewScanner(file)

	for input.Scan() {

		line := strings.Split(input.Text(), " <-> ")
		lineRight := strings.Split(line[1], ", ")

		id, _ := strconv.Atoi(line[0])

		parent := program{}
		if _, ok := programs[id]; ok {
			parent = *programs[id]
		} else {
			parent.id = id
			programs[id] = &parent
		}
		
		children := []*program{}
		for _, v := range lineRight {
			childID, _ := strconv.Atoi(v)
			if _, ok := programs[childID]; ok {
				children = append(children, programs[childID])
			} else {
				child := program{ id: childID }
				children = append(children, &child)
				programs[childID] = &child
			}
		}

		parent.children = children
		programs[id] = &parent

	}

	markPrograms(programs[3])

	for i := 0; i < len(programs); i++ {
		fmt.Println(programs[i].id, programs[i].marked, programs[i].children)
	}
	
}


func markPrograms(parent *program) {

	fmt.Println("Markprograms called for ", parent.id)
	parent.marked = true
	for i := 0; i < len(parent.children); i++ {
		if parent.children[i].marked == false {
			markPrograms(parent.children[i])
		}
	}

}