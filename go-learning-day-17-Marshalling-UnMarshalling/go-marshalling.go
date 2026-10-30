package main

import(
	"fmt"
	"encoding/json"
)

// type User struct{
// 	Name string
// 	Age int
// 	Email string
// }

// func main(){
// 	user:=User{
// 		Name:"Taha Shafiq",
// 		Age:21,
// 		Email:"taha223tata@gmail.com",
// 	}
// 	data,err:=json.Marshal(user)
// 	if err!=nil{
// 		fmt.Println("The err is ",err)
// 		return
// 	}
//  fmt.Println(string(data))
// 	fmt.Println(data)

// 	fmt.Printf("Data type is %T\n",data)
// 	fmt.Printf("Data Type is %T\n",string(data))
// }



// func main(){
// 	jsonData:=[]byte(`{
// 	"Name":"Taha Shafiq",
// 	"Age":21,
// 	"Email":"taha223tata@gmail.com"
// 	}`)
// 	fmt.Println(jsonData)
// 	var user User
// 	err:=json.Unmarshal(jsonData,&user)
// 	if err!=nil {
// 		fmt.Println("The Error is ",err)
// 		return
// 	}
// 	fmt.Println(user.Age)

// }


type User struct{
	Name string `json:"name"`
	Age int `json:"age,omitempty"`
	Email string `json:"email"`
}


func main(){
	user:=User{
		Name:"Taha Shafiq",
		Email:"taha223tata@gmail.com",
	}
	data,err:=json.Marshal(user)
	if err!=nil{
		fmt.Println("The err is ",err)
		return
	}
 fmt.Println(string(data))
	fmt.Println(data)

	fmt.Printf("Data type is %T\n",data)
	fmt.Printf("Data Type is %T\n",string(data))
}
