package datasources

import (
	"os"
	"testing"
	"time"
)

var formatter togglDate

func TestMain(m *testing.M) {
	formatter = newTogglDate()
	os.Exit(m.Run())
}

func TestNewTogglDate(t *testing.T) {
	togglDate := newTogglDate();

	if togglDate == nil {
		t.Error("NewTogglDate() should not return nil")
	}
}

func TestToTogglDate(t *testing.T) {
	stockholm, _ := time.LoadLocation("Europe/Stockholm")
	losAngeles, _ := time.LoadLocation("America/Los_Angeles")

	testData := []struct {
		time time.Time
		expect string
	} {
		{
			time: time.Date(2021, 1, 15, 22, 15, 33, 0, time.UTC),
			expect: "2021-01-15T22:15:33+00:00",
		},
		{
			time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			expect: "2021-01-01T00:00:00+00:00",
		},
		{
			time: time.Date(2021, 3, 14, 5, 23, 12, 0, stockholm),
			expect: "2021-03-14T05:23:12+01:00",
		},
		{
			time: time.Date(2021, 3, 10, 5, 23, 12, 0, losAngeles),
			expect: "2021-03-10T05:23:12-08:00",
		},
	}

	for _, test := range testData {
		result := formatter.toTogglDate(test.time)

		if result != test.expect {
			t.Errorf("toTogglDate(%v) should return %v but returned %v", test.time, test.expect, result)
		}
	}
}

func TestFromTogglDateInvalidFormat(t *testing.T) {
	testData := []string{
		"Tue Mar 3 23:15:05 2006",
		"2018-03-14 18:13",
		"2018-03-14 18:13",
		"2018-03-14T18:13:17 UTC",
		"2018-03-14T18:13:17Z01:00",
	}

	for _, date := range testData {
		_, err := formatter.fromTogglDate(date)
		
		if err == nil {
			t.Errorf("fromTogglDate(%v) should return an error because it's not a valid ISO8601 date", date)
		}
	}
}

func TestFromTogglDate(t *testing.T) {
	stockholm, _ := time.LoadLocation("Europe/Stockholm")
	losAngeles, _ := time.LoadLocation("America/Los_Angeles")

	testData := []struct {
		expect time.Time
		date string
	} {
		{
			date: "2021-01-15T22:15:33+00:00",
			expect: time.Date(2021, 1, 15, 22, 15, 33, 0, time.UTC),
		},
		{
			date: "2021-01-01T00:00:00+00:00",
			expect: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			date: "2021-03-14T05:23:12+01:00",
			expect: time.Date(2021, 3, 14, 5, 23, 12, 0, stockholm),
		},
		{
			date: "2021-03-10T05:23:12-08:00",
			expect: time.Date(2021, 3, 10, 5, 23, 12, 0, losAngeles),
		},
	}

	for _, test := range testData {
		result, _ := formatter.fromTogglDate(test.date)

		if result.Format(time.RFC3339) != test.expect.Format(time.RFC3339) {
			t.Errorf("fromTogglDate(%v) should return %v but returned %v", test.date, test.expect, result)
		}
	}
}