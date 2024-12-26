package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greet()

	result := adder(3, 6)
	fmt.Println("Result of addition is: ", result)

	result = multiAdder(1, 2, 3, 4, 5)
	fmt.Println("Result of multiple addition is: ", result)

	result, msg := multiAdderWithMessage(1, 3, 5, 7, 9)

	fmt.Printf("Result of multiple addition is: %v, with message: %v\n", result, msg)
}

func multiAdder(values ...int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}

	return sum
}

func multiAdderWithMessage(values ...int) (int, string) {
	sum := 0
	for _, val := range values {
		sum += val
	}

	return sum, "Hello"
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

func greet() {
	fmt.Println("Namaste from golang")
}
