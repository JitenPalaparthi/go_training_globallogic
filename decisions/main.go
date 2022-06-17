// Decision making statements in Go
// There is no terinary operatior in Golang;You have to use if else
// If does not have the body (even it is the worst thing) still you must give the scope that is {}

package main

import (
	"math/rand"
)

func main() {
	var age uint8 = 18

	if age >= 18 { // The
		println("Major")
	} else {
		println("Minor")
	}

	// ---------------
	var gender string = "female"

	if age >= 18 && gender == "female" {
		println("She is eligible for marriage according to the law of land")
	} else if age >= 21 && gender == "male" {
		println("He is eligible for marriage according to the law of land")
	} else if gender != "male" && gender != "female" {
		println("no idea")
	} else {
		println("not eligible for marriage")
	}

	// declaring and using conditions
	// a,b:=10,20
	if num1, num2 := rand.Intn(1000), rand.Intn(2000); (num1+num2)%2 == 0 {
		println((num1 + num2), "is an even number")
	} else {
		println((num1 + num2), "is an odd number")
	}
	// fmt.Println(num) this does not work bcz the scope of the number is only confined to if
}

// A simple program to find which is the biggest among a,b,c
// a  b c assign rand numbers
// what is a,b and c are same also check
