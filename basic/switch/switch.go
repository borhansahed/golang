package main

import (
	"fmt"
)

func main() {

	// In golang we don't need to use break statement

	// Normal switch
	// day := "Monday"

	// switch day {
	// case "Monday":
	// 	fmt.Println("Monday")
		
	// }

	// Switch with multiple case
	// switch time.Now().Weekday() {
	// case time.Friday, time.Saturday:
	// 	fmt.Println("Weekday")
	// default:
	// 	fmt.Println("workday")
	// }

	// type switch

	// interface{} is a type in golang that can hold any value
	var i interface{} = 40.4

	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}

	// switch  with expression
	switch i := 10; i {
	case 10:
		fmt.Println("10")
	case 20:
		fmt.Println("20")
	default:
		fmt.Println("unknown")
	}

	// switch without condition
	switch i := 10; {
	case i > 10:
		fmt.Println("i > 10")
	case i < 10:
		fmt.Println("i < 10")
	default:
		fmt.Println("i == 10")
	}




		

}
