package arrays

import (
	"fmt"
)

func Slices() {
	slice := []int{}
	slice = append(slice, 5, 7, 12) // first argument can be slice type only

	fmt.Printf("%T => ", slice)
	// slice[1] = 10 => won't work because index is over the slice capacity
	fmt.Println(slice)
	fmt.Println(cap(slice))

	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[:] // Converting array to slice
	fmt.Printf("%T => ", sl)
	fmt.Println(sl)
	sl2 := arr[1:3] // Converting array to slice with specify indexes. element with starting index will be included to the slice, ending index - no
	fmt.Print("sl2 => ")
	fmt.Printf("%T => ", sl2)
	fmt.Println(sl2)
	sl2[1] = 4 // mutate value of the slice and array
	fmt.Println("sl2 => ", sl2, "arr => ", arr)

	/*
			When creating slice we have 3 main things:
			pointer => arr[1]
			length => 2
		  capacity => number of the indexes that exists from the pointer to the end of array, 4
	*/
	fmt.Println(sl2, len(sl2), cap(sl2))

	sl2 = sl2[:4] // reclicing, length increased, capacity is still based by underlying array
	fmt.Println(sl2, len(sl2), cap(sl2))
	sl2 = append(sl2, 654) // creating completely new underlying array
	fmt.Println(sl2, len(sl2), cap(sl2))
	sl2 = sl2[:8]
	fmt.Println(sl2, len(sl2), cap(sl2))

	strSl := []string{"hello", "world"}
	fmt.Println(strSl)

	mutateSlice(strSl) // function mutates original slice, unlike with arrays
	fmt.Println("Mutated slice => ", strSl)

	for i := range 10 {
		strSl = append(strSl, fmt.Sprintf("hello %d", i+1))
		fmt.Print(cap(strSl), "|")
	}
	fmt.Println(strSl, len(strSl), cap(strSl))

	mkSl := make([]int, 12) // make function can create a slice, map or channel
	fmt.Println(mkSl, len(mkSl), cap(mkSl))

	mkSl2 := make([]string, 12, 24)
	fmt.Println(mkSl2, len(mkSl2), cap(mkSl2))

}

func mutateSlice(sl []string) {
	sl[1] = "changed_value"
}
