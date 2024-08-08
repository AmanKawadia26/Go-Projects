package DailyStatus

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	reset  = "\033[0m"
	cyan   = "\033[36m"
	yellow = "\033[33m"
	green  = "\033[32m"
	red    = "\033[31m"
)

func MarkTaskAsCompleted(todoFile, username string) {
	statusFile := username + "_status.txt"

	// Check if the To-Do file exists
	if _, err := os.Stat(todoFile); os.IsNotExist(err) {
		fmt.Println(red + "\n❌ No tasks have been added. Please add a task to the To-Do list first." + reset)
		return
	}

	todoContent, err := os.ReadFile(todoFile)
	if err != nil {
		fmt.Println(red+"\n❌ Error reading To-Do List:"+reset, err)
		return
	}

	tasks := strings.Split(string(todoContent), "\n")
	if len(tasks) == 0 || (len(tasks) == 1 && tasks[0] == "") {
		fmt.Println(red + "\n❌ No tasks available to mark as completed. Please add tasks to your To-Do list." + reset)
		return
	}

	var completedTask string
	fmt.Println(yellow + "\nSelect a task to mark as completed:" + reset)
	for i, task := range tasks {
		if task != "" {
			fmt.Printf(cyan+"%d. %s\n"+reset, i+1, task)
		}
	}

	var taskNum int
	fmt.Print(yellow + "\nEnter task number: " + reset)
	fmt.Scan(&taskNum)

	if taskNum > 0 && taskNum <= len(tasks) {
		completedTask = tasks[taskNum-1]
		tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)

		err = os.WriteFile(todoFile, []byte(strings.Join(tasks, "\n")), 0644)
		if err != nil {
			fmt.Println(red+"\n❌ Error updating To-Do List:"+reset, err)
			return
		}

		statusFileHandle, err := os.OpenFile(statusFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println(red+"\n❌ Error opening Daily Status file:"+reset, err)
			return
		}
		defer statusFileHandle.Close()

		currentTime := time.Now().Format("02/01/2006 15:04:05")
		_, err = statusFileHandle.WriteString(fmt.Sprintf("Completed: %s - %s\n", completedTask, currentTime))
		if err != nil {
			fmt.Println(red+"\n❌ Error updating Daily Status:"+reset, err)
			return
		}

		fmt.Println(green + "\n✅ Task marked as completed and moved to Daily Status!" + reset)
	} else {
		fmt.Println(red + "\n❌ Invalid task number." + reset)
	}
}
