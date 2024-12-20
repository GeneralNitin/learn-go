package main

import . "fmt"

func main() {

	conferenceName := "Go Conference"
	const conferenceTicket int = 50
	var remainingTickets uint = 50
	var booking [50]string // empty array of size 50

	Printf("Welcome to %v booking application\n", conferenceName)
	Printf("We have total of %v tickets and %v are still available.\n", conferenceTicket, remainingTickets)
	Println("Get your tickets here to attend")

	//var booking = [50]string{"a"} // array of size 50 with 1 element
	//var booking = [50]string{} // empty array of size 50
	//var booking [50]string // empty array of size 50
	booking[0] = "Nitin"

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

	remainingTickets -= userTicket
	booking[0] = firstName + " " + lastName

	Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
	Printf("%v tikcets remaining for %v\n", remainingTickets, conferenceTicket)
}
