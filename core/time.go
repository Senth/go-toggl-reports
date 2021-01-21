package core

import "time"

// StartOfWeek returns the starting time for the specified week. I.e. Always Monday 00:00
func StartOfWeek(t *time.Time) *time.Time {
	// Set to start of day, 00:00
	weekTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0 ,0, t.Location())

	// Set to Monday
	if weekTime.Weekday() != time.Monday {
		var diff int
		if weekTime.Weekday() == time.Sunday {
			diff = 6
		} else {
			diff = int(weekTime.Weekday()) - 1
		}
		weekTime = weekTime.Add(time.Hour * 24 * time.Duration(-diff))
	}

	return &weekTime
}