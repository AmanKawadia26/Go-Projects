package utils

import (
	"fmt"
	"os"
)

func ViewTodoList(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading To-Do List:", err)
		return
	}
	fmt.Println("\nTo-Do List:")
	fmt.Println(string(content))
}
