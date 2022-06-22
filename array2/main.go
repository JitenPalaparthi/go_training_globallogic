package main

import "fmt"

func changeVal(arr [3]int, index uint, val int) {
	if index < uint(len(arr)) {
		arr[index] = val
	}
	fmt.Println("Array inside func", arr)
}
func main() {
	arr := [3]int{10, 11, 12}
	fmt.Println("Array", arr)
	changeVal(arr, 2, 112)
	fmt.Println("Array after changed", arr)
}
