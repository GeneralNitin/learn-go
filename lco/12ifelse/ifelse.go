package main

import (
	"fmt"
)

func main() {
	fmt.Println("if else in GoLang")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Normal user"
	} else if loginCount > 10 {
		result = "Power user"
	} else {
		result = "Exactly 10 logins"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num%2 == 0 {
		fmt.Println("Num is even")
	} else {
		fmt.Println("Num is odd")
	}

	//if err != nil {
	//
	//}
}
