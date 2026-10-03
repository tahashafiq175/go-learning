package main

import (
	"errors"
	"fmt"
)

var errorAgeValidator = errors.New("Age is less than 18")

func main() {
	fmt.Println("Hello World")

	//panic("This is panic message")
	fmt.Println("This is after panic message")

	answer:=divide(10, 0)
if(answer==0){
	fmt.Println("Divide by zero error")
}
fmt.Println("Answer: ", answer)

}
func divide(a, b int) int{
	defer func(){
		r:=recover()
		if r!=nil{
			fmt.Println("Recovered from panic: ", r)
		}
	}()
if(b==0){
	panic("Divide by zero error")
    }

	return a/b
}

