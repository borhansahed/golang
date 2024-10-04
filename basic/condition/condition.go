package main

import "fmt"

// GO doesn't have ternary operator
func main () {

	// age := 12;

	// if age >= 20 {
	// 	fmt.Println(age)
	// } else if age == 12 {
	// 	fmt.Println(age)
	// }else {
	// 	fmt.Println("age", age)
	// }

 // you can also declared the variable with if 

//    if age := 12; age < 4 {
//       fmt.Println(age)
//    }else {
// 	fmt.Println(age)
//    }


   age:= 12;
   isActive := true;

   if age > 20 && isActive {
	  fmt.Print("Good to go")
   }
}