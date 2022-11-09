package main

import (
	"fmt"
	"strings"
)

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

	for {
		fmt.Println("================================================")
		fmt.Print("Enter your Username: ")
		fmt.Scan(&userName)
		fmt.Print("Enter your Ticket Count: ")
		fmt.Scan(&userTickets)
		fmt.Print("Enter your Email: ")
		fmt.Scan(&userEmail)
		fmt.Println("================================================")

		if userTickets > remainingTickets {
			fmt.Printf("Oops! We only have %v tickets available", remainingTickets)
			continue
		} else if userTickets == remainingTickets {
			remainingTickets = 0
			fmt.Println("Our conference is booked out!")
			break
		} else {
			bookings = append(bookings, userName)
			remainingTickets = remainingTickets - userTickets

			fmt.Println("================================================")
			fmt.Printf("Thank you %v for booking %v tickets. We will reach out to you on %v\n", userName, userTickets, userEmail)
			fmt.Printf("Remainign Tickets: %v\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Println("================================================")
			fmt.Printf("These are all our bookings: %v\n", firstNames)
		}

	}
}
