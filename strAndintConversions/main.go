package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var str rune = 'A' // char
	//var name string = "Jiten"
	fmt.Println(int(str))
	var num int = 65
	fmt.Println(string(num))
	// "65"

	// options-1
	nums := fmt.Sprintf("%v", num)
	fmt.Println("Value of nums", nums, "Type of numbs", reflect.TypeOf(nums))

	numstr := strconv.Itoa(num)
	fmt.Println("Value of numstr", numstr, "Type of numstr", reflect.TypeOf(numstr))

	num2, _ := strconv.Atoi("65")
	fmt.Println("num2 using atoi", num2)
}
