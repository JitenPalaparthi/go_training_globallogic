package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Slice1", slice1)

	slice2 := slice1 // it copies the referece.. both the slices are referenced to the same memory

	slice2[0] = 101
	slice2[1] = 102
	slice2[2] = 103
	fmt.Println("Slice1", slice1)
	fmt.Println("Slice2", slice2)

	// clone a slice

	slice3 := make([]int, len(slice1))

	/*for i := 0; i < len(slice1); i++ {
		slice3[i] = slice1[i]
	}*/
	for i, v := range slice1 {
		slice3[i] = v
	}
	slice3[0] = 1011
	slice3[1] = 1022
	slice3[2] = 1033
	fmt.Println("Slice1", slice1)
	fmt.Println("Slice3", slice3)

	// clone a slice using copy function

	slice4 := make([]int, len(slice1))
	copy(slice4, slice1)
	fmt.Println("Slice4", slice4)
	slice4[0] = 1011
	slice4[1] = 1022
	slice4[2] = 1033
	fmt.Println("Slice4", slice4)
	fmt.Println("Slice1", slice1)

	//

	slice5 := slice1[5:] // from fifth element to the end

	fmt.Println("slice1 from 5th element onwards", slice1[5:])

	fmt.Println("slice5", slice5)

	slice6 := slice1[:5] // from 0th element to the 5th element

	fmt.Println("slice1 from 0th element to the 5th elemtn", slice1[:5])

	fmt.Println("slice6", slice6)

	slice6[0] = 11111
	fmt.Println("slice6", slice6, "slice1", slice1)

	slice7 := slice1[3:7] // from fifth element to the end

	fmt.Println("slice1 from 3rd element to 7th element", slice1[3:7])

	fmt.Println("slice7", slice7)

	arr := [5]int{1, 2, 3, 4, 5}
	// covnert an array to the slice
	slice9 := arr[:] // this is a reference
	slice10 := arr   // this is a value copy
	slice9[0] = 10000
	slice10[0] = 100
	fmt.Println("Array", arr)
	fmt.Println("Slice9", slice9)
	fmt.Println("Slice10", slice10)

	// remove an element from a slice
	slice9 = remove(slice9, 0)
	// [1 2 3 4 5 6 7 8 9 10]
	// slice[:i] =  0 [ 2 3 4 5 6 7 8 9 10]
	// slice[i+1:] =[2 3 4 5 6 7 8 9 10]
	fmt.Println(slice9)

}

func remove(slice []int, i int) []int {
	slice = append(slice[:i], slice[i+1:]...)
	return slice
}

// examples and excercises
// write functions to add , update , delete and get elemetns from a slice
// reverse a slice
// find duplicate elements in a slice
// remove duplicate elements from a slice
// sort a slice // dont use sort package
