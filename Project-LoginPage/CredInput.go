package main

import "fmt"

func CredInput() (string, string) {
	fmt.Println("Enter Username: ")
	var username string
	fmt.Scan(&username)
	fmt.Println("Enter Password: ")
	var password string
	fmt.Scan(&password)
	return username, password
}
