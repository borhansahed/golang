package main

import "os"


func main () {


	// Create file
	file, err := os.Create("text.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	 file.WriteString("Hello Golang")



}