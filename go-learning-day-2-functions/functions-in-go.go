package main

import "fmt"

// it is a simple function with parameters the syntax is
// func functionName (variable name datatype,variable datatype){
// function body
// }  //

func calculation(a int, b int) {
	fmt.Println(a + b)
}

// it is a another approach of writing a method but the parameters of both the variable is same//

func operation(a, b int) {
	fmt.Println(a + b)
}


// it is a function with return type the syntax is
// func functionName (variable name datatype,variable datatype) returnType{
// function body
// }  //

//if we does not return anything from the function then we can write the 
// function like this
func returnFunc(a, b int) int {
	return a + b
}


// it is a function with multiple return type the syntax is
// func functionName (variable name datatype,variable datatype) (returnType1, returnType2){
// function body
// }  //
func returnFuncs(a, b int) (int, int) {
	sum := a + b
	difference := a * b
	return sum, difference
}


func manyReturnFuncs(a, b int) (int, int) {
	return a + b, a * b
}

//in this approah we are using named return type such that we are giving the name to the return type
//like if we are just write return the variable that is declare with function
//declaration got retunred automatically
func modReturnFuncs(age int, name string) (difference string) {
	difference = fmt.Sprintf("My name is %s with age of %d ", name, age)
	return
}
