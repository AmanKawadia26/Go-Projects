package auth

import (
	"regexp"
	"strings"
)

func IsStrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false
	}

	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false
	}

	if !regexp.MustCompile(`\d`).MatchString(password) {
		return false
	}

	if !regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password) {
		return false
	}

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
