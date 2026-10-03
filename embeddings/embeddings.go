package main
import "fmt"


type Car struct{
	Brand string
	Model string
	Year int
}


func (c Car) Start() {
	fmt.Println("Starting the car...")
}

func (p Person) own() {
	fmt.Println(p.Name, "owns the car ")
}

type Person struct{
	Name string
	Age int 
	Car
}


func main(){
	fmt.Println("Hello, World!")
	taha:=Person{
		Name: "Taha Shafiq",
		Age: 20,
		Car: Car{
			Brand: "Toyota",
			Model: "Corolla",
			Year: 2020,
		},

	}
	fmt.Println(taha.Name)
	fmt.Println(taha.Age)
	fmt.Println(taha.Brand)
	fmt.Println(taha.Model)
	fmt.Println(taha.Year)
	taha.Start()
	taha.own()
}