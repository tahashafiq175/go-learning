# Go Learning — Variables and Data Types

## 1. Introduction

In Go, **variables** are used to store information such as numbers, names, marks, and other values.

Think of a variable as a **box**. Every box has a name, and it stores a value.

For example:

```go
var age int = 19
```

Imagine a box named `age` that stores the number `19`.

- `var` tells Go that we want to declare a variable.
- `age` is the variable's name.
- `int` is the data type.
- `19` is the value stored in the variable.

## 2. Declaring Variables with Explicit Data Types

We can declare a variable by specifying its data type ourselves.

```go
var age int = 19
var name string = "Taha"

fmt.Println(age)
fmt.Println(name)
```

**Output:**

```text
19
Taha
```

### Understanding the Code

- `int` is used for integer numbers, such as `19`, `100`, and `-5`.
- `string` is used for text, such as `"Taha"` and `"Hello Go"`.
- `fmt.Println()` prints the supplied values.

**Remember:** When we explicitly specify a data type, we are telling Go what kind of value the variable should store.

## 3. Changing the Value of a Variable

After declaring a variable, we can change the value stored inside it.

```go
var age int = 19

age = 20

fmt.Println(age)
```

**Output:**

```text
20
```

Think of the variable as a box. Initially, the box contains `19`. We replace that value with `20`.

### Important Concept: `=` Operator

The `=` operator assigns a value to an existing variable.

```go
age = 20
```

Here, we are not declaring another variable. We are updating the existing variable.

**Important:** Changing a variable's value does not change its data type.

For example, an `int` variable can store another integer, but it cannot store a string such as `"Taha"`.

## 4. Type Inference Using `var`

Go can automatically determine a variable's data type from its initial value.

```go
var university = "Islamia University of Bahawalpur"

fmt.Println(university)
```

**Output:**

```text
Islamia University of Bahawalpur
```

Notice that we did not write `string`.

Go sees the text inside the quotation marks and infers that the variable's type is `string`.

The following declarations are equivalent in their resulting variable type:

```go
var university string = "Islamia University of Bahawalpur"

var university = "Islamia University of Bahawalpur"
```

### What Is Type Inference?

**Type inference** means that Go determines the data type from the value provided during initialization.

This saves us from writing the data type explicitly every time.

However, Go is still a **statically typed language**. The variable's type is determined at compile time and does not change when we assign a new value.

## 5. Short Variable Declaration Using `:=`

Go provides a shorter way to declare and initialize variables.

```go
gpa := 3.9

fmt.Println(gpa)
```

**Output:**

```text
3.9
```

This is equivalent to:

```go
var gpa float64 = 3.9
```

### Understanding `:=`

The `:=` operator declares a new variable and initializes it in one step.

- `gpa` is the variable name.
- `3.9` is the initial value.
- Go infers the type as `float64`.

**Important:** Short variable declarations using `:=` are permitted inside functions, not directly at package level.

For example:

```go
func main() {
    score := 100
    fmt.Println(score)
}
```

## 6. Printing the Data Type Using `%T`

Sometimes, we want to know the data type of a variable.

Go provides the `%T` format verb through `fmt.Printf()`.

```go
gpa := 3.9

fmt.Printf("%T\n", gpa)
```

**Output:**

```text
float64
```

### Common Format Verbs

| Format verb | Purpose | Example |
|---|---|---|
| `%T` | Prints the data type | `float64` |
| `%v` | Prints the value in its default format | `3.9` |
| `%d` | Formats an integer in decimal notation | `19` |
| `%f` | Formats a floating-point number | `3.900000` |
| `%s` | Formats a string | `Taha` |

Example:

```go
age := 19

fmt.Printf("Value: %v\n", age)
fmt.Printf("Type: %T\n", age)
```

**Output:**

```text
Value: 19
Type: int
```

The `\n` character moves the output to the next line.

## 7. Variables Cannot Change Their Data Types

Consider the following code:

```go
gpa := 3.9

gpa = 8.9

fmt.Println(gpa)
```

**Output:**

```text
8.9
```

This works because both values can be assigned to the `float64` variable.

However, the following code causes a compilation error:

```go
gpa := 3.9

gpa = "Taha"
```

Why?

Because `gpa` is a `float64` variable, while `"Taha"` is a string.

Go does not allow us to assign a string directly to a variable of type `float64`.

**Remember:** A variable's value can change, but its declared type cannot change.

## 8. Declaring Multiple Variables Using `var`

Go allows us to declare multiple variables inside a single `var` block.

```go
var (
    firstName = "Taha"
    score     = 100
)

fmt.Println(firstName)
fmt.Println(score)
```

**Output:**

```text
Taha
100
```

### Why Use a `var` Block?

Imagine that we have several variables related to a student.

Instead of writing separate declarations, we can group them together.

Go infers each variable's type independently:

- `firstName` has type `string`.
- `score` has type `int`.

The variables do not need to share the same data type.

## 9. Declaring Multiple Variables Using `:=`

We can also declare multiple variables in a single short declaration.

```go
x, y := 10, 20

fmt.Println(x, y)
```

**Output:**

```text
10 20
```

Here:

- `x` receives `10`.
- `y` receives `20`.
- Both variables have type `int`.

This is called a **multiple short variable declaration**.

### Important Rule

At least one variable on the left side must be newly declared in the current scope.

For example:

```go
x := 10

x, y := 20, 30

fmt.Println(x, y)
```

**Output:**

```text
20 30
```

This works because `y` is new, even though `x` already exists.

However, declaring only existing variables with `:=` is not allowed:

```go
x := 10
y := 20

x, y := 30, 40 // Compilation error
```

Use `=` when all variables already exist:

```go
x, y = 30, 40
```

## 10. Swapping Two Variables

Swapping means exchanging the values of two variables.

Go allows us to swap variables without creating a temporary variable.

```go
x, y := 10, 20

y, x = x, y

fmt.Println(x, y)
```

**Output:**

```text
20 10
```

### Step-by-Step Explanation

Initially:

```text
x = 10
y = 20
```

The statement:

```go
y, x = x, y
```

takes the values from the right side and assigns them to the left side.

After swapping:

```text
x = 20
y = 10
```

**Important:** Go evaluates the right-hand-side expressions before assigning the resulting values to the left-hand variables.

Also, because `x` and `y` already exist, we use `=` rather than `:=`.

## 11. Checking Data Types After Swapping

```go
x, y := 10, 20

y, x = x, y

fmt.Printf("%T %T\n", x, y)
```

**Output:**

```text
int int
```

Swapping changes the values stored in the variables, not their data types.

Both variables were declared using integer values, so their types remain `int`.

## 12. Type Conversion: `int` to `float64`

Sometimes, we need to convert a value from one data type to another.

For example:

```go
age := 20

newAge := float64(age)

fmt.Println(newAge)
fmt.Printf("%T\n", newAge)
```

**Output:**

```text
20
float64
```

### What Happened?

Initially, `age` is an integer.

```go
age := 20
```

Its type is `int`.

Next:

```go
newAge := float64(age)
```

The `float64()` conversion produces a floating-point value, which is stored in `newAge`.

Now:

- `age` remains an `int`.
- `newAge` is a `float64`.

**Important:** Go requires explicit numeric conversions. Converting `age` does not change the type of the original variable.

## 13. Assigning a Decimal Value to `float64`

```go
newAge := float64(20)

newAge = 20.7

fmt.Println(newAge)
```

**Output:**

```text
20.7
```

Because `newAge` is a `float64`, it can represent fractional values such as `20.7`.

We use `=` because the variable has already been declared.

## 14. Type Conversion: `float64` to `int`

Now, suppose we have a floating-point number and want an integer.

```go
newAge := 20.7

age := int(newAge)

fmt.Println(age)
```

**Output:**

```text
20
```

### Why Did We Get `20` Instead of `21`?

When converting a floating-point value to an integer, Go discards the fractional part. It does not round to the nearest integer.

Examples:

| Conversion | Result |
|---|---:|
| `int(20.7)` | `20` |
| `int(20.9)` | `20` |
| `int(20.1)` | `20` |
| `int(-20.7)` | `-20` |

This behavior is called **truncation toward zero**.

The original floating-point value remains unchanged unless you explicitly assign a new value to that variable.

**Important:** Converting a floating-point number to an integer may lose information.

## 15. Understanding `main()` and Custom Functions

Go executable programs use the `main` package and a `main()` function as their entry point.

Consider:

```go
package main

import "fmt"

func datatype() {
    age := 19
    fmt.Println(age)
}

func main() {
    datatype()
}
```

**Output:**

```text
19
```

### What Is Happening?

- `func datatype()` defines a function named `datatype`.
- `func main()` defines the program's entry-point function.
- `datatype()` calls the function, causing its statements to execute.

Defining a function does not automatically execute its body. It must be called, directly or indirectly, for its statements to run.

**Remember:** A function definition describes what to do. A function call actually requests that the work be performed.

## 16. Complete Example

```go
package main

import "fmt"

func datatype() {
    var age int = 19
    var name string = "Taha"

    fmt.Println(age)
    fmt.Println(name)

    age = 20
    fmt.Println(age)

    var university = "Islamia University of Bahawalpur"
    fmt.Println(university)

    gpa := 3.9
    fmt.Println(gpa)
    fmt.Printf("%T\n", gpa)

    gpa = 8.9
    fmt.Println(gpa)

    var (
        firstName = "Taha"
        score     = 100
    )

    fmt.Println(firstName)
    fmt.Println(score)

    x, y := 10, 20
    fmt.Println(x, y)

    y, x = x, y
    fmt.Println(x, y)

    newAge := float64(age)
    fmt.Println(newAge)

    newAge = 20.7
    fmt.Println(newAge)

    age = int(newAge)
    fmt.Println(age)
}

func main() {
    datatype()
}
```

## 17. Quick Revision

| Concept | Key idea |
|---|---|
| `var` | Declares a variable |
| `int` | Stores integer values |
| `string` | Stores text |
| `float64` | Stores floating-point values |
| Type inference | Go determines a variable's type from its initial value |
| `:=` | Declares and initializes a variable inside a function |
| `=` | Assigns a value to an existing variable |
| `%T` | Prints a value's data type |
| `var (...)` | Groups multiple variable declarations |
| `x, y := 10, 20` | Declares multiple variables |
| `y, x = x, y` | Swaps two variable values |
| `float64(age)` | Converts a numeric value to `float64` |
| `int(newAge)` | Converts a floating-point value to an integer, truncating its fractional part |
| `main()` | Entry point for a Go executable |

## 18. Practice Exercises

Try solving these exercises without looking at the complete example.

1. Declare an integer variable named `marks` with the value `95`.
2. Declare a string variable named `studentName` using type inference.
3. Create a variable `temperature` using `:=` and assign it `36.5`. Print its data type.
4. Declare two integer variables and swap their values.
5. Convert an integer value to `float64`.
6. Convert `45.8` to an integer and explain the result.
7. Declare three variables inside a `var` block.
8. Write a `main()` function that calls your own custom function.

**Final takeaway:** Go variables are simple boxes that hold values of specific types. Once you understand declarations, assignment, type inference, swapping, and explicit type conversion, you have a strong foundation for writing Go programs.
