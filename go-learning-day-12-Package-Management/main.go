package main

import "fmt"
import "go-learning-day-12-Package-Management/auth"
import "go-learning-day-12-Package-Management/user"
import "github.com/fatih/color"

func main(){
	auth.LoginUser("TahaShafiq175","12345678")
	session:=auth.SessionLogin()
	fmt.Println("session : ",session)
	user:=user.User{
		Name: "Taha Shafiq",

		Email: "taha.shafiq@example.com",
	}
	fmt.Println("user : ",user)


	color.Blue("This is a blue message")
}


