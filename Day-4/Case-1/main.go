package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {

	totalValid := 0

	input, _ := ioutil.ReadFile("input.txt")

	OUTER:
	for _, passphrase := range strings.Split(string(input), "\n"){

		if string(passphrase) == "" { continue OUTER }

		words := make(map[string]int)
		for _, word := range strings.Split(string(passphrase), " ")  {
			words[string(word)]++
		}

		for _, v := range words {
			if v > 1 {	continue OUTER }
		}

		totalValid++

	}

	fmt.Printf("Total valid passphrases: %v\n", totalValid)

}