package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const hashLength = 256

func main() {

	inputOriginal, _ := ioutil.ReadFile("input.txt")

	total := 0

	for r := 0; r < 128; r++ {

		inputString := fmt.Sprintf("%v-%v", string(inputOriginal), r)
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

		denseHash := []int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
		for i := 0; i < 16; i++ {
			for ii := 0; ii < 16; ii++ {
				denseHash[i] = denseHash[i] ^ sparseHash[i*16+ii]
			}
		}
		denseHashHex := ""
		for _, v := range denseHash {
			denseHashHex = fmt.Sprintf("%s%02x", denseHashHex, v)
		}
		
		denseHashBin := ""
		for d := range denseHashHex {
			i, _ := strconv.ParseInt(string(denseHashHex[d]), 16, 64)
			denseHashBin = fmt.Sprintf("%v%04v", denseHashBin, strconv.FormatInt(i, 2))
		}
		
		total += strings.Count(denseHashBin, "1")

	}

	fmt.Println(total)
}