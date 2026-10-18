# Go Error Handling — Complete Beginner's Guide

A practical guide to understanding how Go reports, returns, checks, wraps, and handles errors.

---

## Table of Contents

1. [Introduction](#1-introduction)
2. [Learning Objectives](#2-learning-objectives)
3. [Project Structure](#3-project-structure)
4. [What Is an Error?](#4-what-is-an-error)
5. [The `error` Type and `nil`](#5-the-error-type-and-nil)
6. [Creating Errors with `errors.New()`](#6-creating-errors-with-errorsnew)
7. [Returning Errors from Functions](#7-returning-errors-from-functions)
8. [Checking Errors with `if err != nil`](#8-checking-errors-with-if-err--nil)
9. [Dynamic Error Messages with `fmt.Errorf()`](#9-dynamic-error-messages-with-fmterrorf)
10. [Wrapping Errors with `%w`](#10-wrapping-errors-with-w)
11. [Checking Errors with `errors.Is()`](#11-checking-errors-with-errorsis)
12. [Checking Error Types with `errors.As()`](#12-checking-error-types-with-errorsas)
13. [The Early Return Pattern](#13-the-early-return-pattern)
14. [Ignoring Errors and Common Mistakes](#14-ignoring-errors-and-common-mistakes)
15. [Errors vs. Panics](#15-errors-vs-panics)
16. [Complete Error Handling Flow](#16-complete-error-handling-flow)
17. [Expected Program Output](#17-expected-program-output)
18. [Best Practices](#18-best-practices)
19. [Practice Exercises](#19-practice-exercises)
20. [Final Takeaways](#20-final-takeaways)

---

## 1. Introduction

Go uses a straightforward approach to error handling.

When a function cannot complete an operation, it can return an error to the code that called it. The caller checks that error and decides what should happen next.

For example, imagine that we want to divide two numbers:

```go
result, err := divide(10, 0)
```

The denominator is zero, so the division cannot succeed.

Instead of silently pretending that the operation worked, the function returns an error.

Go does not automatically stop the entire program when a function returns an error. The caller is responsible for checking and handling it.

### The basic idea

```text
        Call a function
               |
               v
      Function performs work
               |
               v
       Did an error occur?
          /          \
        Yes           No
         |             |
         v             v
   Return error     Return result
         |             |
         v             v
   Caller handles   Caller uses
     the error       the result
```

**Core principle:** Errors are values that your program can inspect and handle.

---

## 2. Learning Objectives

After completing this project, you should understand:

- What errors are and why programs need them.
- The built-in `error` interface and the meaning of `nil`.
- How to create errors with `errors.New()`.
- How functions return both results and errors.
- Why Go commonly uses `if err != nil`.
- How to create dynamic error messages using `fmt.Errorf()`.
- How `%w` wraps an existing error.
- How `errors.Is()` identifies a particular error condition.
- How `errors.As()` extracts a particular error type.
- Why early returns make error handling clearer.
- The difference between errors and panics.
- How to design more useful errors in real applications.

---

## 3. Project Structure

A simple project can use the following structure:

```text
go-error-handling/
│
├── go.mod
├── main.go
└── README.md
```

- `go.mod` defines the Go module and its module path.
- `main.go` contains the error-handling examples.
- `README.md` explains the concepts and demonstrates how to use them.

If you do not already have a Go module, open a terminal in your project directory and run:

```bash
go mod init go-error-handling
```

You can choose a different module path if your project already belongs to a repository. Do not run this command again in a directory that already has a `go.mod` file.

To execute the program:

```bash
go run main.go
```

To build an executable:

```bash
go build
```

---

## 4. What Is an Error?

An error is a value that communicates that an operation failed or could not be completed as expected.

For example:

```go
errors.New("cannot divide by zero")
```

This creates an error with the message `cannot divide by zero`.

Consider a function that divides two integers:

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }

    return a / b, nil
}
```

Let's understand the execution.

### Case A: The operation fails

```go
result, err := divide(10, 0)
```

The function returns:

```text
result = 0
err    = cannot divide by zero
```

The integer `0` is a placeholder. It does not mean the division succeeded with a result of zero.

### Case B: The operation succeeds

```go
result, err := divide(10, 2)
```

The function returns:

```text
result = 5
err    = nil
```

The caller can safely use the result after confirming that the error is `nil`.

---

## 5. The `error` Type and `nil`

Go has a built-in interface named `error`.

Conceptually, it is defined like this:

```go
type error interface {
    Error() string
}
```

Any type that has an `Error()` method returning a string can implement this interface.

For example:

```go
err := errors.New("something went wrong")
fmt.Println(err)
```

The error can be printed because it provides a message through its `Error()` method.

### What does `nil` mean?

For ordinary error-handling code, remember this rule:

- `nil` means no error was returned.
- A non-`nil` error means the function reported a failure.

Example:

```go
err := someFunction()

if err != nil {
    fmt.Println("An error occurred:", err)
}
```

This check asks whether the function returned an error value.

### Important detail

An error is not a special command that automatically terminates your program.

Returning an error and handling an error are two separate actions:

1. The function returns the error.
2. The caller checks it.
3. The caller decides whether to retry, return, log, display a message, or take another action.

---

## 6. Creating Errors with `errors.New()`

The `errors` package provides `errors.New()` to create errors with fixed messages.

```go
package main

import (
    "errors"
    "fmt"
)

func main() {
    err := errors.New("file not found")
    fmt.Println(err)
}
```

Output:

```text
file not found
```

### Reusable errors

Instead of repeatedly creating the same error, you can define a reusable error value:

```go
var errorNotEligible = errors.New("age is less than 18")
```

Now different functions can return that same error:

```go
func checkAge(age int) error {
    if age < 18 {
        return errorNotEligible
    }

    return nil
}
```

This is useful when the calling code needs to recognize a particular error condition.

For example:

```go
err := checkAge(15)

if errors.Is(err, errorNotEligible) {
    fmt.Println("The person is not eligible.")
}
```

### When should you use `errors.New()`?

Use it when you need a fixed error message, especially when you want to reuse the same error value.

---

## 7. Returning Errors from Functions

Go functions can return multiple values.

For example:

```go
func divide(a, b int) (int, error)
```

This function returns:

1. An integer containing the result.
2. An error describing a failure, or `nil` when successful.

### Returning a failure

```go
if b == 0 {
    return 0, errors.New("cannot divide by zero")
}
```

The function returns a placeholder integer and an error.

### Returning success

```go
return a / b, nil
```

The division has succeeded, so the function returns the calculated result and `nil`.

### Why return two values?

The caller needs to know both what happened and what the result was.

```go
result, err := divide(20, 4)
```

Here:

- `result` receives the integer.
- `err` receives the error.

The caller can now decide whether to use the result.

---

## 8. Checking Errors with `if err != nil`

This is one of the most important patterns in Go.

```go
result, err := divide(10, 0)

if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}
```

Read the code like a sentence:

> Call the function. If an error exists, handle it. Otherwise, use the result.

### Flow

```text
       Function call
             |
             v
       Check err != nil
          /       \
        Yes        No
         |          |
         v          v
   Handle error  Use result
```

This pattern is common in file operations, network requests, database calls, parsing, and many other tasks.

### Why check before using the result?

A failed operation might return a placeholder value. If you ignore the error, your program might mistakenly treat that placeholder as a valid result.

Always decide what to do with the error before relying on the result.

---

## 9. Dynamic Error Messages with `fmt.Errorf()`

Sometimes an error needs to include a value, such as a person's age or the name of a file.

`fmt.Errorf()` creates a formatted error message.

```go
func ageValidator(age int) (string, error) {
    if age < 18 {
        return "", fmt.Errorf("age %d is below 18", age)
    }

    return "You are eligible to vote.", nil
}
```

If `age` is `15`, the error message becomes:

```text
age 15 is below 18
```

### Understanding `%d`

The `%d` placeholder is replaced with an integer value.

```go
fmt.Errorf("age %d is below 18", 15)
```

The resulting message is:

```text
age 15 is below 18
```

Other useful formatting placeholders include:

| Placeholder | Meaning |
|---|---|
| `%d` | Integer |
| `%s` | String |
| `%v` | General representation of a value |
| `%q` | Quoted string representation |
| `%w` | Wrap an error inside `fmt.Errorf()` |

### `errors.New()` vs. `fmt.Errorf()`

| `errors.New()` | `fmt.Errorf()` |
|---|---|
| Creates an error from a message. | Creates a formatted error. |
| Useful for fixed messages. | Useful when a message includes variable values or context. |
| Does not wrap an existing error. | Supports error wrapping with `%w`. |

For example:

```go
errors.New("file not found")
```

```go
fmt.Errorf("could not open file %s", filename)
```

---

## 10. Wrapping Errors with `%w`

Error wrapping allows a function to add context to an existing error while preserving the original error's identity.

Consider this reusable error:

```go
var errorNotEligible = errors.New("age is less than 18")
```

A function can wrap it:

```go
func ageValidatorNew(age int) (string, error) {
    if age < 18 {
        return "", fmt.Errorf(
            "age validation failed: %w",
            errorNotEligible,
        )
    }

    return "Age is validated", nil
}
```

For an age of `15`, the error message becomes:

```text
age validation failed: age is less than 18
```

The message has two parts:

- `age validation failed` — extra context.
- `age is less than 18` — the original error.

### What does `%w` preserve?

It preserves the relationship between the new error and the original error. The caller can inspect that relationship using `errors.Is()` or `errors.As()`.

### What if you use `%v` instead?

```go
fmt.Errorf("age validation failed: %v", errorNotEligible)
```

This includes the original error's text, but it does not wrap the original error for inspection.

With `%w`:

```go
fmt.Errorf("age validation failed: %w", errorNotEligible)
```

The original error remains discoverable through the error chain.

### When is wrapping useful?

Wrapping is helpful when a lower-level function knows what failed and a higher-level function knows what operation was being performed.

For example:

```text
loading application configuration
    └── opening configuration file
            └── permission denied
```

Each layer can add useful context without losing the original cause.

---

## 11. Checking Errors with `errors.Is()`

`errors.Is()` checks whether an error matches a target error, including an error wrapped inside it.

```go
if errors.Is(err, errorNotEligible) {
    fmt.Println("The person is not eligible.")
}
```

Read this as:

> Does this error represent `errorNotEligible`, either directly or through wrapping?

### Example 1: Direct error

```go
err := errorNotEligible

fmt.Println(errors.Is(err, errorNotEligible))
```

Output:

```text
true
```

The error is the same reusable error value.

### Example 2: Wrapped error

```go
err := fmt.Errorf("validation failed: %w", errorNotEligible)

fmt.Println(errors.Is(err, errorNotEligible))
```

Output:

```text
true
```

The original error is wrapped inside the new error, so `errors.Is()` can still find it.

### Example 3: A different error

```go
err := errors.New("file not found")

fmt.Println(errors.Is(err, errorNotEligible))
```

Output:

```text
false
```

The error does not match the target error.

### Why not compare error messages?

Avoid relying on comparisons such as:

```go
if err.Error() == "age is less than 18" {
    // Handle the condition.
}
```

Error messages are intended for humans and may change. Use `errors.Is()` to check a known error condition.

---

## 12. Checking Error Types with `errors.As()`

`errors.As()` is different from `errors.Is()`.

- `errors.Is()` checks for a particular error value or condition.
- `errors.As()` checks whether an error chain contains a value assignable to a particular type.

This is useful when an error contains structured information.

### Example: Custom error type

```go
package main

import (
    "errors"
    "fmt"
)

type AgeError struct {
    Age int
}

func (e *AgeError) Error() string {
    return fmt.Sprintf("age %d is below 18", e.Age)
}

func validateAge(age int) error {
    if age < 18 {
        return &AgeError{Age: age}
    }

    return nil
}

func main() {
    err := validateAge(15)

    if err != nil {
        var ageErr *AgeError

        if errors.As(err, &ageErr) {
            fmt.Println("Invalid age:", ageErr.Age)
        }
    }
}
```

Output:

```text
Invalid age: 15
```

### How does this work?

**Step 1: Define the type.**

```go
type AgeError struct {
    Age int
}
```

This error carries an integer field named `Age`.

**Step 2: Implement `Error()`.**

```go
func (e *AgeError) Error() string {
    return fmt.Sprintf("age %d is below 18", e.Age)
}
```

The method gives the custom error a readable message. Because it implements `Error() string`, `*AgeError` satisfies the `error` interface.

**Step 3: Return the error.**

```go
return &AgeError{Age: age}
```

This returns a pointer to an `AgeError`.

**Step 4: Extract the error type.**

```go
var ageErr *AgeError

if errors.As(err, &ageErr) {
    fmt.Println(ageErr.Age)
}
```

`errors.As()` searches the error chain and assigns the matching error to `ageErr`. The program can then inspect its fields.

---

## 13. The Early Return Pattern

The early return pattern handles a failure immediately and avoids running code that depends on a successful result.

```go
result, err := divide(10, 0)

if err != nil {
    fmt.Println("Error:", err)
    return
}

fmt.Println("Result:", result)
```

The `return` statement exits the current function.

Because this example is inside `main()`, it exits `main()` and ends the program's normal execution.

### Why is this useful?

Without early returns, code can become deeply nested:

```go
if err == nil {
    if result > 0 {
        // More logic...
    }
}
```

With early returns, the successful path stays easy to read:

```go
result, err := divide(10, 2)
if err != nil {
    return
}

fmt.Println(result)
```

Use early returns when a failure means the current function cannot sensibly continue.

---

## 14. Ignoring Errors and Common Mistakes

### Mistake 1: Ignoring an error

```go
result, _ := divide(10, 0)
```

The underscore discards the error. This means the caller cannot respond to the failure.

Prefer:

```go
result, err := divide(10, 0)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

### Mistake 2: Using a result after failure

A returned result might be only a placeholder when the operation fails.

Check the error before relying on that result.

### Mistake 3: Assuming an error stops the program

```go
if err != nil {
    fmt.Println(err)
}

fmt.Println("Continuing...")
```

The second print statement still runs. Returning an error is not the same as terminating the program.

### Mistake 4: Using `%v` when wrapping is needed

Use `%w` when you want callers to inspect the original error with `errors.Is()` or `errors.As()`.

### Mistake 5: Comparing only error strings

Use error values and error types to identify conditions rather than relying on exact message text.

### Mistake 6: Repeating the same error context

Add useful context at meaningful boundaries. Avoid wrapping the same message repeatedly without adding information.

---

## 15. Errors vs. Panics

Errors and panics serve different purposes.

| Errors | Panics |
|---|---|
| Represent failures that code can handle normally. | Interrupt normal execution. |
| Usually returned as values. | Triggered with `panic()`. |
| Common for invalid input, file access, and network failures. | Usually reserved for exceptional situations and serious programming problems. |
| The caller decides how to respond. | Execution unwinds the current goroutine unless recovered. |

For example, division by zero in a function accepting arbitrary input should normally be handled by returning an error:

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }

    return a / b, nil
}
```

A panic is not a replacement for ordinary error handling.

---

## 16. Complete Error Handling Flow

The complete pattern used in this project is:

```text
                    CALLER
                       |
                       v
               Call a function
                       |
                       v
             Function performs work
                       |
                       v
                Did it succeed?
                  /         \
                No           Yes
                |             |
                v             v
          Return error    Return result
                |             |
                v             v
          Caller checks   Caller checks
          err != nil      err == nil
                |             |
                v             v
          Handle error    Use the result
                |
                v
       Optionally wrap the error
                |
                v
       Inspect with errors.Is()
       or errors.As() if needed
```

Remember that wrapping and error inspection are optional. A simple function may need only to return an error and have its caller check `err != nil`.

---

## 17. Expected Program Output

For the `main.go` example that demonstrates division, age validation, wrapped errors, and reusable errors, the output should be similar to:

```text
========== GO ERROR HANDLING ==========
Division Error: cannot divide by zero
Successful Division Result: 5
Age Validation Error: age 15 is below 18
Eligibility: You are eligible to vote.
Detailed Validation Error: age validation failed: age is less than 18
Reason: The person is not eligible.
Reusable Error Detected: age is less than 18
Age Check Successful: Age is valid.
========== PROGRAM FINISHED ==========
```

The order follows the order of the statements in `main()`. The messages come from the errors and strings returned by the functions.

---

## 18. Best Practices

1. **Check errors explicitly.** Use `if err != nil` when a function returns an error.
2. **Handle failures before using results.** A result returned alongside an error may be a placeholder.
3. **Use `errors.New()` for reusable fixed errors.**
4. **Use `fmt.Errorf()` for formatted messages.**
5. **Use `%w` to wrap errors when preserving their identity matters.**
6. **Use `errors.Is()` to check known error conditions.**
7. **Use `errors.As()` to extract structured custom errors.**
8. **Add meaningful context.** Explain which operation failed without repeating the same information unnecessarily.
9. **Return errors to the right level.** A function should usually return a failure to its caller when the caller is better placed to decide what to do.
10. **Avoid panics for ordinary expected failures.**
11. **Avoid ignoring errors without a good reason.**
12. **Keep error messages clear and useful for debugging.**

---

## 19. Practice Exercises

Try to solve these exercises without copying the examples immediately.

### Exercise 1 — Division

Create a `multiplyOrDivide()` function that returns a result and an error. Return an error if division is requested with a zero denominator.

### Exercise 2 — Age validation

Create a function that accepts an age and returns an error if the age is below 18. Return `nil` for a valid age.

### Exercise 3 — Reusable errors

Create a reusable error named `ErrNotEligible` and return it from two different validation functions.

### Exercise 4 — Dynamic messages

Use `fmt.Errorf()` to include a username or age in an error message.

### Exercise 5 — Error wrapping

Wrap a reusable error with `%w`, then use `errors.Is()` to verify that the original error can still be identified.

### Exercise 6 — Custom error types

Create a `TemperatureError` struct containing a temperature field. Implement `Error() string` and use `errors.As()` to retrieve the temperature.

### Exercise 7 — Early returns

Rewrite a function with nested `if` statements using early returns to make the successful path easier to read.

---

## 20. Final Takeaways

Go error handling is built around a simple idea: **functions report errors, and callers decide how to respond.**

Remember these core tools:

| Tool | Main purpose |
|---|---|
| `errors.New()` | Create an error with a fixed message. |
| `fmt.Errorf()` | Create a formatted error. |
| `if err != nil` | Check whether an error occurred. |
| `return nil` | Report that no error occurred. |
| `%w` | Wrap an existing error. |
| `errors.Is()` | Check for a particular error value or condition. |
| `errors.As()` | Extract an error of a particular type. |
| Early `return` | Stop the current function when it cannot continue safely. |

Once these concepts are clear, you can apply them to file handling, APIs, databases, network requests, configuration loading, and larger Go applications.

**Learning goal:** Do not merely print an error. Understand where it originated, preserve useful context, and decide deliberately what your program should do next.
