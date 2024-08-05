package auth

import (
	"regexp"
	"strings"
)

// IsStrongPassword validates if the given password is strong based on predefined criteria.
func IsStrongPassword(password string) bool {
	// Criteria 1: Length check
	if len(password) < 8 {
		return false
	}

	// Criteria 2: Uppercase letter check
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false
	}

	// Criteria 3: Lowercase letter check
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false
	}

	// Criteria 4: Digit check
	if !regexp.MustCompile(`\d`).MatchString(password) {
		return false
	}

	// Criteria 5: Special character check
	if !regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password) {
		return false
	}

	// Additional criteria: Check for common patterns (optional)
	commonPatterns := []string{
		"password", "123456", "qwerty", "abc123", "password1",
	}
	for _, pattern := range commonPatterns {
		if strings.Contains(strings.ToLower(password), pattern) {
			return false
		}
	}

	// If all checks pass
	return true
}
