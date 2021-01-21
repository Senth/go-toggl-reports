package entities

// Report generated and structured toggl report
type Report struct {
	Tasks []*Task
	ByProject []*Project
	ByTags map[string]*Task
}