package main

import (
	"fmt"
	"io/ioutil"
	//"math"
	"strconv"
)

const hashLength = 256

type coordinate struct {
	x int
	y int
}

type square struct {
	location coordinate
	value int
	checked bool
}

type region struct {
	squares []*square
}

func main() {

	squares := make(map[coordinate]*square)

	inputOriginal, _ := ioutil.ReadFile("input.txt")

	for y := 0; y < 128; y++ {

		inputString := fmt.Sprintf("%v-%v", string(inputOriginal), y)
		inputBytes := []byte(inputString)
		inputBytes = append(inputBytes, 17, 31, 73, 47, 23)

		sparseHash := []int{}
		for i := 0; i < hashLength; i++ {
			sparseHash = append(sparseHash, i)
		}

		position := 0
		skip := 0

		for run := 0; run < 64; run++ {
			for _, v := range inputBytes {

				tempList := []int{}
				for i := 0; i < int(v); i++ {
					tempList = append(tempList, sparseHash[(position+i)%len(sparseHash)])
				}

				for i := 0; i < int(v); i++ {
					sparseHash[(position+i)%len(sparseHash)] = tempList[(len(tempList)-1-i)%len(sparseHash)]
				}

				position += int(v) + skip
				skip++
			}
		}

		denseHash := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for i := 0; i < 16; i++ {
			for ii := 0; ii < 16; ii++ {
				denseHash[i] = denseHash[i] ^ sparseHash[i*16+ii]
			}
		}
		denseHashHex := ""
		for d := range denseHash {
			denseHashHex = fmt.Sprintf("%s%02x", denseHashHex, denseHash[d])
		}

		denseHashBin := ""
		for d := range denseHashHex {
			i, _ := strconv.ParseInt(string(denseHashHex[d]), 16, 64)
			denseHashBin = fmt.Sprintf("%v%04v", denseHashBin, strconv.FormatInt(i, 2))
		}

		for d := range denseHashBin {
			squares[ coordinate{ d, y } ] = &square{ location: coordinate{ d, y }, value: int(denseHashBin[d]) }
		}

	}

	regions := [][]*square{}
	neighbours := []coordinate{ coordinate{x: 0, y: -1}, coordinate{x: -1, y: 0}, coordinate{x: 1, y: 0}, coordinate{x: 0, y: 1}}

	for len(squares) > 0 {
		
		var checkedSquare *square
		for k, v := range squares {
			checkedSquare = v
			delete(squares, k)
			break
		}

		if checkedSquare.value != 49 {
			continue			
		}

		newRegion := []*square{}
		openList := []*square{}
		openList = append(openList, checkedSquare)

		for len(openList) > 0 {

			childSquare := openList[0]
			openList = append([]*square{}, openList[1:]...)

			if childSquare.value != 49 {
				continue
			}

			newRegion = append(newRegion, childSquare)

			for _, neighbour := range neighbours {

				if _, ok := squares[ coordinate{ childSquare.location.x + neighbour.x, childSquare.location.y + neighbour.y } ]; !ok {
					continue
				}

				openList = append(openList, squares[ coordinate{ childSquare.location.x + neighbour.x, childSquare.location.y + neighbour.y } ])
				delete(squares, coordinate{ childSquare.location.x + neighbour.x, childSquare.location.y + neighbour.y } )

			}

		}

		regions = append(regions, newRegion)

	}

	fmt.Println(len(regions))



}