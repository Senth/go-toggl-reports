package entities

import "time"

// Task Toggl task
type Task struct {
	Project *Project
	Name  string
	Start *time.Time
	End 	*time.Time
	// Duration in seconds
	Duration int
}