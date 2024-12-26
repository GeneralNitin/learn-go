package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	nitin := User{
		"Nitin",
		"nitin@nitin.com",
		true,
		24,
	}

	fmt.Println(nitin)

	// formated toString
	fmt.Printf("Nitin details are: %+v\n", nitin)

	// accessing properties
	fmt.Printf("Nitin's email is: %v", nitin.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
