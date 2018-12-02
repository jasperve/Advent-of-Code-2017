package main

import (
	"strings"
	"io/ioutil"
	"strconv"
	"fmt"
)

func main() {

	input, _ := ioutil.ReadFile("input.txt")

	banks := make(map[int]int)
	for i, blocks := range strings.Split(string(input), "\t"){
		blocksInt, _ := strconv.Atoi(blocks)
		banks[i] = blocksInt
	}

	allBanks := make(map[string]int)
	firstDouble := ""

	for {

		bank := ""
		for i:=0; i < len(banks); i++ {
			bank = fmt.Sprintf("%v %v", bank, banks[i])
		}

		if _, ok := allBanks[bank]; ok {
			firstDouble = bank
			break
		}

		allBanks[bank] = len(allBanks)

		maxBlocksIndex := -1
		maxBlocksValue := -1
		for i:=0; i < len(banks); i++ {
			if banks[i] > maxBlocksValue {
				maxBlocksIndex = i
				maxBlocksValue = banks[i]
			}
		}

		banks[maxBlocksIndex] = 0

		for maxBlocksValue > 0 {

			maxBlocksIndex++
			banks[maxBlocksIndex%len(banks)]++
			maxBlocksValue--

		}

	}

	fmt.Printf("Number of redistrubition: %v\n", len(allBanks))
	fmt.Printf("Cycles between doubles: %v\n", len(allBanks) - allBanks[firstDouble])

}