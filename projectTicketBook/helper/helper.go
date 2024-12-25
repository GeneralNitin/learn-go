package helper

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {
	var isValidNames = len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketQuantities := userTicket > 0 && userTicket <= remainingTickets

	return isValidNames, isValidEmail, isValidTicketQuantities
}
