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
	// ANSI escape codes
	const (
		reset  = "\033[0m"
		cyan   = "\033[36m"
		yellow = "\033[33m"
		red    = "\033[31m"
		green  = "\033[32m"
	)

	// Open the user file
	f, err := os.OpenFile("user.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read existing users from the file
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			user := models.Users{
				Username: parts[0],
				Password: parts[1],
			}
			config.Users = append(config.Users, user)
		}
	}

	// Welcome banner

	for {

		fmt.Println(yellow + "===================================" + reset)
		fmt.Println(yellow + "🌟 Welcome to the Login Page 🌟" + reset)
		fmt.Println(yellow + "===================================" + reset)

		if !auth.Verification() {
			fmt.Println(red + "❌ Verification Failed! Can't Login or SignUp." + reset)
			return
		}

		// Main menu
		fmt.Println("\n" + cyan + "Select an option:" + reset)
		fmt.Println(cyan + "1. 🕵️‍♂️ Login" + reset)
		fmt.Println(cyan + "2. 📝 Signup" + reset)

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			ui.Login()
		case 2:
			ui.SignUp(f)
		default:
			fmt.Println(red + "Invalid option. Please choose 1 or 2." + reset)
			return
		}
	}
}
