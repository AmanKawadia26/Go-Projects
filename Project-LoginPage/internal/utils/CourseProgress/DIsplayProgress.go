package CourseProgress

import (
	"fmt"
	"os"
	"strings"
)

type Progress struct {
	TotalTasks           int
	CompletedTasks       int
	RemainingTasks       int
	CompletionPercentage float64
}

func CalculateProgress(todoFile, statusFile string) (Progress, error) {
	todoContent, err := os.ReadFile(todoFile)
	if err != nil {
		//fmt.Println(err)
		return Progress{}, fmt.Errorf("error reading To-Do List: %w", err)
	}

	statusContent, err := os.ReadFile(statusFile)
	if err != nil {
		return Progress{}, fmt.Errorf("error reading Daily Status: %w", err)
	}

	tasks := strings.Split(string(todoContent), "\n")
	completedTasks := strings.Split(string(statusContent), "\n")

	totalTasks := len(tasks)
	remainingTasks := totalTasks
	completedCount := 0

	for _, task := range completedTasks {
		if strings.HasPrefix(task, "Completed: ") {
			completedCount++
			remainingTasks--
		}
	}

	completionPercentage := 0.0
	if totalTasks > 0 {
		completionPercentage = (float64(completedCount) / float64(totalTasks)) * 100
	}

	return Progress{
		TotalTasks:           totalTasks,
		CompletedTasks:       completedCount,
		RemainingTasks:       remainingTasks,
		CompletionPercentage: completionPercentage,
	}, nil
}

// DisplayProgress prints the progress in a human-readable format
func DisplayProgress(todoFile, statusFile string) {
	progress, err := CalculateProgress(todoFile, statusFile)
	if err != nil {
		fmt.Println("Error calculating progress:", err)
		return
	}

	fmt.Printf("\nCourse Progress:\n")
	fmt.Printf("Total Tasks: %d\n", progress.TotalTasks)
	fmt.Printf("Completed Tasks: %d\n", progress.CompletedTasks)
	fmt.Printf("Remaining Tasks: %d\n", progress.RemainingTasks)
	fmt.Printf("Completion Percentage: %.2f%%\n", progress.CompletionPercentage)
}
