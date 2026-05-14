package main

import "fmt"

func consoleOutput() {
	fmt.Print("anything", "something else")                   // Print something without formatting (plain output)
	fmt.Println("printing new line", "new value after space") // Print formatted (space between arguments) with line ending

	// Formatting to specified format
	fmt.Printf("%T", false) // Print member type
}
