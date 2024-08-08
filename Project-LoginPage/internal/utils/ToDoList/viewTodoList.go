package ToDoList

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ViewTodoList(todoFile string) {
	fmt.Println(yellow + "\n📝 Your To-Do List:" + reset)

	content, err := os.ReadFile(todoFile)
	if err != nil {
		fmt.Println(red + "\n❌ Error reading To-Do List because there is no task history." + reset)
		return
	}

	tasks := strings.Split(string(content), "\n")
	if len(tasks) == 0 || (len(tasks) == 1 && tasks[0] == "") {
		fmt.Println(cyan + "Your To-Do List is empty. Add some tasks to get started!" + reset)
	} else {
		for i, task := range tasks {
			if task != "" {
				fmt.Printf(cyan+"%d. %s\n"+reset, i+1, task)
			}
		}
	}

	fmt.Println(green + "\nPress Enter to continue..." + reset)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
