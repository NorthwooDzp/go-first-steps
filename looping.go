package main

import "fmt"

func looping() {
	for i := 5; i < 10; i += 2 {
		fmt.Println("Current index is", i)
	}

	fmt.Println()

	length := 16
	for i := range length {
		fmt.Println("Current index is", i)
	}

	fmt.Println()
	// While loop in go using for
	i := 1
	for i < 16 {
		i++
		fmt.Println("Index is", i)
	}

	// Looping in the stirngs
	str := "Hello world"
	fmt.Println(str[0])
	fmt.Printf("%T", str[0]) // if string will contain emodzi or any special character this notation will return only partial result, first 8 bits of symbol that might have 3-4 bits (UTF-8)
	fmt.Println()

	for i := range len(str) {
		fmt.Printf("%c", str[i]) // works fine with ASCII symbols
	}
	fmt.Println()

	for _, char := range str {
		fmt.Printf("%c", char)
	}
}
