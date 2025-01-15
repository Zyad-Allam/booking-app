package main

import (
	"fmt"
)

func main() {
	var festival_name = "Radahn Festival"
	const ticket_count = 50
	var remaining_tickets = 49

	fmt.Println("Greetings, young traveller. the ", festival_name, "awaits!!!")
	fmt.Printf("Out of the initial %v only %v remain\n", ticket_count, remaining_tickets)
	fmt.Println("Get your tickets now!, before it is too late")

	var first_name string
	var last_name string
	var user_email string

	var user_tickets int

	fmt.Println("Could we have your name my good lad")
	fmt.Scan(&first_name)
	fmt.Println("And your last name please")
	fmt.Scan(&last_name)
	fmt.Println("And your email if you will")
	fmt.Scan(&user_email)
	fmt.Println("And how many tickets would you like to purchase")
	fmt.Scan(&user_tickets)

	fmt.Printf("Thank you %v %v for purchasing %v tickets, you shall recieve a confirmation email on %v \n", first_name, last_name, user_tickets, user_email)
}
