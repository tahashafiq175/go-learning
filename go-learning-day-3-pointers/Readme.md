# Go Learning — Pointers in Go

## 📌 Table of Contents

1. [Introduction to Pointers](#1-introduction-to-pointers)
2. [Why Do We Need Pointers?](#2-why-do-we-need-pointers)
3. [Memory Addresses](#3-memory-addresses)
4. [The Address Operator (`&`)](#4-the-address-operator-)
5. [Pointer Variables](#5-pointer-variables)
6. [The Dereference Operator (`*`)](#6-the-dereference-operator-)
7. [Changing Values Through Pointers](#7-changing-values-through-pointers)
8. [Understanding `var ptr *int`](#8-understanding-var-ptr-int)
9. [Pointers and Structs](#9-pointers-and-structs)
10. [Accessing Struct Fields Through Pointers](#10-accessing-struct-fields-through-pointers)
11. [Pointers in Functions](#11-pointers-in-functions)
12. [Go Passes Arguments by Value](#12-go-passes-arguments-by-value)
13. [Returning Values and Pointers](#13-returning-values-and-pointers)
14. [The `nil` Pointer](#14-the-nil-pointer)
15. [Pointer vs Value](#15-pointer-vs-value)
16. [Common Mistakes](#16-common-mistakes)
17. [Complete Practical Example](#17-complete-practical-example)
18. [Quick Revision Table](#18-quick-revision-table)
19. [Practice Exercises](#19-practice-exercises)

---

## 1. Introduction to Pointers

A **pointer is a variable that stores the memory address of another variable**.

Before understanding pointers, let's understand a normal variable.

```go
package main

import "fmt"

func main() {
    age := 10

    fmt.Println(age)
}
```

Output:

```text
10
```

Here, `age` is a variable that stores the value `10`.

Imagine that computer memory is a collection of numbered boxes. Each box can store data, and each box has an address.

```text
Computer Memory

Address        Stored Value
+---------+    +-------------+
| 0x100   | -> | age = 10    |
+---------+    +-------------+
```

The address `0x100` is just an example. Actual memory addresses vary between executions.

Now imagine that we want another variable to store the address of `age`, rather than storing the number `10` itself.

That is where pointers come in.

```go
age := 10
ptr := &age
```

Now:

- `age` stores the value `10`.
- `&age` gives us the memory address of `age`.
- `ptr` stores that address.

Conceptually:

```text
age
 |
 v
+----------------+
| Value: 10      |
| Address: 0x100 |
+----------------+
        ^
        |
        | ptr stores this address
        |
      ptr
```

**Remember:** A normal variable stores a value. A pointer stores an address.

---

## 2. Why Do We Need Pointers?

Imagine you have a box containing the number `10`.

You give someone a copy of the number `10`. If they change their copy to `20`, your original number remains `10`.

```go
package main

import "fmt"

func main() {
    age := 10
    copiedAge := age

    copiedAge = 20

    fmt.Println(age)
    fmt.Println(copiedAge)
}
```

Output:

```text
10
20
```

Why?

Because `copiedAge := age` copies the value. The two variables are independent.

But sometimes, we want to modify the original value through another variable or pass access to that value into a function.

Pointers help us do that.

```go
package main

import "fmt"

func main() {
    age := 10
    ptr := &age

    *ptr = 20

    fmt.Println(age)
}
```

Output:

```text
20
```

The original `age` changed because `ptr` points to its memory location.

### Real-world analogy

Think of a house:

- The **house** represents the original variable.
- The **house address** represents its memory address.
- A **piece of paper containing the address** represents a pointer.

If someone has the correct address, they can locate the house. Similarly, a pointer lets your program access the value stored at a particular memory address.

### Important uses of pointers

- Modifying original values through functions.
- Avoiding unnecessary copying of large data structures.
- Sharing access to a struct between different parts of a program.
- Representing the absence of a value using a `nil` pointer, where appropriate.

Pointers are useful, but they are not always necessary. Go automatically manages memory allocation and garbage collection.

---

## 3. Memory Addresses

Every variable stored in memory has a location. In Go, the address operator lets us obtain a variable's address.

```go
package main

import "fmt"

func main() {
    age := 10

    fmt.Println(age)
    fmt.Println(&age)
}
```

Example output:

```text
10
0xc000012090
```

Your actual address will probably be different.

Let's understand each line:

```go
age := 10
```

Creates a variable named `age` and stores `10`.

```go
fmt.Println(age)
```

Prints the value inside the variable.

```go
fmt.Println(&age)
```

Prints the memory address of the variable.

### Value vs address

| Expression | Meaning |
|---|---|
| `age` | The value stored in `age` |
| `&age` | The memory address of `age` |

For example:

```go
fmt.Println(age)  // 10
fmt.Println(&age) // An address, such as 0xc000012090
```

**Remember:** `age` and `&age` are not the same thing. One gives you the value; the other gives you the address.

---

## 4. The Address Operator (`&`)

The `&` operator obtains the memory address of a variable.

Syntax:

```go
&variableName
```

Example:

```go
package main

import "fmt"

func main() {
    number := 50

    address := &number

    fmt.Println(number)
    fmt.Println(address)
}
```

Example output:

```text
50
0xc000012090
```

Explanation:

```go
number := 50
```

Creates an integer variable.

```go
address := &number
```

Obtains the address of `number` and stores it in `address`.

Because Go can infer the type, `address` automatically becomes a pointer to an integer, or `*int`.

Conceptually:

```text
number = 50
    |
    | &number
    v
  Address
    |
    v
address stores the address of number
```

### Important distinction

The `&` operator does not change the value of the variable. It obtains the address where the value is stored.

---

## 5. Pointer Variables

A pointer variable stores the address of a value.

Let's examine this example:

```go
package main

import "fmt"

func main() {
    age := 10

    var ptr *int = &age

    fmt.Println(age)
    fmt.Println(ptr)
}
```

Example output:

```text
10
0xc000012090
```

Now let's understand the important line:

```go
var ptr *int = &age
```

Break it into three parts:

| Part | Meaning |
|---|---|
| `var` | Declares a variable |
| `ptr *int` | Declares `ptr` as a pointer to an integer |
| `&age` | Supplies the address of `age` |

The type `*int` means **pointer to an integer**.

It does not mean that `ptr` directly stores an integer value. It stores an address that points to an integer value.

### Shorter syntax

Go can infer the pointer type:

```go
age := 10
ptr := &age
```

Both versions create a pointer to `age`.

### Common pointer types

| Pointer type | Points to |
|---|---|
| `*int` | An integer |
| `*string` | A string |
| `*bool` | A Boolean |
| `*float64` | A `float64` value |
| `*User` | A `User` struct |

For example:

```go
name := "Taha"
namePtr := &name

fmt.Println(namePtr)
```

Here, `namePtr` has type `*string`.

---

## 6. The Dereference Operator (`*`)

We have learned that a pointer stores an address.

But how do we access the value located at that address?

We use the dereference operator: `*`.

Example:

```go
package main

import "fmt"

func main() {
    age := 10
    ptr := &age

    fmt.Println(ptr)
    fmt.Println(*ptr)
}
```

Example output:

```text
0xc000012090
10
```

Let's understand the difference:

```go
fmt.Println(ptr)
```

Prints the address stored in the pointer.

```go
fmt.Println(*ptr)
```

Follows the address and retrieves the value stored there.

### Visual explanation

```text
age = 10
Address = 0x100

       ptr
        |
        | stores 0x100
        v
      0x100
        |
        v
   +----------+
   |    10    |
   +----------+

ptr  -> prints the address
*ptr -> accesses the value 10
```

The address `0x100` is illustrative.

### A very important difference

The symbol `*` has two related but different uses in Go.

**1. In a type declaration:**

```go
var ptr *int
```

Here, `*int` means a pointer to an integer.

**2. When using a pointer:**

```go
fmt.Println(*ptr)
```

Here, `*ptr` means to dereference the pointer and access its value.

The meaning depends on where the operator appears.

---

## 7. Changing Values Through Pointers

A pointer can be used to modify the original value.

Consider this example:

```go
package main

import "fmt"

func main() {
    age := 10
    ptr := &age

    *ptr = 20

    fmt.Println(age)
    fmt.Println(*ptr)
}
```

Output:

```text
20
20
```

Why did `age` change?

Let's trace the execution.

### Step 1: Create the variable

```go
age := 10
```

Memory contains:

```text
age = 10
```

### Step 2: Create a pointer

```go
ptr := &age
```

The pointer stores the address of `age`.

```text
ptr --------> age = 10
```

### Step 3: Modify the value through the pointer

```go
*ptr = 20
```

This means:

"Go to the memory address stored in `ptr` and replace the value there with `20`."

Therefore:

```text
Before:

ptr --------> age = 10

After:

ptr --------> age = 20
```

### Step 4: Print the original variable

```go
fmt.Println(age)
```

Output:

```text
20
```

The original variable has changed because we modified its stored value through the pointer.

**Remember:**

- `ptr = ...` changes what address the pointer stores, if the assignment is type-correct.
- `*ptr = ...` changes the value at the address stored in the pointer.

These are different operations.

---

## 8. Understanding `var ptr *int`

Consider this code:

```go
var ptr *int
```

A beginner might think that `ptr` already contains an integer.

It does not.

This statement declares a pointer variable without assigning it an address. Its initial value is `nil`.

```go
package main

import "fmt"

func main() {
    var ptr *int

    fmt.Println(ptr)
}
```

Output:

```text
<nil>
```

The pointer does not currently point to a valid integer value.

If we want it to point to a variable:

```go
package main

import "fmt"

func main() {
    age := 10

    var ptr *int = &age

    fmt.Println(*ptr)
}
```

Output:

```text
10
```

Now `ptr` points to `age`, so dereferencing it is valid.

**Important:** Never dereference a `nil` pointer. Doing so causes a runtime panic.

---

## 9. Pointers and Structs

A struct groups related values into a single data structure.

For example:

```go
type User struct {
    Name string
    Age  int
}
```

This defines a new type named `User` with two fields:

- `Name` stores a string.
- `Age` stores an integer.

We can create a value of this type:

```go
user := User{
    Name: "Taha",
    Age:  13,
}
```

The `user` variable contains a `User` struct.

To create a pointer to that struct:

```go
usrptr := &user
```

Now `usrptr` stores the address of `user`.

Complete example:

```go
package main

import "fmt"

type User struct {
    Name string
    Age  int
}

func main() {
    user := User{
        Name: "Taha",
        Age:  13,
    }

    usrptr := &user

    fmt.Println(user)
    fmt.Println(usrptr)
    fmt.Println(*usrptr)
}
```

Example output:

```text
{Taha 13}
&{Taha 13}
{Taha 13}
```

Let's understand each expression:

| Expression | Meaning |
|---|---|
| `user` | The struct value |
| `usrptr` | A pointer to the struct |
| `*usrptr` | The struct value accessed through the pointer |

Go's default formatting explains the output:

- A struct value appears like `{Taha 13}`.
- A pointer to that struct appears like `&{Taha 13}`.
- Dereferencing the pointer gives the struct value again.

The exact memory address is not important for understanding the concept.

---

## 10. Accessing Struct Fields Through Pointers

Look at this example:

```go
package main

import "fmt"

type User struct {
    Name string
    Age  int
}

func main() {
    user := User{
        Name: "Taha",
        Age:  13,
    }

    usrptr := &user

    fmt.Println(usrptr.Name)
}
```

Output:

```text
Taha
```

You might wonder:

"Why did we write `usrptr.Name` instead of `(*usrptr).Name`?"

Go automatically handles this common case for you.

These expressions access the same field:

```go
usrptr.Name
```

```go
(*usrptr).Name
```

The second form explicitly dereferences the pointer before accessing the field. The first form is the simpler, idiomatic way to write it in Go.

### Changing struct fields through a pointer

```go
package main

import "fmt"

type User struct {
    Name string
    Age  int
}

func main() {
    user := User{
        Name: "Taha",
        Age:  13,
    }

    usrptr := &user

    usrptr.Name = "Usama"
    usrptr.Age = 20

    fmt.Println(user.Name)
    fmt.Println(user.Age)
}
```

Output:

```text
Usama
20
```

Why did the original `user` change?

Because `usrptr` points to the original struct. Assigning to `usrptr.Name` or `usrptr.Age` changes the corresponding fields of that struct.

**Remember:** Go automatically dereferences a pointer when you access a struct field with the dot operator.

---

## 11. Pointers in Functions

One of the most important uses of pointers is allowing a function to modify a value owned by the caller.

Let's begin without a pointer.

```go
package main

import "fmt"

func changeAge(age int) {
    age = 20
}

func main() {
    age := 10

    changeAge(age)

    fmt.Println(age)
}
```

Output:

```text
10
```

Why is the output still `10`?

When we call:

```go
changeAge(age)
```

Go passes a copy of the integer value into the function.

Inside `changeAge`, the parameter `age` is a separate variable. Changing it does not change the original variable in `main`.

### Solution: Pass a pointer

```go
package main

import "fmt"

func changeAge(age *int) {
    *age = 20
}

func main() {
    age := 10

    changeAge(&age)

    fmt.Println(age)
}
```

Output:

```text
20
```

Let's break it down.

**Step 1: The function parameter**

```go
func changeAge(age *int)
```

The function receives a pointer to an integer, not an integer value.

**Step 2: The function call**

```go
changeAge(&age)
```

We pass the address of the original variable.

**Step 3: Modify the original value**

```go
*age = 20
```

The parameter contains the address of the original integer. Dereferencing it allows the function to change the original value.

### Execution flow

```text
main()
  |
  | age := 10
  |
  | changeAge(&age)
  v
changeAge(age *int)
  |
  | *age = 20
  v
Original variable becomes 20
  |
  v
main prints 20
```

This pattern is useful when a function needs to update the caller's data.

---

## 12. Go Passes Arguments by Value

This concept is extremely important for interviews.

**Go always passes arguments by value.** This includes pointer arguments.

Let's understand what that means.

```go
func change(user *User) {
    user.Name = "Usama"
}
```

When we call:

```go
change(usrptr)
```

Go copies the pointer value into the function parameter named `user`.

The copied pointer and the original pointer point to the same struct.

```text
Original pointer usrptr
          |
          v
      +--------+
      |  User  |
      +--------+
          ^
          |
Copied pointer user
inside change()
```

The pointer itself is copied, but both pointers refer to the same underlying struct.

Therefore:

```go
user.Name = "Usama"
```

changes the original struct.

However, if we reassign the local pointer itself:

```go
func changePointer(user *User) {
    anotherUser := User{
        Name: "Ali",
        Age:  25,
    }

    user = &anotherUser
}
```

This does not change the caller's pointer. Only the local parameter now points to `anotherUser`.

The difference is:

- Modifying `user.Name` modifies the shared struct.
- Assigning `user = &anotherUser` changes only the local pointer variable.

**Interview tip:** Go is pass-by-value, not pass-by-reference. A pointer is a value that contains an address.

---

## 13. Returning Values and Pointers

A function can return a struct value or a pointer to a struct. These are different return types.

### Example A: Return a struct value

```go
package main

import "fmt"

type User struct {
    Name string
    Age  int
}

func getUser() User {
    return User{
        Name: "Taha",
        Age:  20,
    }
}

func main() {
    user := getUser()

    fmt.Println(user)
}
```

Output:

```text
{Taha 20}
```

The return type is `User`, so the function returns a struct value.

### Example B: Return a pointer to a struct

```go
package main

import "fmt"

type User struct {
    Name string
    Age  int
}

func getUser() *User {
    return &User{
        Name: "Taha",
        Age:  20,
    }
}

func main() {
    user := getUser()

    fmt.Println(user.Name)
}
```

Output:

```text
Taha
```

Here:

```go
func getUser() *User
```

means that the function returns a pointer to a `User`.

Go permits returning a pointer to a local variable or a local struct literal when it is safe to do so. The compiler ensures that the memory remains valid for as long as it is needed.

### Understanding `return *user`

Consider:

```go
func change(user *User) User {
    user.Name = "Usama"
    user.Age = 20

    return *user
}
```

The return type is `User`, not `*User`.

Therefore:

```go
return *user
```

means:

1. Dereference the pointer.
2. Obtain the struct value.
3. Return that struct value to the caller.

Compare the two:

| Code | Meaning |
|---|---|
| `return user` | Returns the pointer, if the function's return type is `*User` |
| `return *user` | Returns the struct value, if the return type is `User` |

The return statement must match the function's declared return type.

---

## 14. The `nil` Pointer

In Go, `nil` represents the zero value of a pointer.

A pointer that has not been initialized with a valid address has the value `nil`.

```go
package main

import "fmt"

func main() {
    var ptr *int

    if ptr == nil {
        fmt.Println("Pointer is nil")
    }
}
```

Output:

```text
Pointer is nil
```

Why is checking `nil` useful?

Imagine a function that sometimes has a user and sometimes does not.

```go
package main

import "fmt"

type User struct {
    Name string
}

func main() {
    var user *User

    if user == nil {
        fmt.Println("No user available")
        return
    }

    fmt.Println(user.Name)
}
```

Output:

```text
No user available
```

The program checks the pointer before accessing the struct field.

### What happens if we dereference `nil`?

This is unsafe:

```go
var ptr *int

fmt.Println(*ptr)
```

It causes a runtime panic because `ptr` does not point to a valid integer.

**Rule:** Check whether a pointer is `nil` before dereferencing it when it might not have a valid value.

---

## 15. Pointer vs Value

Let's compare ordinary variables and pointers.

| Feature | Value | Pointer |
|---|---|---|
| Example | `age := 10` | `ptr := &age` |
| Stores | The value itself | An address |
| Access | `age` | `*ptr` |
| Obtain address | `&age` | `&ptr` gives the pointer variable's own address |
| Change original through it | Directly changes `age` | `*ptr = 20` changes the pointed-to value |
| Can be `nil`? | Not for ordinary `int`, `string`, or `bool` values | Yes |
| Type example | `int` | `*int` |

### A simple comparison

```go
age := 10
copiedAge := age

copiedAge = 20

fmt.Println(age) // 10
```

Here, `copiedAge` is a separate integer variable.

Now compare:

```go
age := 10
ptr := &age

*ptr = 20

fmt.Println(age) // 20
```

Here, `ptr` provides access to the original variable.

**Important:** Pointers are not automatically better than values. Use them when sharing or modifying data through an address is useful.

---

## 16. Common Mistakes

### Mistake 1: Printing the pointer instead of the value

```go
age := 10
ptr := &age

fmt.Println(ptr)
```

This prints an address.

To print the value:

```go
fmt.Println(*ptr)
```

### Mistake 2: Dereferencing a `nil` pointer

Incorrect:

```go
var ptr *int
fmt.Println(*ptr)
```

This causes a runtime panic.

Correct:

```go
var ptr *int

if ptr != nil {
    fmt.Println(*ptr)
}
```

### Mistake 3: Confusing `*int` with `int`

```go
var age int = 10
var ptr *int = &age
```

- `age` has type `int`.
- `ptr` has type `*int`.

They are different types.

### Mistake 4: Expecting a normal function parameter to change the caller's integer

Incorrect expectation:

```go
func change(age int) {
    age = 50
}
```

The function changes its local copy, not the caller's variable.

Correct when the caller's value must be modified:

```go
func change(age *int) {
    *age = 50
}
```

Call it with:

```go
age := 10
change(&age)
fmt.Println(age) // 50
```

### Mistake 5: Thinking `return *user` returns a pointer

```go
func change(user *User) User {
    return *user
}
```

The function returns a `User` struct value because its return type is `User`.

To return a pointer, the function must declare `*User` as its return type.

### Mistake 6: Thinking a pointer parameter breaks pass-by-value

Go still copies the pointer value when calling the function. The copied pointer simply points to the same underlying data.

### Mistake 7: Thinking `usrptr.Name` is invalid

In Go, this is valid:

```go
usrptr.Name
```

Go automatically handles the pointer dereference for struct field access. Writing `(*usrptr).Name` is also valid.

---

## 17. Complete Practical Example

This example combines pointers, structs, functions, return values, and formatted strings.

```go
package main

import "fmt"

// User represents a person.
type User struct {
    Name string
    Age  int
}

// change updates the original User through a pointer.
// It returns a copy of the updated User struct.
func change(user *User) User {
    user.Name = "Usama"
    user.Age = 20

    return *user
}

func main() {
    // Create an ordinary integer variable.
    age := 10

    // Store the address of age in ptr.
    ptr := &age

    fmt.Println("Original address:", ptr)
    fmt.Println("Original value:", *ptr)

    // Update the original integer through the pointer.
    *ptr = 20

    fmt.Println("Updated age:", age)

    // Create a User struct.
    user := User{
        Name: "Taha",
        Age:  13,
    }

    // Create a pointer to the User struct.
    usrptr := &user

    fmt.Println("Original user:", user)
    fmt.Println("User through pointer:", *usrptr)
    fmt.Println("Name through pointer:", usrptr.Name)

    // Change the original User through a pointer parameter.
    updatedUser := change(usrptr)

    fmt.Println("Updated user:", user)
    fmt.Println("Returned user:", updatedUser)

    // Build a formatted string without printing it immediately.
    details := fmt.Sprintf(
        "The user has a name of %s and age of %d",
        user.Name,
        user.Age,
    )

    fmt.Println(details)
}
```

Example output (the memory address will vary):

```text
Original address: 0xc000012090
Original value: 10
Updated age: 20
Original user: {Taha 13}
User through pointer: {Taha 13}
Name through pointer: Taha
Updated user: {Usama 20}
Returned user: {Usama 20}
The user has a name of Usama and age of 20
```

### Understanding the important lines

**1. Creating the pointer**

```go
ptr := &age
```

The pointer stores the address of `age`.

**2. Changing the original integer**

```go
*ptr = 20
```

The value at that address changes to `20`.

**3. Creating a struct pointer**

```go
usrptr := &user
```

The pointer refers to the original `User` struct.

**4. Passing the pointer to a function**

```go
updatedUser := change(usrptr)
```

The function receives a copy of the pointer value, but that pointer refers to the same struct.

**5. Updating the original struct**

```go
user.Name = "Usama"
user.Age = 20
```

Inside `change`, these field assignments modify the original struct through the pointer.

**6. Returning the struct**

```go
return *user
```

The function returns the updated struct value, not the pointer itself.

**7. Formatting the result**

```go
details := fmt.Sprintf(
    "The user has a name of %s and age of %d",
    user.Name,
    user.Age,
)
```

`fmt.Sprintf` builds a string and returns it. It does not print the string by itself.

- `%s` formats a string.
- `%d` formats an integer.

The resulting string is printed later with `fmt.Println(details)`.

---

## 18. Quick Revision Table

| Concept | Syntax | Meaning |
|---|---|---|
| Normal variable | `age := 10` | Stores the integer value |
| Address operator | `&age` | Gets the address of `age` |
| Pointer declaration | `var ptr *int` | Declares a pointer to an integer |
| Pointer initialization | `ptr := &age` | Stores the address of `age` |
| Dereferencing | `*ptr` | Reads the value at the stored address |
| Modifying through pointer | `*ptr = 20` | Updates the pointed-to value |
| Struct pointer | `ptr := &user` | Points to a `User` struct |
| Field access | `ptr.Name` | Accesses the pointed-to struct's field |
| Nil check | `ptr == nil` | Checks whether the pointer is nil |
| Return a struct | `return *ptr` | Returns the struct value |
| Return a pointer | `return ptr` | Returns a pointer when the return type matches |

### Three rules to remember

1. `&` means **get the address**.
2. `*` means **pointer type** in declarations and **access the pointed-to value** when dereferencing.
3. Passing a pointer lets a function access the pointed-to data, but Go still passes the pointer value by value.

---

## 19. Practice Exercises

Try solving these exercises without looking at the solutions immediately.

### Exercise 1: Print the address

Create an integer variable named `number` with value `100`. Print its value and its address.

Expected result:

```text
100
An address will be printed
```

### Exercise 2: Modify through a pointer

Create `score := 50`. Create a pointer to `score`, then use the pointer to change `score` to `90`.

Expected output:

```text
90
```

### Exercise 3: Pointer to a string

Create a string variable:

```go
name := "Taha"
```

Create a pointer to it and print the string using the pointer.

Expected output:

```text
Taha
```

### Exercise 4: Update a struct

Create a `Student` struct with `Name` and `Marks` fields. Create a pointer to a student, then change the student's marks through the pointer.

Expected result: The original struct should contain the updated marks.

### Exercise 5: Pointer function

Write a function named `increase` that accepts a pointer to an integer and increases the original number by `1`.

Example:

```go
number := 5
increase(&number)

fmt.Println(number)
```

Expected output:

```text
6
```

### Exercise 6: Understand the output

Predict the output before running this program:

```go
package main

import "fmt"

func update(ptr *int) {
    *ptr = 100
}

func main() {
    number := 10

    update(&number)

    fmt.Println(number)
}
```

**Question:** Why does the output show `100` instead of `10`?

### Exercise 7: Struct and return value

Write a `User` struct with `Name` and `Age` fields. Write a function that receives a `*User`, updates the user's age, and returns the updated `User` value.

Try to explain the difference between the function's pointer parameter and its returned struct value.

---

## Final Summary

A pointer is a value that stores the address of another value.

For example:

```go
age := 10
ptr := &age

fmt.Println(*ptr) // 10

*ptr = 20

fmt.Println(age) // 20
```

You now know the foundation of pointers in Go:

- How variables and memory addresses relate.
- How to obtain an address with `&`.
- How to declare and use pointers.
- How to read and modify values through pointers.
- How pointers work with structs and functions.
- Why Go still passes arguments by value.
- How to distinguish a struct value from a pointer to a struct.
- Why you should check potentially nil pointers before dereferencing them.

**The most important idea:** A pointer stores an address. Dereferencing it lets you access the value at that address.

Once this idea becomes clear, pointers in Go become much easier to understand.
