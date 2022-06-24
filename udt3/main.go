package main

import "fmt"

func main() {

	p := &Person{}
	//p.Number = new(int)
	//*p.Number = 100
	p.Number = 100
	p.Name = "Jiten"
	p.Email = "Jitenp@outlook.com"
	//p.Address = &Address{Line1: "Flat No 202,Maharaj Apartments", City: "Bangalore", State: "Karnataka", Country: "India", PinCode: "560016"}
	addr := &Address{Line1: "Flat No 202,Maharaj Apartments", City: "Bangalore", State: "Karnataka", Country: "India", PinCode: "560016"}
	p.Address = addr

	// These are promoted fields. These are directly called
	p.LinkedIn = "Jpalaparthi"
	p.Twitter = "jiten_1981"
	p.InstaGram = "JpalaparthiInsta"
	p.FaceBook = "Jiten Palaparthi"

	fmt.Println(p)
	fmt.Println(p.Address)
}

type Person struct {
	//Number  *int
	Number      int
	Name        string
	Email       string
	Address     *Address // Embedding the another struct. This is also called as composition
	SocialMedia          // You create field directly with the type then these can be called directly
	// with the Person object
}

type Address struct {
	Line1, City, State, Country, PinCode string
}
type SocialMedia struct {
	//Name string // This does not work as a promoted field becase Name is already there in Person
	LinkedIn, Twitter, FaceBook, InstaGram string
}

func (p *Person) SetPerson(number int, name, email string, address *Address) {
	p.Number = number
	p.Name = name
	p.Email = email
	p.Address = address
}
func (a *Address) SetAddess(line1, city, state, country, pincode string) {
	a.Line1 = line1
	a.City = city
	a.State = state
	a.Country = country
	a.PinCode = pincode
}
