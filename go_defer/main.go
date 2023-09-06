package main

import "fmt"

type Person struct {
	Name string
	Body MyBody
}

type MyBody struct {
	Foot int
	Skin string
}

var per Person = Person{
	Body: MyBody{Skin: "Yellow"},
}

func main() {
	fmt.Println("my name is: ", per.Name)
	fmt.Println("my Skin is: ", per.Body.Skin)

}
