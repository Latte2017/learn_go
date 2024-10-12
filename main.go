package main

import "fmt"

func main() {
	const conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTickets = conferenceTicket
	var tom string
	fmt.Printf("Welcome conference users to %s\n", conferenceName)
	fmt.Printf("Get your tickets here %d total tickets %d available\n", conferenceTicket, remainingTickets)
	fmt.Printf("Enter Tom's name")
	fmt.Scanf("%s", &tom)
	fmt.Printf("Animal is %s\n", tom)
	var bookings []string
	var i int
	for i = 4; i < 10; i++ {
		bookings = append(bookings, "gtfg")
	}
	fmt.Printf("The first element is %s\n", bookings[0])
	fmt.Printf("The last element is %s\n", bookings[len(bookings)-1])
	bookings = append(bookings, "kjutg")

	fmt.Printf("The whole array is %v\n", bookings)

}
