package main

import "fmt"

func main() {

	var myMap map[string]string
	myMap = make(map[string]string, 1) // length is 1 it doesnot mean that you can give only 1 element

	myMap["522316"] = "Guntur"
	myMap["560086"] = "yeshvantpur Bangalore"
	myMap["560096"] = "Mahalakshmi Layout, Bangalore"
	myMap["560031"] = "Rajaji Nagar,Bangalore"

	fmt.Println(myMap)

	for key, val := range myMap {
		fmt.Println("Key", key, "Value", val)
	}
	// There is not cap to the map , cap(myMap) <-- this does not work
	fmt.Println("Length of the myMap", len(myMap))
	delete(myMap, "560096") // delete to delete an element from the map provided by map and the key
	fmt.Println("After deleting the key")
	fmt.Println("Length of the myMap", len(myMap))
	for key, val := range myMap {
		fmt.Println("Key", key, "Value", val)
	}

	val, ok := myMap["111111"] // it returns two things the value of the map , whether the value for the key ther or not
	// if key is not present ok is false; if key is not there that means there is no value as well
	if ok {
		fmt.Println("Can update an element to the map, becasue there is a key ")
	} else {
		fmt.Println("Cannot update an element but can insert an element to the map, becase ther is not key")
	}
	fmt.Println(ok, val)
}
