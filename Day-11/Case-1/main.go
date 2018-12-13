package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"math"
)

func main() {

	input, _ := ioutil.ReadFile("input.txt")

	x, y := 0.0, 0.0
	maxX, maxY := 0.0, 0.0

	for _, v := range strings.Split(string(input), ",") {
		switch v {
		case "n":
			y-=1
		case "ne":
			x+=0.5
			y-=0.5
		case "se":
			x+=0.5
			y+=0.5
		case "s":
			y+=1
		case "sw":
			x-=0.5
			y+=0.5
		case "nw":
			x-=0.5
			y-=0.5
		}

		fmt.Println(x+y)

		if math.Abs(x) + math.Abs(y) > maxX + maxY {
			maxX = math.Abs(x)
			maxY = math.Abs(y)
		}

	}

	x = math.Abs(maxX)
	y = math.Abs(maxY)

	difference := math.Abs(x - y)

	heighest := 0.0

	if x > y {
		heighest = x
	} else {
		heighest = y
	}

	fmt.Println()
	fmt.Println(heighest)
	fmt.Println(difference)
	fmt.Println()

	steps := (heighest - difference)/0.5 + difference

	fmt.Println(steps)


}