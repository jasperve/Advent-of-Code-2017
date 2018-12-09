package main

import (
	"fmt"
	"io/ioutil"
)

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