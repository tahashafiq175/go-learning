# Go Learning — Slices, Copying & Underlying Arrays

A practical, beginner-friendly guide to understanding slices in Go, including slice creation, copying elements, length, capacity, appending, subslicing, and shared underlying arrays.

## 📌 Overview

In Go, a **slice** is a flexible way to work with a sequence of elements. Unlike an array, whose size is fixed, a slice can grow when new elements are appended.

This module explores how slices work internally and how Go manages their data when we copy, append, or extract elements.

By running the examples in this project, you will understand why two slices can contain the same values but use different underlying arrays, and why modifying one slice can sometimes affect another.

## 🎯 Learning Objectives

By completing this module, you will learn how to:

- Create slices using slice literals and `make()`.
- Understand the difference between length and capacity.
- Copy elements using the built-in `copy()` function.
- Understand what happens when the destination slice is smaller than the source.
- Add elements using `append()`.
- Understand how Go handles capacity when a slice grows.
- Extract a subslice using slicing expressions.
- Understand how multiple slices can share an underlying array.
- Predict the output of slice operations.

## 📂 Project Structure

```text
slices-learning/
├── main.go
└── README.md
```

`main.go` contains the practical examples, while this README explains the concepts demonstrated in the code.

## 1. Creating a Slice

A slice literal lets us create a slice and initialize it with values.

```go
numbers := []int{1, 2, 3, 4, 5}
```

- `[]int` means a slice of integers.
- `{1, 2, 3, 4, 5}` provides the initial values.
- `len(numbers)` returns `5`.
- The slice refers to an underlying array that stores its elements.

Slices are commonly used when working with collections whose size may change.

## 2. Creating Slices with `make()`

The `make()` function creates a slice with a specified length and, optionally, a capacity.

```go
number := make([]int, 5)
```

This creates a slice with:

| Property | Value |
|---|---:|
| Type | `[]int` |
| Length | 5 |
| Capacity | 5 |
| Initial values | `[0 0 0 0 0]` |

We can also specify capacity separately:

```go
newSlice := make([]int, 5, 10)
```

| Property | Value |
|---|---:|
| Length | 5 |
| Capacity | 10 |
| Initial values | `[0 0 0 0 0]` |

The slice has five existing elements and room to grow up to ten elements before it needs a larger underlying array.

**Important:** Capacity does not mean that all ten elements are immediately accessible by indexing. Only indexes `0` through `4` are valid until the slice's length increases.

## 3. Understanding `len()` and `cap()`

Go provides two built-in functions for inspecting a slice.

| Function | Purpose |
|---|---|
| `len(slice)` | Returns the number of elements currently in the slice. |
| `cap(slice)` | Returns the capacity available from the slice's starting position. |

Example:

```go
numbers := make([]int, 5, 10)

fmt.Println(len(numbers)) // 5
fmt.Println(cap(numbers)) // 10
```

Think of length as the number of seats currently available for passengers, while capacity represents how many seats the underlying storage can accommodate before needing expansion.

## 4. Copying Elements with `copy()`

The built-in `copy()` function copies elements from a source slice into a destination slice.

```go
numbers := []int{1, 2, 3, 4, 5}

number := make([]int, 5)
copy(number, numbers)
```

Here:

- `number` is the destination slice.
- `numbers` is the source slice.
- Both slices have five elements, so all five values are copied.

The result is:

```text
numbers = [1 2 3 4 5]
number  = [1 2 3 4 5]
```

The important point is that `copy()` copies the elements into the destination's existing storage. It does not make the two slices share the same underlying array.

### What if the destination is smaller?

```go
number2 := make([]int, 4)
copy(number2, numbers)

fmt.Println(number2) // [1 2 3 4]
```

The destination can hold only four elements, so only the first four values are copied.

The number of elements copied is the smaller of the source and destination lengths.

## 5. Adding Elements with `append()`

The `append()` function adds elements to the end of a slice.

```go
dummynumber := make([]int, 5, 5)

dummynumber[0] = 1
dummynumber[1] = 2
dummynumber[2] = 3
dummynumber[3] = 4
dummynumber[4] = 5

dummynumber = append(dummynumber, 6)
```

Before appending:

```text
Length   = 5
Capacity = 5
Values   = [1 2 3 4 5]
```

After appending `6`:

```text
Length   = 6
Values   = [1 2 3 4 5 6]
```

Because the original capacity was insufficient, Go must make room for the additional element. It may allocate a larger underlying array and copy the existing elements into it.

The new capacity is determined by Go's allocation strategy. **Do not assume that it will always become 10.**

Always assign the result of `append()` back to the slice variable:

```go
dummynumber = append(dummynumber, 6)
```

This matters because `append()` returns the updated slice, which may refer to a different underlying array.

## 6. Slicing — Extracting Part of a Slice

A slicing expression lets us create a new slice that refers to part of an existing array or slice.

```go
part := dummynumber[1:4]
```

The general syntax is:

```go
slice[start:end]
```

- `start` is included.
- `end` is excluded.

For example, if `dummynumber` contains:

```text
[1 2 3 4 5 6]
```

Then:

```go
part := dummynumber[1:4]
```

Produces:

```text
part = [2 3 4]
```

The selected indexes are `1`, `2`, and `3`.

The length of `part` is `3`. Its capacity extends from the selected starting position to the end of the underlying array.

## 7. Understanding the Underlying Array

A slice is not simply an independent collection of values. It is a descriptor that refers to an underlying array and records the slice's length and capacity.

When we create a subslice, Go does not automatically copy the selected elements into a new array.

```go
part := dummynumber[1:4]

part[0] = 10

fmt.Println(part)
fmt.Println(dummynumber)
```

Suppose the original slice contains:

```text
dummynumber = [1 2 3 4 5 6]
```

After changing `part[0]`:

```text
part         = [10 3 4]
dummynumber  = [1 10 3 4 5 6]
```

Why did the original slice change?

Because `part[0]` and `dummynumber[1]` refer to the same element in the underlying array.

**Key lesson:** Two slices can share the same underlying array. Modifying a shared element through one slice can affect the other.

## 8. Appending to a Slice with Existing Elements

Consider this example:

```go
newSlice := make([]int, 5, 10)

newdummySlice := []int{}
newdummySlice = append(newSlice, 1, 2, 3, 4, 5)

fmt.Println(newdummySlice)
```

The first statement creates:

```text
newSlice = [0 0 0 0 0]
Length   = 5
Capacity = 10
```

The five zeros are real elements in the slice. They are not unused capacity.

The `append()` call adds five more elements after the existing five zeros.

The result is:

```text
newdummySlice = [0 0 0 0 0 1 2 3 4 5]
Length        = 10
```

The capacity is at least `10`. Because `newSlice` already had capacity `10`, the append operation can fit these elements into its existing underlying array.

### How to create a slice containing only the appended values

If the intention is to start with an empty slice and add only the five numbers, use:

```go
newdummySlice := []int{}
newdummySlice = append(newdummySlice, 1, 2, 3, 4, 5)

fmt.Println(newdummySlice) // [1 2 3 4 5]
```

Alternatively, create a slice with length zero and capacity ten:

```go
newdummySlice := make([]int, 0, 10)
newdummySlice = append(newdummySlice, 1, 2, 3, 4, 5)
```

This version starts with no existing elements but has room for ten elements.

## 9. Common Mistakes to Avoid

| Mistake | Correct understanding |
|---|---|
| Assuming `make([]int, 5, 10)` creates ten existing elements | It creates five elements and capacity for ten. |
| Assuming `copy()` resizes the destination | `copy()` copies only as many elements as fit in the destination. |
| Assuming `append()` always doubles capacity | Go decides the resulting capacity; the exact growth is not guaranteed. |
| Assuming a subslice is an independent copy | A subslice normally shares the original underlying array. |
| Forgetting to assign the result of `append()` | Use `slice = append(slice, values...)`. |
| Assuming an empty slice and a zero-filled slice are equivalent | An empty slice has length zero; a zero-filled slice can have several existing elements containing zero. |

## 🧪 Practice Exercises

Try these exercises by modifying `main.go`:

1. Create a slice with length `3` and capacity `8`. Print both values.
2. Copy five elements into a destination slice of length `2`. Observe the result.
3. Append an element when length equals capacity. Print the new length and capacity.
4. Create a subslice and modify one of its elements. Check the original slice.
5. Create a slice with length zero and capacity five. Append three elements.
6. Modify the final example so that the result contains only `[1 2 3 4 5]`, without the five leading zeros.

## 📝 Final Takeaways

- **Slices** provide a flexible way to work with collections in Go.
- **Length** counts the elements currently in a slice.
- **Capacity** describes the available space from the slice's starting position in its underlying array.
- **`make()`** creates a slice with a chosen length and optional capacity.
- **`copy()`** copies elements into an existing destination slice.
- **`append()`** adds elements and may cause a new underlying array to be allocated.
- **Subslice expressions** create views into existing storage.
- **Shared underlying arrays** explain why changes through one slice can appear in another.

Understanding these relationships is essential for writing correct and predictable Go programs.

---

**Go Learning Series** · Arrays, Slices, Memory, and Data Management