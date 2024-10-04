package main

import "fmt"

// We can define const variable here
const GLOBAL string = "global"

var yo  string = "sahed"

func main() {

	// Variable with type
	var name string
    name = "sahed"

	var roll int = 2
	
	var discount float32 = 40.4

	var isActive bool = true

	// Variable without type

	var withoutType = 4;

	// Shorthand variable
	shorthand := "shortand string variable"

	// const with type
	const  name1 string  = "sahed"

	// const without type
	const  name3 string  = "sahed"

	// Multiple const at once
	const (
		PORT = 400;
		DEVICE = "MAC"
	)

	fmt.Println(name)
	fmt.Println(roll)
	fmt.Println(discount)
	fmt.Println(isActive)
	fmt.Println(withoutType)
	fmt.Println(shorthand)
	fmt.Println(PORT)
	fmt.Println(yo)
}