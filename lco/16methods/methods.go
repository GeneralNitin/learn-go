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

	//fmt.Println(nitin)

	// formated toString
	//fmt.Printf("Nitin details are: %+v\n", nitin)

	// accessing properties
	fmt.Printf("Nitin's email is: %v\n", nitin.Email)

	nitin.getStatus()
	nitin.updateEmail()
	fmt.Println(nitin)
	nitin.updateEmailWithPtr()
	fmt.Println(nitin)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() {
	fmt.Println("Is User active: ", u.Status)
}

func (u User) updateEmail() {
	u.Email = "test@tets.com"
	fmt.Println("Email updated to:", u.Email)
}

func (u *User) updateEmailWithPtr() {
	u.Email = "test+pointer@tets.com"
	fmt.Println("Email updated to:", u.Email)
}
