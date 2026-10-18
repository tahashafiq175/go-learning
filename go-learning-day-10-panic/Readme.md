# Panic, Defer, and Recover in Go

A beginner-friendly guide to understanding exceptional situations, deferred execution, and panic recovery in Go.

---

## Table of Contents

1. [Introduction](#1-introduction)
2. [What Is a Panic?](#2-what-is-a-panic)
3. [What Is Defer?](#3-what-is-defer)
4. [What Is Recover?](#4-what-is-recover)
5. [Understanding Panic, Defer, and Recover Together](#5-understanding-panic-defer-and-recover-together)
6. [Complete Example](#6-complete-example)
7. [Program Execution Flow](#7-program-execution-flow)
8. [Expected Output](#8-expected-output)
9. [Important Rules of Recover](#9-important-rules-of-recover)
10. [Defer Execution Order](#10-defer-execution-order)
11. [Panic vs Error](#11-panic-vs-error)
12. [Best Practices](#12-best-practices)
13. [Common Mistakes](#13-common-mistakes)
14. [Quick Revision](#14-quick-revision)

---

## 1. Introduction

Go provides three important mechanisms for dealing with exceptional situations and controlling what happens when a function finishes:

| Concept | Simple Meaning |
|---|---|
| `panic` | Signals that normal execution cannot continue in the usual way. |
| `defer` | Schedules a function to run when the surrounding function returns or unwinds. |
| `recover` | Stops an active panic when called correctly from a deferred function. |

Think of a program as a journey.

- **Panic:** Something unexpected happens, and the current journey is interrupted.
- **Defer:** A task is scheduled to happen before the current function finishes.
- **Recover:** A safety mechanism catches a panic and allows the program to continue at the caller.

These three concepts are related, but each has a different responsibility.

## 2. What Is a Panic?

A `panic` interrupts the normal execution of the current goroutine and starts unwinding the call stack.

In simple words, the program says:

"I have encountered a serious situation, so I cannot continue with the normal instructions in this function."

### Basic example

```go
package main

import "fmt"

func main() {
    fmt.Println("Program started")

    panic("Something went wrong")

    fmt.Println("Program finished")
}
```

### Explanation

- `fmt.Println("Program started")` prints the first message.
- `panic("Something went wrong")` starts a panic.
- The final `fmt.Println` is not executed because execution does not resume at the statement after the panic.

Without recovery, the panic eventually reaches the top of the goroutine and the program terminates with a panic report.

### Output

```text
Program started
panic: Something went wrong
```

The exact diagnostic output may include additional runtime information.

### When does panic occur?

A panic can happen because of:

1. An explicit call to `panic()`.
2. Certain runtime errors, such as indexing a slice outside its bounds.
3. Dereferencing a nil pointer in situations where Go cannot perform the operation.

Example of an explicit panic:

```go
panic("Invalid program state")
```

The string is the message associated with the panic.

**Important:** A panic is not the normal way to report every error. Expected failures, such as a file not existing or a user entering invalid input, are usually handled by returning an `error`.

---

## 3. What Is Defer?

The `defer` keyword schedules a function call to execute when the surrounding function is about to finish.

That function may finish because it:

- Reaches the end of its body.
- Executes a `return`.
- Encounters a panic and begins unwinding.

### Basic example

```go
package main

import "fmt"

func main() {
    defer fmt.Println("This runs later")

    fmt.Println("This runs first")
    fmt.Println("This runs second")
}
```

### Explanation

```go
defer fmt.Println("This runs later")
```

Go schedules the print statement for later. It does not print the message immediately.

The two ordinary print statements execute first. When `main()` is about to finish, the deferred call runs.

### Output

```text
This runs first
This runs second
This runs later
```

### Why is defer useful?

Suppose a program opens a file. After finishing its work, the file should be closed.

```go
file, err := os.Open("demo.txt")

if err != nil {
    fmt.Println("Error:", err)
    return
}

defer file.Close()
```

- `os.Open()` attempts to open the file.
- `err != nil` checks whether opening the file failed.
- `defer file.Close()` schedules the file to close when the surrounding function finishes.

This helps prevent resources from remaining open unnecessarily.

Other common uses include:

- Closing files.
- Releasing locks.
- Closing network connections.
- Cleaning up resources.
- Running recovery logic.

---

## 4. What Is Recover?

`recover()` is a built-in Go function that can stop an active panic.

However, it works only when called directly from a deferred function executing in the same goroutine that is panicking.

### Basic example

```go
package main

import "fmt"

func main() {
    defer func() {
        result := recover()

        if result != nil {
            fmt.Println("Recovered from panic:", result)
        }
    }()

    fmt.Println("Before panic")

    panic("Something went wrong")

    fmt.Println("After panic")
}
```

### Explanation

```go
defer func() {
```

This schedules an anonymous function to run when `main()` returns or begins unwinding because of a panic.

```go
result := recover()
```

When the deferred function executes during the panic, `recover()` stops that panic and returns its value.

```go
if result != nil {
```

If `result` is not `nil`, the deferred function recovered an active panic.

```go
fmt.Println("Recovered from panic:", result)
```

The program prints the recovered panic value.

### Output

```text
Before panic
Recovered from panic: Something went wrong
```

The message `After panic` is not printed.

**Why?** Recovering from a panic does not rewind execution to the statement after `panic()`. Instead, the panicking function finishes unwinding, and execution can continue in its caller.

In this example, `main()` itself is the panicking function. Its deferred function recovers the panic, but `main()` does not resume at the next print statement.

---

## 5. Understanding Panic, Defer, and Recover Together

Now combine all three concepts using a division example.

Imagine a function receives two numbers:

```text
divide(10, 0)
```

Division by zero is not allowed in this example. We explicitly call `panic()` when the second number is zero.

### The idea

1. `main()` calls `divide(10, 0)`.
2. `divide()` schedules its deferred recovery function.
3. The function checks whether the divisor is zero.
4. Because the divisor is zero, the function calls `panic()`.
5. Go begins unwinding the call stack.
6. The deferred function executes.
7. `recover()` catches the panic.
8. `divide()` finishes without resuming after the panic.
9. `main()` continues after the function call.

The next section puts this complete flow into code.

---

## 6. Complete Example

Create a file named `main.go`.

```go
package main

import (
    "fmt"
)

func main() {
    // Step 1: Print a message before calling divide().
    fmt.Println("Hello World")

    // Step 2: Call divide() with 10 and 0.
    answer := divide(10, 0)

    // Step 3: Check the returned value.
    // In this example, recovery makes the unnamed int result zero.
    if answer == 0 {
        fmt.Println("Divide by zero error")
    }

    // Step 4: The caller continues after divide() returns.
    fmt.Println("Answer:", answer)
}

func divide(a, b int) int {
    // Step 5: Schedule the anonymous function.
    // It will execute when divide() returns or unwinds.
    defer func() {
        // Step 7: Attempt to recover from an active panic.
        recoveredValue := recover()

        // Step 8: A non-nil value indicates a recovered panic.
        if recoveredValue != nil {
            fmt.Println("Recovered from panic:", recoveredValue)
        }
    }()

    // Step 6: Check the divisor before division.
    if b == 0 {
        // Start a panic because division by zero is not allowed here.
        panic("Divide by zero error")
    }

    // This line executes only if the panic condition is not reached.
    return a / b
}
```

### Important explanation

The function has this signature:

```go
func divide(a, b int) int
```

It returns one `int`, but the return value is **unnamed**.

When a panic occurs and is recovered by a deferred function, `divide()` does not resume at the statement following `panic()`. As the function finishes unwinding, its unnamed integer return value is the zero value of `int`, which is `0`.

That is why `answer` becomes `0` in this particular example.

### A limitation in this example

```go
if answer == 0 {
    fmt.Println("Divide by zero error")
}
```

Checking only whether `answer` equals zero is not a reliable way to detect division failure. For example, `divide(0, 5)` legitimately returns zero.

For real application code, returning an error is usually the better approach. A later section shows how.

---

## 7. Program Execution Flow

The following diagram shows the order of execution.

```text
              main()
                |
                v
        Print "Hello World"
                |
                v
          divide(10, 0)
                |
                v
      Register deferred function
                |
                v
          Is b == 0?
                |
                v
               YES
                |
                v
       panic("Divide by zero error")
                |
                v
       Begin stack unwinding
                |
                v
      Execute deferred function
                |
                v
             recover()
                |
                v
        Panic is recovered
                |
                v
       divide() finishes
       with unnamed int = 0
                |
                v
       Return to main()
                |
                v
        answer = 0
                |
                v
       Print division message
                |
                v
        Print Answer: 0
```

The central idea is that **recover stops the panic; it does not restart the interrupted function at the panic line**.

---

## 8. Expected Output

```text
Hello World
Recovered from panic: Divide by zero error
Divide by zero error
Answer: 0
```

The order matters:

- `Hello World` prints before the function call.
- The recovery message prints while the deferred function executes.
- The caller then receives the zero-valued result.
- `main()` continues with the instructions after the call to `divide()`.

---

## 9. Important Rules of Recover

### Rule 1: Recover must be called from a deferred function

Correct:

```go
defer func() {
    recoveredValue := recover()
    fmt.Println(recoveredValue)
}()
```

This allows `recover()` to catch a panic while the deferred function is executing during panic unwinding.

### Rule 2: Calling recover normally does not catch a later panic

```go
func main() {
    recoveredValue := recover()
    fmt.Println(recoveredValue)

    panic("Problem")
}
```

Here, `recover()` runs before the panic starts. It returns `nil` and does not protect the program from the later panic.

### Rule 3: Recover does not resume the interrupted function

```go
func example() {
    panic("Stop here")

    // This statement is not resumed after recovery.
}
```

Recovery stops the active panic, but execution does not jump back to the next statement after `panic()`.

### Rule 4: Recover works within the same goroutine

A deferred function in another goroutine cannot recover a panic occurring in this goroutine.

Each goroutine has its own execution and panic-unwinding behavior.

### Rule 5: Recover is not a replacement for ordinary error handling

Use returned errors for expected problems. Reserve panic and recovery for exceptional situations and carefully chosen boundaries.

---

## 10. Defer Execution Order

When a function registers multiple deferred calls, Go executes them in **last-in, first-out (LIFO)** order.

This means the last deferred call runs first.

### Example

```go
package main

import "fmt"

func main() {
    defer fmt.Println("First defer")
    defer fmt.Println("Second defer")
    defer fmt.Println("Third defer")

    fmt.Println("Main function")
}
```

### Output

```text
Main function
Third defer
Second defer
First defer
```

### Why?

Think of deferred calls as plates stacked on top of one another.

- The first deferred call is at the bottom.
- The second is placed above it.
- The third is placed at the top.
- When execution finishes, the top plate comes off first.

This LIFO behavior also applies when a panic causes deferred functions to execute during stack unwinding.

---

## 11. Panic vs Error

Panic and error are not the same thing.

| Feature | `error` | `panic` |
|---|---|---|
| Purpose | Report a failure to the caller | Signal an exceptional condition |
| Typical use | Expected or recoverable problems | Serious unexpected states |
| Control flow | The function can return normally | Normal execution is interrupted |
| Handling | Check the returned error | Use recovery where appropriate |
| Example | File not found | An invalid internal program state |

### Recommended division example using error

```go
package main

import (
    "errors"
    "fmt"
)

func main() {
    answer, err := divide(10, 0)

    // Check whether the operation failed.
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // This executes only if division succeeds.
    fmt.Println("Answer:", answer)
}

func divide(a, b int) (int, error) {
    // Return a normal error for an expected failure.
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }

    // Return the answer and nil to indicate success.
    return a / b, nil
}
```

### Output

```text
Error: cannot divide by zero
```

Here, the function explicitly returns two values:

- An integer result.
- An `error` describing a failure, or `nil` if there is no error.

This makes the failure clear without using panic and recover.

---

## 12. Best Practices

1. **Prefer errors for expected failures.** Invalid user input, missing files, and failed network requests usually should return errors.
2. **Use defer for cleanup.** Schedule resource cleanup close to where the resource is acquired.
3. **Use recover deliberately.** Recovery is often appropriate at a boundary where a program needs to prevent one panicking operation from crashing a larger process.
4. **Do not silently ignore recovered panics.** Log or report the failure when appropriate, and avoid pretending the operation succeeded.
5. **Remember that recovery does not resume the interrupted function.** Design return values and cleanup logic with this behavior in mind.
6. **Do not use a zero result as the only failure signal.** Zero can be a valid result; use an error when you need to distinguish success from failure.
7. **Keep deferred functions understandable.** They execute later, so complicated cleanup and recovery logic can make code harder to follow.

---

## 13. Common Mistakes

### Mistake 1: Expecting execution to continue after panic

```go
panic("Problem")
fmt.Println("This will run")
```

The second statement is not executed unless the flow is structured differently. Recovering the panic does not make the same function resume from that line.

### Mistake 2: Calling recover outside a deferred function

```go
recoveredValue := recover()
```

If there is no active panic being recovered from in the correct deferred context, `recover()` returns `nil`.

### Mistake 3: Using panic for every error

```go
if err != nil {
    panic(err)
}
```

This may be appropriate in a few situations, but ordinary application failures are usually better handled by returning or propagating the error.

### Mistake 4: Assuming defer runs immediately

```go
defer fmt.Println("Later")
fmt.Println("Now")
```

Output:

```text
Now
Later
```

The deferred call executes when the surrounding function is finishing, not at the moment `defer` is encountered.

### Mistake 5: Assuming recovery catches panics from every goroutine

A recovery function can recover a panic only in its own goroutine's panic flow. If a goroutine needs its own recovery boundary, that goroutine must establish one.

---

## 14. Quick Revision

| Question | Answer |
|---|---|
| What does `panic()` do? | Interrupts normal execution and begins stack unwinding. |
| What does `defer` do? | Schedules a function call to execute as the surrounding function finishes. |
| What does `recover()` do? | Stops an active panic when called directly from a deferred function in the same goroutine. |
| Does recover resume after the panic line? | No. The panicking function finishes unwinding. |
| In what order do deferred calls run? | Last-in, first-out (LIFO). |
| Can one goroutine recover another goroutine's panic? | No. |
| Should every error cause a panic? | No. Expected failures should generally use returned errors. |
| What does an unnamed `int` return as its zero value? | `0`. |

### Final mental model

```text
defer   = Schedule work for when the function finishes.
panic   = Interrupt normal execution and begin unwinding.
recover = Stop an active panic from a deferred function.
error   = Report an ordinary failure to the caller.
```

**Remember:** `defer`, `panic`, and `recover` are useful Go language features, but they serve different purposes. Use `defer` for cleanup, errors for expected failures, and panic recovery only where it is intentionally needed.

---