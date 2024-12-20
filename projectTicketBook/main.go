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

	for remainingTickets > 0 && len(bookings) < 50 { // for { and for true { is same
		// ask user for input
		Println("Enter your first name:")
		Scan(&firstName)

		Println("Enter your last name:")
		Scan(&lastName)

		Println("Enter your email address:")
		Scan(&email)

		Println("Enter number of tickets:")
		Scan(&userTicket)

		var isValidNames = len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketQuantities := userTicket > 0 && userTicket <= remainingTickets

		if isValidNames && isValidEmail && isValidTicketQuantities {
			remainingTickets -= userTicket
			bookings = append(bookings, firstName+" "+lastName)

			Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
			Printf("%v tikcets remaining for %v\n", remainingTickets, conferenceName)

			// print first names
			var firstNames []string
			//firstNames := []string{} // -- alternate syntax
			for _, fullName := range bookings { // _ here is the index
				var names = strings.Fields(fullName)
				fName := names[0]
				firstNames = append(firstNames, fName)
			}

			Printf("These are all our bookings: %v\n", firstNames)
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
