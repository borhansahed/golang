package main

import (
	"fmt"
	"slices"
)

func main() {
		// Array with dynamic size 
		var slice []int;

	    slice = append(slice, 4)
	    
	    fmt.Println(cap(slice))

		// make function is used to create a slice with a specified length and capacity
		// make function argument is (type, length, capacity)
		var slice2 = make([]int, 3, 10)
		slice2[0] = 1
		slice2[1] = 2
		slice2[2] = 3
		fmt.Println(slice2)


		names := []string{"John", "Doe", "Jane"};

		names = append(names, "sahed")
		// fmt.Println(names)

		// slice operator
		// slice3 := names[1:3]
		// fmt.Println(slice3)

		
		// 2D slices
		slice4 := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		fmt.Println(slice4)

		
		// slice package
		slice5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		// slice6 := []int{}
		// fmt.Println(slice)
		// fmt.Println(len(slice5))
		// fmt.Println(cap(slice5))

		fmt.Println(slices.Compare(slice5, slice5))


		// copy slice
		slice6 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		slice7 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
		copy(slice6, slice7)
		fmt.Println(slice6)
		fmt.Println(slice7)
}
