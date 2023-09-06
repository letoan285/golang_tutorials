package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Engineer struct {
	Name    string
	Age     int
	Project []Project
}

type Project struct {
	Name         string
	Priority     string
	Technolories []string
}

type People struct {
	Address string
	Level   int
}

var movies []string

type Person struct {
	Firstname string
	Lastname  string
	Email     string
}

func (e Engineer) Print() {
	fmt.Println("Im print Method")
	fmt.Println("My name is", e.Name)
	fmt.Println("My Age is", e.Age)
}

func (e *Engineer) UpdateAge() {
	e.Age += 10
}

func (p *Person) GetEmail() {
	p.Email = strings.ToLower(p.Firstname) + strings.ToLower(p.Lastname) + "@gmail.com"
}

func main() {
	movies = append(movies, "Good to be here")
	movies = append(movies, "THis is good place")
	enginer := Engineer{
		Name: "Le Van Toan",
		Age:  33,
		Project: []Project{
			{Name: "Legrand project", Priority: "Urgent", Technolories: []string{"Go", "Javascript"}},
			{Name: "Smart project", Priority: "Urgent", Technolories: []string{"Dotnet", "Angular"}},
			{Name: "Coca project", Priority: "Middle", Technolories: []string{"React", "Native"}},
		},
	}

	myEnginer := Engineer{
		Name:    "Dao Thi Hien",
		Age:     22,
		Project: []Project{},
	}
	fmt.Printf("%+v\n", enginer)
	fmt.Printf("%+v\n", movies)

	enginer.UpdateAge()

	enginer.Print()
	myEnginer.Print()

	newPerson := Person{
		Lastname:  "Toan",
		Firstname: "Le",
		Email:     "Toan@gmail.com",
	}
	newPerson.GetEmail()

	fmt.Println("My Firstname is: ", newPerson.Firstname)
	fmt.Println("My LastName is: ", newPerson.Lastname)
	fmt.Println("My LastName is: ", newPerson.Email)
	file, err := os.Create("app.js")
	if err != nil {
		panic(err)

	}
	io.WriteString(file, "let x = 3")
}
