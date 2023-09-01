package main

import "fmt"

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

func (e Engineer) Print() {
	fmt.Println("Im print Method")
	fmt.Println("My name is", e.Name)
	fmt.Println("My Age is", e.Age)
}

func (e *Engineer) UpdateAge() {
	e.Age += 10
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
}
