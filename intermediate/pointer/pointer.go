package main

import "fmt"

func changeValue(num *int) int {
	*num = 100
	return *num
}

// In go language, the pointer is a variable that stores the memory address of another variable.
// But Slice also passed by value in the function argument. 
// Map also passed by value in the function argument.

func main() {
	num := 10

	changeValue(&num)
	fmt.Println(num)
}
