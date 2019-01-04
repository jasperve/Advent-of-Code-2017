package main

import "fmt"

func main() {

	buffer := []int{0,1}
	position := 1

	bufferAtOne := []int{}

	for i := 2; i <= 500000; i++ {

		if i%10000 == 0 {
			fmt.Println(i)
		}

		position = (position+344)%len(buffer)+1
		buffer = append(buffer[:position], append([]int{i}, buffer[position:]...)...)

		bufferAtOne = append(bufferAtOne, buffer[1])

	}

	encountered := map[int]bool{}
	result := []int{}

	for v := range bufferAtOne {
		if encountered[bufferAtOne[v]] == true {
		} else {
			encountered[bufferAtOne[v]] = true
			result = append(result, bufferAtOne[v])
		}
	}

	fmt.Println(result)

}