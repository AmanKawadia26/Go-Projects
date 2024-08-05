package main

import (
	"bufio"
	"fmt"
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
		//fmt.Println(parts)
		if len(parts) == 2 {
			user := Users{
				username: parts[0],
				password: parts[1],
			}
			users = append(users, user)
		}
	}
	//fmt.Println(users)

	for {
		if !Verification() {
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
			Login()
		case 2:
			SignUp(f)
		default:
			fmt.Println("Invalid option")
			return
		}
	}
}
