package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// Normal index style loop
	//for i := 0; i < len(days); i++ {
	//	fmt.Println(days[i])
	//}

	// for each loop with index
	//for i := range days {
	//	fmt.Println(days[i])
	//}

	// for loop with index and value
	//for index, value := range days {
	//	fmt.Printf("index is: %v, and value is: %v\n", index, value)
	//}

	// while style for loop
	var i = 1

	for i < 10 {
		println("Value is:", i)
		i++

		if i == 2 {
			goto label
		}
	}

	// how to come back from goto ??
label:
	println("Hello")
	println("World")
}
