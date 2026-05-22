package basics

import "fmt"

func ConditionsAndConditionals() {
	x := 64

	if x < 98 {
		fmt.Println("X is less than 98")
	} else if x > 498 {
		fmt.Println("X is more than 498")
	} else {
		fmt.Println("X is in the range between 98 and 498")
	}

	a := 3

	switch a {
	case 1:
		fmt.Println("A is 1")
	case 2:
		fmt.Println("A is 2")
	case 3:
		fmt.Println("A is 3")
	default:
		fmt.Println("A is out of range")
	}

	b := 8
	switch {
	case b < 32:
		fmt.Println("B is less than 32")
		fallthrough
	case b < 16:
		fmt.Println("B is less than 16")
	case b > 64:
		fmt.Println("B is more than 64")
	default:
		fmt.Println("B is in the range 32 - 64")
	}

	char := 'a'

	switch char {
	case 'a', 'b', 'c':
		fmt.Println("Char has proper value")
	default:
		fmt.Println("char is far from proper values")
	}
}
