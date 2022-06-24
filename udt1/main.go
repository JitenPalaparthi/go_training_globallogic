package main

import "fmt"

type Employee struct {
	Number    int
	Name      string
	Address   string
	Email     string
	ContactNo string
}

func main() {
	var empty struct{}
	fmt.Println("Empty struct", empty)

	var emp Employee
	emp.Number = 1001
	emp.Name = "Jiten"
	emp.Address = "Yeshvantpur, Bangalore"
	emp.Email = "Jitenp@outlook.com"
	emp.ContactNo = "961800000"

	fmt.Println("Employee", emp)

	// shorthand declaration
	emp1 := Employee{}
	emp1.Number = 1002
	emp1.Name = "Rajesh"
	emp1.Address = "Yeshvantpur, Bangalore"
	emp1.Email = "RajeshK@outlook.com"
	emp1.ContactNo = "9612320000"
	fmt.Println("Employee1", emp1)
	// shorthand declartion with values
	emp2 := Employee{Number: 1003, Name: "Rahim", Address: "Vashi, Mumbai", Email: "Rahim@gmail.com", ContactNo: "949580989"}
	fmt.Println("Employee2", emp2)

	// creating pointer type
	//emp6 := &Employee{}
	var emp4 *Employee
	if emp4 == nil {
		//emp4 = emp6
		emp4 = &Employee{}
	}
	(*emp4).Number = 1004 // can also give like this but rarely used. Even without that can be accessed
	emp4.Name = "Bob"
	emp4.Address = "London,UK"
	emp4.Email = "Bob.john@gmail.com"
	emp4.ContactNo = "+441234342"
	fmt.Println("Employee4", emp4)

	// Creating pointer type using new
	emp3 := new(Employee)
	emp3.Number = 1005
	emp3.Name = "Anila"
	emp3.Address = "London,UK"
	emp3.Email = "Anila.Mathew@gmail.com"
	emp3.ContactNo = "+441234332"
	fmt.Println("Employee3", emp3)

	// Shorthand pointer declaration
	emp5 := &Employee{}
	emp5.Number = 1005
	emp5.Name = "Saba"
	emp5.Address = "California,USA"
	emp5.Email = "Saba.Thomas@gmail.com"
	emp5.ContactNo = "+120909990"
	fmt.Println("Employee5", emp5)

}
