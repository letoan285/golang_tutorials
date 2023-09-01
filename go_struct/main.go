package main

import "fmt"

type Engineer struct {
	Name    string
	Age     int
	Project string
}

func main() {

	engineer := Engineer{}

	engineer.Age = 33
	engineer.Name = "Khoa Nguyen"
	engineer.Project = "Urry"

	fmt.Println("Enginer name is: ", engineer.Name)
	fmt.Println("Enginer Age is: ", engineer.Age)
	fmt.Println("Enginer Current Project is: ", engineer.Project)

}
