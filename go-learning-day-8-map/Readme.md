# Go Learning — Maps

A beginner-friendly, practical guide to understanding maps in Go, from basic key-value operations to nested maps, slices, structs, pointers, and functions.

## 📌 Overview

A **map** in Go stores information as **key-value pairs**.

Think of a map like a dictionary. You look up a word (the key) to find its meaning (the value). In Go, you can use a key to retrieve, add, update, or delete its associated value.

For example:

```go
age := map[string]int{
    "Taha":  22,
    "Usama": 24,
}
```

Here, `"Taha"` and `"Usama"` are keys, while `22` and `24` are their values.

Maps are useful when you need to look up information using a unique identifier, such as a username, product ID, country code, or configuration key.

## 🎯 Learning Objectives

After completing this module, you should understand how to:

- Create and initialize maps.
- Read, add, update, and delete map entries.
- Check whether a key exists.
- Understand map zero values and initialization.
- Use `len()`, `delete()`, and `clear()`.
- Iterate through maps using `range`.
- Understand map reference-like behavior and copying.
- Pass maps to functions and modify their entries.
- Create nested maps and maps containing slices.
- Store structs and pointers to structs inside maps.
- Understand map limitations, concurrency safety, and common mistakes.

## 📂 Project Structure

```text
go-maps/
├── main.go
└── README.md
```

Your `main.go` file contains the practical demonstrations. This README explains the concepts and provides additional examples.

> **Important:** Your current example defines `mapLearningDemo()` rather than `main()`. To run it as a standalone program, add a `main()` function that calls `mapLearningDemo()`.

For example:

```go
func main() {
    mapLearningDemo()
}
```

## 1. What Is a Map?

A map associates each key with a value.

```go
age := map[string]int{
    "Taha":  22,
    "Usama": 24,
}
```

Understand the declaration:

| Part | Meaning |
|---|---|
| `map` | Declares a map type. |
| `string` | The type of the keys. |
| `int` | The type of the values. |
| `"Taha"` | A key. |
| `22` | The value associated with that key. |

The type `map[string]int` means that every key must be a string and every value must be an integer.

Other examples include:

```go
scores := map[string]float64{
    "Math":    95.5,
    "English": 88.0,
}

isOnline := map[string]bool{
    "Taha": true,
    "Ali":  false,
}
```

A map's keys must be of a comparable type. Strings, integers, and booleans can be keys. Slices, maps, and functions cannot be used directly as map keys.

## 2. Creating Maps

Go provides several ways to create maps.

### Method 1 — Map literal

```go
age := map[string]int{
    "Taha":  22,
    "Usama": 24,
}
```

This creates a map and initializes it with entries.

### Method 2 — Using `make()`

```go
age := make(map[string]int)
```

This creates an empty, initialized map ready to receive entries.

```go
age["Taha"] = 22
age["Usama"] = 24
```

### Method 3 — Declaring a map without initializing it

```go
var age map[string]int
```

This declares a map variable whose value is `nil`.

You can safely read from a nil map, but you cannot add entries to it until it is initialized.

```go
var age map[string]int

fmt.Println(age["Taha"]) // 0

// age["Taha"] = 22
// Runtime panic: assignment to entry in nil map
```

Initialize it before writing:

```go
age = make(map[string]int)
age["Taha"] = 22
```

**Remember:** A nil map and an empty initialized map are different.

| Property | Nil map | Empty initialized map |
|---|---|---|
| Declaration | `var m map[string]int` | `make(map[string]int)` |
| Can read a missing key | Yes | Yes |
| Can add entries | No | Yes |
| `len(m)` | `0` | `0` |
| Is `m == nil` | `true` | `false` |

## 3. Reading Values from a Map

Use a key inside square brackets to retrieve its value.

```go
age := map[string]int{
    "Taha":  22,
    "Usama": 24,
}

fmt.Println(age["Taha"]) // 22
fmt.Println(age["Usama"]) // 24
```

The expression `age["Taha"]` asks Go to retrieve the value associated with `"Taha"`.

### What happens if the key does not exist?

```go
fmt.Println(age["Ali"]) // 0
```

Because the value type is `int`, Go returns its zero value, `0`, when the key is missing.

Other zero values include:

| Value type | Zero value |
|---|---|
| `int` | `0` |
| `string` | `""` |
| `bool` | `false` |
| Pointer | `nil` |
| Slice | `nil` |

A returned zero value does not prove that a key is absent. The key might exist and genuinely store zero.

## 4. Checking Whether a Key Exists

Go provides a two-result lookup to distinguish a missing key from a key that stores its type's zero value.

```go
ages := map[string]int{
    "Taha":  12,
    "Usama": 23,
    "Usman": 25,
}

inAge, exists := ages["Usman"]

fmt.Println(inAge)  // 25
fmt.Println(exists) // true
```

The first result is the value. The second result is a boolean indicating whether the key exists.

For a missing key:

```go
value, exists := ages["Ali"]

fmt.Println(value)  // 0
fmt.Println(exists) // false
```

This pattern is especially useful when a map contains meaningful zero values.

```go
if value, exists := ages["Ali"]; exists {
    fmt.Println("Age:", value)
} else {
    fmt.Println("Student not found")
}
```

This is called the **comma-ok idiom**.

## 5. Adding and Updating Entries

Use assignment to add or update a map entry.

```go
age := map[string]int{
    "Taha": 22,
}

age["Ali"] = 23 // Adds Ali
age["Taha"] = 21 // Updates Taha
```

The behavior depends on whether the key already exists:

- New key: Go adds a new entry.
- Existing key: Go replaces the value associated with that key.

A map cannot contain two separate entries with the same key. Assigning a new value to an existing key updates that entry.

## 6. Deleting Entries

Use the built-in `delete()` function.

```go
age := map[string]int{
    "Taha": 22,
    "Ali":  23,
}

delete(age, "Ali")

fmt.Println(age)
```

`delete()` takes two arguments:

1. The map to modify.
2. The key to remove.

Deleting a key that does not exist is safe:

```go
delete(age, "Ali")
```

This does nothing if `"Ali"` has already been removed.

### Clearing an entire map

In modern Go, you can use `clear()` to remove all entries:

```go
age := map[string]int{
    "Taha": 22,
    "Ali":  23,
}

clear(age)

fmt.Println(len(age)) // 0
```

The map remains initialized and can accept new entries afterward.

## 7. Finding the Number of Entries

Use `len()` to count the entries in a map.

```go
age := map[string]int{
    "Taha":  22,
    "Usama": 24,
}

fmt.Println(len(age)) // 2
```

`len()` returns the number of entries, not the number of possible keys or the map's internal capacity.

## 8. Iterating Through a Map with `range`

Use `range` to visit each key-value pair.

```go
age := map[string]int{
    "Taha":  22,
    "Usama": 24,
}

for key, value := range age {
    fmt.Printf("Name: %s, Age: %d\n", key, value)
}
```

During each iteration:

- `key` receives the current key.
- `value` receives the value associated with that key.

**Important:** Go does not guarantee map iteration order. The entries may appear in a different order between runs.

If you need sorted output, collect the keys into a slice, sort the slice, and then retrieve the values using those keys.

```go
import (
    "fmt"
    "sort"
)

age := map[string]int{
    "Usama": 24,
    "Taha":  22,
    "Ali":   23,
}

names := make([]string, 0, len(age))

for name := range age {
    names = append(names, name)
}

sort.Strings(names)

for _, name := range names {
    fmt.Println(name, age[name])
}
```

This produces output in alphabetical order.

## 9. Maps Behave Like Shared References

Consider the following code:

```go
originalMap := map[string]int{
    "Taha":  21,
    "Usama": 123,
    "Ali":   213,
}

dummyMap := originalMap

dummyMap["Taha"] = 32

fmt.Println(originalMap["Taha"]) // 32
```

Why did the original map change?

A map value refers to internally managed map data. Assigning it to another variable does not copy every entry into a separate map.

Both `originalMap` and `dummyMap` refer to the same underlying map data.

You can think of them as two labels pointing to the same notebook. Changing an entry through either label changes the same notebook.

Technically, maps are not pointers in the same sense as `*Student`. A map value is a descriptor managed by Go, and assigning it copies that descriptor rather than independently copying all entries.

## 10. Making an Independent Copy of a Map

If you need a separate map, create one and copy the entries.

```go
originalMap := map[string]int{
    "Taha":  21,
    "Usama": 123,
    "Ali":   213,
}

newMap := make(map[string]int)

for key, value := range originalMap {
    newMap[key] = value
}

newMap["Taha"] = 100

fmt.Println(originalMap["Taha"]) // 21
fmt.Println(newMap["Taha"])      // 100
```

The two maps now have separate entries.

Go also provides a convenient alternative using the `maps` standard-library package:

```go
import "maps"

newMap := maps.Clone(originalMap)
```

`maps.Clone()` makes a **shallow copy**. For a map with integer values, the entries are copied independently. If values contain pointers, slices, or other references, the objects those references point to may still be shared.

## 11. Passing Maps to Functions

Maps can be passed to functions and modified inside those functions.

```go
func changeAge(age map[string]int) {
    age["Taha"] = 1233
}

func main() {
    newMap := map[string]int{
        "Taha": 21,
    }

    changeAge(newMap)

    fmt.Println(newMap["Taha"]) // 1233
}
```

The parameter `age` receives a copy of the map value, but that value refers to the same map data as `newMap`.

Therefore, changing an entry inside `changeAge()` is visible to the caller.

You do not need a pointer parameter or a dereference operator to update an existing map entry.

If a function needs to replace the caller's map variable with a completely different map, that is a different operation and may require returning the new map or passing a pointer to the map variable.

## 12. Nested Maps

A nested map is a map whose values are themselves maps.

```go
student := map[string]map[string]int{
    "Taha": {
        "Math":    90,
        "English": 100,
    },
    "Usama": {
        "Math":    98,
        "English": 92,
    },
}
```

The outer map stores student names. Each value contains another map of subjects and marks.

To access Taha's Math mark:

```go
fmt.Println(student["Taha"]["Math"]) // 90
```

Read the expression from left to right:

1. `student["Taha"]` retrieves Taha's inner map.
2. `["Math"]` retrieves the Math value from that inner map.

### Adding an inner map

```go
newStudent := make(map[string]map[string]int)

newStudent["Taha"] = map[string]int{
    "Math":    90,
    "English": 100,
}
```

An important detail: assigning a value into a nil inner map causes a panic. Initialize the inner map before adding its entries.

```go
student := make(map[string]map[string]int)

student["Taha"] = make(map[string]int)
student["Taha"]["Math"] = 90
```

## 13. Maps Containing Slices

A map can store slices as values.

```go
mapSlice := map[string][]string{
    "Taha":  {"Math", "English", "Urdu"},
    "Usama": {"History", "English", "Science"},
}
```

The type `map[string][]string` means:

- Keys are strings.
- Values are slices of strings.

Access the whole slice:

```go
fmt.Println(mapSlice["Taha"])
```

Access its first element:

```go
fmt.Println(mapSlice["Taha"][0]) // Math
```

Append another subject:

```go
mapSlice["Taha"] = append(mapSlice["Taha"], "Database")
```

Assign the result back to the map entry because `append()` can return a slice with a new length and potentially a new underlying array.

## 14. Slices of Maps

A slice can store maps as its elements.

```go
arrMap := []map[string]int{
    {
        "Taha":  23,
        "Usama": 32,
    },
    {
        "Math":    90,
        "English": 95,
    },
}
```

Here:

- `[]` means the outer collection is a slice.
- `map[string]int` means each element is a map.

Access the first map:

```go
fmt.Println(arrMap[0])
```

Access a value in that map:

```go
fmt.Println(arrMap[0]["Taha"]) // 23
```

Although the variable is named `arrMap`, the type is a slice of maps, not an array of maps.

## 15. Storing Structs in Maps

Maps can store custom struct values.

```go
type Student struct {
    Name string
    Age  int
    Mark int
}

structMap := map[string]Student{
    "Taha": {
        Name: "Taha Shafiq",
        Age:  21,
        Mark: 35,
    },
}
```

The map type `map[string]Student` means that each key is a string and each value is a `Student` struct.

Access a field:

```go
fmt.Println(structMap["Taha"].Name)
fmt.Println(structMap["Taha"].Age)
fmt.Println(structMap["Taha"].Mark)
```

This pattern is useful for storing structured records such as students, employees, products, and users.

### Updating a struct value stored in a map

Map values that are structs cannot be modified through direct field assignment.

This is invalid:

```go
// structMap["Taha"].Age = 25
```

Instead, retrieve the struct, modify the copy, and assign it back:

```go
student := structMap["Taha"]

student.Age = 25

structMap["Taha"] = student
```

The last assignment replaces the stored struct value.

## 16. Storing Pointers to Structs in Maps

Another approach is to store pointers to structs.

```go
type Student struct {
    Name string
    Age  int
    Mark int
}

studentMap := map[string]*Student{
    "Taha": {
        Name: "Taha Shafiq",
        Age:  23,
        Mark: 32,
    },
}
```

The type `map[string]*Student` means that each value is a pointer to a `Student`.

Now you can update a field directly:

```go
studentMap["Taha"].Age = 34

fmt.Println(studentMap["Taha"].Age) // 34
```

Go automatically follows the pointer when accessing the struct field.

You can explicitly dereference the pointer:

```go
fmt.Println(*studentMap["Taha"])
```

This prints the struct value referenced by the pointer.

**Important:** If a key is missing, `studentMap["Ali"]` returns a nil `*Student`. Trying to access a field through that nil pointer causes a runtime panic. Check that the key exists before dereferencing it.

## 17. Maps Containing Structs with Slice Fields

A struct can itself contain a slice.

```go
type Students struct {
    Name string
    Age  int
    Mark []string
}

mapArrayStruct := map[string]Students{
    "Taha": {
        Name: "Taha Shafiq",
        Age:  34,
        Mark: []string{"English", "Urdu"},
    },
}
```

In this example:

- The map key is a string.
- Each value is a `Students` struct.
- The struct's `Mark` field is a slice of strings.

Access the student's subjects:

```go
fmt.Println(mapArrayStruct["Taha"].Mark)
```

The output contains the two subject names.

If you need to change a field directly on the struct, use the retrieve-modify-assign approach described earlier. Slice elements themselves can sometimes be modified through the slice field, but appending to that field requires care because the updated slice header may need to be stored back into the map.

For example:

```go
student := mapArrayStruct["Taha"]

student.Mark = append(student.Mark, "Math")

mapArrayStruct["Taha"] = student
```

## 18. Important Map Limitations

### Maps are not safe for concurrent writes

A regular Go map should not be read and written concurrently by multiple goroutines without synchronization.

Concurrent map access can cause a runtime failure or a data race.

Use a `sync.Mutex` or `sync.RWMutex` to protect shared map access, or consider `sync.Map` when its specific usage pattern fits your application.

### Map iteration order is unspecified

Never depend on a map returning entries in insertion order or alphabetical order. Sort the keys when deterministic ordering is needed.

### Map keys must be comparable

Valid examples:

```go
map[string]int{}
map[int]string{}
map[bool]string{}
```

Invalid examples:

```go
// map[[]int]string{} // A slice cannot be a map key.
// map[map[string]int]string{} // A map cannot be a map key.
```

### Maps are not directly comparable

You can compare a map with `nil`, but you cannot compare two maps using `==` to check whether their contents are equal.

```go
if age == nil {
    fmt.Println("Map is nil")
}
```

To compare map contents, compare their lengths and entries or use an appropriate helper from the standard library, such as `maps.Equal()` when the key and value types meet its requirements.

## 19. Common Mistakes

| Mistake | Correct approach |
|---|---|
| Writing to a nil map | Initialize it with `make()` or a map literal. |
| Assuming a missing key causes an error | Use the comma-ok idiom to check whether it exists. |
| Assuming map iteration is ordered | Sort the keys if order matters. |
| Expecting a map assignment to copy all entries | Create a new map and copy entries explicitly. |
| Modifying a struct field directly inside a map value | Retrieve the struct, modify it, and assign it back. |
| Writing to an uninitialized nested map | Initialize the inner map first. |
| Assuming maps are safe for concurrent writes | Protect shared access with synchronization. |
| Using slices or maps as map keys | Use a comparable key type. |

## 🧪 Practice Exercises

Try these exercises by modifying `main.go`:

1. Create a map of five students and their marks.
2. Add a new student, update a mark, and delete a student.
3. Check whether a student exists using the comma-ok idiom.
4. Print all entries and then print them in alphabetical order.
5. Copy a map and prove that changing the copy does not change the original.
6. Write a function that updates a map entry.
7. Create a nested map of students and subjects.
8. Create a map where each student has a slice of subjects.
9. Store structs in a map and update a struct field.
10. Store pointers to structs in a map and modify a field through a pointer.
11. Create a nil map and an initialized empty map, then compare their behavior.
12. Write a function that safely checks whether a nested-map entry exists before accessing it.

## 🧠 Final Takeaways

- **Map:** Stores key-value pairs for efficient lookup.
- **`make()`:** Creates an initialized map.
- **`map[key]`:** Retrieves the value associated with a key.
- **Comma-ok:** Checks whether a key exists.
- **Assignment:** Adds a new entry or updates an existing one.
- **`delete()`:** Removes a map entry.
- **`clear()`:** Removes all entries.
- **`range`:** Iterates through map entries in unspecified order.
- **Map assignment:** Copies the map value, not all entries into an independent map.
- **Explicit copying:** Creates a separate map, but copying is shallow if values contain references.
- **Nested maps and slices:** Represent grouped or hierarchical information.
- **Struct values:** Must be retrieved and reassigned when modifying a stored struct.
- **Struct pointers:** Allow direct field modification through the pointer.
- **Concurrency:** Requires synchronization when goroutines share a map and access it concurrently.

The key to mastering Go maps is understanding how keys identify values, how map variables refer to shared map data, and how maps interact with slices, structs, pointers, and functions.

---

**Go Learning Series** · Maps, Collections, and Data Structures