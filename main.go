package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

// declare struct
type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	//infinite loop -> for only or with until condition is met remainingTickets > 0
	//for remainingTickets > 0 {

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isGotTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isGotTickets {

		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email, conferenceName)

		firstNames := getFirstNames()
		fmt.Printf("These are all our bookings %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0

		if noTicketsRemaining {
			//end program
			fmt.Println("Our conference is sold out. Come back next year")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("First Name or Last Name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address is invalid")
		}
		if !isGotTickets {
			fmt.Println("Invalid tickets number")
		}

	}

	wg.Wait()
	//}

}

func greetUsers() {
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("We have total of %v tickets & %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	//for each loop
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//ask user to input all info
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of ticket purchased: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// use struct instead of map
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets with us.\n", firstName, lastName, userTickets)
	fmt.Println("Tickets confirmation have been emailed at", email, ".")
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string, conferenceName string) {
	time.Sleep(3 * time.Second)
	var ticket = fmt.Sprintf("%v %v tickets for %v %v\n", userTickets, conferenceName, firstName, lastName)
	fmt.Println("############################")
	fmt.Printf("Sending ticket:\n %v to %v\n", ticket, email)
	fmt.Println("############################")
	wg.Done()
}
