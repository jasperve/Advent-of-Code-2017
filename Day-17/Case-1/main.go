package main

import "fmt"

func main() {

	buffer := []int{0,1}
	position := 1

	for i := 2; i <= 2017; i++ {
		position = (position+344)%len(buffer)+1
		buffer = append(buffer[:position], append([]int{i}, buffer[position:]...)...)
	}

	fmt.Println(buffer[position+1])

}