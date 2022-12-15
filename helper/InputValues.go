package helper

import "fmt"
// You export functions and variables by capitalizing them

// Exportable Variable
var MyVar string = "somevalue"

// Exportable function
func InputValues() (string, string, string, uint) {
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