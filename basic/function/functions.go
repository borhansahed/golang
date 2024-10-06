package main

import "fmt"

// function declaration
func greet(name string) string {
	return "Hello, " + name
}

// function return multiple values
func getPerson() (string, int) {
	return "John", 20
}

// function return another function
func getPerson2() func(str string) string {
	return func(str string) string  {
		return str
	}
}

func main() {
	fmt.Println(greet("John"))
	fmt.Println(getPerson())
	fmt.Println(getPerson2()("Shaed"))
}
