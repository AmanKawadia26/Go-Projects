package models

type Users struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
	Age       int
	MobileNo  string
}

type Progress struct {
	TotalTasks           int
	CompletedTasks       int
	RemainingTasks       int
	CompletionPercentage float64
}
