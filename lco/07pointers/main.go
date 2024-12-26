package main

import "fmt"

func main() {
	fmt.Println("Class on pointers")

	var ptr *int
	fmt.Println("Value of pointer is", ptr)
	//fmt.Println("Value of pointer is", *ptr) // throws null pointer - panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Println("Value of pointer is", &ptr)

	//*ptr = 30 // not allowed
	//fmt.Println("Value of pointer is", *ptr)

	number := 90
	fmt.Println("Value of ptr2 is", number)
	fmt.Println("Value of ptr2 is", &number)

	ptr = &number
	fmt.Println("Value of pointer is", ptr)
	fmt.Println("Value of pointer is", *ptr)
	fmt.Println("Value of pointer is", &ptr)

	*ptr = *ptr + 10
	fmt.Println("New value is", number)
}
