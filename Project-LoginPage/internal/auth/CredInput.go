package auth

import (
	"fmt"
)

func CredInput() (string, string) {
	fmt.Println("Enter Username: ")
	var username string
	fmt.Scan(&username)
	var password string
	fmt.Println("Enter Password: ")
	fmt.Scan(&password)

	return username, password
}
