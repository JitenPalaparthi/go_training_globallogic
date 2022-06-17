// Empty interface{} and complex type
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num interface{} // empty interface

	var num1 int

	var num2 = 200

	num = 101

	// num1 = int(num) // interface{} type is assigned to the int type; type casting does not work with empty interfaces{}
	// type casting work very well for defined types
	// for empty interface{} , it is called type assertion
	// interface{}.(type)

	num1 = num.(int) // type assertion applies only for interface{}

	fmt.Println("Additon of num+num2", num1+num2)

	fmt.Println("Num is ", num, " type of num ", reflect.TypeOf(num))

	num3 := num.(int)
	fmt.Println("Additon of num3+num2", num3+num2)

	num = 101.11

	fmt.Println("Num is ", num, " type of num ", reflect.TypeOf(num))

	num = "101.11f"

	fmt.Println("Num is ", num, " type of num ", reflect.TypeOf(num))

	num = false

	fmt.Println("Num is ", num, " type of num ", reflect.TypeOf(num))

	// + - * / % ++ --
	// && || != ==
	// >> << & |

	// complex numbers
	cmplx1 := complex(1000.343, 321.12312)
	fmt.Println("Complex number ", cmplx1, "type of complex number", reflect.TypeOf(cmplx1))

	var r1, i1 float32 = 1000.343, 321.12312
	//var i2 float64 = 321.12312
	cmplx2 := complex(r1, i1)

	fmt.Println("Complex number ", cmplx2, "type of complex number", reflect.TypeOf(cmplx2))

	cmplx3 := complex(1000.345, i1)

	fmt.Println("Complex number ", cmplx3, "type of complex number", reflect.TypeOf(cmplx3))

	cmplx4 := cmplx3 + complex64(cmplx1)

	fmt.Println("Complex number ", cmplx4, "type of complex number", reflect.TypeOf(cmplx4))

	cmplx5 := cmplx3 - complex64(cmplx1)

	fmt.Println("Complex number ", cmplx5, "type of complex number", reflect.TypeOf(cmplx5))

	cmplx6 := cmplx3 * complex64(cmplx1)

	fmt.Println("Complex number ", cmplx6, "type of complex number", reflect.TypeOf(cmplx6))

	// constants

	const MIN int = 9

	const MAX = 999

	const (
		MAX2 = 999
		MAX3 = 9999
		MIN2 = 99
	)

	var a = (100 * 2) / 2

	//const A = a // cannot assign normal variable to the const

	fmt.Println("Constant value and normal value", MAX, a)

	const MAX1 = (100 * 9) / MIN

	fmt.Println("Constant value", MAX1)

}
