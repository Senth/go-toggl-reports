package time_test

import (
	"testing"
	"time"

	coretime "github.com/Senth/go-toggl-reports/core/time"
)


type startOfWeekResult struct {
	in time.Time
	expected time.Time
}

var (
	stockholm, _ = time.LoadLocation("Europe/Stockholm")
	
)

var startOfWeekResults = []startOfWeekResult{
	{in: time.Date(2021, 1, 22, 15, 35, 15, 36, stockholm), expected: time.Date(2021, 1, 18, 0, 0, 0, 0, stockholm)},
	{in: time.Date(2021, 1, 24, 23, 59, 15, 36, stockholm), expected: time.Date(2021, 1, 18, 0, 0, 0, 0, stockholm)},
	{in: time.Date(2021, 3, 28, 23, 59, 59, 5, stockholm), expected: time.Date(2021, 3, 22, 0, 0, 0, 0, stockholm)},
	{in: time.Date(2020, 10, 25, 0, 0, 0, 0, stockholm), expected: time.Date(2020, 10, 19, 0, 0, 0, 0, stockholm)},
	{in: time.Date(2020, 10, 25, 23, 59, 59, 0, stockholm), expected: time.Date(2020, 10, 19, 0, 0, 0, 0, stockholm)},
}

func TestStartOfWeek(t *testing.T) {
	for _, test := range startOfWeekResults {
		result := coretime.StartOfWeek(test.in)

		if *result != test.expected {
			t.Fatalf("Expected %v, Got: %v", test.expected, *result)
		}
	}
}