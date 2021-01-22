package entities

// Project containing tasks sorted by tag name
type Project struct {
	Name string
	Color string
	Tasks []*Task
}