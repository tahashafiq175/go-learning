package main

import ("fmt")

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("hello world")

	user := User{
		Name: "Taha",
		Age:  13,
	}

	fmt.Println(user)

	user.Age = 20

	fmt.Println(user)

	// for normal variable assignment it should be like something
	// userptr*=&user like we ar storing the address of user inside the
	// pointer userptr but in below line there is no need for *userptr
	// because in go it is inferred automatically //
	userptr := &user

	// same here for changin the value there is need for *userptr
	// because that is dereferance operator like basically pointer
	// holds the value of the address of a particular variable
	// but if we simply do *userptr=something means it is now defereing
	// means refering the value at the address //
	userptr.Name = "Usama"

	//simple analogy &userptr means the address of userptr and
	// so &userptr simply prints the address of userptr not user
	// *userptr means the value at the address of userptr
	// as told before * is deference operator means userptr
	// simply holds the address  of user and *userptr holds the value
	// inside the user//

	fmt.Println(&userptr)
	fmt.Println(userptr)
	fmt.Println(*userptr)

	introduce(userptr)

	Age(*userptr)
	fmt.Println(userptr.Age)
	userptr.changeAge()
	fmt.Println(user.Age)

	fmt.Println("What is age of user like user.age like user.originalAge()")
	user.changeOriginalAge()
	fmt.Println(user.Age)

	user.changeAge()
	fmt.Println(user.Age)
}

func Age(userptr User) {
	userptr.Age = 30
	fmt.Println(userptr.Age)
}
func introduce(userptr *User) {

	userptr.Age = 30
	fmt.Println("hello my name is ", userptr.Name)

	// userptr.changeOriginalAge()
	// fmt.Println(userptr.Age)

}

// simply this is method of the struct and it is not changing the original value of the struct
// becasue it is not a pointer receiver it is a value receiver so it is changing the copy of the struct
// like //
func (u User) changeAge() {
	u.Age = 30
}

// because it is a pointer receiver it is changing the original value of the struct
// like the change inside the struct
// as it has *User it has pointer so it is receiving the address of the struct
// so the changes made to u will affect the original struct //
func (u *User) changeOriginalAge() {
	u.Age = 20
}
