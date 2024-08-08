package auth

import (
	"fmt"
)

const (
	reset  = "\033[0m"
	yellow = "\033[33m"
)

func CredInput() (string, string) {
	fmt.Println(yellow + "\nEnter Username: " + reset)
	var username string
	fmt.Scan(&username)
	var password string
	fmt.Println(yellow + "Enter Password: " + reset)
	fmt.Scan(&password)

	return username, password
}
