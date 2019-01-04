package main

import "fmt"

type buffer struct {
	left *buffer
	right *buffer
	value int
}

func main() {

	firstBuffer := &buffer{ value: 0 }
	firstBuffer.left = firstBuffer
	firstBuffer.right = firstBuffer
	curBuffer := firstBuffer

	for i := 1; i <= 50000000; i++ {

		if i%10000 == 0 {
			fmt.Println(i)
		}

		nextBuffer := curBuffer.getNeighbour(344)
		newBuffer := &buffer{ left: nextBuffer, right: nextBuffer.right, value: i }
		nextBuffer.right.left = newBuffer
		nextBuffer.right = newBuffer
		curBuffer = newBuffer
	}

	fmt.Println(firstBuffer.getNeighbour(1).value)

}

func (current *buffer) getNeighbour(offset int) *buffer {
	var actual *buffer
	if offset < 0 {
		actual = current.left.getNeighbour(offset + 1)
	} else if offset > 0 {
		actual = current.right.getNeighbour(offset - 1)
	} else if offset == 0 {
		actual = current
	}
	return actual
}