package time_test

import (
	"testing"
	"time"

	coretime "github.com/Senth/go-toggl-reports/core/time"
)


type timeResult struct {
	in time.Time
	expected time.Time
}

var (
	stockholm, _ = time.LoadLocation("Europe/Stockholm")
	
)

func TestStartOfWeek(t *testing.T) {
	var startOfWeekResults = []timeResult{
		{in: time.Date(2021, 1, 22, 15, 35, 15, 36, stockholm), expected: time.Date(2021, 1, 18, 0, 0, 0, 0, stockholm)},
		{in: time.Date(2021, 1, 24, 23, 59, 15, 36, stockholm), expected: time.Date(2021, 1, 18, 0, 0, 0, 0, stockholm)},
		{in: time.Date(2021, 3, 28, 23, 59, 59, 5, stockholm), expected: time.Date(2021, 3, 22, 0, 0, 0, 0, stockholm)},
		{in: time.Date(2020, 10, 25, 0, 0, 0, 0, stockholm), expected: time.Date(2020, 10, 19, 0, 0, 0, 0, stockholm)},
		{in: time.Date(2020, 10, 25, 23, 59, 59, 0, stockholm), expected: time.Date(2020, 10, 19, 0, 0, 0, 0, stockholm)},
	}

	for _, test := range startOfWeekResults {
		result := coretime.StartOfWeek(test.in)

		if *result != test.expected {
			t.Errorf("Expected %v, Got: %v", test.expected, *result)
		}
	}
}

func TestEndOfWeek(t *testing.T) {
	const nanoMax int = int(time.Second - time.Nanosecond)

	var endOfWeekResults = []timeResult{
		{in: time.Date(2021, 1, 22, 15, 35, 15, 36, stockholm), expected: time.Date(2021, 1, 24, 23, 59, 59, nanoMax, stockholm)},
		{in: time.Date(2021, 1, 18, 15, 35, 15, 36, stockholm), expected: time.Date(2021, 1, 24, 23, 59, 59, nanoMax, stockholm)},
		{in: time.Date(2021, 1, 24, 15, 35, 15, 36, stockholm), expected: time.Date(2021, 1, 24, 23, 59, 59, nanoMax, stockholm)},
		{in: time.Date(2021, 3, 22, 0, 0, 0, 0, stockholm), expected: time.Date(2021, 3, 28, 23, 59, 59, nanoMax, stockholm)},
		{in: time.Date(2021, 3, 22, 23, 59, 59, 0, stockholm), expected: time.Date(2021, 3, 28, 23, 59, 59, nanoMax, stockholm)},
		{in: time.Date(2020, 10, 19, 0, 0, 0, 0, stockholm), expected: time.Date(2020, 10, 25, 23, 59, 59, nanoMax, stockholm)},
		{in: time.Date(2020, 10, 19, 23, 59, 59, 0, stockholm), expected: time.Date(2020, 10, 25, 23, 59, 59, nanoMax, stockholm)},
	}

	for _, test := range endOfWeekResults {
		result := coretime.EndOfWeek(test.in)

		if *result != test.expected {
			t.Errorf("Expected %v, Got: %v", test.expected, *result)
		}
	}
}