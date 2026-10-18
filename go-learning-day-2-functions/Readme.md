# Go Learning — Functions in Go

## 1. Introduction to Functions

A **function** is a reusable block of code that performs a specific task.

Imagine you have a calculator. Whenever you want to add two numbers, you do not need to build a new calculator. You simply provide two numbers, and the calculator performs the operation.

Functions work in a similar way. We write the logic once and call the function whenever we need it.

### Example

```go
package main

import "fmt"

func greet() {
    fmt.Println("Hello, Taha!")
}

func main() {
    greet()
}
```

**Output:**

```text
Hello, Taha!
```

### Understanding the Code

- `func` is the keyword used to declare a function.
- `greet` is the function's name.
- `()` means this function does not accept any parameters.
- `{}` contains the function body.
- `greet()` calls the function and executes its code.
- `main()` is the entry point of an executable Go program.

**Important:** Defining a function does not automatically execute it. We must call it for its statements to run.

---

## 2. Basic Syntax of a Function

The general syntax of a function is:

```go
func functionName(parameters) {
    // Function body
}
```

Let's understand each part.

| Part | Meaning |
|---|---|
| `func` | Tells Go that we are declaring a function |
| `functionName` | The name used to call the function |
| `parameters` | Inputs the function can receive |
| `{}` | Contains the code executed by the function |

A function can accept inputs, perform operations, return results, or simply perform an action such as printing a message.

---

## 3. Functions with Parameters

**Parameters** allow us to send information into a function.

Imagine a calculator that needs two numbers before it can add them. Those numbers are the inputs of the operation.

### Example

```go
package main

import "fmt"

func calculation(a int, b int) {
    fmt.Println(a + b)
}

func main() {
    calculation(10, 20)
}
```

**Output:**

```text
30
```

### Understanding the Code

```go
func calculation(a int, b int)
```

- `calculation` is the function's name.
- `a` is the first parameter, and its datatype is `int`.
- `b` is the second parameter, and its datatype is `int`.

When we call:

```go
calculation(10, 20)
```

Go assigns `10` to `a` and `20` to `b`.

The function then calculates `a + b` and prints the result.

### Important Concept: Parameters vs. Arguments

These two terms are related but different.

```go
func calculation(a, b int) { // Parameters
    fmt.Println(a + b)
}

calculation(10, 20) // Arguments
```

- **Parameters** are the named variables in the function definition.
- **Arguments** are the actual values supplied during a function call.

Think of parameters as empty boxes and arguments as the values placed inside those boxes.

---

## 4. Shorthand Parameter Declaration

Go allows us to shorten parameter declarations when consecutive parameters have the same datatype.

These two declarations are equivalent:

```go
func calculation(a int, b int) {
    fmt.Println(a + b)
}
```

```go
func operation(a, b int) {
    fmt.Println(a + b)
}
```

In the second example, the `int` written after `b` applies to both `a` and `b`.

### When Parameters Have Different Types

If the parameters have different datatypes, specify each type where necessary.

```go
func student(name string, age int) {
    fmt.Println(name, age)
}
```

Here:

- `name` has type `string`.
- `age` has type `int`.

**Remember:** The shorthand works when consecutive parameter names share the same datatype.

---

## 5. Functions Without Return Values

Some functions perform an action without returning a result to the caller.

For example:

```go
func calculation(a, b int) {
    fmt.Println(a + b)
}
```

When we call:

```go
calculation(10, 20)
```

The function prints:

```text
30
```

However, the function does not return the number `30`.

### Printing vs. Returning

Printing and returning are two different things.

**Printing:**

```go
func add(a, b int) {
    fmt.Println(a + b)
}
```

This displays the result but does not return it.

**Returning:**

```go
func add(a, b int) int {
    return a + b
}
```

This sends the result back to the caller, allowing it to be stored, printed, or used in another calculation.

**Important:** A function that has no explicit return type can still execute operations and print values. It simply does not return a result value.

---

## 6. Functions with a Return Value

Sometimes, we want a function to calculate something and give the result back to the code that called it.

### Example

```go
package main

import "fmt"

func returnFunc(a, b int) int {
    return a + b
}

func main() {
    sum := returnFunc(100, 30)
    fmt.Println(sum)
}
```

**Output:**

```text
130
```

### Understanding the Syntax

```go
func returnFunc(a, b int) int
```

The first `int` describes the datatype of the parameters.

The final `int`, written after the closing parenthesis, specifies the **return type**.

It tells Go that the function returns an integer.

### How Does `return` Work?

Consider:

```go
return a + b
```

The function calculates the sum and sends the result back to the caller.

When we write:

```go
sum := returnFunc(100, 30)
```

the returned value `130` is assigned to `sum`.

Then:

```go
fmt.Println(sum)
```

prints the result.

### Step-by-Step Execution

1. The program calls `returnFunc(100, 30)`.
2. The parameters receive `100` and `30`.
3. The function calculates `100 + 30`.
4. `return` sends `130` back to the caller.
5. The variable `sum` stores `130`.
6. `fmt.Println(sum)` displays `130`.

**Remember:** `return` sends a value back to the caller. It also ends the current function execution at that point.

---

## 7. Functions with Multiple Return Values

One of Go's useful features is that a function can return more than one value.

For example, a calculator might return both the sum and the product of two numbers.

### Example

```go
func returnFuncs(a, b int) (int, int) {
    sum := a + b
    product := a * b

    return sum, product
}
```

### Understanding the Syntax

```go
func returnFuncs(a, b int) (int, int)
```

The two types inside the parentheses indicate that the function returns two integer values.

The return statement:

```go
return sum, product
```

sends both values back to the caller in the specified order.

### Calling the Function

```go
add, multiply := returnFuncs(20, 50)

fmt.Println(add)
fmt.Println(multiply)
```

**Output:**

```text
70
1000
```

### How Are the Values Assigned?

The first returned value is assigned to `add`.

The second returned value is assigned to `multiply`.

| Variable | Calculation | Value |
|---|---|---:|
| `add` | `20 + 50` | `70` |
| `multiply` | `20 * 50` | `1000` |

**Important:** The order of returned values matters. Go assigns them to the receiving variables in the same order.

---

## 8. Returning Multiple Values Directly

We do not always need to create temporary variables before returning values.

The same function can be written more concisely:

```go
func manyReturnFuncs(a, b int) (int, int) {
    return a + b, a * b
}
```

Calling it:

```go
add, multiply := manyReturnFuncs(20, 50)

fmt.Println(add)
fmt.Println(multiply)
```

**Output:**

```text
70
1000
```

### Difference Between the Two Approaches

**Using intermediate variables:**

```go
func returnFuncs(a, b int) (int, int) {
    sum := a + b
    product := a * b

    return sum, product
}
```

**Returning expressions directly:**

```go
func manyReturnFuncs(a, b int) (int, int) {
    return a + b, a * b
}
```

Both approaches work.

The first approach can be easier to understand when learning or when intermediate results need meaningful names. The second is convenient when the calculations are simple.

---

## 9. Named Return Values

Go allows us to name the values returned by a function.

Consider this example:

```go
func modReturnFuncs(age int, name string) (message string) {
    message = fmt.Sprintf(
        "My name is %s with age of %d",
        name,
        age,
    )

    return
}
```

Here, `message` is a **named return value**.

### What Does `(message string)` Mean?

It tells Go that:

- The function returns a value of type `string`.
- The return value has the name `message`.
- `message` is available as a local variable inside the function.

We can assign a value directly to this variable:

```go
message = fmt.Sprintf(
    "My name is %s with age of %d",
    name,
    age,
)
```

The function then uses:

```go
return
```

This is called a **bare return**. It returns the current value of the named return variable.

### Calling the Function

```go
details := modReturnFuncs(10, "Taha")
fmt.Println(details)
```

**Output:**

```text
My name is Taha with age of 10
```

### Important: Named Return Values Have Zero Values

When a named return variable is declared, Go initializes it to its datatype's zero value.

For example:

- `int` starts at `0`.
- `string` starts as `""`.
- `bool` starts as `false`.

If we do not assign a new value to `message`, a bare `return` returns an empty string.

**Best practice:** Named return values can be helpful when they make a function easier to understand. However, for simple functions, explicit return statements are often clearer.

---

## 10. Understanding `fmt.Sprintf()`

The `fmt.Sprintf()` function creates a formatted string and returns it.

Unlike `fmt.Println()`, it does not print the result automatically.

### Example

```go
name := "Taha"
age := 10

message := fmt.Sprintf(
    "My name is %s with age of %d",
    name,
    age,
)

fmt.Println(message)
```

**Output:**

```text
My name is Taha with age of 10
```

### Understanding the Format Verbs

| Format verb | Meaning |
|---|---|
| `%s` | Formats a string |
| `%d` | Formats an integer in decimal notation |
| `%f` | Formats a floating-point number |
| `%v` | Formats a value using its default representation |
| `%T` | Prints the datatype of a value |

In the example, `%s` is replaced by `"Taha"` and `%d` is replaced by `10`.

**Remember:**

- `fmt.Println()` prints values.
- `fmt.Sprintf()` builds and returns a string.
- `return` sends a function's result back to its caller.

---

## 11. Calling One Function from Another

A function can call other functions.

This helps us organize code into smaller, reusable tasks.

### Example

```go
func main() {
    fmt.Println("hello world")
    greet()
}

func greet() {
    fmt.Println("greeting from taha")
    calculation(10, 20)
    operation(30, 10)
}
```

**Output:**

```text
hello world
greeting from taha
30
40
```

### How Does Execution Flow?

1. Go starts executing `main()`.
2. `hello world` is printed.
3. `greet()` is called.
4. `greet()` prints `greeting from taha`.
5. `calculation(10, 20)` is called and prints `30`.
6. `operation(30, 10)` is called and prints `40`.
7. `greet()` finishes, and execution returns to the caller.

**Important:** Function definitions can appear in different places in the same Go source file. The order of function definitions does not determine the order of execution; function calls do.

---

## 12. Complete Example

The following example combines the concepts covered in this lesson.

```go
package main

import "fmt"

func calculation(a int, b int) {
    fmt.Println(a + b)
}

func operation(a, b int) {
    fmt.Println(a + b)
}

func returnFunc(a, b int) int {
    return a + b
}

func returnFuncs(a, b int) (int, int) {
    sum := a + b
    product := a * b

    return sum, product
}

func manyReturnFuncs(a, b int) (int, int) {
    return a + b, a * b
}

func modReturnFuncs(age int, name string) (message string) {
    message = fmt.Sprintf(
        "My name is %s with age of %d",
        name,
        age,
    )

    return
}

func main() {
    fmt.Println("hello world")
    greet()
}

func greet() {
    fmt.Println("greeting from taha")

    calculation(10, 20)
    operation(30, 10)

    sum := returnFunc(100, 30)

    fmt.Println("Calling the sum function")
    fmt.Println(sum)

    add, multiply := returnFuncs(20, 50)

    fmt.Println("Calling the function with two return values")
    fmt.Println(add)
    fmt.Println(multiply)

    details := modReturnFuncs(10, "Taha")
    fmt.Println(details)

    // Calling the other multiple-return function:
    add2, multiply2 := manyReturnFuncs(5, 6)
    fmt.Println(add2, multiply2)
}
```

### Expected Output

```text
hello world
greeting from taha
30
40
Calling the sum function
130
Calling the function with two return values
70
1000
My name is Taha with age of 10
11 30
```

---

## 13. Quick Revision

| Concept | Key idea |
|---|---|
| Function | A reusable block of code |
| `func` | Declares a function |
| Parameters | Named inputs in a function definition |
| Arguments | Actual values supplied during a function call |
| Return type | Specifies the type of returned value |
| `return` | Sends result values back to the caller |
| Multiple return values | A function can return several values |
| Named return value | Gives a name to a return variable |
| Bare return | Returns the current named return values |
| `fmt.Println()` | Prints values |
| `fmt.Sprintf()` | Creates and returns a formatted string |
| `main()` | Entry point of an executable Go program |
| Function call | Executes a function's statements |

---

## 14. Practice Exercises

Try solving these exercises yourself.

1. Create a function named `greetStudent` that accepts a student's name and prints a greeting.
2. Create a function named `add` that accepts two integers and returns their sum.
3. Create a function named `multiply` that returns the product of two integers.
4. Create a function that returns both the sum and product of two numbers.
5. Create a function that accepts a name and age and returns a formatted message using `fmt.Sprintf()`.
6. Rewrite a function using named return values and a bare `return`.
7. Create a function that returns two values directly without using intermediate variables.
8. Create a `main()` function that calls at least three custom functions.

## Final Takeaway

Functions make Go programs reusable, organized, and easier to maintain.

- Use **parameters** to send data into a function.
- Use **return values** to send results back.
- Use **multiple return values** when a function needs to return more than one result.
- Use **named return values** when they make the code easier to understand.
- Use **function calls** to execute the code inside your functions.

Once you understand these concepts, you will have a strong foundation for writing reusable Go programs.
