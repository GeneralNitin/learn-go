package main

import . "fmt"
import . "time"

func main() {
	var name string
	var year int

	Println("Enter your name")
	Scan(&name)
	Println("Enter your birth year")
	Scan(&year)

	Println("\nHello", name, "\nYou are", Now().Year()-year, "years old")
}
