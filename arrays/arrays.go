package arrays

import "fmt"

func Arrays() {

	// Array is a fixed size instance and [5]int type does not equal [3]int, type is not just array of integers, type is 5 cells array of integers and it is different from 3 cells array of integers
	arr := [5]int{1, 2, 3, 4, 5}
	for i, el := range arr {
		fmt.Println("Current index is => ", i, ", ", "Current element is => ", el)
	}
	fmt.Println(arr)

	arr2 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Printf("%T => ", arr2)
	fmt.Println(arr2)
	changeArr2(arr2)
	fmt.Println(arr2) // Value is not mutated outside the func changeArr2

	arr3 := [...][3]int{{2, 4, 6}, {6, 8, 12}, {9, 7, 32}}
	fmt.Printf("%T => ", arr3)
	fmt.Println(arr3)

	var arr4 [5][2]int // cannot use [...] without assignment a value
	fmt.Printf("%T => ", arr4)
	fmt.Println(arr4)

	// actions with arrays

	arr3[2] = [3]int{5, 9, 12}
	fmt.Printf("%T => ", arr3)
	fmt.Print(" length of array is ", len(arr3), " => ")
	fmt.Println(arr3)

}

func changeArr2(arr [2][2]int) {
	arr[0] = [2]int{100, 100}
}
