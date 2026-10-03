package main

import "fmt"

func EmptyInterfaceFunction(value any) {
	numbers := value
	fmt.Println(value)
	fmt.Println(numbers)

	switch v := value.(type) {
	case int:
		fmt.Println("Integer:", v*2)
	case float64:
		fmt.Println("Float:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Boolean:", v)
	default:
		fmt.Println("Unknown type")
	}

}


func EmptyOkInterface(value any) {
	integer, ok := value.(int)
	if ok {
		fmt.Println("Integer:", integer*2)
	} else {
		fmt.Println("Not an integer")
	}

}



func EmptyInterface() {
	EmptyInterfaceFunction("taha")
	EmptyInterfaceFunction(10)
	EmptyInterfaceFunction(10.5)
	EmptyInterfaceFunction(true)

	EmptyOkInterface(10)
	EmptyOkInterface("taha")
}
