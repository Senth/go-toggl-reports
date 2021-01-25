package repositories

import (
	"time"

	"github.com/Senth/go-toggl-reports/features/reports/data/models"
)

// TaskFetcher fetches reports from somewhere
type TaskFetcher interface {
	// Tasks fetch all tasks between start and end time
	Tasks(workspaceID int, start, end time.Time) ([]*models.Task, error)
}