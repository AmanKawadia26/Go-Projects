package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Users struct {
	username string
	password string
}

var users []Users

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
		fmt.Println(parts)
		if len(parts) == 2 {
			user := Users{
				username: parts[0],
				password: parts[1],
			}
			users = append(users, user)
		}
	}
	fmt.Println(users)

	fmt.Println("Select an option: ")
	fmt.Println("1. Login")
	fmt.Println("2. Signup")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		Login()
	case 2:
		SignUp(f)
	default:
		fmt.Println("Invalid option")
		return
	}
}

func Login() {
	fmt.Println("Enter Username: ")
	var username string
	fmt.Scan(&username)
	fmt.Println("Enter Password: ")
	var password string
	fmt.Scan(&password)
	user := Users{username, password}
	//fmt.Println(user)

	for _, pass := range users {
		if pass.username == user.username && pass.password == user.password {
			fmt.Println("You are logged in")
			return
		}
	}
	fmt.Println("Invalid username or password")
}

func SignUp(f io.Writer) {
	fmt.Println("Enter Username: ")
	var username string
	fmt.Scan(&username)
	fmt.Println("Enter Password: ")
	var password string
	fmt.Scan(&password)
	user := Users{username, password}
	fmt.Println(user)

	for _, pass := range users {
		if pass.username == user.username {
			fmt.Println("This username exists. Can't Signup")
			return
		}
	}
	writingString := fmt.Sprintf("\n%s:%s", username, password)

	_, err := f.Write([]byte(writingString))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Successfully wrote to file")
}
