package ToDoList

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	cyan   = "\033[36m" // Cyan
	reset  = "\033[0m"
	yellow = "\033[33m"
	green  = "\033[32m"
	red    = "\033[31m"
)

func AddTask(todoFile string) {
	file, err := os.OpenFile(todoFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(red+"\n❌ Error opening file:"+reset, err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print(yellow + "\nEnter your new task: " + reset)

	task, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(red+"\n❌ Error reading input:"+reset, err)
		return
	}

	task = strings.TrimSpace(task)

	if task == "" {
		fmt.Errorf(red + "\n❌ Task cannot be empty." + reset)
		return
	}

	_, err = file.WriteString(task + "\n")
	if err != nil {
		fmt.Println(red+"\n❌ Error writing to file:"+reset, err)
		return
	}

	fmt.Println(green + "\n✅ Task added successfully!" + reset)
}
