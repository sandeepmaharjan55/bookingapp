package main

import "fmt"

// The "main" function is the entrypoint of a Go Program
func main() {
	//use of variables in Go
	var confName = "Go to the Conference"
	//var confName string = "Go to the Conference"
	const confTickets int = 100
	var remTickets uint = 50
	confNameTrial := "Go to the Conference"
	// if here confTickets = 50 is done it will not change the value of confTickets as it is defined as const.
	bookings := []string{}

	fmt.Println("Welcome to this booking application!!")
	fmt.Println("Get your tickets from here.")

	fmt.Println(confName)
	fmt.Println("Welcome to", confName, "booking application!!")
	fmt.Println("Welcome to Trial ", confNameTrial, "booking application!!")
	fmt.Println("We have total of ", confTickets, "tickets and ", remTickets, " are remaining")

	// use of print formatted data
	fmt.Printf("Welcome to %v booking application!! \n", confName)
	fmt.Printf("We have total of %v tickets and %v are remaining", confTickets, remTickets)

	var userName string
	var userTickets uint

	//uint for positive number of tickets.
	fmt.Println("\n Your Name?")
	fmt.Scan(&userName)
	fmt.Println("\nHello ", userName, " How many tickets you want?")
	fmt.Scan(&userTickets)
	fmt.Println("\nNumber of tickets", userTickets)

	//arrays
	remTickets = remTickets - userTickets
	bookings = append(bookings, userName)
	// var bookings = []string{"Sandy", "Mandy", "Handy", "Katappa"}
	// for i := 0; i < len(bookings); i++ {
	// 	fmt.Println(i+1, ".", bookings[i])
	// }
	fmt.Printf("\nThank you for booking %v tickets. Now remaining tickets are %v", userTickets, remTickets)
	fmt.Printf("\nThese are all bookings %v", bookings)
}
