package main

import (
	"fmt"
)

const searchedValue = 347991

func main() {

	gridSize := 20
	grid := make([][]int, gridSize)

	for i := 0; i < gridSize; i++ {
		grid[i] = make([]int, gridSize)
	}

	x := gridSize / 2
	y := gridSize / 2

	totalSteps := 1
	stepNumber := 0
	direction := 90
	secondTime := false

	iNumber:=1

	for {

		grid[y][x] = iNumber

		if iNumber > searchedValue {
			break
		}

		if stepNumber < totalSteps {
			switch direction {
			case 90:
				x += 1
			case 180:
				y += 1
			case 270:
				x -= 1
			case 360:
				y -= 1
			}
			stepNumber++
		}

		if stepNumber == totalSteps {
			switch direction {
			case 90:
				direction = 360
			case 180:
				direction = 90
			case 270:
				direction = 180
			case 360:
				direction = 270
			}

			stepNumber = 0

			if secondTime {
				secondTime = false;
				totalSteps++
			} else if !secondTime {
				secondTime = true
			}
		}

		iNumber = grid[y+1][x+1] + grid[y+1][x] + grid[y+1][x-1] + grid[y][x-1] + grid[y-1][x-1] + grid[y-1][x] + grid[y-1][x+1] + grid[y][x+1]

	}

	for _, v := range grid {
		fmt.Println(v)
	}

	fmt.Println("First higher value found: ", iNumber)

}