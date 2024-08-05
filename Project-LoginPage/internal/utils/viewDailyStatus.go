package utils

import (
	"fmt"
	"os"
)

func ViewDailyStatus(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading Daily Status:", err)
		return
	}
	fmt.Println("\nDaily Status:")
	fmt.Println(string(content))
}
