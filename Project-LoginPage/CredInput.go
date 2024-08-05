package main

import (
	"LoginPage/Project-LoginPage/auth"
	"fmt"
)

func CredInput() (string, string) {
	fmt.Println("Enter Username: ")
	var username string
	fmt.Scan(&username)
	var password string
	for {
		fmt.Println("Enter Password: ")
		fmt.Scan(&password)
		if !auth.IsStrongPassword(password) {
			fmt.Println("Weak Password. Enter a stronger password (min. length 8 characters, 1 uppercase, 1 lowercase, 1 digit and 1 special character)")
		} else {
			break
		}
	}

	return username, password
}
