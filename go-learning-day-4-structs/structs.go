# Go Learning — Structs in Go

## Complete Code


package main

import "fmt"

// CONCEPT 1: DEFINING A STRUCT
//
// A struct is like a form containing related information.
//
// User is a blueprint that defines two fields:
// Name -> stores text
// Age  -> stores an integer
//
// Defining a struct does not create a user record.
// It only defines the structure of that record.

type User struct {
    Name string
    Age  int
}

func main() {
    fmt.Println("hello world")

    // CONCEPT 2: CREATING A STRUCT VALUE
    //
    // We create a variable named user.
    //
    // User{} creates a struct value.
    // Name receives "Taha".
    // Age receives 13.
    //
    // This is called a struct literal.

    user := User{
        Name: "Taha",
        Age:  13,
    }

    fmt.Println(user)

    // CONCEPT 3: ACCESSING STRUCT FIELDS
    //
    // A field is an individual piece of information
    // inside a struct.
    //
    // user.Name accesses the Name field.
    // user.Age accesses the Age field.
    //
    // The dot operator connects the struct variable
    // to the field we want to access.

    fmt.Println(user.Name)
    fmt.Println(user.Age)

    // CONCEPT 4: MODIFYING STRUCT FIELDS
    //
    // We can update a field using the dot operator.
    //
    // The Name field changes from "Taha" to "Usama".
    // The Age field changes from 13 to 20.
    //
    // We are modifying the existing struct value.

    user.Name = "Usama"
    user.Age = 20

    fmt.Println(user)

    // CONCEPT 5: CREATING MULTIPLE STRUCT VALUES
    //
    // The User type is a blueprint.
    // We can create many independent User values
    // using the same blueprint.
    //
    // user2 is a separate User value from user.

    user2 := User{
        Name: "Ali",
        Age:  18,
    }

    fmt.Println(user)
    fmt.Println(user2)

    // CONCEPT 6: CREATING A STRUCT WITH ZERO VALUES
    //
    // We can create a User without providing field values.
    //
    // Go automatically assigns each field its zero value:
    // string -> ""
    // int    -> 0

    var user3 User

    fmt.Println(user3)

    // CONCEPT 7: INITIALIZING ONLY SOME FIELDS
    //
    // We do not have to provide every field.
    //
    // The Name field receives "Ahmed".
    // The Age field is not provided, so it remains 0.

    user4 := User{
        Name: "Ahmed",
    }

    fmt.Println(user4)

    // CONCEPT 8: COMPARING STRUCT VALUES
    //
    // User contains string and int fields.
    // Both fields are comparable types.
    //
    // Therefore, two User values can be compared using ==.
    //
    // They are equal only when their corresponding
    // field values are equal.

    firstUser := User{
        Name: "Taha",
        Age:  20,
    }

    secondUser := User{
        Name: "Taha",
        Age:  20,
    }

    fmt.Println(firstUser == secondUser)

    // CONCEPT 9: USING STRUCTS IN CONDITIONALS
    //
    // We can access struct fields inside an if statement.
    //
    // If the user's age is at least 18,
    // the user is considered an adult in this example.

    if user2.Age >= 18 {
        fmt.Println(user2.Name, "is an adult")
    } else {
        fmt.Println(user2.Name, "is not an adult")
    }
}
```

---

## Concept 1: Defining a Struct

```go
type User struct {
    Name string
    Age  int
}
```

Imagine you want to maintain records for students. Every record needs a name and an age.

Instead of defining a completely different type for every student, you create one blueprint called `User`.

- `type` defines a new named type.
- `User` is the name of that type.
- `struct` groups related fields.
- `Name string` defines a field that stores text.
- `Age int` defines a field that stores an integer.

Think of a struct as a blank registration form. It defines which information belongs on the form, but it does not create an individual record.

## Concept 2: Creating a Struct Value

```go
user := User{
    Name: "Taha",
    Age:  13,
}
```

Now we create an actual record using the `User` blueprint.

- `user` is the variable name.
- `User{...}` creates the struct value.
- `Name: "Taha"` initializes the name field.
- `Age: 13` initializes the age field.
- `:=` declares the variable and lets Go infer its type.

The result is conceptually:

```text
user
┌──────────────────┐
│ Name: "Taha"     │
│ Age:  13         │
└──────────────────┘
```

This way of creating a struct value is called a **struct literal**.

## Concept 3: Accessing Struct Fields

```go
fmt.Println(user.Name)
fmt.Println(user.Age)
```

The dot (`.`) operator lets us access a field inside a struct.

- `user.Name` means the name stored in `user`.
- `user.Age` means the age stored in `user`.

Output:

```text
Taha
13
```

Imagine your registration form has two boxes. `user.Name` opens the name box, while `user.Age` opens the age box.

## Concept 4: Modifying Struct Fields

```go
user.Name = "Usama"
user.Age = 20
```

We can change the information stored in an existing struct value.

Before:

```text
Name: Taha
Age:  13
```

After:

```text
Name: Usama
Age:  20
```

The variable `user` still represents the same struct value; we have updated its fields.

When we print it:

```go
fmt.Println(user)
```

Output:

```text
{Usama 20}
```

By default, `fmt.Println` prints the field values in declaration order, without the field names.

## Concept 5: Creating Multiple Struct Values

```go
user := User{
    Name: "Usama",
    Age:  20,
}

user2 := User{
    Name: "Ali",
    Age:  18,
}
```

Both variables use the same `User` type, but they hold different values.

```text
User blueprint
      │
      ├── user
      │   ├── Name: Usama
      │   └── Age:  20
      │
      └── user2
          ├── Name: Ali
          └── Age:  18
```

Changing a field in `user` does not automatically change the corresponding field in `user2`.

**Key idea:** One struct type can describe thousands of different struct values.

## Concept 6: Struct Zero Values

```
var user3 User

fmt.Println(user3)
```

What happens when we create a struct without providing any field values?

Go initializes each field with that field type's **zero value**.

| Field type | Zero value |
|---|---|
| `string` | `""` — empty string |
| `int` | `0` |
| `bool` | `false` |
| Pointer | `nil` |

Because `User` contains a string and an integer, the output is:

```text
{ 0}
```

The space appears because the string is empty and the integer is zero.

There is no need to initialize every field manually if its zero value is suitable.

## Concept 7: Initializing Only Some Fields

```
user4 := User{
    Name: "Ahmed",
}
```

You can provide values for only the fields you need.

Here:

- `Name` becomes `"Ahmed"`.
- `Age` is not specified, so Go initializes it to `0`.

Output:

```text
{Ahmed 0}
```

This is useful when some information is available now and other information will be added later.

## Concept 8: Comparing Struct Values

```go
firstUser := User{
    Name: "Taha",
    Age:  20,
}

secondUser := User{
    Name: "Taha",
    Age:  20,
}

fmt.Println(firstUser == secondUser)
```

Both struct values have identical field values.

Therefore, the output is:

```text
true
```

Now imagine `secondUser.Age` were `21`. The comparison would return `false`.

For this comparison to work, every field in the struct must have a comparable type. Strings and integers are comparable, but a struct containing a slice cannot be compared directly using `==`.

## Concept 9: Using Struct Fields in Conditions

```
if user2.Age >= 18 {
    fmt.Println(user2.Name, "is an adult")
} else {
    fmt.Println(user2.Name, "is not an adult")
}
```

Struct fields work just like ordinary variables when used in conditions.

Go checks whether `user2.Age` is greater than or equal to `18`.

Because `user2.Age` is `18`, the output is:

```text
Ali is an adult
```

The age rule is just an example for learning; real-world age requirements depend on the context and applicable laws.

---

## Quick Revision

| Concept | Example | Meaning |
|---|---|---|
| Define a struct | `type User struct {...}` | Create a type blueprint |
| Create a value | `User{Name: "Taha", Age: 13}` | Build a struct value |
| Access a field | `user.Name` | Read a field |
| Modify a field | `user.Age = 20` | Update a field |
| Create a zero-valued struct | `var user User` | Initialize fields to their zero values |
| Initialize selected fields | `User{Name: "Ali"}` | Leave other fields at zero values |
| Compare values | `firstUser == secondUser` | Compare corresponding fields, when the type is comparable |
| Use fields in conditions | `user.Age >= 18` | Make decisions using struct data |

## Practice Exercises

1. Create a `Student` struct with `Name`, `RollNumber`, and `Marks` fields.
2. Create two students and print their details.
3. Change the marks of the first student.
4. Create a student without initializing any fields and observe the output.
5. Create two `User` values with identical fields and compare them.
6. Print `"Passed"` if a student's marks are at least `50`; otherwise, print `"Failed"`.

**Remember:** A struct is a way to group related data under one type. Define the structure once, create as many values as you need, and access or modify individual fields using the dot operator.