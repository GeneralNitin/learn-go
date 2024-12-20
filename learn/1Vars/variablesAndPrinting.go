package main

// Ways to import packages
import . "fmt"
import (
	osAlias "os"
	. "time"
)

func main() {

	// Ways to declare variables
	var fName string = "Nitin"
	var mName = "Kumar"
	lName := "Singh"

	const companyName = "NetApp"
	var companyAge int
	companyAge = Now().Year() - 1992

	Println("Hello", fName, mName, lName)
	Printf("Hello %v %v %v, Welcome to %v\n", fName, mName, lName, companyName) // %v is a placeholder for the value of the variable
	Printf("We are %v years into the game\n", companyAge)

	// Declare multiple variables in a single line
	var a, b, c int = 1, 2, 3
	Println(a, b, c)

	// Declare multiple variables in a single line with different types
	var x, y, z = 1, "two", 3.0
	x1, y1, z1 := 1, "two", 3.0
	Println(x, y, z, x1, y1, z1)

	// Declare multiple variables in a single line with different types
	var (
		name = "Nitin"
		age  = 25
	)
	Println(name, age)

	var osName, _ = osAlias.Hostname()
	Println("OS Hostname:", osName)
}
