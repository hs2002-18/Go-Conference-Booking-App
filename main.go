package main

import (
	"fmt"
	"strings"
	"time"
)

type UserData struct {
	firstName    string
	lastName     string
	emailAddress string
	userTickets  uint
}

func main() {
	conferenceName := "Go Conference"
	const totalTickets uint = 50

	remainingTickets := totalTickets
	bookings := make([]UserData, 0)

	greetUser(conferenceName, totalTickets, remainingTickets)

	var firstName string
	var lastName string
	var emailAddress string
	var userTickets uint

	const bookingLimit uint = 4

	for {
		firstName, lastName, emailAddress = userInput(
			firstName,
			lastName,
			emailAddress,
		)

		userTickets = userBookingTicket(bookingLimit)

		isValidName, isValidEmail, isValidTicketNumber := userValidation(
			firstName,
			lastName,
			emailAddress,
			userTickets,
			remainingTickets,
		)

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets, bookings = userBookingConfirmation(
				firstName,
				lastName,
				emailAddress,
				userTickets,
				remainingTickets,
				bookings,
			)

			go sendTicket(
				userTickets,
				firstName,
				lastName,
				emailAddress,
			)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked. Please come back next year!")
				break
			}

		} else {
			userValidationError(
				isValidName,
				isValidEmail,
				isValidTicketNumber,
			)
		}
	}

	// Prevent program from exiting before goroutines complete
	time.Sleep(11 * time.Second)
}

func greetUser(
	confName string,
	totalTickets uint,
	remainingTickets uint,
) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf(
		"We have a total of %v tickets and %v are still available for sale.\n",
		totalTickets,
		remainingTickets,
	)
	fmt.Printf("Get your tickets here to attend the %v\n", confName)
}

func getFirstName(bookings []UserData) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func userValidation(
	firstName string,
	lastName string,
	emailAddress string,
	userTickets uint,
	remainingTickets uint,
) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	isValidEmail := len(emailAddress) > 5 &&
		strings.Contains(emailAddress, "@") &&
		strings.Contains(emailAddress, ".")

	isValidTicketNumber := userTickets > 0 &&
		userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func userInput(
	firstName string,
	lastName string,
	emailAddress string,
) (string, string, string) {

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&emailAddress)

	return firstName, lastName, emailAddress
}

func userBookingTicket(bookingLimit uint) uint {
	var userTickets uint

	for {
		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets > bookingLimit {
			fmt.Printf(
				"One person can only book %v tickets\n",
				bookingLimit,
			)
			fmt.Println("Please enter again")
		} else {
			return userTickets
		}
	}
}

func userBookingConfirmation(
	firstName string,
	lastName string,
	emailAddress string,
	userTickets uint,
	remainingTickets uint,
	bookings []UserData,
) (uint, []UserData) {

	remainingTickets -= userTickets

	booking := UserData{
		firstName:    firstName,
		lastName:     lastName,
		emailAddress: emailAddress,
		userTickets:  userTickets,
	}

	bookings = append(bookings, booking)

	fmt.Printf("\nList of bookings: %v\n", bookings)

	fmt.Printf(
		"Thank you %v %v for booking %v tickets.\n",
		firstName,
		lastName,
		userTickets,
	)

	fmt.Printf(
		"You will receive a confirmation on %v\n",
		emailAddress,
	)

	fmt.Printf("%v remaining tickets\n", remainingTickets)

	firstNames := getFirstName(bookings)
	fmt.Printf("Current bookings: %v\n\n", firstNames)

	return remainingTickets, bookings
}

func userValidationError(
	isValidName bool,
	isValidEmail bool,
	isValidTicketNumber bool,
) {

	if !isValidName {
		fmt.Println(
			"First name or last name entered is too short.",
		)
	}

	if !isValidEmail {
		fmt.Println(
			"Email address entered is not valid.",
		)
	}

	if !isValidTicketNumber {
		fmt.Println(
			"Number of tickets entered is invalid.",
		)
	}
}

func sendTicket(
	userTickets uint,
	firstName string,
	lastName string,
	emailAddress string,
) {

	time.Sleep(10 * time.Second)

	ticket := fmt.Sprintf(
		"%v tickets for %v %v",
		userTickets,
		firstName,
		lastName,
	)

	fmt.Println("\n#################")
	fmt.Printf(
		"Sending ticket:\n%v\nto email address %v\n",
		ticket,
		emailAddress,
	)
	fmt.Println("#################")
}