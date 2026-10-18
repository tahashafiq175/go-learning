package main

import (
	"errors"
	"fmt"
)

// ============================================================
// CONCEPT 1: CREATING OUR OWN REUSABLE ERROR
// ============================================================

// errors.New() creates an error containing a fixed message.
//
// Think of this as creating a reusable error label.
// Whenever our program needs to report that a person is too
// young, we can return this same error.
//
// We declare it outside main() so that all functions in this
// file can access it.
//
// The variable has the type error, which is Go's built-in
// interface used to represent errors.
var errorNotEligible = errors.New("age is less than 18")

func main() {

	fmt.Println("========== GO ERROR HANDLING ==========")

	// ========================================================
	// CONCEPT 2: CALLING A FUNCTION THAT RETURNS AN ERROR
	// ========================================================

	// divide() returns two values:
	// 1. The calculated result.
	// 2. An error describing a failure, or nil if successful.
	//
	// We store those returned values in result and err.
	result, err := divide(10, 0)

	// nil means that no error is present.
	// Therefore, err != nil means that an error occurred.
	//
	// Go does not automatically stop the program when a
	// function returns an error. We must decide what to do.
	if err != nil {
		fmt.Println("Division Error:", err)
	} else {
		// We only use the result as a successful answer
		// when the function reports no error.
		fmt.Println("Division Result:", result)
	}

	// Try the same function with valid inputs.
	// Since 20 divided by 4 is valid, the error will be nil.
	result, err = divide(20, 4)

	if err != nil {
		fmt.Println("Division Error:", err)
	} else {
		fmt.Println("Successful Division Result:", result)
	}

	// ========================================================
	// CONCEPT 3: CREATING DYNAMIC ERRORS WITH fmt.Errorf()
	// ========================================================

	// ageValidator() returns:
	// 1. A string containing an eligibility message.
	// 2. An error if the age is below 18.
	//
	// Here, age is 15, so the function returns an error.
	eligibilityMessage, err := ageValidator(15)

	if err != nil {
		fmt.Println("Age Validation Error:", err)
	} else {
		// This branch runs only when validation succeeds.
		fmt.Println("Eligibility:", eligibilityMessage)
	}

	// Let's check an eligible age.
	eligibilityMessage, err = ageValidator(21)

	if err != nil {
		fmt.Println("Age Validation Error:", err)
	} else {
		fmt.Println("Eligibility:", eligibilityMessage)
	}

	// ========================================================
	// CONCEPT 4: WRAPPING AN EXISTING ERROR WITH %w
	// ========================================================

	// ageValidatorNew() wraps errorNotEligible with additional
	// information about where the failure happened.
	//
	// The returned error contains both:
	// 1. The additional context.
	// 2. The original error that was wrapped.
	status, err := ageValidatorNew(16)

	if err != nil {
		fmt.Println("Detailed Validation Error:", err)

		// errors.Is() checks whether err is the specified
		// error or wraps an error matching that value.
		//
		// Because ageValidatorNew() uses %w to wrap
		// errorNotEligible, this condition evaluates to true.
		if errors.Is(err, errorNotEligible) {
			fmt.Println("Reason: The person is not eligible.")
		}
	} else {
		fmt.Println("Status:", status)
	}

	// ========================================================
	// CONCEPT 5: RETURNING A REUSABLE ERROR DIRECTLY
	// ========================================================

	// isError() returns errorNotEligible when age < 18.
	// Otherwise, it returns nil.
	err = isError(14)

	// Here we check whether the returned error matches our
	// reusable error value.
	if errors.Is(err, errorNotEligible) {
		fmt.Println("Reusable Error Detected:", err)
	} else if err != nil {
		// Another error may have occurred, so we handle
		// any non-nil error that wasn't the expected one.
		fmt.Println("A Different Error Occurred:", err)
	} else {
		fmt.Println("Age Check Successful.")
	}

	// Check a valid age.
	err = isError(25)

	if err != nil {
		fmt.Println("Age Check Error:", err)
	} else {
		fmt.Println("Age Check Successful: Age is valid.")
	}

	fmt.Println("========== PROGRAM FINISHED ==========")
}

// ============================================================
// CONCEPT 6: RETURNING A RESULT AND AN ERROR
// ============================================================

// The function accepts two integers and returns:
// 1. An integer result.
// 2. An error.
//
// The return type (int, error) means that both values must
// be supplied whenever this function returns.
func divide(a int, b int) (int, error) {

	// Before dividing, we check whether the denominator is zero.
	//
	// Division by zero is not a valid operation, so we return
	// a placeholder result (0) and an error.
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	// If b is not zero, division can proceed.
	//
	// We return the calculated result and nil.
	// nil indicates that no error occurred.
	return a / b, nil
}

// ============================================================
// CONCEPT 7: FORMATTED ERRORS WITH fmt.Errorf()
// ============================================================

// This function receives an age and returns:
// 1. An eligibility message as a string.
// 2. An error if the age is below 18.
func ageValidator(age int) (string, error) {

	// We check the supplied age.
	if age < 18 {

		// fmt.Errorf() creates a formatted error.
		//
		// %d is a formatting placeholder for an integer.
		// Go replaces %d with the value stored in age.
		//
		// For age = 15, the error message becomes:
		// "age 15 is below 18"
		//
		// The empty string is returned because validation
		// did not succeed.
		return "", fmt.Errorf("age %d is below 18", age)
	}

	// If age is 18 or above, validation succeeds.
	// We return the message and nil.
	return "You are eligible to vote.", nil
}

// ============================================================
// CONCEPT 8: WRAPPING ERRORS WITH %w
// ============================================================

// This function returns a status message or an error.
func ageValidatorNew(age int) (string, error) {

	// Check whether the age is below 18.
	if age < 18 {

		// %w wraps an existing error.
		//
		// errorNotEligible is the original error.
		// "Age validation failed" adds useful context.
		//
		// The returned error message will be:
		// "age validation failed: age is less than 18"
		//
		// Because we used %w, errors.Is() can still
		// recognize errorNotEligible inside this error.
		return "", fmt.Errorf(
			"age validation failed: %w",
			errorNotEligible,
		)
	}

	// If the age is valid, return a success message
	// and nil to indicate that no error occurred.
	return "Age has been validated successfully.", nil
}

// ============================================================
// CONCEPT 9: RETURNING AN ERROR DIRECTLY
// ============================================================

// This function returns only an error.
//
// Notice that the return type is error, not (string, error).
// That is because the function only needs to report whether
// the age check succeeded or failed.
func isError(age int) error {

	// If the age is below 18, return our reusable error.
	if age < 18 {
		return errorNotEligible
	}

	// If the age is valid, return nil.
	return nil
}
