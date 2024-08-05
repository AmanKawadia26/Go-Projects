package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Login() {
	var user Users
	user.username, user.password = CredInput()
	for _, pass := range users {
		if pass.username == user.username && pass.password == user.password {
			fmt.Println("You are logged in")
			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Println("1. Check To-Do List and Daily Status")
				fmt.Println("2. Go back to main menu")
				fmt.Print("Enter your choice: ")

				choice, _ := reader.ReadString('\n')
				choice = strings.TrimSpace(choice)

				switch choice {
				case "1":
					TodoList(user.username)
				case "2":
					return
				default:
					fmt.Println("Invalid choice. Please try again.")
				}
			}
		}
	}
	fmt.Println("Invalid username or password")
}
