package ui

import (
	"LoginPage/internal/auth"
	"LoginPage/internal/config"
	"LoginPage/internal/models"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Login() {
	var user models.Users
	user.Username, user.Password = auth.CredInput()

	user.Password = hashPassword(user.Password)

	for _, pass := range config.Users {
		if pass.Username == user.Username && pass.Password == user.Password {
			fmt.Println("You are logged in.")
			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Println("\n1. Check To-Do List and Daily Status")
				fmt.Println("2. Go back to main menu")
				fmt.Print("Enter your choice: ")

				choice, _ := reader.ReadString('\n')
				choice = strings.TrimSpace(choice)

				switch choice {
				case "1":
					TodoList(user.Username)
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
