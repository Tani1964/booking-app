package main

import (
	"fmt"
	"strings"
)

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func main() {
	conferenceName := "Tanisty Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	// var bookings [50]string
	var bookings = []string{}

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 && len(bookings) < conferenceTickets {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email: ")
		fmt.Scan(&email)

		fmt.Print("Enter your number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if !isValidName {
			fmt.Println("First name or last name you entered is too short")
			continue
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign")
			continue
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
			continue
		}

		if userTickets <= remainingTickets {
			remainingTickets -= userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

			firstNames := []string{}

			// foreach
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("These are all our bookings: %v \n", firstNames)
			var noTicketsremaining bool = remainingTickets == 0

			if noTicketsremaining {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else if userTickets == remainingTickets {
			fmt.Printf("You have booked the last %v tickets. Our conference is booked out. Come back next year.\n", userTickets)
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Println(bookings)
			break
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets \n", remainingTickets, userTickets)
		}
	}

}
