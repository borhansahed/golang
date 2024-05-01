package main

import (
	"bufio"
	"fmt"
	"os"
)



func main (){

	fmt.Println("Write you name");

	var name string;
     
	reader := bufio.NewReader(os.Stdin)

	name, _ = reader.ReadString('\n')
	// fmt.Scan(&name)

	fmt.Println("My name is ", name)
}