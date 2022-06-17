package main

import (
	"demo/vars"
	"fmt"
	"time"
)

// // global variables. can be used any where in main package
// var (
// 	number    int    = 100
// 	name      string = "Jiten"
// 	isMarried bool   = false
// )

// to run or build go main package in multiple files , you need to give all main package file banes

// go run main.go vars.go
func main() {
	fmt.Println("Hello World")
	fmt.Println(time.Now())

	// integers --> int,int8,int16,int32,int64,uint8,uint16,uint32,uint64
	// floats   --> float32, float64
	// bools    --> bool
	// strings  --> string
	// byte     --> byte
	// rune     --> rune

	var num int
	fmt.Println("number is ", num)

	var num1 int8 = 10
	fmt.Println("int8 number is ", num1)

	var decimal float32
	fmt.Println("float number is ", decimal)

	var ok bool
	fmt.Println("bool value is  ", ok)

	var bt byte
	fmt.Println("byte value is  ", bt)

	var rn rune
	fmt.Println("rune value is  ", rn)
	var str string

	fmt.Println("Value of string", str)

	// type inference
	// int is asisgned 0
	// float is assigned 0
	// bool is assign false
	// string is assigned ""

	fmt.Printf("Number: %d, Name: %s Married: %v \n", vars.Number, vars.Name, vars.IsMarried)

}

// go.dev
// effective go
