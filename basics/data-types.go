package basics

import (
	"fmt"
	"math"
)

const SimplePi float32 = 3.14 // defininf const

func DataTypes() {
	const MathPi float32 = math.Pi // local const in fumction
	// uint - unsigned integer, uint32 or uint64 based on value (uint8 - unit 64)
	// int - signed integer, same as uint
	// float32 and float64 - value with floating point
	// byte = int8, can be used for store individual character e.g. 'c'
	// rune = int32
	// bool = boolean value true/false
	// string - text enclosed in double quotes ""
	// nil = zero value for reference-like types such as pointers, slices, maps, and interfaces

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
