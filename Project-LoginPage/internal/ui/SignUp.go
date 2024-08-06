package ui

import (
	"LoginPage/internal/auth"
	"LoginPage/internal/config"
	"LoginPage/internal/models"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func SignUp(f io.Writer) {
	var dummyUser, user models.Users
	for {
		dummyUser.Username, dummyUser.Password = auth.CredInput()
		if !auth.IsStrongPassword(dummyUser.Password) {
			fmt.Println("Weak Password. Enter a stronger password (min. length 8 characters, at least 1 uppercase, at least 1 lowercase, at least 1 digit and at least 1 special character)")
		} else {
			break
		}
	}
	user.Username, user.Password = dummyUser.Username, dummyUser.Password

	user.Password = hashPassword(user.Password)

	//fmt.Println(user)

	for _, pass := range config.Users {
		if pass.Username == user.Username {
			fmt.Println("This username exists. Can't Signup")
			return
		}
	}
	writingString := fmt.Sprintf("\n%s:%s", user.Username, user.Password)

	_, err := f.Write([]byte(writingString))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Successfully Signed Up.")
}

func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
