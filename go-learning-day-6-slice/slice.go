# Go Arrays and Slices — Simple Learning Notes

package main

import "fmt"

// CONCEPT 1: STRUCT DEFINITION
//
// A struct is like a form containing related information.
//
// Here, arrayStruct is a blueprint with two fields:
// Name -> stores text
// Age  -> stores an integer
//
// We will use this struct to store multiple people's
// information inside a slice later.

type arrayStruct struct {
    Name string
    Age  int
}

func main() {
    fmt.Println("Hello World")

    // CONCEPT 2: CREATING A FIXED-SIZE ARRAY
    //
    // numbers is an array that can store exactly 5 integers.
    //
    // [5]int means:
    // 5   -> the array has 5 elements
    // int -> every element stores an integer
    //
    // The elements are stored at indexes 0, 1, 2, 3, and 4.
    //
    // IMPORTANT:
    // An array's length is fixed. We cannot append another
    // element to this array to increase its length.

    numbers := [5]int{1, 2, 3, 4, 5}

    fmt.Println(numbers)

    // CONCEPT 3: CREATING A SLICE
    //
    // Here, we write []int instead of [5]int.
    //
    // A slice is a flexible view of elements backed by an array.
    // Unlike an array, a slice can grow when we append elements.
    //
    // We do not specify the length inside [].
    // Go creates a slice containing these five integers.
    //
    // IMPORTANT:
    // A slice is not simply a dynamically sized array.
    // It refers to an underlying array and keeps track
    // of its length and capacity.

    number := []int{1, 2, 3, 4, 5}

    fmt.Println(number)

    // CONCEPT 4: APPENDING ONE ELEMENT
    //
    // append() adds an element to the end of a slice.
    //
    // Here, we add 6 after the existing elements.
    //
    // append() returns the resulting slice, so we assign
    // that result back to number.
    //
    // The slice now contains six elements.

    number = append(number, 6)

    fmt.Println(number)

    // CONCEPT 5: APPENDING MULTIPLE ELEMENTS
    //
    // append() can add several elements in one operation.
    //
    // Here, we add 7, 8, and 9 to the end of number.
    //
    // The slice now contains nine elements.

    number = append(number, 7, 8, 9)

    fmt.Println(number)

    // CONCEPT 6: FINDING THE LENGTH OF A SLICE
    //
    // len() tells us how many elements a slice currently contains.
    //
    // number contains nine elements, so its length is 9.
    //
    // Length counts the elements that are currently
    // part of the slice, not the total available capacity.

    size := len(number)

    fmt.Println("Size of the slice:", size)

    // CONCEPT 7: CREATING A SLICE USING make()
    //
    // make() can create a slice without writing all
    // its elements manually.
    //
    // make([]int, 5, 10) means:
    // []int -> create a slice of integers
    // 5     -> set its length to 5
    // 10    -> set its capacity to 10
    //
    // All five existing elements initially contain 0.
    //
    // IMPORTANT:
    // Capacity is not the same as length.
    // We can access indexes 0 through 4 right now.
    // We cannot directly access index 5 until the slice grows.

    numbers2 := make([]int, 5, 10)

    fmt.Println("The capacity of numbers2 is ", cap(numbers2))

    // CONCEPT 8: CHANGING A SLICE ELEMENT
    //
    // numbers2[0] accesses the first element.
    //
    // Slice indexes start at 0, not 1.
    //
    // We replace the first element's value with 1.
    // The other four elements remain 0.

    numbers2[0] = 1

    fmt.Println("Numbers2:", numbers2)

    // CONCEPT 9: MODIFYING AN ELEMENT BY INDEX
    //
    // We can also modify an existing element in number.
    //
    // number[0] means the first element of number.
    //
    // We change that element from 1 to 0.
    // This does not add or remove any elements.

    number[0] = 0

    fmt.Println("Numbers:", number)

    // CONCEPT 10: CREATING A SUBSLICE
    //
    // numbers[1:4] creates a slice starting at index 1
    // and stopping just before index 4.
    //
    // Start index 1 -> included
    // End index 4   -> excluded
    //
    // So the selected elements are at indexes 1, 2, and 3.
    //
    // IMPORTANT:
    // This slice usually shares the same underlying array
    // as numbers. It does not automatically copy the elements.

    part := numbers[1:4]

    fmt.Println("Part of the array:", part)

    // CONCEPT 11: CHANGING AN ELEMENT THROUGH A SUBSLICE
    //
    // part[0] refers to the first element of the part slice.
    //
    // That element originally came from numbers[1].
    //
    // Because part and numbers share the same underlying array,
    // changing part[0] also changes numbers[1].
    //
    // This is why modifying a subslice can affect
    // the original slice.

    part[0] = 10

    fmt.Println("Part of the array:", part)

    // CONCEPT 12: FINDING THE CAPACITY OF A SUBSLICE
    //
    // cap() tells us how many elements the slice can cover
    // from its starting position to the end of its
    // underlying array.
    //
    // For numbers[1:4], the capacity is measured starting
    // at index 1 of numbers.
    //
    // numbers has a length of 9 and, in this example,
    // its capacity is at least 9.
    //
    // Therefore, part has a capacity of 8.

    fmt.Println("capacity of the part arr is ", cap(part))

    fmt.Println("Numbers:", numbers)

    // CONCEPT 13: LOOPING THROUGH A SLICE USING range
    //
    // range lets us visit every element in numbers.
    //
    // index receives the element's index.
    // value receives the element's value.
    //
    // The loop visits each element one by one.
    //
    // IMPORTANT:
    // fmt.Sprintf() creates a formatted string but does not
    // print it. Since we do not store or print that string,
    // the formatted result is discarded.
    //
    // Use fmt.Printf() to print the formatted string.

    for index, value := range numbers {
        fmt.Printf(
            "The index of numbers array is %d and value is %d\n",
            index,
            value,
        )
    }

    // CONCEPT 14: A SLICE OF STRUCTS
    //
    // We can store structs inside a slice, just as we
    // store integers inside a slice.
    //
    // []arrayStruct means a slice whose elements have
    // the type arrayStruct.
    //
    // Each element stores a person's name and age.
    //
    // This slice contains three records:
    // Taha, Ali, and Ahmed.

    structarray := []arrayStruct{
        {Name: "Taha", Age: 20},
        {Name: "Ali", Age: 25},
        {Name: "Ahmed", Age: 30},
    }

    fmt.Println("Struct of array of structure is :", structarray)

    // CONCEPT 15: LOOPING THROUGH A SLICE OF STRUCTS
    //
    // range visits every struct in structarray.
    //
    // The underscore (_) means we do not need the index.
    // value receives the current arrayStruct value.
    //
    // value.Name accesses that person's name.
    // value.Age accesses that person's age.
    //
    // fmt.Printf() replaces %s with the name
    // and %d with the integer age, then prints the result.

    for _, value := range structarray {
        fmt.Printf(
            "The name is %s and age is %d\n",
            value.Name,
            value.Age,
        )
    }
}
