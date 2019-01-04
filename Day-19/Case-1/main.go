package main

import (
	"os"
	"bufio"
	"fmt"

)

const (
	up = 0
	left = 1
	right = 2
	down = 3
)

type coordinate struct {
	x int
	y int
}

type track struct {
	form rune
	direction int
}

func main() {

	grid := make(map[coordinate]*track)

	file, _ := os.Open("input.txt")
	input := bufio.NewScanner(file)

	y := 0
	for input.Scan() {
		for x, v := range input.Text() {
			grid[ coordinate{ x, y } ] = &track{ form: v }
		}
		y++
	}

	nextCoordinate := coordinate{ 15, 0 }
	grid[nextCoordinate].direction = down
	code := []rune{}

	for {


		currCoordinate := nextCoordinate

		switch grid[currCoordinate].form {
		
		case 43:
			switch grid[currCoordinate].direction {
			case up, down:
				if grid[ coordinate{ currCoordinate.x - 1, currCoordinate.y } ].form != 32 {
					nextCoordinate = coordinate{ currCoordinate.x - 1, currCoordinate.y }
					grid[nextCoordinate].direction = left
				} else if grid[ coordinate{ currCoordinate.x + 1, currCoordinate.y } ].form != 32 {
					nextCoordinate = coordinate{ currCoordinate.x + 1, currCoordinate.y }
					grid[nextCoordinate].direction = right
				}
			case left, right:
				if grid[ coordinate{ currCoordinate.x, currCoordinate.y - 1 } ].form != 32 {
					nextCoordinate = coordinate{ currCoordinate.x, currCoordinate.y - 1 }
					grid[nextCoordinate].direction = up
				} else if grid[ coordinate{ currCoordinate.x, currCoordinate.y + 1 } ].form != 32 {
					nextCoordinate = coordinate{ currCoordinate.x, currCoordinate.y + 1 }
					grid[nextCoordinate].direction = down
				}
			}
		
		default:
			switch grid[currCoordinate].direction {
			case up:
				nextCoordinate = coordinate{ currCoordinate.x, currCoordinate.y - 1 }
			case left:
				nextCoordinate = coordinate{ currCoordinate.x - 1, currCoordinate.y }
			case right:
				nextCoordinate = coordinate{ currCoordinate.x + 1, currCoordinate.y }
			case down:
				nextCoordinate = coordinate{ currCoordinate.x, currCoordinate.y + 1 }
			}
			grid[nextCoordinate].direction = grid[currCoordinate].direction

			if grid[currCoordinate].form == 32  {
				fmt.Println("nooit hier?")
				break
			} else if grid[currCoordinate].form != 43 && grid[currCoordinate].form != 45 && grid[currCoordinate].form != 124 {
				code = append(code, grid[currCoordinate].form)
			}

		}

		if nextCoordinate == currCoordinate { break }

	}

	fmt.Println(string(code))

}