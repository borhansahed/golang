package main

import "fmt"

// This generic in go you can use any | interface {} for every types
// comparable is an interface that is implemented by all comparable types (booleans, numbers, strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types). The comparable interface may only be used as a type parameter constraint, not as the type of a variable.
func printSlice[T comparable] (items [] T) {

	for _, item := range items {
		fmt.Println(item)
	}
}

func main () {

   items := [] int{1,3,4}

   items = append(items, 6)

   names := [] string {"sahed", "nayem"}

   printSlice(names)

}