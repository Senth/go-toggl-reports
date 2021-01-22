package usecases

import (
	"time"

	coretime "github.com/Senth/go-toggl-reports/core/time"
	"github.com/Senth/go-toggl-reports/features/reports/data/repositories"
	"github.com/Senth/go-toggl-reports/features/reports/domain/entities"
)

// Reports contains the access repositories
type Reports struct {
	repo repositories.TaskFetcher
}

// NewReports Creates a new reports function
func NewReports(r repositories.TaskFetcher) *Reports {
	return &Reports{
		repo: r,
	}
}

// Between Get the report (tasks) between two dates
func (r *Reports) Between(start, end *time.Time) (report *entities.Report, err error) {
	return &entities.Report{}, nil
}

// ThisWeek Get the report (tasks) for this week
func (r *Reports) ThisWeek() (report *entities.Report, err error) {
	now := time.Now()
	return r.Between(coretime.StartOfWeek(now), coretime.EndOfWeek(now))
}