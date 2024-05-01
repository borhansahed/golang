package main

import "fmt"

// func returnOne(message string) (string) {
//     return message;
// }
// func returnTwo(message string) (string, int) {
//     return message, 30;
// }

func main (){
	
	// str := returnOne("Hello bhai")
	// str , _= returnTwo("Hello bhai")
	str , _, _:= returnThree("Hello bhai")
	
	fmt.Println(str)
	
}
func returnThree(message string) (string, int , bool) {
	return message+"working", 30, true;
}