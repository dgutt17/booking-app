package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	greetUsers(conferenceTickets, remainingTickets, conferenceName)

	// Setting a slice. Slices are like arrays in ruby
	var bookings = []string{}

	// while loop
	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := inputValues()

		if userTickets <= remainingTickets {
			bookings = append(bookings, firstName+" "+lastName)
			remainingTickets = remainingTickets - userTickets

			bookingStats(bookings)

			fmt.Printf("Thank you %v %v for booking %v ticket you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			var firstNames = firstNames(bookings)
			fmt.Printf("first name of bookings %v\n", firstNames)

		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}
	}
	fmt.Println("Our conference is booked out. Come back next year")
}

func greetUsers(conferenceTickets int, remainingTickets uint, conferenceName string) {
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

}

func bookingStats(bookings []string) {
	fmt.Printf("The whole slice %v\n", bookings)
	fmt.Printf("First value: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice length: %v\n", len(bookings))
}

func firstNames(bookings []string) []string {
	var firstNames = []string{}
	// For loop
	for _, booking := range bookings {
		// this turns the element into a string?
		var name = strings.Fields(booking)
		firstNames = append(firstNames, name[0])
	}
	return firstNames
}

func inputValues() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
