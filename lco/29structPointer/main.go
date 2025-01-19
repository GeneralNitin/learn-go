package main

import (
	"fmt"
)

func main() {
	obj1 := MyStruct{
		"Nitin",
		Address{},
		1,
	}

	obj2 := MyStructWithInnerRef{
		"Nitin",
		nil,
		1,
	}

	if &obj1.Address == nil {
		fmt.Println("obj1.Address is nil")
	} else {
		fmt.Println("obj1.Address is not nil")
	}

	if obj2.Address == nil {
		fmt.Println("obj2.Address is nil")
	} else {
		fmt.Println("obj2.Address is not nil")
	}
}

type MyStructWithInnerRef struct {
	Name    string
	Address *Address
	Age     int
}

type MyStruct struct {
	Name    string
	Address Address
	Age     int
}

type Address struct {
	Line1 string
	Line2 string
	Pin   int
}
