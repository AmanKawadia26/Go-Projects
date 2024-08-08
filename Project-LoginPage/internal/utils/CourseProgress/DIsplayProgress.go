package CourseProgress

import (
	"LoginPage/internal/models"
	"fmt"
	"os"
	"strings"
)

func CalculateProgress(username, courseFile string) (models.Progress, error) {
	courseContent, err := os.ReadFile(courseFile)
	if err != nil {
		return models.Progress{}, fmt.Errorf("error reading course file: %w", err)
	}

	completedFile := username + "_completed_modules.txt"
	completedContent, err := os.ReadFile(completedFile)
	if err != nil {
		return models.Progress{}, fmt.Errorf("error reading completed modules file: %w", err)
	}

	modules := strings.Split(string(courseContent), "\n")
	completedModules := strings.Split(string(completedContent), "\n")

	totalModules := len(modules)
	completedCount := 0
	for _, module := range completedModules {
		if module != "" {
			completedCount++
		}
	}

	completionPercentage := 0.0
	if totalModules > 0 {
		completionPercentage = (float64(completedCount) / float64(totalModules)) * 100
	}

	return models.Progress{
		TotalTasks:           totalModules,
		CompletedTasks:       completedCount,
		RemainingTasks:       totalModules - completedCount,
		CompletionPercentage: completionPercentage,
	}, nil
}

func DisplayProgress(username, courseFile string) {
	progress, err := CalculateProgress(username, courseFile)
	if err != nil {
		fmt.Println("Error calculating progress because there is no task history.")
		return
	}

	fmt.Printf("\nCourse Progress for %s:\n", username)
	fmt.Printf("Total Modules: %d\n", progress.TotalTasks)
	fmt.Printf("Completed Modules: %d\n", progress.CompletedTasks)
	fmt.Printf("Remaining Modules: %d\n", progress.RemainingTasks)
	fmt.Printf("Completion Percentage: %.2f%%\n", progress.CompletionPercentage)
}
