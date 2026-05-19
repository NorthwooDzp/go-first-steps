package arrays

import "fmt"

func Arrays() {
	arr := [5]int{1, 2, 3, 4, 5}
	for _, el := range arr {
		fmt.Println("Current element value is => ", el)
	}
	fmt.Println(arr)
}
