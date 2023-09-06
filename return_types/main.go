package main

import "fmt"

type Person struct {
	Name string
	Age  int
}
type Apple struct {
	Title string
	Date  string
	Price int
}

func getValue() (Person, Apple) {
	var myPer Person = Person{Name: "Toan", Age: 12}
	var myApp Apple = Apple{"Epple", "20-009-9", 2000}
	return myPer, myApp
}

func main() {
	myVar, myVar2 := getValue()

	fmt.Println("My Value 1 is: ", myVar)
	fmt.Println("My Value 2 is: ", myVar2)

}
