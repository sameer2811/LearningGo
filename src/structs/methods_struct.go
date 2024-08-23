package main

import "fmt"

type user struct {
	firstName string
	lastName  string
}

func (appUser *user) inputUserDetails() {
	fmt.Println("\nWELCOME TO THE SECTION OF TAKING THE USER INPUTS--")
	fmt.Println("Please enter your firstName ---->")
	fmt.Scan(&((*appUser).firstName))
	fmt.Println("Please enter your lastName ---->")
	fmt.Scan(&((*appUser).lastName))
	fmt.Println("Thanks For Entering the Details ---->")
}

func (appUser *user) outputUserDetails() {
	fmt.Println("\nHERE ARE THE DETAILS OF THE USER INPUTS ----> ", (*appUser).firstName, (*appUser).lastName)
}

func main() {
	var appUser user
	appUser.inputUserDetails()
	appUser.outputUserDetails()

}
