package main

import (
	. "fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTicket int = 50
	var remainingTickets uint = 50
	var bookings []string // dynamic array
	//var bookings = []string{} // --alternative syntax
	//bookings := []string{} // --alternative syntax

	Printf("Welcome to %v booking application\n", conferenceName)
	Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTickets)
	Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	for {
		// ask user for input
		Println("Enter your first name:")
		Scan(&firstName)

		Println("Enter your last name:")
		Scan(&lastName)

		Println("Enter your email address:")
		Scan(&email)

		Println("Enter number of tickets:")
		Scan(&userTicket)

		remainingTickets -= userTicket
		bookings = append(bookings, firstName+" "+lastName)

		Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
		Printf("%v tikcets remaining for %v\n", remainingTickets, conferenceName)

		// print first names
		firstNames := []string{}
		//var firstNames []string // -- alternate syntax
		for _, fullName := range bookings { // _ here is the index
			var names = strings.Fields(fullName)
			fName := names[0]
			firstNames = append(firstNames, fName)
		}

		Printf("These are all our bookings: %v\n", firstNames)
	}

}
