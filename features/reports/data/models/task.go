package models

import "time"

// Task gotten from toggl
type Task struct {
	Description string
	Duration int
	Start time.Time
	End time.Time
	Project string
	ProjectColor string
	Tags []string
}