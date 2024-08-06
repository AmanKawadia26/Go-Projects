package ui

import (
	"LoginPage/internal/utils/CourseProgress"
	"LoginPage/internal/utils/DailyStatus"
	"LoginPage/internal/utils/ToDoList"
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func TodoList(username string) {
	todoFile := username + "_todo.txt"
	statusFile := username + "_status.txt"
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n1. To-Do List")
		fmt.Println("2. Daily Status")
		fmt.Println("3. View Course Progress")
		fmt.Println("4. Go Back")

		choice1, _ := reader.ReadString('\n')
		choice1 = strings.TrimSpace(choice1)

		switch choice1 {
		case "1":
			fmt.Println("\n1. View To-Do List")
			fmt.Println("2. Add Task to To-Do List")
			fmt.Println("3. Go back")
			fmt.Print("Enter your choice: ")

			choice2, _ := reader.ReadString('\n')
			choice2 = strings.TrimSpace(choice2)

			var wg sync.WaitGroup

			switch choice2 {
			case "1":
				wg.Add(1)
				go func() {
					defer wg.Done()
					ToDoList.ViewTodoList(todoFile)
				}()
			case "2":
				wg.Add(1)
				go func() {
					defer wg.Done()
					ToDoList.AddTask(todoFile)
				}()
			case "3":
				wg.Wait()
				break
			default:
				fmt.Println("Invalid choice. Please try again.")
			}
			wg.Wait()
		case "2":
			fmt.Println("1. View Daily Status")
			fmt.Println("2. Mark Task as Completed")
			fmt.Println("3. Go back")
			fmt.Print("Enter your choice: ")
			choice2, _ := reader.ReadString('\n')
			choice2 = strings.TrimSpace(choice2)

			var wg sync.WaitGroup

			switch choice2 {
			case "1":
				wg.Add(1)
				go func() {
					defer wg.Done()
					DailyStatus.ViewDailyStatus(statusFile)
				}()
			case "2":
				wg.Add(1)
				go func() {
					defer wg.Done()
					DailyStatus.MarkTaskAsCompleted(todoFile, statusFile)
				}()
			case "3":
				wg.Wait() // Wait for any ongoing tasks to complete
				break
			default:
				fmt.Println("Invalid choice. Please try again.")
			}
			wg.Wait()
		case "3":
			CourseProgress.DisplayProgress(todoFile, statusFile)
		}
	}
}
