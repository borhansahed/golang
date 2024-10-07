package main

import "fmt"

// Here Person Uppercase means it is a public struct and lowercase means it is a private struct.
type Person struct {
	Name string
	Age  int
	Wife *Person

}
// Method is a function that is associated with a type.
func (p Person) PrintName() {
	fmt.Println(p.Name)
}

// method to write the value of the struct
func (p *Person) WriteName(name string) {
	p.Name = name
}

func main() {
    wife := Person{Name: "Who", Age: 22}
	person := Person{Name: "John", Age: 23, Wife: &wife}
	fmt.Println(person.Wife.Name)
	person.PrintName()
	person.WriteName("WhoUpdated")
	wife.WriteName("WifeUpdated")
   fmt.Println(wife)
	// struct in one line

	language := struct {
		Name string
		Year int
	}{Name: "Go", Year: 2009}
	fmt.Println(language)
}
