package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"regexp"
	"strings"
	"strconv"
)

type program struct {
	id string
	parent *program
	weight int
	weightChildren int
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

		stringId := line[0][1]
		intWeight, _ := strconv.Atoi(line[0][2])
		stringChildren := line[0][4]

		var parent *program
		children := []*program{}

		if programs[stringId] == nil {
			parent = &program{id: stringId, weight: intWeight, children: children}
			programs[stringId] = parent
		} else {
			parent = programs[stringId]
			parent.weight = intWeight
			children = parent.children
		}

		if stringChildren != "" {

			for _, v := range strings.Split(stringChildren, ",") {

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

		parent.children = children

	}

	topId := "aapssr"
	//topId := "tknk"

	for _, v := range programs {
		if v.id == topId {
			calculateWeightChildren(v)
			compareWeightChildren(v)
		}
	}

}

func calculateWeightChildren(parent *program) {

	if parent.children == nil {
		return
	} else {
		for _, child := range parent.children {
			calculateWeightChildren(child)
			parent.weightChildren = parent.weightChildren + child.weight + child.weightChildren
		}
	}

}

func compareWeightChildren(parent *program) {

	if parent.children == nil {
		return
	} else {

		for _, child := range parent.children {
			compareWeightChildren(child)
		}

		weights := make(map[int]int)
		for _, child := range parent.children {
			weights[child.weight + child.weightChildren]++
		}

		for i, v := range weights {

			if len(weights) > 1 && v == 1 {

				for _, child := range parent.children {
					if(child.weight + child.weightChildren == i) {
						fmt.Printf("We have found a problem at tower %v with total value %v and unit value %v. Below you can see the difference in value\n", child.id, i, child.weight)
						fmt.Println(weights)
						os.Exit(1)
					}
				}

			}
		}

	}

}