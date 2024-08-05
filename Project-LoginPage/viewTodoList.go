package main

import (
	"fmt"
	"os"
)

func viewTodoList(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading To-Do List:", err)
		return
	}
	fmt.Println("\nTo-Do List:")
	fmt.Println(string(content))
}
