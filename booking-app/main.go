package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have %v/%v tickets available.\n", remainingTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your E-mail address:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of Tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - uint(userTickets)

	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets) # Print the location in the memory where this variale is saved.
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation e-mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Println("We have %v ticker remaininig for the %$v.", remainingTickets, conferenceName)
	// ask user for their name
	/* This is a comment */
}
