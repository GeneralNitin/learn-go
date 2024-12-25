package main

import (
	"awesomeProject/projectTicketBook/helper"
	. "fmt"
	"strconv"
	"strings"
)

// package level variable can't use the short-hand syntax
// conferenceName := "Go Conference"
var conferenceName = "Go Conference"

const conferenceTicket int = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0) // dynamic array

func main() {

	//var bookings = []string{} // --alternative syntax
	//bookings := []string{} // --alternative syntax

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 { // for { and for true { is same

		firstName, lastName, email, userTicket := getUserInput()
		isValidNames, isValidEmail, isValidTicketQuantities := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTickets)

		if isValidNames && isValidEmail && isValidTicketQuantities {

			bookTicket(userTicket, firstName, lastName, email)

			//printFirstName(bookings)
			firstNames := getFirstNames()
			Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidNames {
				Println("First or Last Name is too short")
			}
			if !isValidEmail {
				Println("Email address you entered doesn't contains '@'")
			}
			if !isValidTicketQuantities {
				Println("Number of tickets you entered is invalid")
			}
		}
	}
}

func bookTicket(userTicket uint, firstName string, lastName string, email string) {
	remainingTickets -= userTicket

	// create a map for a use
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["userTicket"] = strconv.FormatUint(uint64(userTicket), 10)

	bookings = append(bookings, userData)
	Printf("List of bookings is %v\n", bookings)

	Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
	Printf("%v tikcets remaining for %v\n", remainingTickets, conferenceName)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	// ask user for input
	Println("Enter your first name:")
	Scan(&firstName)

	Println("Enter your last name:")
	Scan(&lastName)

	Println("Enter your email address:")
	Scan(&email)

	Println("Enter number of tickets:")
	Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func printFirstName(bookings []string) {
	// print first names
	var firstNames []string
	//firstNames := []string{} // -- alternate syntax
	for _, fullName := range bookings { // _ here is the index
		var names = strings.Fields(fullName)
		fName := names[0]
		firstNames = append(firstNames, fName)
	}

	Printf("The first names of bookings are: %v\n", firstNames)
}

func getFirstNames() []string {
	// print first names
	var firstNames []string
	//firstNames := []string{} // -- alternate syntax
	for _, booking := range bookings { // _ here is the index
		name := booking["firstName"] + booking["lastName"]
		firstNames = append(firstNames, name)
	}

	return firstNames
}

func greetUsers() {
	Printf("Welcome to %v booking application\n", conferenceName)
	Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTickets)
	Println("Get your tickets here to attend")
}

//func validateUserInput(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool) {
//	var isValidNames = len(firstName) >= 2 && len(lastName) >= 2
//	isValidEmail := strings.Contains(email, "@")
//	isValidTicketQuantities := userTicket > 0 && userTicket <= remainingTickets
//
//	return isValidNames, isValidEmail, isValidTicketQuantities
//}
