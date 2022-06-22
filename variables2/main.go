package main

import "fmt"

func main() {

	var num int8 = 10

	var num1 = 100

	// var num2, num3, isNum = 100.45, "200", false
	var (
		num2, num3, isNum = 100.45, "200", false
	)

	numStr := "Number is 1000" // shorthand notation of creating a variable by assigning a value

	// numSr1 // This is not valid becase you mush assign a value to create a variable in shorthand notation

	fmt.Printf("Type of num %T value of num %v\n", num, num)
	fmt.Printf("Type of num1 %T value of num1 %v\n", num1, num1)
	fmt.Printf("Type of num2 %T value of num2 %v\n", num2, num2)
	fmt.Printf("Type of num3 %T value of num3 %v\n", num3, num3)
	fmt.Printf("Type of isNum %T value of isNum %v\n", isNum, isNum)
	fmt.Printf("Type of numStr %T value of numStr %v\n", numStr, numStr)

	{
		var a, b, c = 10, 20, 30
		/*var t = b
		b = a
		a = t*/
		// simple swapping in go

		console.log("Before swap a,b,c", a, b, c)
		a, b, c = c, a, b

		console.log("After swap a,b,c", a, b, c)
	}

	// Type Casting
	var long int64
	var small int32 = 32345

	// long = small // This is not possible in go bease long and small int64 and int32.
	// According to go type system, it can assign only the same type
	// fmt.Println(long)

	long = int64(small) // type casting in go

	fmt.Println("Long and small", long, small)

	long1 := 10003 // type is int (even int is of 8 bytes on 64 bit arch)

	var small1 int32 // type is int32
	small1 = int32(long1)
	fmt.Println("Long1 and small1", long1, small1)

	long1 = 2000

	num = 127

	fmt.Println("Mutation of long1 and num", long1, num)

}
