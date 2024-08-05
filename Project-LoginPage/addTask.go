package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func addTask(filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task: ")

	// Read the entire line, including any spaces
	task, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Remove the trailing newline and any other whitespace
	task = strings.TrimSpace(task)

	if task == "" {
		fmt.Println("Task cannot be empty.")
		return
	}

	// Write the task to the file
	_, err = file.WriteString(task + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Task added successfully!")
}
