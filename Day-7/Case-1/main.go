package main

import (
	"bufio"
	"fmt"
	"os"
	"log"

)

type program struct {
	parent *program
	children []*program
}

func main() {

	//programs := make(map[string]program)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	input := bufio.NewScanner(file)

	for input.Scan() {
		fmt.Println(input.Text())
	}

}