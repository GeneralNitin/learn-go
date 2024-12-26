package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var list1 = []string{}
	//var list1 []string // alternate syntax
	fmt.Printf("Type of list1 is: %T\n", list1)
	list1 = append(list1, "Hello")
	fmt.Println("Fruit List 1:", list1)

	// create a new list2
	var list2 = []string{"Watermelon", "grapes", "Berries"}
	fmt.Println("Fruit List 2:", list2)

	// add new values to list1, with empty list
	list1 = append([]string{}, "Mango", "Banana", "DragonFruit", "Strawberry")
	fmt.Println("Fruit List 1:", list1)

	// remove first value from list 1 and add to list2
	list2 = append(list1[1:])
	fmt.Println("Fruit List 2:", list2)

	// change the second value in the list
	list2[1] = "XXXXXXX"
	fmt.Println("Fruit List 2:", list2)
	fmt.Println("Fruit List 1:", list1) // It's 2nd index is also affected

	// trim list from 1st index to 2nd index (non inclusive)
	list2 = append(list2[1:2])
	fmt.Println("Fruit List 2:", list2)
	fmt.Println("Fruit List 1:", list1) // It's 2nd index is also affected

	// adding list1 values back to list2
	list2 = append(list1[1:4])
	fmt.Println("Fruit List 2:", list2)
	fmt.Println("Fruit List 1:", list1) // It's 2nd index is also affected

	// this will limit the list 2 till index n-1 i.e. 1 [0,1]
	list2 = append(list2[:2])
	fmt.Println("Fruit List 2:", list2)

	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("Slices using make")

	highScores := make([]int, 4) // type and size

	highScores[0] = 123
	highScores[1] = 789
	highScores[2] = 234
	highScores[3] = 345
	//highScores[4] = 456 // throws index out of bound
	fmt.Println("HighScores:", highScores)
	highScores = append(highScores, 456, 567, 678)
	fmt.Println("HighScores:", highScores)

	fmt.Println("HighScores is Sorted?:", sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("HighScores:", highScores)

	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("Remove an element from a slices")
	courses := []string{"Java", "SpringBoot", "React", "Angular", "Go"}
	fmt.Println(courses)
	var index = 2
	fmt.Printf("Removed %v from the courses\n", index)
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
