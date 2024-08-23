package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	createdAt time.Time
}

func main() {

	//  way1 :- declared and initialized
	var appUser user = user{
		firstName: "Sameer",
		lastName:  "singh",
		createdAt: time.Now(),
	}

	//  way2 declaration and then initialization
	var appUser2 user
	appUser2 = user{}

	//  way3 declaration and then initialization
	appUser3 := user{}

	fmt.Println(appUser)
	fmt.Println(appUser2)
	fmt.Println(appUser3)

	takeUserInputDetails(&appUser2)
	showUserOutputDetails(&appUser2)
}

func takeUserInputDetails(userDetaitls *user) {
	fmt.Println("Enter the firstName of the User")
	fmt.Scan(&(*userDetaitls).firstName)
	fmt.Println("Enter the SecondName of the User")
	fmt.Scan(&(*userDetaitls).lastName)
	fmt.Println("Thanks For entering the Details")
}

func showUserOutputDetails(userDetaitls *user) {
	fmt.Println("Your User Details are here ", (*userDetaitls).firstName, (*userDetaitls).lastName)
}
