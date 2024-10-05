package main

import "fmt"

func main() {

	// Map

	var myMap map[string]int = map[string]int {"a":1, "b":2, "c":3};
	
	myMap["d"] = 4;
  
	v, ok := myMap["z"] // get map value and check if the key is present in the map

	if ok {
		fmt.Println(v);
	} else {
		fmt.Println("Key not found");
	}
	
	delete(myMap, "a"); // delete the element with the key "a"
	clear(myMap); // clear the map means delete all the elements in the map
	fmt.Println(myMap);
	//  myMap["a"] = 1;
	//  myMap["b"] = 2;
	//  myMap["c"] = 3;

	//  fmt.Println(myMap["a"]);
	
}		