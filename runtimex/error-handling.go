package runtimex

import (
	"errors"
	"fmt"
)

func divideWithPanic(a, b int) float64 {
	if b == 0 {
		panic("Division by zero not allowed!") // Manually triggered panic
	}
	return float64(a) / float64(b)
}

func divideWithError(a, b int) (float64, error) { // commonly used error handling practice; avoid panic
	if b == 0 {
		return 0, errors.New("Divide by 0 not allowed!")
	}

	return float64(a) / float64(b), nil
}

func logDefer() {
	fmt.Println("Deferred function log")
	r := recover() // Catch the panic and continue program execution

	if r == nil {
		fmt.Println("No panic, normal execution")
	} else {
		fmt.Println("Panic! Something went wrong!")
	}
	fmt.Println("Recover =>", r)
}

func ErrorHandling() {
	divideByZero := func() {
		defer logDefer() // deferred function will be executed after all other lines have been executed or a panic occurs
		/* A deferred function can be useful when a panic may occur but teardown logic should still run.
		For example, opening a file can trigger a panic if it is already opened by another program, yet you need to close the file before stopping execution. */
		divideWithPanic(55, 0)
	}
	divideByZero() // function triggers a panic; code below won't execute without recover

	fmt.Println(divideWithPanic(186, 11))

	result, err := divideWithError(54, 0)
	if err != nil {
		fmt.Println("Error occurred =>", err)
	} else {
		fmt.Println("Result =>", result)
	}

	result, err = divideWithError(54, 7)
	if err != nil {
		fmt.Println("Error occurred =>", err, "|", "err.Error() result =>", err.Error())
	} else {
		fmt.Println("Result =>", result)
	}
}
