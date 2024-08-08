package ui

import (
	"LoginPage/internal/utils/CourseProgress"
	"LoginPage/internal/utils/DailyStatus"
	"LoginPage/internal/utils/ToDoList"
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
		fmt.Println(yellow + "\n📋 Task Management Menu:" + reset)
		fmt.Println(cyan + "1. 📝 Manage To-Do List" + reset)
		fmt.Println(cyan + "2. 📊 View Daily Status" + reset)
		fmt.Println(cyan + "3. 📈 View Daily Status Progress" + reset)
		fmt.Println(cyan + "4. 🔙 Go Back to Main Menu" + reset)

		fmt.Print(yellow + "\nPlease enter your choice (1-4): " + reset)
		choice1, _ := reader.ReadString('\n')
		choice1 = strings.TrimSpace(choice1)

		switch choice1 {
		case "1":
			handleManualTodoList(username, todoFile)
		case "2":
			handleDailyStatus(username, statusFile)
		case "3":
			CourseProgress.DisplayProgress(todoFile, statusFile)
		case "4":
			fmt.Println(green + "\n👋 Returning to main menu..." + reset)
			return
		default:
			fmt.Println(red + "\n❌ Invalid choice. Please try again." + reset)
		}
	}
}

func handleManualTodoList(username, todoFile string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(yellow + "\n📝 To-Do List Management:" + reset)
		fmt.Println(cyan + "1. 👀 View To-Do List" + reset)
		fmt.Println(cyan + "2. ➕ Add Task to To-Do List" + reset)
		fmt.Println(cyan + "3. 🔙 Go Back" + reset)

		fmt.Print(yellow + "\nPlease enter your choice (1-3): " + reset)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			ToDoList.ViewTodoList(todoFile)
		case "2":
			ToDoList.AddTask(todoFile)
		case "3":
			fmt.Println(green + "\n👋 Returning to Task Management Menu..." + reset)
			return
		default:
			fmt.Println(red + "\n❌ Invalid choice. Please try again." + reset)
		}
	}
}

func handleDailyStatus(username, statusFile string) {
	todoFile := username + "_todo.txt"
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(yellow + "\n📊 Daily Status Management:" + reset)
		fmt.Println(cyan + "1. 👀 View Daily Status" + reset)
		fmt.Println(cyan + "2. ✅ Mark Task as Completed" + reset)
		fmt.Println(cyan + "3. 🔙 Go Back" + reset)

		fmt.Print(yellow + "\nPlease enter your choice (1-3): " + reset)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			DailyStatus.ViewDailyStatus(statusFile)
		case "2":
			DailyStatus.MarkTaskAsCompleted(todoFile, username)
		case "3":
			fmt.Println(green + "\n👋 Returning to Task Management Menu..." + reset)
			return
		default:
			fmt.Println(red + "\n❌ Invalid choice. Please try again." + reset)
		}
	}
}
