package ui

import (
	"LoginPage/internal/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TodoList(username string) {
	todoFile := username + "_todo.txt"
	statusFile := username + "_status.txt"
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n1. View To-Do List")
		fmt.Println("2. Add Task to To-Do List")
		fmt.Println("3. View Daily Status")
		fmt.Println("4. Mark Task as Completed")
		fmt.Println("5. Go back")
		fmt.Print("Enter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			utils.ViewTodoList(todoFile)
		case "2":

			utils.AddTask(todoFile)
		case "3":

			utils.ViewDailyStatus(statusFile)
		case "4":

			utils.MarkTaskAsCompleted(todoFile, statusFile)
		case "5":

			return
		default:

			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
