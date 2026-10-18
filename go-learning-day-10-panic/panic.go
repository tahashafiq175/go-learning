package main

// package main tells Go that this is an executable program.
// Go starts executing this program from the main() function.

import (
	"errors"
	"fmt"
)

// errors is a standard Go package for creating and inspecting errors.
// fmt is used to print messages to the terminal.

// This creates a reusable error with a fixed message.
// In the current program, this variable is not used.
// errors.New() creates an error value, but it does not cause a panic.
var errorAgeValidator = errors.New("Age is less than 18")

func main() {

	fmt.Println("Hello World")
	// Prints "Hello World" to the terminal.
	// Execution then continues to the next statement.

	// panic("This is panic message")
	// If uncommented, this would trigger a panic.
	// A panic interrupts normal execution and starts unwinding
	// the current goroutine's function calls.
	// A deferred function may recover from the panic.
	// If nobody recovers it, the program normally terminates.

	fmt.Println("This is after panic message")
	// Prints this message because the panic statement above
	// is commented out and therefore does not execute.

	answer := divide(10, 0)
	// Calls divide() with a = 10 and b = 0.
	// Because b is zero, divide() triggers a panic.
	// Its deferred function recovers from that panic.
	// The function then returns the zero value of int, which is 0.
	// That returned value is stored in answer.

	if answer == 0 {
		// Checks whether answer equals zero.
		// In this example, answer is zero after the recovered panic.
		fmt.Println("Divide by zero error")
		// Prints the message inside the if block.
	}

	fmt.Println("Answer: ", answer)
	// Prints the value stored in answer.
	// Because divide() recovered from the panic and returned,
	// execution continues here.
}

func divide(a, b int) int {
	// This function accepts two integers and returns one integer.
	// Its return type is int, so it does not return an error value.

	defer func() {
		// defer schedules this anonymous function to execute
		// when divide() is about to return or when a panic
		// causes the function to unwind.

		r := recover()
		// recover() checks whether this deferred function is
		// handling an active panic.
		// If it catches a panic, r contains the panic value.
		// If there is no active panic to recover, r is nil.

		if r != nil {
			// A non-nil r means a panic was recovered.
			fmt.Println("Recovered from panic: ", r)
			// Prints the value that caused the panic.
		}
		// After this deferred function finishes, divide() can
		// return to its caller because the panic was recovered.
	}()

	if b == 0 {
		// Checks whether the denominator is zero.
		// Dividing by zero is invalid, so the function deliberately
		// triggers a panic instead of performing the division.
		panic("Divide by zero error")
		// panic() starts panic handling.
		// The deferred function above runs during unwinding.
		// recover() catches the panic, so it does not escape
		// from divide() into main().
	}

	return a / b
	// This line executes only if b is not zero.
	// For example, divide(10, 2) returns 5.
}