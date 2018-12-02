package main

import (
	"fmt"
	"math"
)

const highestNumber = 347991

func main() {

	x := 0
	y := 0

	totalSteps := 1
	stepNumber := 0
	direction := 90
	secondTime := false

	for iNumber:=1; iNumber < highestNumber; iNumber++ {

		if stepNumber < totalSteps {
			switch direction {
			case 90:
				y += 1
			case 180:
				x += 1
			case 270:
				y -= 1
			case 360:
				x -= 1
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
				secondTime = false;	totalSteps++
			} else if !secondTime {
				secondTime = true
			}

		}

}

	fmt.Printf( "Amount of steps needed: %v", math.Abs(float64(x)) + math.Abs(float64(y)))

	}