package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

type component struct {
	leftPort int
	rightPort int
}

type bridge struct {
	configuration []*component
	lastInverted bool
}

var components []*component
var bridges []*bridge

func main() {

	file, _ := os.Open("input.txt")
	input := bufio.NewScanner(file)

	for input.Scan() {

		ports := strings.Split(input.Text(), "/")
		leftPort, _ := strconv.Atoi(ports[0])
		rightPort, _ := strconv.Atoi(ports[1])
		components = append(components, &component{ leftPort: leftPort, rightPort: rightPort })

	}

	buildBrige(bridge{})

	strongestBridgeStrength := 0
	var strongestBridge *bridge

	for b := 0; b < len(bridges); b++ {

		strength := 0
		for c := 0 ; c < len(bridges[b].configuration); c++ {
			strength += bridges[b].configuration[c].leftPort + bridges[b].configuration[c].rightPort
		}

		if strength > strongestBridgeStrength { 
			strongestBridgeStrength = strength 
			strongestBridge = bridges[b]
		}

	}

	for c := 0; c < len(components); c++ {
		fmt.Println(components[c])
	}


	if strongestBridge != nil {
		fmt.Printf("Strongest bridge has a strength of %v and consists of the following components:\n", strongestBridgeStrength)

		for c := 0; c < len(strongestBridge.configuration); c++ {
			fmt.Println(&strongestBridge.configuration[c], strongestBridge.configuration[c].leftPort, strongestBridge.configuration[c].rightPort)
		}

	}
	
}


func buildBrige(possibleBridge bridge) {

	// Define the left port we're searching for
	searchedPort := 0
	if len(possibleBridge.configuration) > 0 {
		if possibleBridge.lastInverted {
			searchedPort = possibleBridge.configuration[len(possibleBridge.configuration)-1].rightPort
		} else {
			searchedPort = possibleBridge.configuration[len(possibleBridge.configuration)-1].leftPort
		}
	}

	OUTER:
	for _, c := range components {

		// Check if either one of the ports matches the search port
		if c.leftPort != searchedPort && c.rightPort != searchedPort { continue }

		// Check if this component has already been used
		for i := 0; i < len(possibleBridge.configuration); i++ {
			if possibleBridge.configuration[i] == c { continue OUTER }
		}

		// Check if we need to set the lastInverted to true/false
		inverted := false
		if c.rightPort != searchedPort {
			inverted = true
		}

		newPossibleBridge := bridge { configuration: append(possibleBridge.configuration, c), lastInverted: inverted }
		//fmt.Println(newPossibleBridge)
		buildBrige(newPossibleBridge)
		
	}

	bridges = append(bridges, &possibleBridge)

}