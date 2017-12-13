package adventofcode2017

import (
	"fmt"
	"strings"
	"strconv"
)

const SCANNING_UP = -1
const SCANNING_DOWN = 1

type fireWallLayer struct {
	depth int
	scanningRange int
	pos int
	scanningDirection int
}

func atoi(input string) int {
	value,_ := strconv.Atoi(input)
	return value
}

func updateFirewallPositions(firewall map[int]fireWallLayer)  {
	for _,fwl :=range  firewall{

		if (fwl.scanningDirection == SCANNING_UP) {

			fwl.pos++

			if fwl.pos == fwl.scanningRange {
				fwl.scanningDirection = SCANNING_DOWN
			}

		} else {
			fwl.pos--

			if fwl.pos == 1 {
				fwl.scanningDirection = SCANNING_UP
			}
		}

		firewall[fwl.depth] =fwl
	}
}

func resetFirewall(firewall map[int]fireWallLayer)  {
	for _,fwl :=range  firewall{
		fwl.pos = 1
		fwl.scanningDirection = SCANNING_UP
		firewall[fwl.depth] =fwl
	}
}

func calculateSeverity(firewall map[int]fireWallLayer, secondsToRun int, startDelay int) (severity int, caught bool) {
	currentPosition := 0

	resetFirewall(firewall)
	started := false
	//severity := 0
	//caught := false
	for i := 0 ; i <= secondsToRun ; i++ {

		if i == startDelay {
			started = true
		}
		//fmt.Println("Pico second", i)
		//fmt.Println(firewall[6])

		if fwl, exists := firewall[currentPosition]; exists && started {
			if fwl.pos == 1 {
				caught = true
				//fmt.Println("Caught at pos", currentPosition)
				//fmt.Println("Current Severity:", severity)
				//fmt.Println("Additional Severity: ", fwl.depth, " * ", fwl.scanningRange)
				//fmt.Println("Additional Severity:", (fwl.depth * fwl.scanningRange))
				severity += (fwl.depth * fwl.scanningRange)
				//fmt.Println("New Severity:", severity)

			}
		}

		updateFirewallPositions(firewall)
		//fmt.Println("Firewall updated")
		//fmt.Println(firewall[6])
		if started {
			currentPosition++
		}
	}

	return severity, caught
}

func DayThirteenExample() {

	fmt.Println("Day 13 - Example")

	input := "0: 3\n1: 2\n4: 4\n6: 4"

	fireWallDefinitions := strings.Split(input, "\n")

	firewall := make(map[int]fireWallLayer)

	maxDepth := 0

	for i :=0 ; i <len(fireWallDefinitions) ; i++ {

		parts := strings.Split(fireWallDefinitions[i], ":")

		fwl := fireWallLayer{atoi(parts[0]), atoi(strings.Trim(parts[1], " ")), 1,SCANNING_UP}

		if fwl.depth > maxDepth {
			maxDepth = fwl.depth
		}

		firewall[fwl.depth] = fwl
	}

	//fmt.Println("Firewall Depth:", maxDepth)
	//
	severity,caught := calculateSeverity(firewall, maxDepth, 0)
	fmt.Println("Severity:", severity)
	fmt.Println("Caught:", caught)

	//fmt.Println("Firewall Depth:", maxDepth)

	//fmt.Println(firewall)
	for i :=0 ; i < 20 ; i++ {
		delay := i
		_, caught := calculateSeverity(firewall, maxDepth + delay, delay)
		//fmt.Println("Severity:", severity)
		//fmt.Println(caught)

		if !caught {
			fmt.Println("Not caught at delay", delay)
			break
		}
	}
}

func DayThirteenPartOne() {

	fmt.Println("Day 13 - Part One")

	input := ReadFile("day13-input.txt")

	fireWallDefinitions := strings.Split(input, "\n")

	firewall := make(map[int]fireWallLayer)

	maxDepth := 0

	for i :=0 ; i <len(fireWallDefinitions) ; i++ {

		parts := strings.Split(fireWallDefinitions[i], ":")

		fwl := fireWallLayer{atoi(parts[0]), atoi(strings.Trim(parts[1], " ")), 1,SCANNING_UP}

		if fwl.depth > maxDepth {
			maxDepth = fwl.depth
		}

		firewall[fwl.depth] = fwl
	}

	fmt.Println("Firewall Depth:", maxDepth)
	severity,_ := calculateSeverity(firewall, maxDepth,0)
	fmt.Println("Severity:", severity)
}

func DayThirteenPartTwo() {

	fmt.Println("Day 13 - Part Two")

	input := ReadFile("day13-input.txt")

	fireWallDefinitions := strings.Split(input, "\n")

	firewall := make(map[int]fireWallLayer)

	maxDepth := 0

	for i :=0 ; i <len(fireWallDefinitions) ; i++ {

		parts := strings.Split(fireWallDefinitions[i], ":")

		fwl := fireWallLayer{atoi(parts[0]), atoi(strings.Trim(parts[1], " ")), 1,SCANNING_UP}

		if fwl.depth > maxDepth {
			maxDepth = fwl.depth
		}

		firewall[fwl.depth] = fwl
	}

	fmt.Println("Firewall Depth:", maxDepth)

	//fmt.Println(firewall)
	for i :=50000 ; i < 100000 ; i++ {
		delay := i
		//fmt.Println("Trying delay", delay)
		_, caught := calculateSeverity(firewall, maxDepth + delay, delay)
		//fmt.Println("Severity:", severity)
		//fmt.Println(caught)

		if !caught {
			fmt.Println("Not caught at delay", delay)
			break
		}
	}
}


