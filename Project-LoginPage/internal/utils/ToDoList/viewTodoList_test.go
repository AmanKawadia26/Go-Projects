package ToDoList

import (
	"os"
	"testing"
)

func TestViewTodoList(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		expected    string
		wantErr     bool
	}{
		{
			name:        "View list with tasks",
			fileContent: "Task 1\nTask 2\nTask 3\n",
			expected:    "📝 Your To-Do List:\n1. Task 1\n2. Task 2\n3. Task 3\nPress Enter to continue...",
			wantErr:     false,
		},
		{
			name:        "View empty list",
			fileContent: "",
			expected:    "📝 Your To-Do List:\nYour To-Do List is empty. Add some tasks to get started!\nPress Enter to continue...",
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile, err := os.CreateTemp("", "todo_test_*.txt")
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer os.Remove(tmpFile.Name())

			if !tt.wantErr {
				if _, err := tmpFile.WriteString(tt.fileContent); err != nil {
					t.Fatalf("Failed to write to temp file: %v", err)
				}
				tmpFile.Close()
			} else {
				os.Remove(tmpFile.Name())
			}

		})
	}
}
