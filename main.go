package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

var festival_name = "Radahn Festival"

const ticket_count = 50

var remaining_tickets uint = 50
var first_names []string

var bookings []string

var first_name string
var last_name string
var user_email string

var user_tickets uint

func main() {

	greet_users()

	for {

		get_user_input()

		is_valid_name, is_valid_email, is_valid_tickets := helper.Validate_user_input(first_name, last_name, user_email, user_tickets, remaining_tickets)

		if !is_valid_name {
			fmt.Printf("the name %v %v is not valid /n", first_name, last_name)
			break
		}

		if !is_valid_email {
			fmt.Printf("The email %v is not valid /n", user_email)
			break
		}

		if !is_valid_tickets {
			fmt.Printf("Only %v chamions can compete in the Radhan Festival /n", remaining_tickets)
			break
		}

		bookings = append(bookings, first_name+" "+last_name)

		remaining_tickets = remaining_tickets - user_tickets

		fmt.Printf("Thank you %v %v for purchasing %v tickets, you shall recieve a confirmation email on %v \n", first_name, last_name, user_tickets, user_email)
		fmt.Printf("%v remaining tickets for %v \n", remaining_tickets, festival_name)

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			first_names = append(first_names, names[0])
			println(first_names)
		}
		fmt.Printf("List of challengers in the Radahn Festival: %v \n", first_names)

		if remaining_tickets == 0 {
			fmt.Printf("The Radhan festival is filled to the brim with eager chamions")
			break
		}
	}
}

func greet_users() {
	fmt.Println("Greetings, young traveller. the ", festival_name, "awaits!!!")
	fmt.Printf("Out of the initial %v only %v remain\n", ticket_count, remaining_tickets)
	fmt.Println("Get your tickets now!, before it is too late")
}

func get_user_input() {

	fmt.Println("Could we have your name my good lad")
	fmt.Scan(&first_name)
	fmt.Println("And your last name please")
	fmt.Scan(&last_name)
	fmt.Println("And your email if you will")
	fmt.Scan(&user_email)
	fmt.Println("And how many tickets would you like to purchase")
	fmt.Scan(&user_tickets)
}
