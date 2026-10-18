package main

import "fmt"

// CONCEPT 1: Functions
// A function is a block of code that performs a specific task.
// We can write code inside a function and execute it when needed.
// This function is named datatype.
//
// IMPORTANT:
// Go does not automatically execute every function.
// Because our function is not named main, its code will not run
// automatically when we execute this program.
// To run datatype(), we must call it from main().

func datatype() {

	// CONCEPT 2: Declaring variables with an explicit datatype

	var age int = 19
	var name string = "Taha"

	// Think of a variable as a box that stores information.
	//
	// age is a box that stores an integer (whole number).
	// int is its datatype.
	// 19 is the value stored inside the box.
	//
	// name is a box that stores text.
	// string is the datatype used for text.
	//
	// General syntax:
	// var variableName datatype = value

	fmt.Println(age)
	fmt.Println(name)

	// CONCEPT 3: Changing a variable's value

	age = 20

	fmt.Println(age)

	// We can change the value stored inside age because
	// age was declared using var.
	//
	// IMPORTANT:
	// We use = here because age already exists.
	// We are changing its value, not declaring a new variable.
	//
	// The datatype remains int.
	// We cannot assign a string like "Taha" to age.


	// CONCEPT 4: Type inference with var

	var university = "Islamia University of Bahawalpur"

	fmt.Println(university)

	// Here, we did not write the datatype explicitly.
	// Go looks at the initial value and determines the datatype.
	//
	// Because the value is text, Go infers the type as string.
	//
	// This feature is called TYPE INFERENCE.
	//
	// It is similar to writing:
	// var university string = "Islamia University of Bahawalpur"
	//
	// IMPORTANT:
	// Go determines the datatype at compile time.
	// The variable does not become dynamically typed.


	// CONCEPT 5: Short variable declaration using :=

	gpa := 3.9

	fmt.Println(gpa)

	// The := operator declares a new variable and initializes it.
	// Go infers the type from the value on the right.
	//
	// Because 3.9 is an untyped floating-point constant,
	// Go infers the default type float64 for gpa.
	//
	// This is equivalent to:
	// var gpa float64 = 3.9
	//
	// IMPORTANT:
	// We can use := inside a function.
	// We cannot use := to declare a package-level variable.


	// CONCEPT 6: Checking a variable's datatype with %T

	fmt.Printf("%T\n", gpa)

	// %T is a format verb used by fmt.Printf.
	// It prints the datatype of the supplied value.
	//
	// Output:
	// float64
	//
	// The \n moves the cursor to a new line.
	//
	// Remember:
	// %T prints the type.
	// %v prints the value.
	// %d is commonly used for decimal integers.
	// %f is commonly used for floating-point numbers.


	// CONCEPT 7: A variable cannot change its datatype

	gpa = 8.9

	fmt.Println(gpa)

	// gpa was inferred as float64 when it was declared.
	// We can change its value to another float64-compatible value.
	//
	// But this will cause a compile-time error:
	// gpa = "Taha"
	//
	// Why?
	// Because "Taha" is text, while gpa is a float64 variable.
	//
	// IMPORTANT:
	// Variables can change their values, but their declared
	// datatypes do not change.


	// CONCEPT 8: Declaring multiple variables using var

	var (
		firstName = "Taha"
		score     = 100
	)

	fmt.Println(firstName)
	fmt.Println(score)

	// We can declare several variables inside one var block.
	//
	// Go infers the datatype of each variable separately.
	//
	// firstName becomes string.
	// score becomes int.
	//
	// Both variables do not need to have the same datatype.
	// This makes related declarations easier to organize.


	// CONCEPT 9: Declaring multiple variables using :=

	x, y := 10, 20

	fmt.Println(x, y)

	// Go lets us declare multiple variables in one statement.
	//
	// x receives 10.
	// y receives 20.
	//
	// Both are inferred as int.
	//
	// This is called a MULTIPLE SHORT VARIABLE DECLARATION.
	//
	// IMPORTANT:
	// When using :=, at least one variable on the left side
	// must be new in the current scope.
	//
	// For example:
	// x, z := 10, 30
	//
	// Here, x already exists, but z is new, so this is valid.


	// CONCEPT 10: Swapping two variables

	y, x = x, y

	fmt.Println(x, y)

	// Before swapping:
	// x = 10
	// y = 20
	//
	// After swapping:
	// x = 20
	// y = 10
	//
	// Go evaluates the right-hand-side expressions and then
	// assigns the resulting values to the left-hand variables.
	//
	// We do not need a temporary variable.
	//
	// IMPORTANT:
	// We use = because x and y already exist.
	// Using := here would be incorrect because both variables
	// have already been declared in this scope.


	// CONCEPT 11: Checking the datatypes after swapping

	fmt.Printf("%T %T\n", x, y)

	// Both x and y were initialized using integer values.
	// Swapping their values does not change their datatypes.
	//
	// Output:
	// int int


	// CONCEPT 12: Type conversion from int to float64

	new_age := float64(age)

	fmt.Println(new_age)

	// age is an int variable containing 20.
	//
	// float64(age) converts the integer value into float64.
	//
	// Before conversion:
	// age = 20, type int
	//
	// After conversion:
	// new_age = 20, type float64
	//
	// IMPORTANT:
	// Go does not automatically convert between int and float64.
	// We must explicitly request the conversion.
	//
	// Also, this does not change age into a float64.
	// It creates a converted value that is stored in new_age.


	// CONCEPT 13: Assigning a decimal value after conversion

	new_age = 20.7

	fmt.Println(new_age)

	// new_age is already a float64 variable.
	// Therefore, it can store a decimal value such as 20.7.
	//
	// We use = because new_age already exists.
	// We do not need to declare it again.


	// CONCEPT 14: Type conversion from float64 to int

	age = int(new_age)

	fmt.Println(age)

	// new_age contains 20.7.
	//
	// int(new_age) converts the floating-point value to an integer.
	//
	// Result:
	// 20
	//
	// IMPORTANT:
	// Converting a float to an integer removes the fractional
	// part. It does not round to the nearest whole number.
	//
	// For example:
	// int(20.7)  becomes 20
	// int(20.9)  becomes 20
	// int(20.1)  becomes 20
	// int(-20.7) becomes -20
	//
	// The original new_age value remains 20.7.
	// Only the converted result is assigned to age.
}


// CONCEPT 15: The main function
// Go starts executing a normal executable program from main()
// in package main.
// Calling datatype() makes the code above execute.

