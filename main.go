package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	// Setting array with values
	// var bookings = [50]string{"Dan", "Sam", "Brian"}

	// Setting an empty array
	// var bookings = [50]string{}

	// Setting a slice. Slices are like arrays in ruby
	var bookings = []string{}

	// while loop
	// specifically while true
	for {
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

		bookings = append(bookings, firstName+" "+lastName)
		remainingTickets = remainingTickets - userTickets

		// fmt.Printf("The whole array %v\n", bookings)
		// fmt.Printf("First value: %v\n", bookings[0])
		// fmt.Printf("Array type: %T\n", bookings)
		// fmt.Printf("Array length: %v\n", len(bookings))

		// fmt.Printf("The whole slice %v\n", bookings)
		// fmt.Printf("First value: %v\n", bookings[0])
		// fmt.Printf("Slice type: %T\n", bookings)
		// fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v ticket you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		var firstNames = []string{}
		// For loop
		for _, booking := range bookings {
			var name = strings.Fields(booking)
			firstNames = append(firstNames, name[0])
		}
		fmt.Printf("first name of bookings %v\n", firstNames)
	}
}
