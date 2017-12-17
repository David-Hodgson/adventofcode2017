package adventofcode2017

import (
	"fmt"
)

type listItem struct {
	value int
	next *listItem
}

func (li *listItem) insertAfter(newValue int) {
	newlistItem := listItem{value:newValue}

	newlistItem.next = li.next
	li.next = &newlistItem
}

func DaySeventeenExample() {

	fmt.Println("Day 17 - Example")

	spinlock := 3

	list := listItem{value:0}
	list.next = &list

	currentPos := &list


	for i:= 1 ; i < 2018 ; i++ {
		for j :=0 ; j < spinlock ;j++ {
			currentPos = currentPos.next
		}

		currentPos.insertAfter(i)
		currentPos = currentPos.next
	}


	startPos := &list

	for i:=0;i < 2018;i++ {
		if startPos.value == 2017 {
			fmt.Println("Found value is:",startPos.next.value)
			break
		}
		startPos = startPos.next
	}
}

func DaySeventeenPartOne() {

	fmt.Println("Day 17 - Part One")

	spinlock := 363

	list := listItem{value:0}
	list.next = &list

	currentPos := &list


	for i:= 1 ; i < 2018 ; i++ {
		for j :=0 ; j < spinlock ;j++ {
			currentPos = currentPos.next
		}

		currentPos.insertAfter(i)
		currentPos = currentPos.next
	}


	startPos := &list

	for i:=0;i < 2018;i++ {
		if startPos.value == 2017 {
			fmt.Println("Found value is:",startPos.next.value)
			break
		}
		startPos = startPos.next
	}
}


func DaySeventeenPartTwo() {

	fmt.Println("Day 17 - Part Two")

	spinlock := 363

	list := listItem{value:0}
	list.next = &list

	currentPos := &list


	for i:= 1 ; i < 500 000 001 ; i++ {

		if i % 1000 == 0 {
			fmt.Println("inserting:", i)
		}
		for j :=0 ; j < spinlock ;j++ {
			currentPos = currentPos.next
		}

		currentPos.insertAfter(i)
		currentPos = currentPos.next
	}

	fmt.Println("Value is:", currentPos.value)


}








