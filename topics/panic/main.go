package main

import (
	"fmt"
)

func main() {
	resultChan := make(chan []string, 1)
	results := []string{"result1", "result2"}

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
			resultChan <- results
		}()

		// Uncomment the next line to simulate a panic
		//panic("simulated panic")

		// Normal execution
		fmt.Println("Function executed normally")
	}()

	fmt.Println("Results Chanel:", <-resultChan)
}
