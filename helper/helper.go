package helper

import "strings"

func Validate_user_input(first_name string, last_name string, user_email string, user_tickets uint, remaining_tickets uint) (bool, bool, bool) {
	is_valid_name := len(first_name) >= 2 && len(last_name) >= 2
	is_valid_email := strings.Contains(user_email, "@")
	is_valid_tickets := user_tickets > 0 && user_tickets <= remaining_tickets

	return is_valid_name, is_valid_email, is_valid_tickets
}
