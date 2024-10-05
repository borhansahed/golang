package main

import "fmt"

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}


	for i, name := range names {
		fmt.Println(i, name)
	}

	class := map[string]string{
		"John": "Math",
		"Paul": "Science",
		"George": "History",
		"Ringo": "English",
	}

	
	for name, subject := range class{
		fmt.Println(name, "is studying", subject) 
	}

	for i, char := range "golang"{
		fmt.Println(i, string(char))
	}
}
