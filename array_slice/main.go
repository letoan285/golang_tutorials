package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       int
	Email     string
}

func (p *Person) GetEmail() {
	p.Email = p.Firstname + p.Lastname + "@gmail.com"

}

func main() {

	var person Person = Person{
		Firstname: "Toan",
		Lastname:  "le",
		Age:       30,
		Email:     "toan@gmail.com",
	}

	person.GetEmail()
	// var x = 4

	// var y = &x
	// var x = 2
	// var y = &x
	// // y = 8

	// fmt.Println("Value of Y", *y)
	// fmt.Println("Value of x", x)

	fmt.Println("My Firs Nane is: ", person.Firstname)
	fmt.Println("My Last Nane is: ", person.Lastname)
	fmt.Println("My  Email is: ", person.Email)
	fmt.Println("My Age is: ", person.Age)

}
