package main

import (
	"fmt"
	"io/ioutil"
)

<<<<<<< HEAD
func main() {

	list := make([]int, 256)
	for i := 0; i < len(list); i++ {
		list[i] = i
	}
	lengths := []int{14,58,0,116,179,16,1,104,2,254,167,86,255,55,122,244}
	skips := 0

	listIndex := 0

	for _, v := range lengths {

		subList := append(list[:0])


	}

}
=======
const hashLength = 256

func main() {

	sparseHash := []int{}
	for i := 0; i < hashLength; i++ {
		sparseHash = append(sparseHash, i)
	}

	input, _ := ioutil.ReadFile("input.txt")
	input = append(input, 3, 4, 1, 5, 17, 31, 73, 47, 23)

	position := 0
	skip := 0

	for run := 0; run < 64; run++ {
		for _, v := range input {

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
			fmt.Println(sparseHash[i*16+ii])
			denseHash[i] = denseHash[i] ^ sparseHash[i*16+ii]
		}
	}

	for _, v := range denseHash {
		fmt.Printf("%02x", v)
	}

	fmt.Println(denseHash)

}
>>>>>>> 219b139bdf690b051985cdd4ad19daf82128f845
