package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("111")
	defer fmt.Println("222")
	fmt.Println("Hello")

	myDeferredFunction()
}

func myDeferredFunction() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
