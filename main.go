package main

import (
	"fmt"
	// "strings"
	"booking-app/helper"
	// "strconv"
)

// Package level variables
// Think of them like constants in ruby
var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

// slice of strings
// var bookings = []string{}

// slice of maps
// var bookings = make([]map[string]string, 0)

// slice of UserData structs
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

func main() {
	greetUsers()

	// Setting a slice. Slices are like arrays in ruby

	// while loop
	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := helper.InputValues()

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets

			// map[key_type]value_type
			// empty map
			// var userData = make(map[string]string)

			// userData["firstName"] = firstName
			// userData["lastName"] = lastName
			// userData["email"] = email
			// // Converting uint to string
			// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

			var userData = UserData{
				firstName: firstName,
				lastName:  lastName,
				email:     email,
				// userTickets: userTickets,
			}

			userData.userTickets = userTickets

			// bookings = append(bookings, firstName+" "+lastName)
			bookings = append(bookings, userData)

			bookingStats()

			fmt.Printf("Thank you %v %v for booking %v ticket you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			var firstNames = firstNames()
			fmt.Printf("first name of bookings %v\n", firstNames)

			fmt.Printf("Current Bookings %v\n", bookings)

		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}
	}
	fmt.Println("Our conference is booked out. Come back next year")
}

func greetUsers() {
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

}

func bookingStats() {
	fmt.Printf("The whole slice %v\n", bookings)
	fmt.Printf("First value: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice length: %v\n", len(bookings))
}

// func firstNames() []string {
// 	var firstNames = []string{}
// 	// For loop
// 	for _, booking := range bookings {
// 		// this turns the element into a string?
// 		var name = strings.Fields(booking)
// 		firstNames = append(firstNames, name[0])
// 	}
// 	return firstNames
// }

// Map firstNames
// func firstNames() []string {
// 	var firstNames = []string{}
// 	// For loop
// 	for _, booking := range bookings {
// 		firstNames = append(firstNames, booking["firstName"])
// 	}
// 	return firstNames
// }

func firstNames() []string {
	var firstNames = []string{}
	// For loop
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
