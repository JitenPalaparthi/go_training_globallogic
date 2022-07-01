package main

import "fmt"

func main() {
	defer fmt.Println("From Main:This gets deffered enen though this is called at the beginning")
	fmt.Println("From Main:Calling firstname and lastname function")
	fn := new(string)
	//ln := new(string)
	*fn = "Jiten"
	//*ln = "Palaparthi"
	fullName(fn, nil)
	fmt.Println("From Main: main ends calling without defer")
}

func fullName(firstName, lastName *string) {
	defer recoverFullName()
	defer fmt.Println("FirstName and LastName function ends")
	if firstName == nil {
		panic("runetime error:firstName cannot be nil")
	}
	func() {
		fmt.Println("This is a anon func; scoped in panic")
		if lastName == nil {
			panic("runetime error:lastName cannot be nil")
		}
		defer fmt.Println("This is a anon func; scoped in panic but called at the end")
	}()
	fmt.Println(*firstName, *lastName)
}

func recoverFullName() {
	// r := recover()
	// if r != nil {
	// 	fmt.Println("recovered from ", r)
	// }
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}
