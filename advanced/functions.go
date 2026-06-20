package advanced

import (
	"fmt"
)

// Basic function declaration example
func add(a, b int) int { // Private function, lowercase naming
	return a + b
}

// Multiple return

func addMult(a, b int, text string) (int, string) {
	sum := a + b
	newText := text + " changed"
	return sum, newText
}

// Accepting function as a param
func callFunc(callback func(a, b int) int, x, y int) int {
	return callback(x, y)
}

// Returning function
func getFunc(text string) func(string) string {
	return func(text2 string) string {
		return text + " <=> " + text2
	}
}

// Variadic function with named return
func sum(nums ...int) (res int) {
	for i := range len(nums) {
		res += nums[i]
	}
	return
}

func Functions() {
	fmt.Println("Basic function declaration result =>", add(4, 6))

	// Basic function expression example
	outVar := "Outer scope variable"
	testfunc := func(arg string) string {

		fmt.Println("Function argument =>", arg, "|| outer variable =>", outVar) // outer variables are available inside the function
		return "test2"                                                           // Hardcoded return value
	}

	fmt.Println("Basic function expression result =>", testfunc("test"))

	multSum, multText := addMult(5, 7, "test text string") // should use variable for each return value, if not needed use _ as a var name
	fmt.Println("Function with multiple return result =>", multSum, "|", multText)

	fmt.Println("Function with function as a param result =>", callFunc(add, 7, 12))
	fmt.Println("Anonymous function result =>", callFunc(func(a, b int) int { // Anonymous function as a function param
		return a*10 + b*12
	}, 15, 25))

	fmt.Println("Returned function result =>", getFunc("First text part")("Second text part"))

	fmt.Println("Variadic function with default return result =>", sum(5, 12, 65, 89, 77, 14, 56, 54))
	fmt.Println("Variadic function with slice as a param result =>", sum([]int{54, 16, 36, 89, 17}...))
}
