package calculator

import "fmt"

// FIX: The main problem is not the uppercase/lowercase names.
// The real issue is that helper functions must not be declared inside another function.
// In Go, functions are normally written at package level, outside other functions.
// If you define add(), subtract(), multiply(), divide() inside Operation(), it is not valid Go.
//
// Also: uppercase means exported/public, lowercase means unexported/private.
// Example: Operation() is exported, add() is private.
func Operation(a int, b int, operator string) (int, error) {

	switch operator {
	case "+":	
		return add(a,b), nil
	case "-":
		return subtract(a,b), nil
	case "*":
		return multiply(a,b), nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return divide(a,b), nil		
	default:
		return 0, fmt.Errorf("invalid operator")
	}
	// These helper functions are private because their names start with lowercase letters.
	// The same rule is true for variables, structs, and interfaces:
	// uppercase = exported/public, lowercase = unexported/private.

}

// FIX: These functions are placed outside Operation() so the package can use them correctly.
func add(a int, b int) int {
		return a + b
	}
	func subtract(a int, b int) int {
		return a - b
	}
	func multiply(a int, b int) int {
		return a * b
	}
	func divide(a int, b int) int {
		return a / b
	}