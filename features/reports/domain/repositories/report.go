package repositories

import (
	"time"

	"github.com/Senth/go-toggl-reports/features/reports/data/models"
)

// ReportInteractor Repository for getting reports
type ReportInteractor interface {
	GetTasks(start time.Time, end time.Time) (models.Task, error)
}