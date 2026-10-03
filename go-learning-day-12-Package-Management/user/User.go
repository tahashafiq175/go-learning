package user

type User struct {
	Name string
	Email string
}


//the logic is simple we are creating a package and we have defined a 
//struct such that whose name or variable name starts with capital lette
//means it is exported and can be used outside of the package
// and same goes with the fields inside the struct if they start with capital letter then they are exported and can be used outside of the package