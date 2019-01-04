package main

import ( 
	"os"
	"bufio"
	"fmt"
)

const (
	up = 0
	right = 1
	down = 2
	left = 3	
)

type coordinate struct {
	x int
	y int
}


func main() {

	grid := make(map[coordinate]rune)
	minX, maxX, minY, maxY := 0, 0, 0, 0

	file, _ := os.Open("input.txt")
	input := bufio.NewScanner(file)

	y := 0
	for input.Scan() {
		for x, infected := range input.Text() {
			grid[coordinate{ x, y }] = infected
			if x > maxX { maxX = x }
		}	
		y++ 
	}
	maxY = y-1

	currentNode := coordinate{ (maxX+1-minX)/2, (maxY+1-minY)/2 }
	currentDirection := up

	infectionCounter := 0

	for i := 0; i < 10000; i++ {
		
		if _, ok := grid[currentNode]; !ok || grid[currentNode] == 46 {
			currentDirection = modulo(currentDirection-1, 4)
			grid[currentNode] = 35
			infectionCounter++
		} else {
			currentDirection = modulo(currentDirection+1, 4)
			grid[currentNode] = 46
		}

		switch currentDirection {
		case up:
			currentNode = coordinate{ currentNode.x, currentNode.y-1 }
		case right: 
			currentNode = coordinate{ currentNode.x+1, currentNode.y }
		case down:
			currentNode = coordinate{ currentNode.x, currentNode.y+1 }
		case left:
			currentNode = coordinate{ currentNode.x-1, currentNode.y }
		}
		
	}

	for k := range grid {
		if k.x < minX { minX = k.x}
		if k.x > maxX { maxX = k.x}
		if k.y < minX { minY = k.y}
		if k.y > maxX { maxY = k.y}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, ok := grid[coordinate{ x, y}]; !ok || grid[coordinate{ x, y}] == 46 {
				fmt.Printf(".")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Printf("\n")
	}



	fmt.Println(infectionCounter)

}


func modulo(operand, modulus int) (result int) {
	result = operand % modulus
	if result < 0 && modulus > 0 {
	   return result + modulus
	}
	return result
 }