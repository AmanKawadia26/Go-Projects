package main

import (
	"fmt"
	"os"
)

func viewDailyStatus(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading Daily Status:", err)
		return
	}
	fmt.Println("\nDaily Status:")
	fmt.Println(string(content))
}
