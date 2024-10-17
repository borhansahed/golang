package main

import (
	"fmt"

	"github.com/borhansahed/golang/auth"
)



func main () {

	auth.Signup("sahed@gmail.com", "sahed124", 
"Mohammad", "sahed")

	fmt.Print("Hello")
}