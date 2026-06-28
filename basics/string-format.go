package basics

import "fmt"

func StringFormat() {

	// Print without formatting (plain output)
	fmt.Print("anything", "something else", "\n")

	// Print formatted (space between arguments) with line ending
	fmt.Println("printing new line", "new value after space")

	// Formatting with fmt.Printf
	fmt.Printf("%T %T", false, 654) // Print member type
	fmt.Println()

	name := "John"
	age := 38

	fmt.Printf("Person %s is %d years old", name, age) // Formatting where %s means string and %d - number
	fmt.Println()
	fmt.Printf("%v", true) // %v - print the value
	fmt.Println()
	fmt.Printf("%b", 255) // %b - print the value in binary representation
	fmt.Println()
	fmt.Printf("%e", 25565585.22544) // %e - print the value in scientific notation
	fmt.Println()
	fmt.Printf("%f", 85.22) // %f - print the floating-point value with 6 decimal places
	fmt.Println()
	fmt.Printf("%.2f", 85.22544) // represent only 2 decimals with math rounding
	fmt.Println()
	fmt.Printf("%20.2f", 85.22544) // represent only 2 decimals with math rounding and adds n spaces before
	fmt.Println()
	fmt.Printf("%.2f%%", 78.22544) // 2 decimals precision with % sign after (without convert to percentage)
	fmt.Println()
	fmt.Printf("\"The %s is the capital of Great Britain\"", "London") // \ - escape character
	str := fmt.Sprintf("%20.2f", 64.22544)                             // format string and assign to variable
	fmt.Println()
	fmt.Println(str)
}
