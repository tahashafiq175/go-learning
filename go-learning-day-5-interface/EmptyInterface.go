# Go Learning Notes — Interfaces and Empty Interfaces

## Complete Code With Simple Explanations


package main

import "fmt"

// CONCEPT 1: DEFINING AN INTERFACE
//
// An interface describes what a type must be able to do.
//
// Imagine we have different animals, such as cats and dogs.
// A cat makes a sound and moves.
// A dog also makes a sound and moves.
//
// We want both animals to provide two methods:
// Speak() -> makes a sound
// Move()  -> describes movement
//
// IMPORTANT:
// An interface describes behavior, not the actual data
// stored inside a struct.
//
// In Go, a type implements an interface automatically
// by providing all the required methods with matching
// names and signatures.
//
// NOTE:
// "Animal" is the conventional spelling of this name.
// Go is case-sensitive, so Animal and Anmal are different names.

type Animal interface {
    Speak()
    Move()
}

// CONCEPT 2: DEFINING TYPES THAT WILL IMPLEMENT THE INTERFACE
//
// Cat and Dog are two different types.
//
// They are empty structs because they contain no fields.
// We are using them to focus on their behavior.
//
// An empty struct is written as struct{}.

type Cat struct{}
type Dog struct{}

// CONCEPT 3: IMPLEMENTING SPEAK() FOR CAT
//
// The Cat type has a method named Speak().
//
// When a Cat receives a call to Speak(),
// it prints "Meow".
//
// Because the method name and signature match
// the interface requirement, Cat provides this behavior.

func (c Cat) Speak() {
    fmt.Println("Meow")
}

// CONCEPT 4: IMPLEMENTING SPEAK() FOR DOG
//
// Dog also has a method named Speak().
//
// The method has the same name and signature
// as Cat's Speak(), but it performs different behavior.
//
// Cat says "Meow".
// Dog says "Woof".
//
// This is how different types can provide their own
// implementations of the same interface method.

func (d Dog) Speak() {
    fmt.Println("Woof")
}

// CONCEPT 5: IMPLEMENTING MOVE() FOR CAT
//
// The Animal interface also requires Move().
//
// Cat provides this method, so it satisfies the
// Move() requirement.
//
// Together, Cat's Speak() and Move() methods satisfy
// every method required by Animal.

func (c Cat) Move() {
    fmt.Println("Runs on all fours")
}

// CONCEPT 6: IMPLEMENTING MOVE() FOR DOG
//
// Dog also provides the required Move() method.
//
// Dog now has both required methods:
// Speak()
// Move()
//
// Therefore, Dog also satisfies the Animal interface.

func (d Dog) Move() {
    fmt.Println("Runs on all fours")
}

// CONCEPT 7: ACCEPTING AN INTERFACE AS A PARAMETER
//
// The parameter a has type Animal.
//
// This means makeSound() accepts values whose types
// implement the Animal interface.
//
// It does not need to know whether the actual value
// is a Cat or a Dog.
//
// The function simply asks the animal to Speak().
//
// The actual type's Speak() method determines the sound.

func makeSound(a Animal) {
    a.Speak()
}

// CONCEPT 8: THE EMPTY INTERFACE — ANY
//
// The any type can hold a value of any type.
//
// For example:
// int     -> 10
// string  -> "Taha"
// float64 -> 10.5
// bool    -> true
//
// Unlike Animal, which requires Speak() and Move(),
// any does not require any methods.
//
// Therefore, almost any ordinary Go value can be passed
// to a function whose parameter type is any.
//
// any is an alias for interface{}.

func EmptyInterfaceFunction(value any) {

    // CONCEPT 9: ASSIGNING AN INTERFACE VALUE
    //
    // Here, numbers receives the value stored in value.
    //
    // Despite its name, numbers does not automatically
    // become a numeric variable.
    //
    // Its inferred type is any because value has type any.
    //
    // The actual value and its dynamic type are preserved.

    numbers := value

    // Both statements print the stored value.
    //
    // For example, if value contains "taha",
    // both lines print taha.

    fmt.Println(value)
    fmt.Println(numbers)

    // CONCEPT 10: TYPE SWITCH
    //
    // An any value can contain different types.
    //
    // A type switch lets us check the actual type
    // of the value currently stored inside the interface.
    //
    // v becomes the value with the type matched by a case.
    //
    // If value contains an int, the int case runs.
    // If value contains a string, the string case runs.
    // Only the matching case executes.

    switch v := value.(type) {

    case int:
        // If the value is an integer, multiply it by 2.
        // Example: 10 becomes 20.
        fmt.Println("Integer:", v*2)

    case float64:
        // If the value is a float64, print it.
        fmt.Println("Float:", v)

    case string:
        // If the value is text, print it as a string.
        fmt.Println("String:", v)

    case bool:
        // If the value is true or false, print it as a Boolean.
        fmt.Println("Boolean:", v)

    default:
        // This runs if none of the listed types match.
        fmt.Println("Unknown type")
    }
}

// CONCEPT 11: TYPE ASSERTION WITH THE OK IDIOM
//
// value has type any, so we do not know whether it
// currently contains an integer.
//
// The expression value.(int) checks whether the
// dynamic type inside value is int.
//
// The two results are:
// integer -> the integer value, if the assertion succeeds
// ok      -> true if successful, otherwise false
//
// This two-result form is called the "comma-ok" idiom.
// It avoids a panic when the stored value is not an int.

func EmptyOkInterface(value any) {

    integer, ok := value.(int)

    // If value contains an int, ok is true.
    // We can safely use integer in this branch.

    if ok {
        fmt.Println("Integer:", integer*2)

    } else {
        // If value contains something else, such as a string,
        // ok is false and this branch executes.
        fmt.Println("Not an integer")
    }
}

// CONCEPT 12: CALLING THE FUNCTIONS
//
// We pass different values to EmptyInterfaceFunction().
//
// The parameter type is any, so these calls are valid:
// a string, an integer, a float64, and a Boolean.
//
// We then test the comma-ok type assertion with
// an integer and a string.

func EmptyInterface() {

    EmptyInterfaceFunction("taha")
    EmptyInterfaceFunction(10)
    EmptyInterfaceFunction(10.5)
    EmptyInterfaceFunction(true)

    EmptyOkInterface(10)
    EmptyOkInterface("taha")
}

func main() {

    // CONCEPT 13: INTERFACE POLYMORPHISM
    //
    // Create one Cat and one Dog.
    //
    // Both types implement Animal, so both can be
    // passed to makeSound().
    //
    // The function does not need separate versions
    // for cats and dogs.
    //
    // Each type provides its own Speak() behavior.

    cat := Cat{}
    dog := Dog{}

    makeSound(cat)
    makeSound(dog)

    // We can also call their methods directly.

    cat.Speak()
    dog.Speak()

    cat.Move()
    dog.Move()

    // CONCEPT 14: TESTING ANY AND TYPE ASSERTIONS
    //
    // This function demonstrates how any can hold
    // values of different types and how we can inspect them.

    EmptyInterface()
}
```

---

## Expected output

```

```
text

Meow
Woof
Meow
Woof
Runs on all fours
Runs on all fours
taha
taha
String: taha
10
10
Integer: 20
10.5
10.5
Float: 10.5
true
true
Boolean: true
Integer: 20
Not an integer
```
