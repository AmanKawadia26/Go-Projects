package main

import (
	"fmt"
	"os"
	"strings"
)

func markTaskAsCompleted(todoFile, statusFile string) {
	// Read To-Do List
	todoContent, err := os.ReadFile(todoFile)
	if err != nil {
		fmt.Println("Error reading To-Do List:", err)
		return
	}

	tasks := strings.Split(string(todoContent), "\n")
	var completedTask string

	// Display tasks and get user input
	fmt.Println("\nSelect a task to mark as completed:")
	for i, task := range tasks {
		if task != "" {
			fmt.Printf("%d. %s\n", i+1, task)
		}
	}

	var taskNum int
	fmt.Print("Enter task number: ")
	fmt.Scan(&taskNum)

	if taskNum > 0 && taskNum <= len(tasks) {
		completedTask = tasks[taskNum-1]
		// Remove task from To-Do List
		tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)

		// Write updated To-Do List
		err = os.WriteFile(todoFile, []byte(strings.Join(tasks, "\n")), 0644)
		if err != nil {
			fmt.Println("Error updating To-Do List:", err)
			return
		}

		// Append completed task to Daily Status
		statusFile, err := os.OpenFile(statusFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("Error opening Daily Status file:", err)
			return
		}
		defer statusFile.Close()

		_, err = statusFile.WriteString("Completed: " + completedTask + "\n")
		if err != nil {
			fmt.Println("Error updating Daily Status:", err)
			return
		}

		fmt.Println("Task marked as completed and moved to Daily Status!")
	} else {
		fmt.Println("Invalid task number.")
	}
}
