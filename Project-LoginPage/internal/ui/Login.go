package ui

import (
	"LoginPage/internal/auth"
	"LoginPage/internal/models"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Login() {
	var user models.Users
	fmt.Println(cyan + "\n🔐 Welcome to the Login Portal 🔐" + reset)
	user.Username, user.Password = auth.CredInput()

	user.Password = hashPassword(user.Password)

	loggedInUser, found := findUser(user.Username, user.Password)
	if !found {
		fmt.Println(red + "\n❌ Invalid username or password. Please try again." + reset)
		return
	}

	fmt.Println(green + "\n✅ Login successful! Welcome back, " + loggedInUser.FirstName + "!" + reset)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(yellow + "\n📋 Main Menu:" + reset)
		fmt.Println(cyan + "1. 📝 Manage To-Do List and Daily Status" + reset)
		fmt.Println(cyan + "2. 📚 Manage Course Progress" + reset)
		fmt.Println(cyan + "3. 👤 View Profile" + reset)
		fmt.Println(cyan + "4. 🚪 Log Out" + reset)
		fmt.Print(yellow + "\nPlease enter your choice (1-4): " + reset)

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			TodoList(loggedInUser.Username)
		case "2":
			manageCourseProgress(loggedInUser.Username)
		case "3":
			viewProfile(loggedInUser)
		case "4":
			fmt.Println(green + "\n👋 Thank you for using our application. Have a great day!" + reset)
			return
		default:
			fmt.Println(red + "\n❌ Invalid choice. Please try again." + reset)
		}
	}
}

func findUser(username, password string) (models.Users, bool) {
	file, err := os.ReadFile("user.txt")
	if err != nil {
		fmt.Println(red + "Error reading user file: " + err.Error() + reset)
		return models.Users{}, false
	}

	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) == 6 && fields[0] == username && fields[1] == password {
			age, _ := strconv.Atoi(fields[4])
			return models.Users{
				Username:  fields[0],
				Password:  fields[1],
				FirstName: fields[2],
				LastName:  fields[3],
				Age:       age,
				MobileNo:  fields[5],
			}, true
		}
	}
	return models.Users{}, false
}

func viewProfile(user models.Users) {
	fmt.Println(yellow + "\n👤 User Profile:" + reset)
	fmt.Printf(cyan+"Username: "+reset+"%s\n", user.Username)
	fmt.Printf(cyan+"First Name: "+reset+"%s\n", user.FirstName)
	fmt.Printf(cyan+"Last Name: "+reset+"%s\n", user.LastName)
	fmt.Printf(cyan+"Age: "+reset+"%d\n", user.Age)
	fmt.Printf(cyan+"Mobile Number: "+reset+"%s\n", user.MobileNo)
	fmt.Println(green + "\nPress any key to return to the main menu..." + reset)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
