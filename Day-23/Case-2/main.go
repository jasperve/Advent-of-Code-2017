package main

import "fmt"

func main() {

	reg := map[string]int{	"b": 109900, "c": 126900 }

	for reg["b"] != reg["c"] {
		reg["d"] = 2
		for reg["d"] != reg["b"] {
			reg["e"] = 2
			for reg["e"] != reg["b"] {
				if reg["d"] * reg["e"] == reg["b"] {
					reg["h"]++			// NUMBER IS NOT A PRIME
					reg["e"] = reg["b"] - 1 // QUICK WAY TO EXIT LOOP FOR THIS NUMBER
					reg["d"] = reg["b"] - 1 // QUICK WAY TO EXIT LOOP FOR THIS NUMBER
				}
				reg["e"]++
			}
			reg["d"]++
		}
		reg["b"] += 17
	}

	fmt.Println(reg["h"])

}

