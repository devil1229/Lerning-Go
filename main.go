package main

import (
	"fmt"
	// "strings"
	"sync"
	"time"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remaingTickets uint= 50
var bookings = make([]User , 0)



type User struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {


	//practice Question Calls
	printPascals()
	addTwoMatrix()
	fibbonacci()

	startBooking()

	
}

func startBooking(){
	//declaring the variables at packge level so that each function the package can use them 
	// conferenceName := "Go Conference"
	// const conferenceTickets = 50
	// var remaingTickets uint= 50
	// //var bookingsArray [50]string  we will be using slice for this example
	// var bookingsSlice []string

	greetUser()

	//Moving it to the fuction for more clean code
	//fmt.Println("welcome to", conferenceName, "booking Application")
	// fmt.Printf("welcome to %v booking Application\n" , conferenceName)
	// //fmt.Println("we have total of", conferenceTickets, "tickets and", remainggTickets , "are still available.")
	// fmt.Printf("we have total of %v tickets and %v are still available.", conferenceTickets, remaingTickets)
	// fmt.Println("Get your tickets here to attend")

	// for {
		// var firstName string
		// var lastName string
		// var email string
		// var userTickets uint  moving it to the func getUserInput

		firstName, lastName, email, userTickets := getUserInput()

		// fmt.Printf("Enter your first name: ")
		// fmt.Scan(&firstName)

		// fmt.Printf("Enter your last name: ")
		// fmt.Scan(&lastName)

		// fmt.Printf("Enter your email address: ")
		// fmt.Scan(&email)

		// fmt.Printf("Enter Number of tickets: ")
		// fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicketNuber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidEmail && isValidName && isValidTicketNuber {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

            firstNames := getFirstNames()
			fmt.Printf("These are all our Bookings: %v\n", firstNames)

			if remaingTickets == 0 {
				fmt.Println("Our confrrence is booking out. Come back next year.")
				// break
			}
			// fmt.Printf("we only have %v tickets remaining, so you can't book %v tickets\n" , remaingTickets , userTickets)
			// continue
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}

			if !isValidTicketNuber {
				fmt.Println("number of tickets you entered is invalid")
			}

		}
        
		wg.Wait()
		// remaingTickets = remaingTickets - userTickets

		//bookingsArray[0] = firstName + " " + lastName
		//bookingsSlice = append(bookingsSlice, firstName + " " + lastName)

		//use of arrays
		// fmt.Printf("The whole Array: %v\n", bookingsArray)
		// fmt.Printf("The first value: %v\n", bookingsArray[0])
		// fmt.Printf("Array Type: %T\n", bookingsArray)
		// fmt.Printf("Array length: %v\n", len(bookingsArray))
        
		//use of slice
		// fmt.Printf("The whole Slice: %v\n", bookingsSlice)
		// fmt.Printf("The first value: %v\n", bookingsSlice[0])
		// fmt.Printf("Slice Type: %T\n", bookingsSlice)
		// fmt.Printf("Slice length: %v\n", len(bookingsSlice))


		// fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName,userTickets, email)
		// fmt.Printf("%v tickets remaining for %v\n", remaingTickets, conferenceName)

		// firstNames := []string{}

		// for _, booking := range bookingsSlice{
		// 	var names = strings.Fields(booking)
		// 	firstNames = append(firstNames, names[0])
		// }

		// fmt.Printf("These are all our Bookings: %v\n", firstNames)

		// if remaingTickets == 0 {
		// 	fmt.Println("Our confrrence is booking out. Come back next year")
		// 	break
		// }
	// }
}

func greetUser() {
	fmt.Printf("welcome to %v booking Application\n" , conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.", conferenceTickets, remaingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

    fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Enter Number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remaingTickets = remaingTickets - userTickets

	var newUser = User{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, newUser)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName,userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remaingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets , firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###################")
	wg.Done()
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings{
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}