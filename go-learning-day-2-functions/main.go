package main

import "fmt"

func main() {
	fmt.Println("hello world")
	greet()
}

func greet() {
	fmt.Println("greeting from taha")
	calculation(10, 20)
	operation(30, 10)
	sum := returnFunc(100, 30)
	add, multiply := returnFuncs(20, 50)
	fmt.Println("calling the sum function")
	fmt.Println(sum)
	fmt.Println("calling the add with one return type")
	fmt.Println(add)
	fmt.Println(multiply)
	details := modReturnFuncs(10, "Taha")
	fmt.Println(details)
}
