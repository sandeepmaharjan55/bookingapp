package main

import "fmt"

// The "main" function is the entrypoint of a Go Program
func main() {
	//use of variables in Go
	var confName = "Go to the Conference"
	//var confName string = "Go to the Conference"
	const confTickets = 100
	var remTickets = 50
	confNameTrial := "Go to the Conference"
	// if here confTickets = 50 is done it will not change the value of confTickets as it is defined as const.

	fmt.Println("Welcome to this booking application!!")
	fmt.Println("Get your tickets from here.")

	fmt.Println(confName)
	fmt.Println("Welcome to", confName, "booking application!!")
	fmt.Println("Welcome to Trial ", confNameTrial, "booking application!!")
	fmt.Println("We have total of ", confTickets, "tickets and ", remTickets, " are remaining")

	// use of print formatted data
	fmt.Printf("Welcome to %v booking application!! \n", confName)
	fmt.Printf("We have total of %v tickets and %v are remaining", confTickets, remTickets)

}
