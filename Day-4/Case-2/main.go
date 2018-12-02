package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func main() {

	totalValid := 0

	input, _ := ioutil.ReadFile("input.txt")

	OUTER:
	for _, passphrase := range strings.Split(string(input), "\n"){

		if string(passphrase) == "" { continue OUTER }

		words := make(map[string]int)
		for _, word := range strings.Split(string(passphrase), " ")  {
			words[SortString(string(word))]++
		}

		for _, v := range words {
			if v > 1 {	continue OUTER }
		}

		totalValid++

	}

	fmt.Printf("Total valid passphrases: %v\n", totalValid)

}