package main

import "fmt"

// enumerator is set of constant value that are used all over the project

type Status string

const (

	 ACTIVE Status = "active"
	 INACTIVE Status  = "inactive"
	 BLOCKED Status = "blocked"
) 


func myStatus (status Status) {

	fmt.Println("my status is ", status )
}
func main () {

	myStatus(ACTIVE)

}