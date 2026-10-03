package main

import "fmt"

func mapLearningDemo() {
	age := map[string]int{
		"Taha":  22,
		"Usama": 24,
	}
	fmt.Println(age)
	//in this approach we are simply printing the value corresponding to key of "Taha"
	fmt.Println(age["Taha"])

	//simple logic is if the key exists just replace its value
	age["Ali"] = 23
	//simple if key does not exists donot replace its value but create new entry in map
	age["Taha"] = 21
	fmt.Println(age["Taha"])

	fmt.Println("map is ", age)
	//simply delete the the entry Ali from the map
	delete(age, "Ali")
	fmt.Println("Map is ", age)
	delete(age, "Ali")

	ages := map[string]int{
		"Taha":  12,
		"Usama": 23,
		"Usman": 25,
	}
	fmt.Println(ages["Usman"])
	inAge, exists := ages["Usman"]
	fmt.Println(inAge)
	fmt.Println(exists)

	//going advances with maps

	//map reference behaviour

	originalMap := map[string]int{
		"Taha":  21,
		"Usama": 123,
		"Ali":   213,
	}

	//much like array map in go is also reference type like it basically stores
	//the address of the map or reference like memory reference
	//nut when when we print it go automatically defernce it
	//to print the value stored inside the map so if make or create another map like the line below the changes in new map reflects in prvious ones
	dummyMap := originalMap
	fmt.Println("The outout of the dummyMap is ", dummyMap)
	dummyMap["Taha"] = 32
	fmt.Println("The content of original map after mapping changes to the dummyMap is ", originalMap)

	//to make new map that is not referring to the same map belwo is the syntax

	newMap := make(map[string]int)

	fmt.Println("New Map is ", newMap)
	for key, value := range originalMap {
		newMap[key] = value
	}

	for key, value := range originalMap {
		fmt.Printf("the value of the key is %s and %d is \n", key, value)
	}
	fmt.Println(newMap)

	//12 functions and maps
	//as maps store the reference so we can change the
	//value of the map using function without writing or using
	// dereference operator//

	changeAge(newMap)
	
	fmt.Println(newMap["Taha"])

	//13 Nested Map

	student := map[string]map[string]int{
		"Taha": {
			"Math":    90,
			"English": 100,
		},
		"Usama": {
			"Math":    98,
			"English": 92,
		},
	}

	fmt.Println("The marks of student taha is maths is ", student["Taha"]["Math"])

	newStudent := make(map[string]map[string]int)
	fmt.Println("the value of new map is ", newStudent)
	newStudent["Taha"] = map[string]int{
		"Math":    90,
		"English": 100,
	}
	fmt.Println(newStudent["Taha"]["English"])
	//map and slicing

	mapSlice := map[string][]string{
		"Taha": {
			"Math",
			"English",
			"Urdu",
		},
		"Usama": {
			"History",
			"English",
			"Science",
		},
	}
	fmt.Println("The slice inside the map is like this : ", mapSlice["Taha"])
	fmt.Println("The Slice inside the map is like this and then first index of tha slice is : ", mapSlice["Taha"][0])

	//Adding value inside the slice inside the map
	mapSlice["Taha"] = append(mapSlice["Taha"], "Database")
	fmt.Println("The slice inside the map is like this : ", mapSlice["Taha"])

	//Array or Slice of Maps

	arrMap := []map[string]int{
		{
			"Taha":  23,
			"Usama": 32,
		},
		{
			"Math":    90,
			"English": 95,
		},
	}

	fmt.Println("The array of Slices is like that ", arrMap[0])
	fmt.Println("The array of Slices is like that ", arrMap[0]["Taha"])

	//Struct inside the Map

	structMap := map[string]Student{
		"Taha": {
			Name: "Taha Shafiq",
			Age:  21,
			Mark: 35,
		},
		"Usama": {
			Name: "Usama Shafiq",
			Age:  25,
			Mark: 100,
		},
	}

	fmt.Println("The structure of the map is like that the age of Taha is : ", structMap["Taha"].Age)
	fmt.Println("The Structure of the Map is like that the Name of Taha is : ", structMap["Taha"].Name)
	fmt.Println("The structure of the map is like that the Mark of Taha is : ", structMap["Taha"].Mark)

	//Modification for values inside the Map

	students := structMap["Taha"]
	students.Name = "Muhammad Ali"
	structMap["Taha"] = students

	fmt.Println("The value of structure after modification is ", structMap["Taha"])

	//if u want to modify the content of map with struct value follow below syntax

	mapPointerStruct := map[string]*Student{
		"Taha": {
			Name: "Taha Shafiq",
			Age:  23,
			Mark: 32,
		},
		"Usama": {
			Name: "Usama Shafiq",
			Age:  25,
			Mark: 34,
		},
	}

	fmt.Println("The taha value of mapPointerStruct is : ", mapPointerStruct["Taha"])

	mapPointerStruct["Taha"].Age = 34

	fmt.Println("The taha value of mapPointerStruct is : ", mapPointerStruct["Taha"])
	fmt.Println("The taha value of mapPointerStruct is : ", *mapPointerStruct["Taha"])
	fmt.Println("The taha value Age of mapPointerStruct is : ", mapPointerStruct["Taha"].Age)

	mapArrayStruct := map[string]Students{
		"Taha": {
			Name: "Taha Shafiq",
			Age:  34,
			Mark: []string{
				"English",
				"Urdu",
			},
		},
	}

	fmt.Println(mapArrayStruct["Taha"])
}

type Student struct {
	Name string
	Age  int
	Mark int
}

type Students struct {
	Name string
	Age  int
	Mark []string
}

func changeAge(age map[string]int) {
	age["Taha"] = 1233
}
