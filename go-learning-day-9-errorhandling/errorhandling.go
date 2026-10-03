package main

import (
	"errors"
	"fmt"
)

//we have defined Error of our own choice such that if we need to call that error or simply need
//to through that error we can simply do that

var errorAgeValidator = errors.New("Age is less than 18")

func main() {
	fmt.Println("Hello World")

	result, err := divide(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Result:", result)

	Eligibale, error := ageValidator(9)
	if error != nil {
		fmt.Println("Error: ", error)
	}
	fmt.Println("is Eligible: ", Eligibale)
	age, ageError := ageValidatorNew(15)

	if ageError != nil {
		fmt.Println("Error: ", ageError)
	}
	fmt.Println("age", age)

	//u can also check the error by using errors.Is() method like if the error is equal to the errorNotEligible then we can simply print the message

	fmt.Println("This is the focus point of this program ")
	isErr := isError(15)
	if errors.Is(isErr, errorNotEligible) {
		fmt.Println("Error: ", isErr)
	}

	//12 
}

// we simply create a method to divide two numbers like if one of the number is or simply denominator is
// 0 there is an error so we made our method to return two things one is integer result and other is
// error like if the denominator is 0 then there is an error ok
func divide(a int, b int) (int, error) {

	if b == 0 {
		//It is only useful for static messages
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil

}

func ageValidator(age int) (string, error) {
	if age < 18 {
		//It is only useful for for dynamic variables if u want messages to be static
		return "", fmt.Errorf("Age %d is smaller then 18", age)
	}
	return "U can Cast Vote", nil
}

func ageValidatorNew(age int) (string, error) {
	if age < 18 {
		return "", fmt.Errorf("The Program Breaks due to the erro : %w", errorNotEligible)
	}
	return "Age is Validated", nil
}

//11

func isError(age int) error {
	if age < 18 {
		return errorNotEligible
	}
	return nil
}
