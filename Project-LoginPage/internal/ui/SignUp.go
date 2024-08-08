package ui

import (
	"LoginPage/internal/auth"
	"LoginPage/internal/config"
	"LoginPage/internal/models"
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SignUp(f io.Writer) {
	var user models.Users
	var err error
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(cyan + "\n🚀 Welcome to the Sign Up Process! 🚀" + reset)
	fmt.Println(yellow + "Let's get you set up with a new account." + reset)

	// Get username and password
	for {
		user.Username, user.Password = auth.CredInput()
		if !auth.IsStrongPassword(user.Password) {
			fmt.Println(red + "\n❌ Your password is not strong enough." + reset)
			fmt.Println(yellow + "Please ensure your password meets the following criteria:" + reset)
			fmt.Println(cyan + "- At least 8 characters long" + reset)
			fmt.Println(cyan + "- Contains at least one uppercase letter" + reset)
			fmt.Println(cyan + "- Contains at least one lowercase letter" + reset)
			fmt.Println(cyan + "- Contains at least one digit" + reset)
			fmt.Println(cyan + "- Contains at least one special character" + reset)
		} else {
			break
		}
	}

	// Get first name
	for {
		fmt.Print(yellow + "\nEnter your first name: " + reset)
		user.FirstName, _ = reader.ReadString('\n')
		user.FirstName = strings.TrimSpace(user.FirstName)
		if isValidName(user.FirstName) {
			break
		}
		fmt.Println(red + "❌ Invalid first name. Please use only letters and avoid spaces." + reset)
	}

	// Get last name
	for {
		fmt.Print(yellow + "\nEnter your last name: " + reset)
		user.LastName, _ = reader.ReadString('\n')
		user.LastName = strings.TrimSpace(user.LastName)
		if isValidName(user.LastName) {
			break
		}
		fmt.Println(red + "❌ Invalid last name. Please use only letters and avoid spaces." + reset)
	}

	// Get age
	for {
		fmt.Print(yellow + "\nEnter your age: " + reset)
		ageStr, _ := reader.ReadString('\n')
		ageStr = strings.TrimSpace(ageStr)
		user.Age, err = strconv.Atoi(ageStr)
		if err == nil && user.Age > 0 && user.Age <= 150 {
			break
		}
		fmt.Println(red + "❌ Invalid age. Please enter a number between 1 and 150." + reset)
	}

	// Get mobile number
	for {
		fmt.Print(yellow + "\nEnter your mobile number (10 digits): " + reset)
		user.MobileNo, _ = reader.ReadString('\n')
		user.MobileNo = strings.TrimSpace(user.MobileNo)
		if isValidMobileNumber(user.MobileNo) {
			break
		}
		fmt.Println(red + "❌ Invalid mobile number. Please enter a 10-digit number." + reset)
	}

	user.Password = hashPassword(user.Password)

	for _, pass := range config.Users {
		if pass.Username == user.Username {
			fmt.Println(red + "\n❌ This username already exists. Please choose a different one." + reset)
			return
		}
	}

	writingString := fmt.Sprintf("\n%s:%s:%s:%s:%d:%s", user.Username, user.Password, user.FirstName, user.LastName, user.Age, user.MobileNo)
	_, err = f.Write([]byte(writingString))
	if err != nil {
		fmt.Println(red + "\n❌ Error writing to file: " + err.Error() + reset)
		return
	}

	courseFileName := user.Username + "_course.txt"
	courseFile, err := os.Create(courseFileName)
	if err != nil {
		fmt.Println(red + "\n❌ Error creating course file: " + err.Error() + reset)
		return
	}
	defer courseFile.Close()

	modules := []string{
		"Module 1: Introduction to the Course",
		"Module 2: Basics",
		"Module 3: Intermediate Concepts",
		"Module 4: Advanced Topics",
		"Module 5: Expert Techniques",
		"Module 6: Practical Applications",
		"Module 7: Case Studies",
		"Module 8: Future Trends",
		"Module 9: Final Project",
		"Module 10: Course Review",
	}

	_, err = courseFile.WriteString("Modules:\n")
	if err != nil {
		fmt.Println(red + "\n❌ Error writing to course file: " + err.Error() + reset)
		return
	}

	for _, module := range modules {
		_, err := courseFile.WriteString(module + "\n")
		if err != nil {
			fmt.Println(red + "\n❌ Error writing to course file: " + err.Error() + reset)
			return
		}
	}

	fmt.Println(green + "\n✅ Sign up successful! Welcome aboard, " + user.FirstName + "! 🎉" + reset)
	fmt.Println(yellow + "You can now log in with your new account." + reset)
}

func isValidName(name string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(name)
}

func isValidMobileNumber(number string) bool {
	return regexp.MustCompile(`^\d{10}$`).MatchString(number)
}

func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
