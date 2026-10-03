package main

import "fmt"

type Anmal interface {
	Speak()
	Move()
}

type Cat struct{}
type Dog struct{}

func (c Cat) Speak() {
	fmt.Println("Meow")
}

func (d Dog) Speak() {
	fmt.Println("Woof")
}

func (c Cat) Move() {
	fmt.Println("Runs on all fours")
}

func (d Dog) Move() {
	fmt.Println("Runs on all fours")
}

func makeSound(a Anmal) {
	a.Speak()
}
func main() {

	// cat := Cat{}
	// dog := Dog{}
	// makeSound(cat)
	// makeSound(dog)

	// cat := Cat{}
	// dog := Dog{}
	// cat.Speak()
	// dog.Speak()

	// fmt.Println("hello world")

	EmptyInterface()
}
