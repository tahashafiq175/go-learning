package main
import "fmt"
type User struct{
Name string
Age int
}
func main(){
	fmt.Println("hello world")

	var age int=10
	var ptr *int=&age
	fmt.Println(ptr)

	*ptr=20
	fmt.Println(*ptr)

user:=User{
Name:"Taha",
Age:13,
}


usrptr :=&user


fmt.Println(*usrptr)

fmt.Println(usrptr.Name)

fmt.Println(change(usrptr))



fmt.Println("taha is duffer")
details:=fmt.Sprintf("the user has a name of %s and age of %d ",user.Name,user.Age)


fmt.Println(details)

}

func change(user *User) User{
user.Name="Usama"
user.Age=20
return *user
}