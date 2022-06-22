package main

import (
	"fmt"
	_ "time"
)

func main() {
	Greet()
	GreetBy("Jiten")
	GreetWith("Hello", "Jiten")

	//----- Return types
	//a,b:= 10,20
	a, p := AreaAndPerimeterOfRect(100.45, 108.74)
	fmt.Println("Area", a, "Perimeter", p)

	a1, _ := AreaAndPerimeterOfRect(100.45, 108.74) // _ is a blank identifier
	fmt.Println("Area", a1)
	_, p1 := AreaAndPerimeterOfRect(100.45, 108.74)
	fmt.Println("Perimeter", p1)

	ai1, pi1 := AreaAndPerimeterOfRecIfc(100.56, 107.46)
	fmt.Println("Area", ai1, "Perimeter", pi1)

	var a2, p2 float32
	a2, p2 = ai1.(float32), pi1.(float32)
	fmt.Println("Area", a2, "Perimeter", p2)
}

func Greet() {
	fmt.Println("Hello World")
}

func GreetBy(name string) {
	fmt.Println("Hello", name)
}

func GreetWith(greet string, name string) {
	fmt.Println(greet, name)
}

// return type

func AreaOfRect(l, b float32) float32 { // unnamed return type
	return l * b
}

func PerimeterOfRect(l float32, b float32) float32 { //named return parameter
	return 2 * (l + b)
}

func AreaAndPerimeterOfRect(l, b float32) (area float32, perimeter float32) {
	area = l * b
	perimeter = 2 * (l + b)
	return area, perimeter
}

func AreaAndPerimeterOfRecIfc(l, b float32) (area interface{}, perimeter interface{}) {
	area = l * b
	perimeter = 2 * (l + b)
	return area, perimeter
}
