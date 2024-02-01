package main

import (
	"fmt"
	"strings"
)

// The "main" function is the entrypoint of a Go Program
func main() {
	//use of variables in Go
	var confName = "Go to the Conference"
	const confTickets int = 100
	var remTickets uint = 50
	// confNameTrial := "Go to the Conference"
	// if here confTickets = 50 is done it will not change the value of confTickets as it is defined as const.
	bookings := []string{}

	// fmt.Println(confName)
	// fmt.Println("Welcome to Trial ", confNameTrial, "booking application!!")
	// use of print formatted data
	fmt.Printf("Welcome to %v booking application!! \n", confName)
	fmt.Println("Get your tickets from here.")
	fmt.Printf("We have total of %v tickets and %v are remaining", confTickets, remTickets)


	for remTickets > 0 && len(bookings) < 50 {
	for {
		var userFirstName string
		var userLastName string
		var userTickets uint
		var email string

		//uint for positive number of tickets.
		fmt.Println("\n Enter your first name: ")
		fmt.Scan(&userFirstName)
		fmt.Println("\n Enter your last name: ")
		fmt.Scan(&userLastName)
		fmt.Println("\n Enter your email id: ")
		fmt.Scan(&email)
		fmt.Println("\nHello ", userFirstName, "!!\nHow many tickets do you want?")
		fmt.Scan(&userTickets)
		fmt.Println("\nEnter the number of tickets", userTickets)

		isValidName := len(userFirstName)>=2 && len(userLastName)>=2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets>0 && userTickets <= remTickets

		//isValidCity := city == "kathmandu" || city == "delhi"
		if isValidName && isValidEmail && isValidTicketNumber {

			//arrays
			remTickets = remTickets - userTickets
			bookings = append(bookings, userFirstName+" "+userLastName)
			// var bookings = []string{"Sandy", "Mandy", "Handy", "Katappa"}
			// for i := 0; i < len(bookings); i++ {
			// 	fmt.Println(i+1, ".", bookings[i])
			// }
			fmt.Printf("\nThank you for booking %v tickets. Confirmation is sent to %v", userTickets, email)

			//splitting the full names and taking only first names in array
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("\nThe first names og all bookings %v \n", firstNames)

			noTicketsRemaining := remTickets == 0
			if noTicketsRemaining {
				fmt.Println("\nThe Conference is full. All Tickets has been booked.")
				break
			}
			//OR
			// if remTickets == 0 {
			// 	fmt.Println("\nThe Conference is full. All Tickets has been booked.")
			// 	break
			// }
		} else {
			fmt.Printf("We only have %v tickets available!!", remTickets)
		}
	}
}
