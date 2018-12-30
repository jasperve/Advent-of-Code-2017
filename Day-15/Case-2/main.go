package main

import (
	"fmt"
	"strconv"
)

func main() {

	matches := 0

	a := int64(116) //116
	b := int64(299) //299
	aOld := a
	bOld := b

	for i := 0; i < 5000000; i++ {

		if i%10000 == 0 {
			fmt.Println(i)
		}

		for a%4 != 0 || a == aOld {
			a = (a*16807)%2147483647
		}
		aOld = a

		for b%8 != 0 || b == bOld {
			b = (b*48271)%2147483647
		}
		bOld = b
		
		aBin := strconv.FormatInt(a, 2)
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