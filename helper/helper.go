package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickents uint) (bool, bool, bool) {
	isValideName := len(firstName) >= 2 && len(lastName) >= 2
	isValidateEmail := strings.Contains(email, "@")
	isValidTickerNumber := userTickets > 0 && userTickets <= remainingTickents

	return isValideName, isValidateEmail, isValidTickerNumber
}
