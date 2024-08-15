package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	reset  = "\033[0m"
	cyan   = "\033[36m"
	yellow = "\033[33m"
	red    = "\033[31m"
	green  = "\033[32m"
)

func manageCourseProgress(username string) {
	courseFileName := username + "_course.txt"

	_, err := os.Stat(courseFileName)
	if os.IsNotExist(err) {
		fmt.Println(red + "Course file does not exist. Please sign up first." + reset)
		return
	}

	modules := loadCourseModules(courseFileName)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println()
		fmt.Println(cyan + "1. View Completed Modules" + reset)
		fmt.Println(cyan + "2. Mark Module as Completed" + reset)
		fmt.Println(cyan + "3. View Uncompleted Modules" + reset)
		fmt.Println(cyan + "4. Check Course Progress" + reset)
		fmt.Println(cyan + "5. Go Back" + reset)

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			viewCompletedModules(courseFileName)
		case "2":
			markModuleAsCompleted(courseFileName)
		case "3":
			viewUncompletedModules(modules, courseFileName)
		case "4":
			checkCourseProgress(courseFileName)
		case "5":
			return
		default:
			fmt.Println(red + "Invalid choice. Please try again." + reset)
		}
	}
}

func loadCourseModules(courseFileName string) map[string]bool {
	file, err := os.ReadFile(courseFileName)
	if err != nil {
		fmt.Println(red + "Error reading course file: " + err.Error() + reset)
		return nil
	}

	modules := make(map[string]bool)
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if line != "" && !strings.HasPrefix(line, "Completed: ") && line != "Modules:" && line != "Uncompleted Modules:" {
			modules[line] = false
		}
	}

	return modules
}

func viewCompletedModules(courseFileName string) {
	file, err := os.ReadFile(courseFileName)
	if err != nil {
		fmt.Println(red + "Error reading course file: " + err.Error() + reset)
		return
	}

	fmt.Println(green + "\nCompleted Modules:" + reset)
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Completed: ") {
			fmt.Println(strings.TrimPrefix(line, "Completed: "))
		}
	}
}

func markModuleAsCompleted(courseFileName string) {
	file, err := os.ReadFile(courseFileName)
	if err != nil {
		fmt.Println(red + "Error reading course file: " + err.Error() + reset)
		return
	}

	lines := strings.Split(string(file), "\n")
	modules := make(map[string]bool)
	for _, line := range lines {
		if line != "" && !strings.HasPrefix(line, "Completed: ") && line != "Modules:" && line != "Uncompleted Modules:" {
			modules[line] = false
		}
	}

	completedModules := loadCompletedModules(courseFileName)

	fmt.Println(cyan + "\nSelect a module to mark as completed:" + reset)
	var availableModules []string
	i := 1
	for _, module := range lines {
		if module != "" && !strings.HasPrefix(module, "Completed: ") && module != "Modules:" && module != "Uncompleted Modules:" {
			if _, completed := completedModules[module]; !completed {
				fmt.Printf(cyan+"%d. %s\n"+reset, i, module)
				availableModules = append(availableModules, module)
				i++
			}
		}
	}

	var choice int
	fmt.Print(cyan + "Enter module number: " + reset)
	fmt.Scan(&choice)

	if choice > 0 && choice <= len(availableModules) {
		selectedModule := availableModules[choice-1]
		if _, exists := modules[selectedModule]; exists {
			if err := appendToCourseFile(courseFileName, "Completed: "+selectedModule+"\n"); err != nil {
				fmt.Println(red + "Error updating course file: " + err.Error() + reset)
				return
			}
			fmt.Println(green + "Module marked as completed!" + reset)
		} else {
			fmt.Println(red + "Invalid module number." + reset)
		}
	} else {
		fmt.Println(red + "Invalid choice." + reset)
	}
}

func appendToCourseFile(courseFileName, content string) error {
	file, err := os.OpenFile(courseFileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func viewUncompletedModules(allModules map[string]bool, courseFileName string) {
	file, err := os.ReadFile(courseFileName)
	if err != nil {
		fmt.Println(red + "Error reading course file: " + err.Error() + reset)
		return
	}

	lines := strings.Split(string(file), "\n")
	completedModules := loadCompletedModules(courseFileName)

	fmt.Println(cyan + "\nUncompleted Modules:" + reset)
	var availableModules []string
	i := 1
	for _, module := range lines {
		if module != "" && !strings.HasPrefix(module, "Completed: ") && module != "Modules:" && module != "Uncompleted Modules:" {
			if _, completed := completedModules[module]; !completed {
				fmt.Printf(cyan+"%d. %s\n"+reset, i, module)
				availableModules = append(availableModules, module)
				i++
			}
		}
	}
}

func loadCompletedModules(courseFileName string) map[string]bool {
	file, err := os.ReadFile(courseFileName)
	if err != nil {
		fmt.Println(red + "Error reading course file: " + err.Error() + reset)
		return nil
	}

	completedModules := make(map[string]bool)
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Completed: ") {
			completedModules[strings.TrimPrefix(line, "Completed: ")] = true
		}
	}

	return completedModules
}

func checkCourseProgress(courseFileName string) {
	modules := loadCourseModules(courseFileName)
	completedModules := loadCompletedModules(courseFileName)

	totalModules := len(modules)
	if totalModules == 0 {
		totalModules = 10
	}
	completedCount := 0

	for module := range modules {
		if completedModules[module] {
			completedCount++
		}
	}

	progressPercentage := (float64(completedCount) / float64(totalModules)) * 100
	fmt.Printf(green+"Course Progress: %.2f%%\n"+reset, progressPercentage)
}
