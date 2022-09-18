package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have %v/%v tickets available.\n", remainingTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
	// ask user for their name
	/* This is a comment */
}
