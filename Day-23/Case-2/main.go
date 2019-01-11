package main

import "fmt"

func main() {

	reg := map[string]int{	"b": 109900, "c": 126900 }

	OUTER:
	for {

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
				if reg["e"] == reg["b"] {
					break ELOOP
				}

			}

			reg["d"]++
			if reg["d"] == reg["b"] {
				break DLOOP
			}

		}

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

