package auth

import "fmt"

func Verification() bool {
	fmt.Println("\nAre you a WatchGuard employee or intern?")
	var answer string
	fmt.Scanf("%s", &answer)
	if answer == "y" || answer == "Y" || answer == "yes" || answer == "YES" {
		return true
	}
	return false
}
