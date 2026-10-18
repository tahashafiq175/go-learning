# Go Learning — Arrays and Slices

## 📌 Overview

This module introduces **Arrays and Slices in Go**, two ways to store and work with multiple values of the same type.

You will learn how arrays differ from slices, how slices grow, how `len()` and `cap()` work, how to extract parts of a collection, and how to store structs inside slices.

The goal is to understand not only the syntax but also **how Go manages the data behind a slice**.

## 🎯 Learning Objectives

After completing this module, you should be able to:

- Understand arrays and their fixed size.
- Create and initialize slices.
- Use `append()` to add elements.
- Understand `len()` and `cap()`.
- Access and modify slice elements.
- Extract a smaller slice from an array or slice.
- Understand how slices share an underlying array.
- Iterate over collections using `range`.
- Create and work with slices of structs.
- Avoid common mistakes when using arrays and slices.

## 📂 Project Structure

```text
arrays-and-slices/
├── main.go
└── README.md
```

> Replace `main.go` with your actual Go filename if your file has a different name.

## 1. Arrays in Go

An array stores a fixed number of elements of the same type.

```go
numbers := [5]int{1, 2, 3, 4, 5}
```

Here:

- `[5]` means the array contains exactly five elements.
- `int` means every element is an integer.
- `{1, 2, 3, 4, 5}` provides the initial values.

The array's length is part of its type. An array of five integers and an array of six integers are different types.

```go
fmt.Println(numbers[0]) // First element
fmt.Println(len(numbers)) // Number of elements
```

**Important:** An array cannot grow or shrink after it is created. You cannot use `append()` to increase the size of an array.

## 2. Slices in Go

A slice is a flexible view of elements stored in an underlying array.

```go
numbers := []int{1, 2, 3, 4, 5}
```

Notice the difference:

```go
[5]int{} // Array with a fixed length of 5
[]int{}  // Slice literal with no explicitly specified length
```

Slices are commonly used when the number of elements may change while a program runs.

For example:

```go
numbers = append(numbers, 6)
numbers = append(numbers, 7, 8, 9)
```

The `append()` function adds elements to a slice. Always use the slice returned by `append()`, because Go may need to create a larger underlying array.

## 3. Length and Capacity

Two important built-in functions are `len()` and `cap()`.

| Function | Meaning |
|---|---|
| `len(slice)` | Number of elements currently in the slice |
| `cap(slice)` | Number of elements available from the slice's starting position to the end of its underlying array |

Example:

```go
numbers := make([]int, 3, 5)

fmt.Println(len(numbers)) // 3
fmt.Println(cap(numbers)) // 5
```

This creates a slice with:

- **Length:** 3 elements.
- **Capacity:** space for up to 5 elements before the underlying array needs to grow.

The slice initially contains three zero-valued integers.

Remember: capacity is not the same as length. You cannot access an element using an index greater than or equal to `len(slice)`, even if that index is within the capacity.

## 4. Creating Slices with `make()`

The `make()` function can create a slice with a specified length and capacity.

```go
numbers := make([]int, 5, 10)
```

This means:

- `[]int`: create a slice of integers.
- `5`: initialize the slice with a length of five.
- `10`: give the slice a capacity of ten.

All five existing elements initially contain `0`.

```go
numbers[0] = 100 // Valid
numbers[4] = 500 // Valid
// numbers[5] = 600 // Invalid: index 5 is outside the current length
```

To add another element, use `append()`:

```go
numbers = append(numbers, 600)
```

## 5. Accessing and Modifying Elements

Slice indexing starts at zero.

```go
numbers := []int{10, 20, 30}

fmt.Println(numbers[0]) // 10
fmt.Println(numbers[1]) // 20

numbers[0] = 100
fmt.Println(numbers) // [100 20 30]
```

Changing an element through a slice changes the corresponding element in its underlying array.

An index outside the valid range causes a runtime panic.

## 6. Slicing — Extracting a Portion

You can create a smaller slice from an array or another slice using the slicing expression:

```go
collection[start:end]
```

The `start` index is included, but the `end` index is excluded.

Example:

```go
numbers := [5]int{10, 20, 30, 40, 50}

part := numbers[1:4]

fmt.Println(part) // [20 30 40]
fmt.Println(len(part)) // 3
fmt.Println(cap(part)) // 4
```

Why is the capacity `4`?

The new slice starts at array index `1`. From that position to the end of the original array, there are four elements: `20`, `30`, `40`, and `50`.

The slice has a length of three because the expression ends before index `4`.

## 7. Understanding the Underlying Array

A slice does not necessarily own a separate copy of its elements. It refers to an underlying array.

```go
numbers := [5]int{10, 20, 30, 40, 50}

part := numbers[1:4]
part[0] = 200

fmt.Println(numbers) // [10 200 30 40 50]
fmt.Println(part)    // [200 30 40]
```

Changing `part[0]` also changes `numbers[1]` because both refer to the same element in the underlying array.

**Key lesson:** Creating a slice from an array does not automatically copy the selected elements.

When you need an independent copy, use `copy()` or another explicit copying method.

## 8. Iterating with `range`

The `range` keyword lets you iterate over the elements of an array or slice.

```go
numbers := []int{10, 20, 30}

for index, value := range numbers {
    fmt.Println(index, value)
}
```

The loop provides:

- `index`: the position of the current element.
- `value`: a copy of the current element's value.

If you do not need the index, use the blank identifier `_`:

```go
for _, value := range numbers {
    fmt.Println(value)
}
```

The blank identifier tells Go that the value is intentionally unused.

## 9. Slices of Structs

Slices can store structs, allowing you to keep multiple records together.

```go
type User struct {
    Name string
    Age  int
}

users := []User{
    {Name: "Ali", Age: 20},
    {Name: "Sara", Age: 22},
}
```

Here:

- `User` defines a custom data type.
- Each `User` value contains a name and an age.
- `[]User` creates a slice that can hold multiple `User` values.

You can iterate over the users:

```go
for _, user := range users {
    fmt.Println(user.Name, user.Age)
}
```

This pattern is useful for storing users, products, orders, employees, and other collections of structured data.

## 10. Common Mistakes

### Mistake 1: Confusing arrays and slices

```go
array := [3]int{1, 2, 3}
slice := []int{1, 2, 3}
```

The array has a fixed length. The slice can be extended using `append()`.

### Mistake 2: Confusing length with capacity

```go
numbers := make([]int, 3, 10)
```

Although the capacity is ten, only indexes `0`, `1`, and `2` are currently valid.

### Mistake 3: Ignoring the result of `append()`

```go
numbers = append(numbers, 4)
```

Assign the result back to the slice variable because `append()` may return a slice referring to a new underlying array.

### Mistake 4: Assuming a subslice is an independent copy

```go
part := numbers[1:3]
```

The subslice may share the original underlying array. Changes to shared elements can be visible through both slices.

### Mistake 5: Forgetting that the end index is excluded

```go
numbers[1:4]
```

This selects indexes `1`, `2`, and `3`, not index `4`.

## 🧠 Final Takeaway

- **Array:** A fixed-size collection of elements of the same type.
- **Slice:** A flexible descriptor that refers to a section of an underlying array.
- **`append()`:** Adds elements and may allocate a new underlying array.
- **`len()`:** Reports the number of elements currently in the slice.
- **`cap()`:** Reports the capacity available from the slice's starting position.
- **`make()`:** Creates a slice with a chosen length and capacity.
- **Slicing:** Creates a view over part of an array or slice.
- **`range`:** Iterates over elements.
- **Slice of structs:** Stores multiple structured records in one collection.

Understanding the relationship between a slice, its length, its capacity, and its underlying array is the foundation for using slices correctly in Go.

## 📚 Practice Tasks

1. Create an array of five integers and print its length.
2. Create a slice and append three additional elements.
3. Create a slice with `make()` and compare its length with its capacity.
4. Extract a subslice and modify one of its elements. Observe whether the original collection changes.
5. Create a slice of three `User` structs and print each user's details.
6. Experiment with `copy()` to create an independent slice.

**Learning goal:** Be able to explain why slices are more flexible than arrays and how Go manages the data behind them.