package datasources

import "time"

const iso8601DateFormat string = "2006-01-02T15:04:05-07:00"

type togglDate interface {
	// toTogglDate returns a formatted string that can be used in Toggl API requests
	toTogglDate(t time.Time) string

	// fromTogglDate converts a Toggl API date string into golang's time.Time
	// Returns an error if the date couldn't be parsed
	fromTogglDate(date string) (time.Time, error)
}

type togglDateFormatter struct {}

func newTogglDate() togglDate {
	return &togglDateFormatter{}
}

func (togglDateFormatter) toTogglDate(t time.Time) string {
	return t.Format(iso8601DateFormat)
}

func (togglDateFormatter) fromTogglDate(date string) (time.Time, error) {
	return time.Parse(iso8601DateFormat, date)
}