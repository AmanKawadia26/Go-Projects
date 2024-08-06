package auth

import "fmt"

func Verification() bool {
	fmt.Println("\nAre you a WatchGuard employee or intern? (Enter 'y' or 'yes' to confirm): ")
	var answer string
	fmt.Scanf("%s", &answer)
	if answer == "y" || answer == "Y" || answer == "yes" || answer == "YES" {
		return true
	}
	return false
}
