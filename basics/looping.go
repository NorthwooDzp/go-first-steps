package basics

import "fmt"

func Looping() {
	for i := 5; i < 10; i += 2 {
		fmt.Println("Current index is", i)
	}

	fmt.Println()

	length := 16
	for i := range length {
		fmt.Println("Current index is", i)
	}

	fmt.Println()
	// While-style loop in Go using for
	i := 1
	for i < 16 {
		i++
		fmt.Println("Index is", i)
	}

	// Looping over strings
	str := "Hello world"
	fmt.Println(str[0])
	fmt.Printf("%T", str[0]) // indexing a string returns a byte, so non-ASCII characters may appear truncated
	fmt.Println()

	for i := range len(str) {
		fmt.Printf("%c", str[i]) // works for ASCII characters
	}
	fmt.Println()

	for _, char := range str {
		fmt.Printf("%c", char)
	}
}
