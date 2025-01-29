package main

import (
	"awesomeProject/goByMatt/02simpleexample/hello"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(hello.Say(os.Args[1]))
	} else {
		fmt.Println("Hello, World!")
	}
}

//go run main.go
//go run main.go Nitin
