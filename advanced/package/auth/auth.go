package auth

import "fmt"

// Scoping you put the method name lowercase then its mean that method work only in packages, if you put method name Uppercase it means anyone can access it
func Signin (email, password string) {
	fmt.Println("Login successfully", email)
}

func Signup (email, password , firstName , lastName  string) {
	fmt.Println("Signup successfully", email, password, firstName, lastName)
	
}