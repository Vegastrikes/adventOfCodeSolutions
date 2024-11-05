package _2015

import (
	"adventOfCodeSolutions/utility"
	"fmt"
)

func D1_P1() {
	input, err := utility.GetInput("./2015/inputs/d1_p1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var floor = 0

	for _, c := range input {
		switch c {
		case 40:
			floor++
		case 41:
			floor--
		default:
			fmt.Println("Wrong Input")
			return
		}
	}

	fmt.Printf("D1_P1: %d\n", floor)
}

func D1_P2() {
	input, err := utility.GetInput("./2015/inputs/d1_p1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var floor = 0
	var position = 0

	for i, c := range input {
		switch c {
		case 40:
			floor++
		case 41:
			floor--
		}

		if floor == -1 {
			position = i + 1
			break
		}
	}

	if position == 0 {
		fmt.Println("NO RESULT: Never enter the basement")
		return
	}

	fmt.Printf("D1_P2: %d\n", position)
}
