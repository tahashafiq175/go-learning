package main
// Every Go program that runs independently starts with a package declaration.
// package main tells Go that this program is an executable program.

import "fmt"
// We import the fmt package because we want to print information on the screen.
// fmt provides functions such as Println() and Printf().

func main() {
// main() is the starting point of our program.
// When we run this program, Go starts executing the statements inside main().

    fmt.Println("Hello World")
    // Println() prints Hello World on the screen.
    // After printing, it moves to the next line.

    numbers := []int{1, 2, 3, 4, 5}
    // We create a slice named numbers.
    // []int means this is a slice that stores integers.
    // {1, 2, 3, 4, 5} means we initially store five integers inside it.
    // Its length is 5.
    // Go automatically creates an underlying array to store these values.
    // We do not explicitly specify the capacity here.

    fmt.Println(numbers)
    // We print all the elements of numbers.
    // Output: [1 2 3 4 5]


    // Now we will create another slice of integers.
    // We will give it a length of 5.
    // We will then copy values from numbers into this new slice.

    number := make([]int, 5)
    // make() creates a new slice of integers.
    // The first argument, []int, tells Go what type of elements to store.
    // The second argument, 5, sets the length to 5.
    // Because we did not specify capacity separately, capacity is also 5.
    // The slice initially contains five zero values: [0 0 0 0 0].
    // Go creates an underlying array for this slice.

    copy(number, numbers)
    // copy() copies elements from the source slice into the destination slice.
    // number is the destination.
    // numbers is the source.
    // Both slices have length 5, so all five elements are copied.
    // The values in number become [1 2 3 4 5].
    // This copies the values; it does not make both slices share the same array.

    fmt.Println(number)
    // We print the copied slice.
    // Output: [1 2 3 4 5]


    // Now we create another slice with a length of 4.
    // We will copy values from numbers into it.
    // Because the destination can hold only four elements,
    // only the first four values will be copied.

    number2 := make([]int, 4)
    // We create a slice named number2.
    // Its length is 4.
    // Its capacity is also 4 because we did not specify a different capacity.
    // Initially, its elements are [0 0 0 0].
    // Important: this code does NOT create a slice with capacity 10.

    copy(number2, numbers)
    // We copy elements from numbers into number2.
    // numbers contains five elements, but number2 can hold only four.
    // copy() copies the smaller of the source and destination lengths.
    // Therefore, only 1, 2, 3, and 4 are copied.
    // The fifth element, 5, is not copied.

    fmt.Println(number2)
    // Output: [1 2 3 4]


    // Now we create a slice with a length of 5 and a capacity of 5.

    dummynumber := make([]int, 5, 5)
    // make([]int, 5, 5) means:
    // []int = a slice of integers.
    // First 5 = length is 5.
    // Second 5 = capacity is 5.
    // The slice initially contains [0 0 0 0 0].
    // Its underlying array has room for five elements.

    dummynumber[0] = 1
    // We place 1 at index 0, the first position.

    dummynumber[1] = 2
    // We place 2 at index 1, the second position.

    dummynumber[2] = 3
    // We place 3 at index 2, the third position.

    dummynumber[3] = 4
    // We place 4 at index 3, the fourth position.

    dummynumber[4] = 5
    // We place 5 at index 4, the fifth position.
    // The slice now contains [1 2 3 4 5].

    // Initially, the capacity of dummynumber is 5.

    fmt.Println(cap(dummynumber))
    // cap() tells us the capacity of the slice.
    // Output: 5


    // Now we append a new element to dummynumber.
    // The slice's length is already 5, and its capacity is also 5.
    // Therefore, the existing underlying array has no extra capacity.
    // Go must obtain a larger underlying array to accommodate the new element.
    // Go copies the existing values into the new array and adds 6.
    // The exact new capacity is decided by Go and is not guaranteed to be 10.

    dummynumber = append(dummynumber, 6)
    // append() adds 6 to the end of the slice.
    // We assign the returned slice back to dummynumber.
    // This is important because append() may return a slice
    // that refers to a newly allocated underlying array.
    // The slice now contains [1 2 3 4 5 6].
    // Its length is 6.

    fmt.Println(dummynumber)
    // Output: [1 2 3 4 5 6]

    fmt.Println("The capacity of dummynumber is ", cap(dummynumber))
    // We print the capacity after appending.
    // The capacity is at least 6.
    // Its exact value depends on Go's allocation strategy.
    // Do not assume it will always become 10.


    // Now we create a new slice from dummynumber.
    // This new slice will refer to a section of the existing underlying array.
    // Slices are often described as reference-like because they refer to an array.
    // More precisely, a slice is a descriptor containing a pointer to an array,
    // along with its length and capacity.
    // Creating a subslice does not automatically copy the elements.


    part := dummynumber[1:4]
    // We create a slice named part.
    // The starting index is 1, and the ending index is 4.
    // The starting index is included, but the ending index is excluded.
    // Therefore, we select indexes 1, 2, and 3.
    // These values are 2, 3, and 4.
    // part contains [2 3 4].
    // Its length is 3.
    // It shares the underlying array with dummynumber.
    // Its capacity extends from index 1 to the end of that underlying array.

    fmt.Println("Part of the array:", part)
    // Output: Part of the array: [2 3 4]

    part[0] = 10
    // part[0] refers to the first element of part.
    // That element originally contained 2.
    // We replace 2 with 10.
    // Because part shares the underlying array with dummynumber,
    // the corresponding element in dummynumber also changes from 2 to 10.

    fmt.Println("Part of the array:", part)
    // Output: Part of the array: [10 3 4]

    fmt.Println("capacity of the part arr is ", cap(part))
    // We print the capacity of part.
    // Its capacity is the number of elements available from its starting
    // position to the end of the shared underlying array.
    // The exact number depends on the capacity of dummynumber.

    fmt.Println("Numbers:", dummynumber)
    // We print dummynumber to see the effect of changing part.
    // Output: Numbers: [1 10 3 4 5 6]
    // The second element changed because both slices share the same array.


    // Now we create a new empty slice.

    newdummySlice := []int{}
    // This creates an empty slice of integers.
    // Its length is 0.
    // Its capacity is 0 for this empty slice literal.
    // We can later add elements using append().


    // Now we create another slice with a length of 5 and capacity of 10.

    newSlice := make([]int, 5, 10)
    // make([]int, 5, 10) means:
    // []int = a slice of integers.
    // 5 = the initial length is 5.
    // 10 = the initial capacity is 10.
    // The slice initially contains [0 0 0 0 0].
    // It has five existing elements, with room to grow up to ten elements
    // before it needs a larger underlying array.


    // Now we append the values 1, 2, 3, 4, and 5 to newSlice.
    // Important: newSlice already contains five elements.
    // Those elements are initially all zero.
    // append() adds the new values AFTER those five existing elements.

    newdummySlice = append(newSlice, 1, 2, 3, 4, 5)
    // append() takes newSlice as its starting slice.
    // It adds 1, 2, 3, 4, and 5 to the end of newSlice.
    // Therefore, the result is:
    // [0 0 0 0 0 1 2 3 4 5]
    // The result has a length of 10.
    // Its capacity is at least 10.
    //
    // Notice that we assign the result to newdummySlice.
    // The variable newdummySlice now refers to the slice returned by append().
    // It does not remain empty.


    fmt.Println("New dummy Slice:", newdummySlice)
    // Output: New dummy Slice: [0 0 0 0 0 1 2 3 4 5]
    // The five zeros are the original elements of newSlice.
    // The five numbers after them are the values added by append().
}
