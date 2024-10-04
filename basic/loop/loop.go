package main

import "fmt"

// Go has only one loop called for loop! you can use this loop in many ways like: while

func main (){

   // while loop
	var i int = 0;

	for i <=3 {
		fmt.Println(i);
		i++;
	}

	// infinite loop

	// for {
	// 	fmt.Printf("infinite loop")
	// }


	// normal for loop
	for i := 0; i<=3; i++ {

		// break;

		// continue
		if i == 2{
			continue
		}
         fmt.Println(i)
	}

	for i:= range(3){
		fmt.Print(i)
	}

}