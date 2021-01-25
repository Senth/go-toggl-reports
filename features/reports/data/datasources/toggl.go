package datasources

import (
	"time"

	"github.com/Senth/go-toggl-reports/features/reports/data/models"
)

// Toggl API
type Toggl struct{}

func (api Toggl) Tasks(start, end time.Time) (*[]models.Task, error) {
	// TODO
	return nil, nil
}