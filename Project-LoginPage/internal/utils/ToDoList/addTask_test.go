package ToDoList

import (
	"os"
	"strings"
	"testing"
)

func TestAddTask(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "Add valid task",
			input:    "Test Task\n",
			expected: "Test Task\n",
			wantErr:  false,
		},
		{
			name:     "Add empty task",
			input:    "\n",
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile, err := os.CreateTemp("", "todo_test_*.txt")
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer os.Remove(tmpFile.Name())

			r, w, _ := os.Pipe()
			go func() {
				defer w.Close()
				w.Write([]byte(tt.input))
			}()

			oldStdin := os.Stdin
			os.Stdin = r

			defer func() { os.Stdin = oldStdin }()

			if tt.wantErr && tt.name == "Simulate file write error" {
				AddTask("/invalid/path/to/todo_test.txt")
			} else {
				AddTask(tmpFile.Name())
			}

			fileContent, err := os.ReadFile(tmpFile.Name())
			if err != nil && !tt.wantErr {
				t.Fatalf("Failed to read temp file: %v", err)
			}

			if strings.TrimSpace(string(fileContent)) != strings.TrimSpace(tt.expected) {
				t.Errorf("Expected file content %q, but got %q", tt.expected, fileContent)
			}
		})

	}
}
