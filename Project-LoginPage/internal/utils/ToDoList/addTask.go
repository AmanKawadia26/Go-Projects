package ToDoList

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AddTask(filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task: ")

	task, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	task = strings.TrimSpace(task)

	if task == "" {
		fmt.Println("Task cannot be empty.")
		return
	}

	_, err = file.WriteString(task + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Task added successfully!")
}
