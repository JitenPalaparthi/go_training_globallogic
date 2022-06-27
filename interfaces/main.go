package main

import (
	"demo/convert"
	"fmt"
)

func main() {
	var IConvert convert.Converter
	emp1 := &Employee{Number: 12345, Name: "Jiten", Email: "JitenP@Outlook.Com"}
	var num CInt = 1000
	IConvert = emp1
	fmt.Println(IConvert.ToString())
	fmt.Println(IConvert.ToByte())
	IConvert = num
	fmt.Println(IConvert.ToString())

	var sm Something

	IConvert = sm
	fmt.Println(IConvert.ToString())
	fmt.Println(IConvert.ToByte())

}

// nil check can be done on slice, maps , chan , pointers and also interface

type Employee struct {
	Number int
	Name   string
	Email  string
}

func (e *Employee) ToString() string {
	return fmt.Sprintf("Number:%v Name:%v Email: %v", e.Number, e.Name, e.Email)
}

func (e *Employee) ToByte() []byte {
	return []byte(fmt.Sprintf("Number:%v Name:%v Email: %v", e.Number, e.Name, e.Email))
}

type CInt int

func (c CInt) ToString() string {
	return fmt.Sprintf("%v", c)
}

func (c CInt) ToByte() []byte {
	return []byte(fmt.Sprintf("%v", c))
}

type Something struct {
}

func (s Something) ToString() string {
	return "Hello World"
}

func (s Something) ToByte() []byte {
	return []byte("Hello World")
}
