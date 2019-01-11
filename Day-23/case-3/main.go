package main

import (
	"math/big"
	"fmt"
)

func main() {

	notPrimeCounter := 0

	for i := 109900; i <= 126900; i +=17 {
		if !big.NewInt(int64(i)).ProbablyPrime(0) {
			notPrimeCounter++
		}
	}

	fmt.Println(notPrimeCounter)

}
