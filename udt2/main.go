package main

import (
	"fmt"
	"udt/shape"
	"udt/shape/rect"
	"udt/shape/square"
)

// Create a separate package called Shape
// Perform operations like Area and Perimeter etc
// Access that pacakge in the main function
func main() {

	shape.Infomation()
	fmt.Println()

	rct := rect.Rect{Length: 123.45, Bredth: 146.75}

	area := rect.AreaOfRect(rct) // this contains its own copy of rect

	fmt.Println("Area of rect", area)

	perimeter := rect.PerimeterOfRect(rct)

	fmt.Println("Perimeter of rect", perimeter)

	fmt.Println("Area of Rect that is stored inside rect", rct.Area)
	fmt.Println("Perimeter of Rect that is stored inside rect", rct.Perimeter)
	fmt.Println()
	rectPtr := &rect.Rect{Length: 123.45, Bredth: 146.75}

	area1 := rect.AreaOfRectPtr(rectPtr) // this contains its own copy of rect

	fmt.Println("Area of rect", area1)

	perimeter1 := rect.PerimeterOfRectPtr(rectPtr)

	fmt.Println("Perimeter of rect", perimeter1)

	fmt.Println("Area of Rect that is stored inside rectPtr", rectPtr.Area)
	fmt.Println("Perimeter of Rect that is stored inside rectPtr", rectPtr.Perimeter)
	fmt.Println()
	rect1 := &rect.Rect{}

	rect1.Length = 12.45
	rect1.Bredth = 15.67

	fmt.Println("Area Of Rect", rect1.AreaOf())
	fmt.Println("Perimter of Rect", rect1.PerimeterOf())
	fmt.Println("Area of Rect that is stored inside rect", rect1.Area)
	fmt.Println("Perimeter of Rect that is stored inside rect", rect1.Perimeter)

	fmt.Println("Rect unwanted method just to say that there is a method", rect1.Add(10, 20))
	fmt.Println()
	sq := new(square.Square)

	sq.Side = 25.27

	fmt.Println("Area Of Square", sq.AreaOf())
	fmt.Println("Perimter of Square", sq.PerimeterOf())
	fmt.Println("Area of Square that is stored inside rect", sq.Area)
	fmt.Println("Perimeter of Square that is stored inside rect", sq.Perimeter)
	fmt.Println()
}
