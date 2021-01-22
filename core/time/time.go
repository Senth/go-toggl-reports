package time

import "time"

// StartOfWeek returns the starting time for the specified week. I.e. Always Monday 00:00
func StartOfWeek(t time.Time) *time.Time {
	// Set to start of day, 00:00 (needs to be done before because of daylight savings time)
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0 ,0, t.Location())

	// Set to Monday
	if t.Weekday() != time.Monday {
		var diff int
		if t.Weekday() == time.Sunday {
			diff = 6
		} else {
			diff = int(t.Weekday()) - 1
		}
		t = t.Add(time.Hour * 24 * time.Duration(-diff))
	}

	// Set to start of day, 00:00
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0 ,0, t.Location())

	return &t
}