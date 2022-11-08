package main

import "fmt"

func main() {
	conferenceName := "GO Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to our booking application for the", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and remaining tickets are:", remainingTickets)
	fmt.Println("Get your", conferenceName, "tickets here!")

	var userName string
	var userTickets int
	fmt.Print("Enter your username: ")
	fmt.Scan(&userName)
	fmt.Print("Enter your ticket count: ")
	fmt.Scan(&userTickets)
	fmt.Println(userName, userTickets)
}
