package main

import (
	"fmt"
	"math"
	"strconv"
)

func arithmeticOperations() {
	/**
	Main operators:
	+
	-
	*
	/
	++
	--
	%
	*/

	x := uint8(65)
	y := 7
	res1 := int(x) + y                // proper type conversion, smaller type converted to the larger
	res2 := float64(x) / float64(y)   // need to convert to float, otherwise only whole part will be in the result
	res3 := "hi" + string(2)          // wrong formatting
	res4 := "hello" + fmt.Sprint(654) // proper format
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)
	/* The math package */
	fmt.Println(math.Pi)

	/*Converting string to numbers*/
	numstr := "6558542"
	res5, err := strconv.Atoi(numstr)
	fmt.Println("res =>", res5, "error =>", err) // works for convert to integer

	numstr2 := "5566.555654"
	res6, err := strconv.ParseFloat(numstr2, 32) // converting to float with bitsize, if bitsize = 32 => float32, any other => float64
	fmt.Println("result => ", res6, " | error => ", err)
	numstr3 := "54445"
	res7, err := strconv.ParseInt(numstr3, 10, 32)
	fmt.Println("result => ", res7, " | error => ", err)

}
