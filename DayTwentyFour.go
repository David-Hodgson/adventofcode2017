package adventofcode2017

import (
	"fmt"
	"strings"
	"strconv"
)

type port struct { 
	inputPins int
	outputPins int
}

func buildPortList(input string) []port {

	inputList := strings.Split(input, "\n")
	portList := make([]port, len(inputList))

	for i := 0 ; i < len(inputList); i++ {
		portParts := strings.Split(inputList[i], "/")

		inputPins, _ := strconv.Atoi(portParts[0])
		outputPins, _ := strconv.Atoi(portParts[1])
		portList[i] = port{inputPins, outputPins}
	}

	return portList
}

func buildBridges(pinSizeRequired int, ports []port) [][]port {

	bridges := make([][]port, 0)
	bridgeCount := 0
	//find nodes that have pin size
	for i := 0 ; i < len(ports) ; i++ {

		prt := ports[i]

		if prt.inputPins == pinSizeRequired || prt.outputPins == pinSizeRequired {

			//we've found a starting point
			nextPinSize := -1
			if prt.inputPins == pinSizeRequired {
				nextPinSize = prt.outputPins
			} else {
				nextPinSize = prt.inputPins
			}

			var remainingPorts []port
			for j := 0; j<len(ports) ; j++ {
				if i != j {
					remainingPorts = append(remainingPorts, ports[j])
				}
			}
			otherBridges := buildBridges(nextPinSize, remainingPorts)

			if len(otherBridges) > 0 {

				for k :=0 ; k < len(otherBridges) ; k++ {
					newBridge := make([]port, 1)
					newBridge[0] = prt
					newBridge = append(newBridge, otherBridges[k]...)
					bridges = append(bridges,newBridge )
					bridgeCount++

				}
			} else {
				newBridge:= make([]port,1)
				newBridge[0] = prt
				bridges = append(bridges,newBridge)
				bridgeCount++
			}
		}
	}
	return bridges
}

func DayTwentyFourExample() {

	fmt.Println("Day 24 - Example")

	input := "0/2\n2/2\n2/3\n3/4\n3/5\n0/1\n10/1\n9/10"

	ports := buildPortList(input)

	fmt.Println(ports)

	pinSizeRequired := 0
	bridges := buildBridges(pinSizeRequired,ports)

	fmt.Println(bridges)

	for i:=0 ; i< len(bridges); i++ {
		fmt.Println(bridges[i])
	
		size:= 0
		for j:=0; j < len(bridges[i]) ; j++ {
			size += bridges[i][j].inputPins + bridges[i][j].outputPins

		}

		fmt.Println("Size:", size)
	}
}

func DayTwentyFourPartOne() {

	fmt.Println("Day 24 - Part One")

	input := ReadFile("day24-input.txt")

	ports := buildPortList(input)

	fmt.Println(ports)

	pinSizeRequired := 0
	bridges := buildBridges(pinSizeRequired,ports)
	maxSize := 0

	for i:=0 ; i< len(bridges); i++ {
		size:= 0
		for j:=0; j < len(bridges[i]) ; j++ {
			size += bridges[i][j].inputPins + bridges[i][j].outputPins

		}

		if size > maxSize {
			maxSize = size
		}
	}

	fmt.Println("Max Size:", maxSize)
}


func DayTwentyFourPartTwo() {

	fmt.Println("Day 24 - Part Two")

	input := ReadFile("day24-input.txt")

	ports := buildPortList(input)

	fmt.Println(ports)

	pinSizeRequired := 0
	bridges := buildBridges(pinSizeRequired,ports)

	bridgeStrengths := make(map[int]int)

	for i:=0 ; i< len(bridges); i++ {

		bridgeLength := len(bridges[i])
		size:= 0
		for j:=0; j < len(bridges[i]) ; j++ {
			size += bridges[i][j].inputPins + bridges[i][j].outputPins
		}

		maxSize := bridgeStrengths[bridgeLength]
		if size > maxSize {
			bridgeStrengths[bridgeLength] = size
		}
	}

	fmt.Println("Sizes:", bridgeStrengths)
}

