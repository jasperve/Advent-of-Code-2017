package main

import "fmt"

func main() {

	reg := map[string]int{	"b": 109900, "c": 126900 }

	reg["c"] = 110274

	numberLoops := 0

	OUTER:
	for {
		numberLoops++

		reg["f"] = 1
		reg["d"] = 2

		DLOOP: 
		for {

			reg["e"] = 2

			ELOOP:
			for {
	
				if reg["d"] * reg["e"] == reg["b"] {
					reg["f"] = 0
					break DLOOP
				}

				reg["e"]++
				if reg["e"] - reg["b"] == 0 {
					break ELOOP
				}

			}

			reg["d"]++
			if reg["d"] - reg["b"] == 0 {
				break DLOOP
			}

		}

		fmt.Println(reg["b"], reg["h"])

		if reg["f"] == 0 {
			reg["h"]++
		}

		if reg["b"] - reg["c"] == 0 {
			break OUTER
		}

		reg["b"] += 17
 
	}

	fmt.Println(reg["h"])
	
}

