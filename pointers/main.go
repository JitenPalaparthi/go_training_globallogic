package main

import (
	"fmt"
)

func main() {
	var ptr *int
	num := 100
	fmt.Println("num", num, "ptr", ptr) //nil pointer becase it has not assigned any address
	ptr = &num
	fmt.Println("Value of num", num, "value of num through pointer", *ptr, "address of num", &num, "address os num through pointer", ptr)
	fmt.Println("Address of pointer", &ptr)
	fmt.Println("Value of num through its address", *(&num))
	var ptrOfptr **int

	ptrOfptr = &ptr

	fmt.Println("Pointer of pointer", ptrOfptr)
	fmt.Println("Value of Pointer of pointer ", **ptrOfptr)
	//

	num1 := 100
	fmt.Println("Value before changing", num1)
	changeVal(num1, 200) //all normal values are call by value so that they are passed as values .. a separate copy is maintained
	fmt.Println("after before changing", num1)
	// using pointer
	fmt.Println("Value before changing", num1)
	//ptr2 := &num1
	//changeValptr(ptr2, 200) //all normal values are call by value so that they are passed as values .. a separate copy is maintained
	changeValptr(&num1, 200)
	fmt.Println("after before changing", num1)

	//num2 := new(int)

	//var ptr3 *int
	//*ptr3 = 500 // This gives an error becase the pointer has not been assigned to any address

	ptr3 := new(int)
	fmt.Println("The value of pointer ptr3", *ptr3)

	// var num4 int  // the default value is 0
	// var ptr4 *int // default is nil
	// ptr4 = &num4  // addres of num4 is assigned to ptr4.So ptr4 contains the address of num4

	changeValptr(ptr3, 600)

	/*arr := [3]int{10, 11, 12}
	ptr3 := &arr[1]
	fmt.Println(&arr[0], &arr[1], &arr[2])
	fmt.Println(ptr3)
	fmt.Println(*ptr3)*/

	/*824633819592 +8
	824633819600 + 8
	824633819608
	ptr3++
	ptr3++
	ptr3++*/

	/*var ptr4 uintptr = 0xc0000181d8

	fmt.Println(ptr4)

	var upter unsafe.Pointer*/

}

// If the pointer is not assigned to any address then it is a nil pointer.
// pointer, slice, map , chan can be nil

func changeVal(variable, value int) {
	variable = value
	fmt.Println("Value after changing inside function", variable)
}

func changeValptr(variable *int, value int) {
	*variable = value
	fmt.Println("Value after changing inside function", *variable)
}

// golang does not support pointer arithmetic

//
