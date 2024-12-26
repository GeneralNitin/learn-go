package main

import "fmt"

const LoginToken = "Hello" // Package level exported variable

func main() {
	var userName string = "Nitin"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n", userName)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.3142857131428571
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var float float64 = 255.3142857131428571
	fmt.Println(float)
	fmt.Printf("Variable is of type: %T \n", float)

	// default values and some aliases
	var someInt int
	fmt.Println(someInt)
	fmt.Printf("Variable is of type: %T \n", someInt)

	// implicit type
	var website = "nitin"
	fmt.Println(website)

	// no var style
	numberOfUser := 391891 // Walrus Operator
	fmt.Println(numberOfUser)
}
