package DailyStatus

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ViewDailyStatus(statusFile string) {
	fmt.Println(cyan + "\n📊 Your Daily Status:" + reset)

	content, err := os.ReadFile(statusFile)
	if err != nil {
		fmt.Println(red + "\n❌ Error reading Daily Status. No tasks were completed previously." + reset)
		return
	}

	lines := strings.Split(string(content), "\n")

	if len(lines) == 0 || (len(lines) == 1 && lines[0] == "") {
		fmt.Println(cyan + "No tasks completed today. Keep up the good work!" + reset)
	} else {
		for i, line := range lines {
			if line != "" {
				// Extract the completed task and timestamp
				parts := strings.SplitN(line, " - ", 2)
				if len(parts) == 2 {
					task := strings.TrimPrefix(parts[0], "Completed: ")
					timestamp := parts[1]
					fmt.Printf(cyan+"%d. %s (Completed at: %s)\n"+reset, i+1, task, timestamp)
				}
			}
		}
	}

	fmt.Println(green + "\nPress Enter to continue..." + reset)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
