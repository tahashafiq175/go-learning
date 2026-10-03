package main

import ("fmt"
"go-learning-day-11-FileHandling/calculator"
)

func main() {
	fmt.Println("Hello World")
	operation, err := calculator.Operation(10, 10, "*")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", operation)
	}
}
