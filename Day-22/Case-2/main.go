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

	cleaned = 46
	weakened = 87
	infected = 35
	flagged = 70
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

	for i := 0; i < 10000000; i++ {
	
		if _, ok := grid[currentNode]; !ok || grid[currentNode] == cleaned {
			currentDirection = modulo(currentDirection-1, 4)
			grid[currentNode] = weakened
		} else if grid[currentNode] == weakened {
			grid[currentNode] = infected
			infectionCounter++
		} else if grid[currentNode] == infected {
			currentDirection = modulo(currentDirection+1, 4)
			grid[currentNode] = flagged
		} else if grid[currentNode] == flagged {
			currentDirection = modulo(currentDirection+2, 4)
			grid[currentNode] = cleaned
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

	fmt.Println(infectionCounter)

}


func modulo(operand, modulus int) (result int) {
	result = operand % modulus
	if result < 0 && modulus > 0 {
	   return result + modulus
	}
	return result
 }