package main

import (
	"fmt"
	"strconv"
)

func main() {

	matches := 0

	a := int64(116) //116
	b := int64(299) //299

	for i := 0; i < 40000000; i++ {

		if i%10000 == 0 {
			fmt.Println(i)
		}

		a = (a*16807)%2147483647
		aBin := strconv.FormatInt(a, 2)
		b = (b*48271)%2147483647
		bBin := strconv.FormatInt(b, 2)

		if len(aBin) < 16 || len(bBin) < 16 {
			continue
		}

		if aBin[len(aBin)-16:] == bBin[len(bBin)-16:] {
			matches++
		}

	}

	fmt.Println(matches)

}