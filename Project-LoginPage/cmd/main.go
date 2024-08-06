package main

import (
	"LoginPage/internal/auth"
	"LoginPage/internal/config"
	"LoginPage/internal/models"
	"LoginPage/internal/ui"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	f, err := os.OpenFile("user.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {

		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	//fmt.Println(scanner)

	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		parts := strings.Split(line, ":")
		//fmt.Println(parts)
		if len(parts) == 2 {
			user := models.Users{
				Username: parts[0],
				Password: parts[1],
			}
			config.Users = append(config.Users, user)
		}
	}
	//fmt.Println(users)

	for {
		if !auth.Verification() {
			fmt.Println("Verification Failed! Can't Login or SignUp.")
			return
		}
		fmt.Println("\nSelect an option: ")
		fmt.Println("1. Login")
		fmt.Println("2. Signup")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			ui.Login()
		case 2:
			ui.SignUp(f)
		default:
			fmt.Println("Invalid option")
			return
		}
	}
}
