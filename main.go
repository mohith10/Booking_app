package main

import "fmt"

func main() {
	conferenceName := "GO Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Println("Welcome to our booking application for the", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and remaining tickets are:", remainingTickets)
	fmt.Println("Get your", conferenceName, "tickets here!")

	bookings := []string{}

	var userName string
	var userTickets uint
	var userEmail string
	fmt.Print("Enter your Username: ")
	fmt.Scan(&userName)
	fmt.Print("Enter your Ticket Count: ")
	fmt.Scan(&userTickets)
	fmt.Print("Enter your Email: ")
	fmt.Scan(&userEmail)

	bookings = append(bookings, userName)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v for booking %v tickets. We will reach out to you on %v\n", userName, userTickets, userEmail)
	fmt.Printf("Remainign Tickets: %v\n", remainingTickets)

	fmt.Printf("These are all our bookings: %v\n", bookings)
}
