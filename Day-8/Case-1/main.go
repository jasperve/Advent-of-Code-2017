package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"regexp"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	input := bufio.NewScanner(file)

	register := make(map[string]int)
	maxRegisterValue := 0

	for input.Scan() {
		line := regexp.MustCompile("([a-z]+)\\s(inc|dec)\\s(-?\\d+)\\sif\\s+([a-z]+)\\s+(.+)+\\s+(-?\\d+)").FindAllStringSubmatch(input.Text(), -1)

		operand := line[0][1]
		operator := line[0][2]
		value, _ := strconv.Atoi(line[0][3])
		conditionVariable := line[0][4]
		conditionOperator := line[0][5]
		conditionValue, _ := strconv.Atoi(line[0][6])

		conditionValidates := false

		switch conditionOperator {
		case "==":
			if register[conditionVariable] == conditionValue {
				conditionValidates = true
			}
		case ">=":
			if register[conditionVariable] >= conditionValue  {
				conditionValidates = true
			}
		case "<=":
			if register[conditionVariable] <= conditionValue {
				conditionValidates = true
			}
		case "!=":
			if register[conditionVariable] != conditionValue {
				conditionValidates = true
			}
		case ">":
			if register[conditionVariable] > conditionValue  {
				conditionValidates = true
			}
		case "<":
			if register[conditionVariable] < conditionValue {
				conditionValidates = true
			}
		}

		if conditionValidates {

			switch operator {
			case "inc":
				register[operand] += value
				if register[operand] > maxRegisterValue {
					maxRegisterValue = register[operand]
				}
			case "dec":
				register[operand] -= value
			}


		}

	}

	largestRegister := ""
	registerValue := 0

	for i, v := range register {
		if v > registerValue {
			largestRegister = i
			registerValue = v
		}
	}

	fmt.Printf("The largest register in the end is %v with the value %v. Heighest overall value reached is %v\n", largestRegister, registerValue, maxRegisterValue)

}