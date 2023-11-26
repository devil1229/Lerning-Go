package main

import "strings"

//we can declare some function in helper which we can use in multiple places in same package

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool , bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNuber := userTickets > 0 && userTickets <= remaingTickets

	return isValidName, isValidEmail, isValidTicketNuber
}