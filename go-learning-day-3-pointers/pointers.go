package main

import "fmt"

// CONCEPT 1: STRUCTS
//
// A struct is like a form that contains multiple related
// pieces of information.
//
// Imagine a student form with two fields:
// Name -> stores text
// Age  -> stores an integer
//
// The User struct is a blueprint.
// It defines what information every User value can contain.
//
// IMPORTANT:
// Defining a struct does not create a particular user.
// We create an actual User value later.

type User struct {
    Name string
    Age  int
}

func main() {
    fmt.Println("hello world")

    // CONCEPT 2: A NORMAL VARIABLE
    //
    // We create a variable named age and store 10 in it.
    //
    // Think of age as a box containing the value 10.
    //
    // int means the variable stores an integer.

    var age int = 10

    // CONCEPT 3: POINTERS
    //
    // A pointer is a variable that stores a memory address.
    //
    // &age means:
    // "Give me the memory address of the variable age."
    //
    // *int means:
    // "This variable is a pointer to an int."
    //
    // So ptr stores the address of age, not the value 10 itself.

    var ptr *int = &age

    // Printing ptr displays the address it contains.
    // The exact address varies between program executions.

    fmt.Println(ptr)

    // CONCEPT 4: DEREFERENCING A POINTER
    //
    // *ptr means:
    // "Go to the address stored in ptr and access the value there."
    //
    // Because ptr points to age, *ptr accesses age's value.
    //
    // Here, we replace the value at that address with 20.
    // We are changing age itself through its pointer.

    *ptr = 20

    // age is now 20, even though we assigned through *ptr.

    fmt.Println(*ptr)

    // CONCEPT 5: CREATING A STRUCT VALUE
    //
    // User is the struct type we defined above.
    //
    // We create a user variable and initialize its fields.
    //
    // Name receives "Taha".
    // Age receives 13.
    //
    // This is called a struct literal.

    user := User{
        Name: "Taha",
        Age:  13,
    }

    // CONCEPT 6: A POINTER TO A STRUCT
    //
    // &user gets the memory address of the user variable.
    //
    // usrptr now stores a pointer to that User value.
    //
    // Because we used :=, Go infers the type as *User.
    // The * means pointer, and User is the pointed-to type.

    usrptr := &user

    // CONCEPT 7: DEREFERENCING A STRUCT POINTER
    //
    // *usrptr means the User value stored at the address.
    //
    // fmt.Println prints the struct's field values.
    // It does not print the pointer's address here.

    fmt.Println(*usrptr)

    // CONCEPT 8: AUTOMATIC POINTER DEREFERENCING FOR FIELDS
    //
    // usrptr.Name looks like we are accessing Name directly
    // through a pointer.
    //
    // Go automatically handles the dereferencing for struct fields.
    //
    // usrptr.Name is equivalent to (*usrptr).Name.
    //
    // Both expressions access the Name field of the User value.

    fmt.Println(usrptr.Name)

    // CONCEPT 9: PASSING A POINTER TO A FUNCTION
    //
    // We pass usrptr to change().
    //
    // The change function receives a *User pointer.
    // This allows it to access and modify the original User value.
    //
    // change() changes the name and age, then returns a User value.
    //
    // fmt.Println prints that returned value.

    fmt.Println(change(usrptr))

    // CONCEPT 10: FORMATTED STRINGS
    //
    // fmt.Sprintf builds a string using the supplied values.
    // It does not print the string by itself.
    //
    // %s is replaced by the user's name.
    // %d is replaced by the user's integer age.
    //
    // At this point, change() has already modified user.
    // Therefore, user.Name is "Usama" and user.Age is 20.

    fmt.Println("taha is duffer")

    details := fmt.Sprintf(
        "the user has a name of %s and age of %d",
        user.Name,
        user.Age,
    )

    fmt.Println(details)
}

// CONCEPT 11: A FUNCTION THAT ACCEPTS A POINTER
//
// user *User means the parameter user is a pointer to a User.
//
// When main passes usrptr, this parameter receives the same
// address stored in usrptr.
//
// Therefore, user and usrptr point to the same original struct.
//
// Changing user.Name or user.Age changes the original struct.

func change(user *User) User {
    user.Name = "Usama"
    user.Age = 20

    // CONCEPT 12: RETURNING A STRUCT VALUE
    //
    // *user accesses the User value at the pointed-to address.
    //
    // The return statement returns that struct value.
    // The function's return type is User, not *User.
    //
    // IMPORTANT:
    // The original struct has already been modified.
    // Returning *user gives the caller a copy of its current
    // struct value.

    return *user
}
