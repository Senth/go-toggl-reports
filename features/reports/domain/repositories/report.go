package repositories

import (
	"github.com/Senth/go-toggl-reports/core"
	"github.com/Senth/go-toggl-reports/features/reports/data/models"
)

// Report Repository for getting reports
type Report interface {
	GetTasks(start int, end int) (models.Task, core.Error)
}