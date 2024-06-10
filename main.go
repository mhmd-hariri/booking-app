package main

import (
	"Booking-APP/helper"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var conferenceName = "Go conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var bookings = make([]UserData, 0)

func main() {

	//bookings :=[]string{}
	greetUsers()
	//fmt.Printf("conferenceName %T conferenceTickets %T remainingTickets %T \n", conferenceName, conferenceTickets, remainingTickets)
	//fmt.Printf("Welcome to %v booking application\n", conferenceName)

	//fmt.Println("Get your tickets here to attend")

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		firtNames := getFirstNames()
		fmt.Printf("the first names of bookings: %v\n", firtNames)
		if remainingTickets == 0 {
			// end program
			fmt.Printf("Our conference is booked out. Come back next year.")
		}
	} else {
		if !isValidName {
			fmt.Printf("your first name or last name you enetred  is too short \n")
		}
		if !isValidEmail {
			fmt.Printf("your email address  you enetred  doesn't conatin @ sign  \n")
		}
		if !isValidTicketNumber {
			fmt.Printf("the number of tickets you entered is invalid\n")
		}
	}
	wg.Wait()
}
func greetUsers() {
	fmt.Printf("Welcome to  %v  booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are stille available\n", conferenceTickets, remainingTickets)
}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var firstName = booking.firstName
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// aks user about the name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	// ask user about the number of tickets
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	/* var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) */

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	/* fmt.Printf("the whole slice : %v\n", bookings)
	fmt.Printf("the first value %v\n", bookings[0])
	fmt.Printf("slice type %T\n", bookings)
	fmt.Printf("slice Length %v\n", len(bookings)) */
	fmt.Printf("thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("############################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("############################")
	wg.Done()
}
