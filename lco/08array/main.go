package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")

	var fruitList [4]string

	fruitList[0] = "Watermelon"
	fruitList[1] = "Grapes"
	fruitList[3] = "Berries"

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Fruit list length:", len(fruitList))

	var veggies = [3]string{"Beans", "Carrot", "Peas"}
	fmt.Println("Veggies list is:", veggies)
	fmt.Println("Veggies list length:", len(veggies))
}
