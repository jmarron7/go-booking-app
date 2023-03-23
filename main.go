package main

import (
	"fmt"
	"go-booking-app/helper"
	"sync"
	"time"
)

const conferenceName = "Go Conference"
const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName  string
	lastName   string
	email      string
	numTickets uint
}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketAmount := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidTicketAmount && isValidEmail {

		bookTickets(firstName, lastName, email, userTickets)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("These are the first names of all of our bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year!")
		}
	} else {
		if !isValidName {
			fmt.Println("First or Last name entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address entered does not contain @ sign")
		}
		if !isValidTicketAmount {
			fmt.Println("Number of tickets entered is invalid")
		}
	}

	wg.Wait()
}

var wg = sync.WaitGroup{}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
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

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets to order: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		numTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###############")
	wg.Done()
}
