package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// Variables with camelCase
const conferenceTickets = 50

var conferenceName = "Go Conference"

// Specifically identifying type helps throw error in case of mismatched value assigned
// i.e. when we don't want negative value, use uint that only expects positive
var remainingTickets uint = conferenceTickets
var bookings = make([]UserData, 0)

// structs are like classes
// allows key value pair with mixed data types
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for remainingTickets > 0 {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTickets {

			bookTickets(userTickets, firstName, lastName, email)
			fmt.Printf("All bookings: %v\n\n", getFirstNames())

		} else {
			if !isValidName {
				fmt.Printf("Invalid name! More than two characters are required.\n")
			}

			if !isValidEmail {
				fmt.Printf("Invalid email! Please provide a valid email address.\n")
			}

			if !isValidTickets {
				fmt.Printf("Invalid tickets! Please provide a valid number of tickets.\n")
			}

			fmt.Println()
		}

	}
	fmt.Println("All conference tickets are booked! Come back next year!")
	wg.Wait()
}

func greetUsers() {

	fmt.Printf("conferenceName is %T; conferenceTicket is %T; remainingTickets is %T", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v conference's booking application.\n", conferenceName)

	// Vars and imports need to be used when defined, otherwise compile error will be thrown
	// Space automatically added after each print item
	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Printf("We only have %v remainingTickets tickets remaining\n", remainingTickets)
	fmt.Println("Get your tickets now!")
}

func getFirstNames() []string {
	firstNames := []string{}

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

	// ask User First Name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	// ask User Last Name
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	// ask User Email
	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	// ask User Tickets
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Map of user data
	/*var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)*/
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v for booking %v tickets!\nYou will receive a confirmation email at %v in a few minutes.\n", firstName, userTickets, email)
	fmt.Printf("There are now only %v tickets remaining!\n", remainingTickets)

	wg.Add(1)
	go sendTicket(userTickets, firstName, lastName, email)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("===========================")
	fmt.Printf("Sent: %v\nSent To: %v\n\n", ticket, email)
	fmt.Println("===========================")
	wg.Done()
}
