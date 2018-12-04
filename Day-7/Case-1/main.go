package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"regexp"
	"strings"
)

type program struct {
	id string
	parent *program
	weight int
	children []*program
}

var programs = map[string]*program{}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	input := bufio.NewScanner(file)

	for input.Scan() {
		line := regexp.MustCompile("(.*)\\s\\((\\d*)\\)(\\s->\\s(.*))?").FindAllStringSubmatch(input.Text(), -1)

		var parent *program
		children := []*program{}

		if programs[line[0][1]] == nil {
			parent = &program{id: line[0][1], children: children}
			programs[line[0][1]] = parent
		} else {
			parent = programs[line[0][1]]
			children = parent.children
		}

		if line[0][4] != "" {
			for _, v := range strings.Split(line[0][4], ",") {

				var child *program
				v = strings.TrimSpace(v)
				if programs[v] == nil {
					child = &program{id: v, parent: parent,}
					programs[v] = child
				} else {
					child = programs[v]
					child.parent = parent
				}
				children = append(children, child)

			}
		}

	}

	currentId := "iwdjjf"

	for {

		for _, v := range programs {

			if v.id == currentId {

				fmt.Println(currentId)

				if v.parent == nil {
					fmt.Printf("Original starting point found: %v\n", currentId)
					return
				} else {
					currentId = (*v.parent).id
					continue
				}
			}

		}
	}

}