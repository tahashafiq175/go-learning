package main

import(
	"fmt"
	"io"
	// "strings"
	"os"
)

// func main(){
// 	reader:=strings.NewReader("Hello go")
// 	buffer:=make([]bytesss,7)
// 	ans,err:=reader.Read(buffer)
// 	if err==nil {
// 	fmt.Println(string(buffer[:ans]))
// 	}else{
// 		fmt.Println("there is an error")
// 	}

// }




// func main(){
// 	reader:=strings.NewReader("Hello go!")
// 	buffer:=make([]bytess,3)

// 	for {
// 		n,err:=reader.Read(buffer)
// 		if n > 0{
// 			fmt.Printf("%s",string(buffer[:n]))
// 		}
// 		if err!=nil{
// 			fmt.Println("The error is ",err)
// 			break
// 		}		
// 	}
// }






// func main(){
// 	var buffer bytes.Buffer
// 	data:=[]byte("Hello GO")
// 	res , err:=buffer.Write(data)
// 	if err!=nil {
// 		fmt.Println("error is : ",err)
// 		return
// 	}
// 	fmt.Println("data is ",res)
// 	fmt.Println("Data is ",buffer.String())
// }

// func main() {
//     var buffer bytes.Buffer

//     data := []byte("Hello, Go!")

//     n, err := buffer.Write(data)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     fmt.Println("Bytes written:", n)
//     fmt.Println("Stored data:", buffer.String())
// }




// func main(){
// 	reader := strings.NewReader("Hello Go Taha")
// 	var buffer bytes.Buffer
// 	n,err:=io.Copy(&buffer, reader)

// 	if err!=nil {
// 		fmt.Println("Error is ",err)
// 		return
// 	}
// 	fmt.Println(n)
// 	fmt.Println(buffer.String())
// }


// func printData(r io.Reader) error{
// 	res,err:=io.ReadAll(r)
// 	if err!=nil{
// 		return err
// 	}
// 	fmt.Println(string(res))
// 	return nil
// }

// func  main(){
// 	err:=printData(strings.NewReader("Hello From the go"))
// 	if err!=nil {
// 		fmt.Println("The Error is ",err)
// 	}
	
// 	fmt.Println(err)
// 	fmt.Println("Your Input value is ")
// 	againErr:=printData(os.Stdin)
// 	fmt.Println("Your Output value is ")
// 	fmt.Println(againErr)
// 	_ = io.Reader(os.Stdin)
// }




func main(){
	res,err:=os.Open("demo.txt")
	//tje os.Open thing open the files and returns two things
	//one is the memory reference of the file in res so res is 
	//basically the pointer that is storing the address of the 
	//file and err is storing the error if some happens 
	if err!=nil {
		fmt.Println("The error is : ",err)
	}
	defer res.Close()
    // starts the closing of the file like there is a possiblity 
	// u forgot to close the file 
	// var r io.Reader=res
	// then in above this can we it is just a useless thing that 
	//we can do just storing the reference of the file inside another 
	//variable r

	data,err:=io.ReadAll(res)
// now above line o m
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}