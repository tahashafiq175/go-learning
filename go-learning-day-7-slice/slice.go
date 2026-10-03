package main

import "fmt"

func main() {
	fmt.Println("Hello World")

 numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Now we will create a slice of integers
	// with a length of 5 and a capacity of 10
	//such that we are copying the values of the numbers slice into the new slice
	//
	number:=make([]int, 5)
	copy(number, numbers)
	fmt.Println(number)
// Now we will create a slice of integers
	// with a length of 4 and a capacity of 10
	//such that we are copying the values of the numbers slice into the new slice
	//but this time we will create a new slice with a length of 4 and a capacity of 10
	//so that we can see the difference between the two slices
	//such that number2 will have a length of 4 and a capacity of 10 so that 
	//is receiving the first 4 elements of the numbers slice and the last element will be
	number2:=make([]int, 4)
	copy(number2, numbers)
	fmt.Println(number2)



	dummynumber := make([]int, 5,5)
	dummynumber[0] = 1
	dummynumber[1] = 2
	dummynumber[2] = 3
	dummynumber[3] = 4
	dummynumber[4] = 5

	//initially the capacity of the dummynumber slice is 5
	fmt.Println(cap(dummynumber))
	//now we will append a new element to the dummynumber slice
	//since the capacity of the dummynumber slice is 5, it will create a new slice with a capacity of 10
	//and copy the values of the dummynumber slice into the new slice and then append the new element to the new slice
	dummynumber = append(dummynumber, 6)
	fmt.Println(dummynumber)	
	fmt.Println("The capacity of dummynumber is ", cap(dummynumber))


	//now we will create a new slice from the dummynumber slice
	//basically slice is a reference type, so when we create a new slice from the dummynumber slice, it will point
	//  to the same underlying array as the dummynumber slice
    //we can say part slice is basically  a pointer that is pointing to array underlying array of the dummynumber slice
	//so when we create another slice from it it will also points
	//to the same underlying array as the dummynumber slice
	//so changes made to the new slice will also be reflected in the dummynumber slice
	part := dummynumber[1:4]
	fmt.Println("Part of the array:", part)
part[0] = 10
	fmt.Println("Part of the array:", part)
	fmt.Println("capacity of the part arr is ", cap(part))
	fmt.Println("Numbers:", dummynumber)


	//now we are creating a new slice by appending the values of the dummynumber slice to a new slice


	//like we are simply appenidng 1,2,3,4,5 to newdummySlice and new dummyslice has first newSlice
	//newSlice is a slice of integers with a length of 5 and a capacity of 10 and it has only 0,0,0,0,0 in it and we are appending 1,2,3,4,5 to it and then we are printing the newdummySlice
newdummySlice :=[]int{}
	newSlice:=make([]int, 5, 10)
	newdummySlice=append(newSlice, 1, 2, 3, 4, 5)
	fmt.Println("New dummy Slice:", newdummySlice)
}