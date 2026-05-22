package basics

import "fmt"

func DataTypes() {
	// uint - unsigned integer, uint32 or uint64 based on value
	// uint8 - unit 64
	// int - signed integer, same as uint
	// float32 and float64 - value with floating point
	// byte = int8, can be used for store individual character e.g. 'c'
	// rune = int32
	// bool = boolean value true/false
	// string - only in ""
	// nil = undefined or null

	// Declaring variables:

	var x string = "hello world"

	fmt.Println(x)

	const y int32 = 2554557

	var (
		a int16
		b int16
		c bool
	)

	c = false
	a = 5

	b = 'c'

	// implicit assignment
	d := "something"
	d = "something else"

	e := uint(64)

	fmt.Println(a, b, c, d, e)

	fmt.Printf("%T", d)
}
