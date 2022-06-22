package main

import "fmt"

func main() {

	slice1 := make([]int, 10) // type inference has given [0 0 0 0 0 0 0 0 0 0]
	//slice2[0] = 1
	slice1 = append(slice1, 1)
	// [0 0 0 0 0 0 0 0  0 0 1]

	var slice2 []string

	slice2 = append(slice2, "Hello!")

	slice2 = append(slice2, "How", "Are", "You?", "Doing") // variadic parameter

	fmt.Println("Slice2:", slice2)

	// for range loop
	// range loop in general returns two parameters
	// index , value for slices
	// key , value for maps
	// For channel it returns only one parameter
	// val for channels
	// for i:=0; i<10; i= i+1 {}
	//for i, val := range slice2 {
	//for i, _ := range slice2 {
	for _, val := range slice2 {
		//fmt.Println("Index", i, "Value", val)
		fmt.Println("Value", val)
		//fmt.Println("Index", i)
	}

	// fmt.Println("How", "Are", "You", "Doing")

	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Slice3", slice3)
	changeVal(slice3, 2, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300)
	changeVal(slice3, 5, 606)
	fmt.Println("slice3", slice3)
	vslice := []int{300, 400, 500, 606, 707, 808, 909, 1001, 1100, 1200, 1300}
	changeVal(slice3, 2, vslice...) // automatically vslice is converted as individual parameters by using ... to the end of the slice
	fmt.Println("slice3", slice3)

}

// variadic parameter and function
// the variadic parameter must be the last parameter
// the variadic paramter is defined using ... eclipse symbol
func changeVal(slice []int, sindex int, vals ...int) {
	if slice != nil {
		if sindex < len(slice) && len(vals) >= 1 {
			for i, j := sindex, 0; i < len(slice) && j < len(vals); i, j = i+1, j+1 {
				slice[i] = vals[j]
			}
		}
	}
}
