# Go Maps — Line-by-Line Explanation
package main
// package main tells Go that this program is an executable program.
// Go starts running this program from the main() function.
// In this example, mapLearningDemo() contains our map demonstrations,
// so main() must call mapLearningDemo() for these examples to run.

import "fmt"
// We import fmt because we want to print values on the screen.
// fmt.Println() prints values and moves to the next line.
// fmt.Printf() prints formatted output.

// This function demonstrates different ways to use maps in Go.
func mapLearningDemo() {

    // =========================================================
    // 1. CREATING A MAP
    // =========================================================

    age := map[string]int{
        "Taha":  22,
        "Usama": 24,
    }
    // We create a map named age.
    // A map stores information as KEY-VALUE pairs.
    //
    // map[string]int means:
    // string = the type of each key.
    // int    = the type of each value.
    //
    // "Taha" is a key, and 22 is its value.
    // "Usama" is a key, and 24 is its value.
    //
    // Think of a map like a dictionary:
    // Key "Taha"  -> Value 22
    // Key "Usama" -> Value 24
    //
    // We use a key to find its corresponding value.
    // Each key must be unique within the map.

    fmt.Println(age)
    // We print the entire map.
    // The output contains both keys and values.
    // Map output order is not guaranteed.

    // Now we access the value belonging to the key "Taha".

    fmt.Println(age["Taha"])
    // age["Taha"] asks Go to find the value stored under key "Taha".
    // The value is 22.
    // Output: 22


    // =========================================================
    // 2. ADDING AND UPDATING MAP VALUES
    // =========================================================

    age["Ali"] = 23
    // We assign the value 23 to the key "Ali".
    // Because "Ali" does not exist yet, Go creates a new entry.
    // The map now contains Taha, Usama, and Ali.

    age["Taha"] = 21
    // The key "Taha" already exists.
    // Therefore, Go replaces its old value, 22, with 21.
    // It does not create a second "Taha" entry.
    //
    // Important:
    // New key      -> add a new entry.
    // Existing key -> replace its value.

    fmt.Println(age["Taha"])
    // We retrieve Taha's updated value.
    // Output: 21

    fmt.Println("map is ", age)
    // We print the complete map after adding Ali
    // and updating Taha.
    // The order of the entries is not guaranteed.


    // =========================================================
    // 3. DELETING A MAP ENTRY
    // =========================================================

    delete(age, "Ali")
    // delete() removes an entry from a map.
    // First argument: the map we want to change.
    // Second argument: the key we want to remove.
    // Here, we remove the entry whose key is "Ali".

    fmt.Println("Map is ", age)
    // We print the map after deleting Ali.
    // Ali is no longer present.

    delete(age, "Ali")
    // We try to delete Ali again.
    // This is safe.
    // Deleting a key that does not exist does nothing.
    // Go does not produce an error just because the key is missing.


    // =========================================================
    // 4. CHECKING WHETHER A KEY EXISTS
    // =========================================================

    ages := map[string]int{
        "Taha":  12,
        "Usama": 23,
        "Usman": 25,
    }
    // We create another map named ages.
    // It contains three string keys and their integer values.
    // This map is separate from the earlier age map.

    fmt.Println(ages["Usman"])
    // We access the value stored under "Usman".
    // Output: 25

    inAge, exists := ages["Usman"]
    // Go allows us to receive TWO results when reading a map.
    //
    // inAge receives the value associated with the key.
    // exists receives a boolean: true or false.
    //
    // Because "Usman" exists:
    // inAge = 25
    // exists = true
    //
    // This is useful when we need to know whether a key exists,
    // instead of simply receiving its value.

    fmt.Println(inAge)
    // Prints: 25

    fmt.Println(exists)
    // Prints: true
    //
    // If we checked a missing key, such as ages["Ali"],
    // the returned integer would be 0 and exists would be false.
    // The value 0 alone cannot tell us whether the key is missing,
    // because a map could also store an actual value of 0.


    // =========================================================
    // 5. MAP REFERENCE BEHAVIOUR
    // =========================================================

    originalMap := map[string]int{
        "Taha":  21,
        "Usama": 123,
        "Ali":   213,
    }
    // We create a map containing three entries.
    //
    // Go manages the map's internal data for us.
    // A map variable refers to the map data rather than holding
    // a separate, independent copy of every entry.

    dummyMap := originalMap
    // We assign originalMap to dummyMap.
    // This does NOT create an independent copy of all map entries.
    // Both variables refer to the same underlying map data.
    //
    // Think of two labels pointing to the same notebook.
    // Changing an entry through either label changes that same notebook.

    fmt.Println("The output of the dummyMap is ", dummyMap)
    // We print the entries accessible through dummyMap.
    // Initially, they are the same entries as originalMap.

    dummyMap["Taha"] = 32
    // We change Taha's value through dummyMap.
    // Because both variables refer to the same map,
    // the change is also visible through originalMap.

    fmt.Println("The content of original map after mapping changes to the dummyMap is ", originalMap)
    // We print originalMap.
    // Taha's value is now 32, even though we changed it through dummyMap.
    //
    // This demonstrates that assigning one map variable to another
    // does not independently copy the map's entries.


    // =========================================================
    // 6. CREATING AN INDEPENDENT MAP COPY
    // =========================================================

    newMap := make(map[string]int)
    // make(map[string]int) creates a new, initially empty map.
    // Its keys are strings, and its values are integers.
    // Unlike dummyMap, newMap is a separate map.

    fmt.Println("New Map is ", newMap)
    // We print the empty map.
    // Output: map[]

    for key, value := range originalMap {
        // range visits each key-value pair in originalMap.
        // key receives the current key.
        // value receives the corresponding value.
        // Go does not guarantee the order of map iteration.

        newMap[key] = value
        // We insert the current key and value into newMap.
        // For example, if key is "Taha" and value is 32,
        // we add newMap["Taha"] = 32.
        //
        // After the loop, newMap contains the same entries
        // as originalMap, but it is a separate map.
    }

    for key, value := range originalMap {
        // We loop through originalMap again.
        // This time, we print each key-value pair.

        fmt.Printf("the value of the key is %s and %d is \n", key, value)
        // Printf() lets us choose how values are displayed.
        //
        // %s is a placeholder for a string.
        // %d is a placeholder for an integer.
        // key replaces %s, and value replaces %d.
        // \n moves the cursor to the next line.
    }

    fmt.Println(newMap)
    // We print the copied map.
    // Its entries match originalMap at the time of copying.
    //
    // Important:
    // If we later change originalMap, newMap does not automatically
    // receive that change because these are separate maps.
    // This is a shallow copy of the map's key-value entries;
    // if values themselves contain references, those references
    // may still point to shared data.


    // =========================================================
    // 7. PASSING A MAP TO A FUNCTION
    // =========================================================

    changeAge(newMap)
    // We call the changeAge() function and pass newMap to it.
    // The function receives a copy of the map value.
    // However, that copied map value still refers to the same map data.
    // Therefore, the function can update entries in newMap.
    //
    // We do not need a dereference operator (*) to change a map entry.

    fmt.Println(newMap["Taha"])
    // changeAge() changes the value of the "Taha" key to 1233.
    // Because it modifies the same map data, we can see that change here.
    // Output: 1233


    // =========================================================
    // 8. NESTED MAPS
    // =========================================================

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
    // This is a map inside another map.
    //
    // The outer map:
    // Key   = student's name (string).
    // Value = another map containing subjects and marks.
    //
    // The inner map:
    // Key   = subject name (string).
    // Value = marks (int).
    //
    // Think of it as:
    // Taha  -> Math: 90, English: 100
    // Usama -> Math: 98, English: 92
    //
    // Nested maps are useful when one key leads to a collection
    // of related information.

    fmt.Println("The marks of student taha is maths is ", student["Taha"]["Math"])
    // First, student["Taha"] retrieves Taha's inner map.
    // Then ["Math"] retrieves the Math mark from that inner map.
    // Output: 90


    // =========================================================
    // 9. CREATING AND POPULATING A NESTED MAP
    // =========================================================

    newStudent := make(map[string]map[string]int)
    // We create an empty outer map.
    // Each key must be a string.
    // Each value must be another map[string]int.
    //
    // At this point, there are no students in the map.

    fmt.Println("the value of new map is ", newStudent)
    // Prints the empty outer map.

    newStudent["Taha"] = map[string]int{
        "Math":    90,
        "English": 100,
    }
    // We create an inner map and store it under the key "Taha".
    // The inner map contains Math and English marks.
    // Now newStudent has a Taha entry.

    fmt.Println(newStudent["Taha"]["English"])
    // First, retrieve Taha's inner map.
    // Then retrieve the value under "English".
    // Output: 100


    // =========================================================
    // 10. A MAP THAT STORES SLICES
    // =========================================================

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
    // Here, each map value is a slice of strings.
    //
    // map[string][]string means:
    // The key is a string.
    // The value is a slice of strings.
    //
    // Taha's slice contains his subjects.
    // Usama's slice contains his subjects.
    //
    // This structure is useful when each person has multiple values,
    // such as subjects, skills, hobbies, or interests.

    fmt.Println("The slice inside the map is like this : ", mapSlice["Taha"])
    // Retrieves and prints the entire slice associated with Taha.
    // Output: [Math English Urdu]

    fmt.Println("The Slice inside the map is like this and then first index of tha slice is : ", mapSlice["Taha"][0])
    // mapSlice["Taha"] retrieves Taha's slice.
    // [0] retrieves the first element of that slice.
    // Slice indexing starts at zero.
    // Output ends with: Math


    // =========================================================
    // 11. ADDING AN ELEMENT TO A SLICE STORED IN A MAP
    // =========================================================

    mapSlice["Taha"] = append(mapSlice["Taha"], "Database")
    // First, mapSlice["Taha"] retrieves Taha's current slice.
    // append() adds "Database" to the end of that slice.
    // We assign the returned slice back to the map entry.
    //
    // This assignment is important because append() may return
    // a slice with a different length or a new underlying array.
    // The map entry must be updated to store the returned slice.

    fmt.Println("The slice inside the map is like this : ", mapSlice["Taha"])
    // Taha's subjects are now:
    // [Math English Urdu Database]


    // =========================================================
    // 12. A SLICE OF MAPS
    // =========================================================

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
    // This is a slice whose elements are maps.
    //
    // []map[string]int means:
    // The outer collection is a slice.
    // Each element in the slice is a map.
    // Each map has string keys and integer values.
    //
    // The first map stores people and their numbers.
    // The second map stores subjects and their marks.
    //
    // This structure is useful when we need a list of map objects.
    // In real applications, each map in one slice often follows
    // the same data format.

    fmt.Println("The array of Slices is like that ", arrMap[0])
    // arrMap[0] retrieves the first map in the slice.
    // Despite the variable name arrMap, this is a slice, not an array.

    fmt.Println("The array of Slices is like that ", arrMap[0]["Taha"])
    // First, arrMap[0] retrieves the first map.
    // Then ["Taha"] retrieves Taha's value from that map.
    // Output: 23


    // =========================================================
    // 13. STORING STRUCTS INSIDE A MAP
    // =========================================================

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
    // This map stores Student structs.
    //
    // map[string]Student means:
    // The key is a string.
    // The value is a Student struct.
    //
    // Each Student struct contains Name, Age, and Mark.
    // The map lets us retrieve a complete student record using a key.
    //
    // For example:
    // "Taha" -> {Name: "Taha Shafiq", Age: 21, Mark: 35}

    fmt.Println("The structure of the map is like that the age of Taha is : ", structMap["Taha"].Age)
    // structMap["Taha"] retrieves Taha's Student value.
    // .Age accesses the Age field of that struct.
    // Output: 21

    fmt.Println("The Structure of the Map is like that the Name of Taha is : ", structMap["Taha"].Name)
    // Retrieves Taha's struct and accesses its Name field.
    // Output: Taha Shafiq

    fmt.Println("The structure of the map is like that the Mark of Taha is : ", structMap["Taha"].Mark)
    // Retrieves Taha's struct and accesses its Mark field.
    // Output: 35


    // =========================================================
    // 14. MODIFYING A STRUCT STORED IN A MAP
    // =========================================================

    students := structMap["Taha"]
    // We retrieve Taha's struct value from the map.
    // Structs are value types, so this gives us a copy of the struct.
    // The variable students is not a direct editable reference
    // to the struct value stored inside the map.

    students.Name = "Muhammad Ali"
    // We change the Name field in our local struct copy.
    // At this point, the struct stored in structMap has not changed yet.

    structMap["Taha"] = students
    // We put the modified struct back into the map.
    // This replaces the old struct value associated with "Taha".
    //
    // This assignment is necessary because Go does not allow
    // direct field assignment to a struct value inside a map,
    // such as structMap["Taha"].Name = "Muhammad Ali".

    fmt.Println("The value of structure after modification is ", structMap["Taha"])
    // We print Taha's updated struct.
    // Its Name is now Muhammad Ali.


    // =========================================================
    // 15. STORING POINTERS TO STRUCTS IN A MAP
    // =========================================================

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
    // This map stores pointers to Student structs.
    //
    // map[string]*Student means:
    // The key is a string.
    // The value is a pointer to a Student.
    //
    // A pointer identifies where a value is stored in memory.
    // The * before Student means the map values have type *Student.
    //
    // Go allows these struct literals to be used here and creates
    // pointers to the corresponding Student values.
    //
    // Because the map stores pointers, we can modify a student's
    // fields through the pointer.

    fmt.Println("The taha value of mapPointerStruct is : ", mapPointerStruct["Taha"])
    // Prints the pointer's value using Go's default formatting.
    // The output includes the fields and may display pointer-related
    // formatting depending on the value being printed.

    mapPointerStruct["Taha"].Age = 34
    // We access the Student pointer stored under "Taha".
    // Then we update the Age field to 34.
    //
    // Go automatically dereferences the pointer for field access.
    // We do not need to write (*mapPointerStruct["Taha"]).Age.
    // The shorter expression is valid and commonly used.

    fmt.Println("The taha value of mapPointerStruct is : ", mapPointerStruct["Taha"])
    // Prints the Student pointer's value after changing the age.

    fmt.Println("The taha value of mapPointerStruct is : ", *mapPointerStruct["Taha"])
    // The * operator dereferences the pointer.
    // It gives us the Student struct value that the pointer refers to.
    // This prints the struct value rather than just the pointer itself.

    fmt.Println("The taha value Age of mapPointerStruct is : ", mapPointerStruct["Taha"].Age)
    // Accesses the Age field through the pointer.
    // Output: 34


    // =========================================================
    // 16. A MAP CONTAINING STRUCTS WITH SLICES
    // =========================================================

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
    // Here, the map stores values of type Students.
    // Notice that Students is a different struct type from Student.
    //
    // Student contains:
    // Name string
    // Age  int
    // Mark int
    //
    // Students contains:
    // Name string
    // Age  int
    // Mark []string
    //
    // In this second struct, Mark is a slice of strings,
    // so it can hold multiple subject names.
    //
    // Taha's record stores:
    // Name = "Taha Shafiq"
    // Age  = 34
    // Mark = ["English", "Urdu"]

    fmt.Println(mapArrayStruct["Taha"])
    // Retrieves Taha's Students struct from the map.
    // Prints the struct and its fields.
}


// =============================================================
// 17. DEFINING THE STUDENT STRUCT
// =============================================================

type Student struct {
    Name string
    // Name stores the student's name as text.

    Age int
    // Age stores the student's age as an integer.

    Mark int
    // Mark stores one integer value representing a mark.
}
// Student is a custom data type.
// It groups three related pieces of information into one value.
// We can store Student values in variables, slices, or maps.


// =============================================================
// 18. DEFINING A STUDENTS STRUCT WITH A SLICE OF MARKS
// =============================================================

type Students struct {
    Name string
    // Name stores the student's name.

    Age int
    // Age stores the student's age.

    Mark []string
    // Mark is a slice of strings.
    // It can store multiple text values, such as subject names.
}
// Students is a separate type from Student.
// The names are similar, but their field types differ.
// Student.Mark is an int.
// Students.Mark is a []string.


// =============================================================
// 19. CHANGING A MAP VALUE THROUGH A FUNCTION
// =============================================================

func changeAge(age map[string]int) {
    // changeAge receives a map whose keys are strings
    // and whose values are integers.
    //
    // The parameter name age is local to this function.
    // It is not the same variable as the earlier age map
    // simply because they have the same name.
    //
    // A map value is passed to the function by value,
    // but the copied map value refers to the same underlying map data.
    // Therefore, updating a map entry here changes the shared map.

    age["Taha"] = 1233
    // We set the value associated with the key "Taha" to 1233.
    // If the key exists, its value is replaced.
    // If the key does not exist, a new entry is created.
    //
    // The caller can see this change because both map values
    // refer to the same underlying map data.
}


