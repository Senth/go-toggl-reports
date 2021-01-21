package usecases

import (
	"time"

	"github.com/Senth/go-toggl-reports/core"
	"github.com/Senth/go-toggl-reports/features/reports/domain/entities"
)

// GetReports Use case that get all reports between two dates
func GetReports(start *time.Time, end *time.Time) (report *entities.Report, err *core.Error) {


	return
}

func GetReportsThisWeek() (report *entities.Report, err *core.Error) {
	// TODO
	// now := time.Now()

	return
}