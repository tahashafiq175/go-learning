package main
import (
	"fmt"
	"net/http"
)


// func main(){
// 	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
// 		fmt.Fprintf(w, "Hello, World!")
// 		fmt.Println(r.Method)
// 		fmt.Println(r.URL.Path)
// 	})
// 	http.ListenAndServe(":8080", nil)		
// }


// func getAll(w http.ResponseWriter,r *http.Request){
// 	fmt.Fprintf(w, "Get All")
// }

// func postOne(w http.ResponseWriter,r *http.Request){
// 	fmt.Fprintf(w, "Post One")
// }

// func putOne(w http.ResponseWriter,r *http.Request){
// 	fmt.Fprintf(w, "Put One")
// }

// func deleteOne(w http.ResponseWriter,r *http.Request){
// 	fmt.Fprintf(w, "Delete One")
// }
// func main(){
// 	http.HandleFunc("/get",getAll)
// 	http.HandleFunc("/post",postOne)
// 	http.HandleFunc("/put",putOne)
// 	http.HandleFunc("/delete",deleteOne)
// 	http.ListenAndServe(":8080", nil)
// }


func getAll(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Get All")
}


func main(){
	http.HandleFunc("/get",getAll)
	http.ListenAndServe(":8080", nil)
}