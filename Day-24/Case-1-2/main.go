package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type component struct {
	leftPort  int
	rightPort int
}

type bridge struct {
	configuration []*component
	lastInverted  bool
}

var components []component

var strongestBridgeStrength int
var longestBridgeLength int
var longestBridgeStrength int

func main() {

	file, _ := os.Open("input.txt")
	input := bufio.NewScanner(file)

	for input.Scan() {

		ports := strings.Split(input.Text(), "/")
		leftPort, _ := strconv.Atoi(ports[0])
		rightPort, _ := strconv.Atoi(ports[1])
		components = append(components, component{leftPort: leftPort, rightPort: rightPort})

	}

	buildBrige(bridge{})

	fmt.Printf("Strongest bridge has a strength of %v.\n", strongestBridgeStrength)
	fmt.Printf("Longest bridge has a strength of %v.\n", longestBridgeStrength)

}

func buildBrige(possibleBridge bridge) {

	searchedPort := 0
	if len(possibleBridge.configuration) > 0 {
		if possibleBridge.lastInverted {
			searchedPort = possibleBridge.configuration[len(possibleBridge.configuration)-1].rightPort
		} else {
			searchedPort = possibleBridge.configuration[len(possibleBridge.configuration)-1].leftPort
		}
	}

	connectingComponentFound := false

OUTER:
	for c := 0; c < len(components); c++ {

		if components[c].leftPort != searchedPort && components[c].rightPort != searchedPort {
			continue
		}

		for i := 0; i < len(possibleBridge.configuration); i++ {
			if possibleBridge.configuration[i] == &components[c] {
				continue OUTER
			}
		}

		connectingComponentFound = true

		inverted := false
		if components[c].rightPort != searchedPort {
			inverted = true
		}

		newPossibleBridge := bridge{configuration: append(possibleBridge.configuration, &components[c]), lastInverted: inverted}
		buildBrige(newPossibleBridge)

	}

	if !connectingComponentFound {

		strength := 0
		for c := 0; c < len(possibleBridge.configuration); c++ {
			strength += possibleBridge.configuration[c].leftPort + possibleBridge.configuration[c].rightPort
			if strength > strongestBridgeStrength {
				strongestBridgeStrength = strength
			}
			if len(possibleBridge.configuration) >= longestBridgeLength {
				longestBridgeLength = len(possibleBridge.configuration)
				if strength >= longestBridgeStrength {
					longestBridgeStrength = strength
				}
			}
		}

	}

}
