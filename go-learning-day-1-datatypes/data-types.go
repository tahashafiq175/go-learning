package main

import "fmt"

func datatype() {
	var age int = 19
	var name string = "Taha"
	// this is the way you are declaring the variable along with its own datatype
	fmt.Println(age)
	fmt.Println(name)

	age = 20

	fmt.Println(age)

	//just like other languages go also provide this feature such that there is no need to declare the
	// datatype of the  variable

	var university = "islamia University of Bahawalpur"

	print(university)

	// below line is same as var gpa=3.9
	gpa := 3.9

	fmt.Print(gpa)

	// we can also print the datatype of the variable using the %T format specifier
	fmt.Printf("%T\n", gpa)
	// we cannot do something like that with it like above we have saud that gpa is a float so we c
	// cannot do that like gpa has float datatype so we cannot assign 
	// a string to it like gpa="taha" this will give an error
	gpa = 8.9
	fmt.Print(gpa)

	//multiple variale declaration simultanuously

	var (
		firstName = "taha"
		score     = 100
	)

	fmt.Println(firstName)
	fmt.Println(score)

	// if we want to declare multple variables of same datatype them

	x, y := 10, 20

	fmt.Println(x, "\n", y)

	//swapping logic

	y, x = x, y

	fmt.Printf("%T %T\n", x, y)

	//type conversion from int to float
	new_age := float64(age)

	fmt.Println(new_age)

	new_age = 20.7

	age = int(new_age)

	fmt.Println(age)

}


