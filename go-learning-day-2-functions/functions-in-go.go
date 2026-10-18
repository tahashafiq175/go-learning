
package main

import "fmt"

// CONCEPT 1: FUNCTION WITH PARAMETERS
//
// A function is a reusable block of code.
// Parameters are inputs that a function receives.
//
// Syntax:
// func functionName(parameterName datatype) {
//     function body
// }
//
// Here, a and b are parameters.
// Both parameters have the datatype int.

func calculation(a int, b int) {
	fmt.Println(a + b)
}

// CONCEPT 2: SHORTHAND PARAMETER DECLARATION
//
// When multiple consecutive parameters have the same datatype,
// we can write the datatype only once.
//
// These two declarations mean the same thing:
//
// func calculation(a int, b int)
// func operation(a, b int)
//
// Go understands that both a and b are integers.

func operation(a, b int) {
	fmt.Println(a + b)
}

// CONCEPT 3: FUNCTION WITH A RETURN VALUE
//
// Sometimes, a function should calculate something and give
// the result back to the code that called it.
//
// The int after the parameter list is the return type.
//
// Syntax:
// func functionName(parameters) returnType {
//     return value
// }
//
// Here, the function returns an integer.
// The return keyword sends the calculated result back.

func returnFunc(a, b int) int {
	return a + b
}

// CONCEPT 4: MULTIPLE RETURN VALUES
//
// A Go function can return more than one value.
//
// Syntax:
// func functionName(parameters) (type1, type2) {
//     return value1, value2
// }
//
// This function returns:
// 1. The sum of a and b.
// 2. The product of a and b.
//
// IMPORTANT:
// difference is a misleading variable name here because
// a * b calculates multiplication, not subtraction.

func returnFuncs(a, b int) (int, int) {
	sum := a + b
	product := a * b

	return sum, product
}

// CONCEPT 5: RETURNING MULTIPLE VALUES DIRECTLY
//
// We do not always need temporary variables.
// We can calculate and return both results directly.
//
// The first returned value is the sum.
// The second returned value is the product.
//
// This produces the same kinds of results as returnFuncs.

func manyReturnFuncs(a, b int) (int, int) {
	return a + b, a * b
}

// CONCEPT 6: NAMED RETURN VALUE
//
// We can give a name to a return value in the function signature.
//
// Here, difference is the name of the returned string.
// The name is confusing because the function creates a message,
// so message would be a better name.
//
// A named return value behaves like a local variable.
// It starts with its datatype's zero value.
// We can assign a result to it and return it.
//
// A bare return returns the current value of the named result.

func modReturnFuncs(age int, name string) (message string) {
	message = fmt.Sprintf(
		"My name is %s with age of %d",
		name,
		age,
	)

	return
}

// CONCEPT 7: THE main FUNCTION
//
// In an executable Go program, main() in package main
// is where execution begins.
//
// We call greet() to start our example.

func main() {
	fmt.Println("hello world")
	greet()
}

// CONCEPT 8: CALLING OTHER FUNCTIONS
//
// greet() is a function without parameters or return values.
// It calls other functions and prints their results.
//
// When main() calls greet(), the statements inside greet()
// execute from top to bottom.

func greet() {
	fmt.Println("greeting from taha")

	// Calling a function that prints the sum itself.
	calculation(10, 20)

	// Another function that prints the sum itself.
	operation(30, 10)

	// Store the returned sum inside a variable.
	sum := returnFunc(100, 30)

	fmt.Println("Calling the sum function")
	fmt.Println(sum)

	// Receive two returned values into two variables.
	add, multiply := returnFuncs(20, 50)

	fmt.Println("Calling the function with two return values")
	fmt.Println(add)
	fmt.Println(multiply)

	// Call the function that returns a named string value.
	details := modReturnFuncs(10, "Taha")
	fmt.Println(details)
}
