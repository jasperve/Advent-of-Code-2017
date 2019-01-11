package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	off = 46
	on  = 35
)

type coordinate struct {
	x int
	y int
}

type pattern struct {
	size        int
	match       map[coordinate]rune
	replacement map[coordinate]rune
}

var patterns []pattern
var grid = map[coordinate]rune{
	coordinate{0, 0}: 46,
	coordinate{1, 0}: 35,
	coordinate{2, 0}: 46,
	coordinate{0, 1}: 46,
	coordinate{1, 1}: 46,
	coordinate{2, 1}: 35,
	coordinate{0, 2}: 35,
	coordinate{1, 2}: 35,
	coordinate{2, 2}: 35,
}

func main() {

	loadPatterns()

	printGrid()
	fmt.Println()

	enhanceGrid(5)
	printGrid()
	fmt.Println()

	totalCount := 0

	for _, v := range grid {
		if v == on {
			totalCount++
		}
	}

	fmt.Println(totalCount)

}

func loadPatterns() {

	file, _ := os.Open("input.txt")
	input := bufio.NewScanner(file)

	for input.Scan() {

		patternSplit := strings.Split(input.Text(), " => ")
		patternMatchS := strings.Split(patternSplit[0], "/")
		patternReplacementS := strings.Split(patternSplit[1], "/")

		match := make(map[coordinate]rune)
		for y := 0; y < len(patternMatchS); y++ {
			for x, v := range patternMatchS[y] {
				match[coordinate{x, y}] = v
			}
		}

		replacement := make(map[coordinate]rune)
		for y := 0; y < len(patternReplacementS); y++ {
			for x, v := range patternReplacementS[y] {
				replacement[coordinate{x, y}] = v
			}
		}

		newPattern := pattern{
			size:        len(patternMatchS),
			match:       match,
			replacement: replacement,
		}

		patterns = append(patterns, newPattern)

	}

}

func enhanceGrid(amount int) {

	for i := 0; i < amount; i++ {

		newGrid := make(map[coordinate]rune)

		gridDivisibleBy := 2
		if len(grid)%3 == 0 {
			gridDivisibleBy = 3
		}

		for x := 0; x < int(math.Sqrt(float64(len(grid))))/gridDivisibleBy; x++ {

			for y := 0; y < int(math.Sqrt(float64(len(grid))))/gridDivisibleBy; y++ {

				section := make(map[coordinate]rune)
				for subX := 0; subX < gridDivisibleBy; subX++ {
					for subY := 0; subY < gridDivisibleBy; subY++ {
						section[coordinate{subX, subY}] = grid[coordinate{x*gridDivisibleBy+subX, y*gridDivisibleBy+subY}]
					}
				}

				possibleSections := generateAlternatives(section, gridDivisibleBy)
				
				SECTIONSLOOP:
				for _, s := range possibleSections {

					PATTERNLOOP:
					for kp, p := range patterns {

						if p.size != gridDivisibleBy {
							continue
						}

						for k, v := range s {

							if p.match[k] != v {
								continue PATTERNLOOP
							}

						}

						fmt.Println("pattern", kp)

						for subY := 0; subY < gridDivisibleBy+1; subY++ {
							for subX := 0; subX < gridDivisibleBy+1; subX++ {
								newGrid[coordinate{x*(gridDivisibleBy+1)+subX, y*(gridDivisibleBy+1)+subY}] = p.replacement[coordinate{subX, subY}]
							}
						}

						break SECTIONSLOOP

					}

				}

			}
		}

		grid = newGrid
		
	}

}

func generateAlternatives(section map[coordinate]rune, size int) (sections []map[coordinate]rune) {

	sections = append(sections, section)
	sections = append(sections, rotate(section, size, 1))
	sections = append(sections, rotate(section, size, 2))
	sections = append(sections, rotate(section, size, 3))

	flippedSection := flip(section, size)
	sections = append(sections, flippedSection)
	sections = append(sections, rotate(flippedSection, size, 1))
	sections = append(sections, rotate(flippedSection, size, 2))
	sections = append(sections, rotate(flippedSection, size, 3))

	return sections

}

func flip(section map[coordinate]rune, size int) map[coordinate]rune {

	result := make(map[coordinate]rune)
	for k, v := range section {
		result[k] = v
	}
	for i := 0; i < size; i++ {
		result[coordinate{i, 0}] = section[coordinate{i, size - 1}]
	}
	for i := 0; i < size; i++ {
		result[coordinate{i, size - 1}] = section[coordinate{i, 0}]
	}
	return result

}

func rotate(section map[coordinate]rune, size int, times int) map[coordinate]rune {

	result := make(map[coordinate]rune)
	for k, v := range section {
		result[k] = v
	}

	if size == 2 {
		for i := 0; i < times; i++ {
			result[coordinate{0, 0}], result[coordinate{1, 0}],
				result[coordinate{1, 1}], result[coordinate{0, 1}] =
				result[coordinate{0, 1}], result[coordinate{0, 0}],
				result[coordinate{1, 0}], result[coordinate{1, 1}]
		}
	}

	if size == 3 {
		for i := 0; i < times; i++ {
			result[coordinate{0, 0}], result[coordinate{1, 0}], result[coordinate{2, 0}],
				result[coordinate{2, 1}],
				result[coordinate{2, 2}], result[coordinate{1, 2}], result[coordinate{0, 2}],
				result[coordinate{0, 1}] =
				result[coordinate{0, 2}], result[coordinate{0, 1}], result[coordinate{0, 0}],
				result[coordinate{1, 0}], 
				result[coordinate{2, 0}], result[coordinate{2, 1}], result[coordinate{2, 2}], 
				result[coordinate{1, 2}]

		}
	}

	return result

}

func printGrid() {

	for y := 0; y < int(math.Sqrt(float64(len(grid)))); y++ {
		for x := 0; x < int(math.Sqrt(float64(len(grid)))); x++ {
			if grid[coordinate{x, y}] == off {
				fmt.Printf(".")
			} else if grid[coordinate{x, y}] == on {
				fmt.Printf("#")
			}
		}
		fmt.Printf("\n")
	}

}