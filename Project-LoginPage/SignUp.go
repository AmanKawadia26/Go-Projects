package main

import (
	"fmt"
	"io"
)

func SignUp(f io.Writer) {
	var user Users
	user.username, user.password = CredInput()
	fmt.Println(user)

	for _, pass := range users {
		if pass.username == user.username {
			fmt.Println("This username exists. Can't Signup")
			return
		}
	}
	writingString := fmt.Sprintf("\n%s:%s", user.username, user.password)

	_, err := f.Write([]byte(writingString))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Successfully wrote to file")
}
