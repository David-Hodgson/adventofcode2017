package adventofcode2017

import (
	"fmt"
	"strings"
	"strconv"
)

type vector struct {
	x,y,z int
}

type particle struct {
	position, velocity,acceleration vector
	destroyed bool
}

func (p *particle) update() {

	//Increase the X velocity by the X acceleration.
	p.velocity.x += p.acceleration.x
	//Increase the Y velocity by the Y acceleration.
	p.velocity.y += p.acceleration.y
	//Increase the Z velocity by the Z acceleration.
	p.velocity.z += p.acceleration.z
	//Increase the X position by the X velocity.
	p.position.x += p.velocity.x
	//Increase the Y position by the Y velocity.
	p.position.y += p.velocity.y
	//Increase the Z position by the Z velocity.
	p.position.z += p.velocity.z
}

func (p *particle) getDistanceFromOrigin() int{

	return abs(p.position.x) + abs(p.position.y) + abs(p.position.z)
}

func parseCoordStringToVector(coords string) vector {

	parts := strings.Split(coords, ",")

	x,_ := strconv.Atoi(parts[0])
	y,_ := strconv.Atoi(parts[1])
	z,_ := strconv.Atoi(parts[2])

	return vector{x,y,z}
}

func parseInputLine(vectorInput string) particle{

	posStart := strings.Index(vectorInput, "p")
	posString := vectorInput[posStart+3:]
	posString = posString[0:strings.Index(posString, ">")]

	velStart := strings.Index(vectorInput, "v")
	velString := vectorInput[velStart+3:]
	velString = velString[0:strings.Index(velString, ">")]

	accStart := strings.Index(vectorInput, "a")
	accString := vectorInput[accStart+3:]
	accString = accString[0:strings.Index(accString, ">")]

	return particle{parseCoordStringToVector(posString), parseCoordStringToVector(velString), parseCoordStringToVector(accString),false}
}

func DayTwentyExample() {

	fmt.Println("Day 20 - Example")

	vectorInput := "p=<-1027,-979,-188>, v=<7,60,66>, a=<9,1,-7>"

	newParticle := parseInputLine(vectorInput)

	fmt.Println(newParticle)
	fmt.Println("Distance from origin:", newParticle.getDistanceFromOrigin())
	newParticle.update()
	fmt.Println(newParticle)
	fmt.Println("Distance from origin:", newParticle.getDistanceFromOrigin())
	newParticle.update()
	fmt.Println("Distance from origin:", newParticle.getDistanceFromOrigin())
	newParticle.update()
	fmt.Println("Distance from origin:", newParticle.getDistanceFromOrigin())
	newParticle.update()
	fmt.Println("Distance from origin:", newParticle.getDistanceFromOrigin())
	newParticle.update()
	fmt.Println("Distance from origin:", newParticle.getDistanceFromOrigin())

}

func DayTwentyPartOne() {

	fmt.Println("Day 20 - Part One")

	input := ReadFile("day20-input.txt")

	inputList := strings.Split(input, "\n")

	particleList := make([]particle, len(inputList))

	for i := 0; i<len(inputList); i++ {

		newParticle := parseInputLine(inputList[i])
		particleList[i] = newParticle

		if newParticle.acceleration.x == 0 && newParticle.acceleration.y == 0 && newParticle.acceleration.z == 0 {
			fmt.Println("Particle",i, "is not moving")
			fmt.Println("Distance:",newParticle.getDistanceFromOrigin())
		}
	}


	for j := 0; j < 1000 ;j++ {

		minDistance := -1
		minPos := -1

		for k := 0; k < len(particleList); k++ {

			currentParticle := &particleList[k]
			currentParticle.update()
			currentParticleDistance := currentParticle.getDistanceFromOrigin()

			if minDistance == -1 {
				minDistance = currentParticleDistance
				minPos = k
			} else {
				if currentParticleDistance < minDistance {
					minDistance = currentParticleDistance
					minPos = k
				}
			}

		}

		fmt.Println("At tick", j, "particle", minPos, " is closet at", minDistance)
	}

}

func DayTwentyPartTwo() {

	fmt.Println("Day 20 - Part Two")

	input := ReadFile("day20-input.txt")

	inputList := strings.Split(input, "\n")

	particleList := make([]particle, len(inputList))

	for i := 0; i<len(inputList); i++ {

		newParticle := parseInputLine(inputList[i])
		particleList[i] = newParticle

	}


	for j := 0; j < 10000 ;j++ {

		collisonMap := make(map[vector]*particle)

		for k := 0; k < len(particleList); k++ {

			currentParticle := &particleList[k]

			if currentParticle.destroyed {
				continue
			}

			currentParticle.update()

			if collidedParticle, occupied := collisonMap[currentParticle.position]; occupied {
				collidedParticle.destroyed = true
				currentParticle.destroyed = true
			}

			collisonMap[currentParticle.position] = currentParticle

		}

		fmt.Println("At tick", j," there were",len(collisonMap), "enteries in the collison map")
	}

	particleCount := 0

	for x := 0 ; x < len(particleList) ; x++ {
		if !particleList[x].destroyed {
			particleCount++
		}
	}

	fmt.Println("Particles left", particleCount)
}



