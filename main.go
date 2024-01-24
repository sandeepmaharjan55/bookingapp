package main

import "fmt"

// The "main" function is the entrypoint of a Go Program
func main() {
	//Add fmt to declare from where Print is coming from
	//fmt.Print returns Hello World%
	//fmt.Print("Hello World")

	//fmt.Println returns Hello World in new line
	//println prints in a new line
	fmt.Println("Welcome to this booking application!!")
	fmt.Println("Get your tickets from here.")
	// everything is organized into packages
	// In this case, we are using the built-in "fmt" package which contains functions

	//use of variables in Go
	var confName = "Go to the Conference."
	fmt.Println(confName)

	//end of use of variables in Go
}
