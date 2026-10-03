package main

import "fmt"

type arrayStruct struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Hello World")

	numbers := [5]int{1, 2, 3, 4, 5}
	//so the simple syntax for the creation of fixed size array is like it is extreamly simple 
	//it is size of array datatype of array and then elements stored inside the array 
	//but this has a disadvantage like i cannot add more then 5 elemetns so appending is not allowed i
	// the creation of fixed size array
	fmt.Println(numbers)

	number := []int{1, 2, 3, 4, 5}
	fmt.Println(number)
//where as if we just not speicy the size then we r basically creating the dynamic size array like
//we can simply append any elements as we can 
	number = append(number, 6)
	fmt.Println(number)

	number = append(number, 7, 8, 9)
	fmt.Println(number)

	size := len(number)
	fmt.Println("Size of the slice:", size)

	// Now we will create a slice of integers
	// with a length of 5 and a capacity of 10
	numbers2 := make([]int, 5, 10)
	//simply if u want to create array without declaration then make() is used 

	fmt.Println("The capacity of numbers2 is ", cap(numbers2))
	numbers2[0] = 1
	fmt.Println("Numbers2:", numbers2)

	numbers[0] = 0

	fmt.Println("Numbers:", numbers)
//i thing it is much like the slicing like we are creating new array with elements form numbers array
//such that from the 1st index element t the 3 index elemetn like last index is not included
	part := numbers[1:4]
	fmt.Println("Part of the array:", part)

	//now if we change the value of the part slice, it will also change the value
	// of the original array

	part[0] = 10
	fmt.Println("Part of the array:", part)
	fmt.Println("capacity of the part arr is ", cap(part))
	fmt.Println("Numbers:", numbers)

	for index, value := range numbers {
		fmt.Sprintf("the index of numbers array is %d and value is %d\n", index, value)
	}

	structarray := []arrayStruct{
		{Name: "Taha", Age: 20},
		{Name: "Ali", Age: 25},
		{Name: "Ahmed", Age: 30},
	}
	fmt.Println("Struct of array of structure is :", structarray)

	for _, value := range structarray {
		fmt.Printf("The name is %s and age is %d\n", value.Name, value.Age)

	}
}
