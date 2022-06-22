package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr [3]int // length of the array is 3
	// array is a value type.
	// once declared array cannot be extended/ appended or cannot be shrunk
	fmt.Println(arr) // due t type inference the values are given as the default value of an underlining type of an array
	// so values are [0,0,0]

	var arr1 [3]int
	arr1[0] = 100
	arr1[1] = 200
	arr1[2] = 300
	//arr1[3] = 400 since the length of the array is 3, you cannot assign 4th element
	fmt.Println(arr1)
	arr2 := [3]int{10, 11, 12}
	fmt.Println(arr2)

	for i := 0; i < len(arr2); i++ {
		fmt.Print(" ", arr2[i])
	}
	fmt.Println("\nLength of an array", len(arr2))
	for i := 1; i < len(arr2); i++ {
		arr2[i] = 100 + i
	}
	fmt.Println("Values of arr3 after reassiging", arr2)
	// type of an array
	fmt.Println("Type of arr2", reflect.TypeOf(arr2)) // type of an array is [2]int but not []int

	var arr3 = [5]int{10, 11, 13, 100, 200}
	fmt.Println("Type of arr3 and value of arr3", reflect.TypeOf(arr3), arr3) // type of an array is [2]int but not []int
	//arr6 = arr2 // type if arr2 [2]int , type of arr6 is []int
	arr = arr1 // because type of arr is [3]int and type of arr1 is [3]int
	//arr2 = [3]int(arr3) // This operation is not valid

	fmt.Println("Values of arr", arr)
	fmt.Println("Values of arr1", arr1)
	arrCopy(arr3, arr1)
	//
	// Multi dimentional arrays
	fourDarr := [2][2][2][3]int{{{{1, 2, 3}, {3, 4, 5}}, {{11, 12, 13}, {14, 15, 16}}}, {{{1, 2, 3}, {3, 4, 5}}, {{11, 12, 13}, {14, 15, 16}}}}
	fmt.Println("4D array", fourDarr)
	for i1 := 0; i1 < 2; i1++ {
		for i2 := 0; i2 < 2; i2++ {
			for i3 := 0; i3 < 2; i3++ {
				for i4 := 0; i4 < 3; i4++ {
					fmt.Println(fourDarr[i1][i2][i3][i4])
				}
			}
		}
	}
}

func arrCopy(src [5]int, dst [3]int) {
	if len(src) <= len(dst) {
		for i := 0; i < len(src); i++ {
			dst[i] = src[i]
		}
	}
}

// Matrics Multuiplication
// Matrics Addition
// Matrics Substraction
// Do everything using Functions
