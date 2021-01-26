package datasources

import (
	"io/ioutil"
	"math"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Senth/go-toggl-reports/features/reports/data/models"
)

func TestNewTogglAPI(t *testing.T) {
	testData := TogglAPI{
		token: "my-token",
		userAgent: "my.email@example.com",
		dateFormatter: newTogglDate(),
	}

	result := NewTogglAPI(testData.token, testData.userAgent)

	if result.token != testData.token {
		t.Errorf("Token not equal. Expected: %v, Got: %v", testData.token, result.token)
	}

	if result.userAgent != testData.userAgent {
		t.Errorf("UserAgent not equal. Expected: %v, Got %v", testData.userAgent, result.token)
	}

	if result.dateFormatter == nil {
		t.Error("Dateformatter should be set but is nil")
	}
}

type testClient struct {}

func (c testClient) Do(req *http.Request) (*http.Response, error) {
	resp := &http.Response{
		StatusCode: 200,
		Body: ioutil.NopCloser(strings.NewReader(tasksJSON)),
	}

	return resp, nil
}

func TestTasks(t *testing.T) {
	Client = testClient{}
	stockholm, _ := time.LoadLocation("Europe/Stockholm")
	togglAPI := NewTogglAPI("my-token", "my.email@example.com")

	test := struct{
		start time.Time
		end time.Time
		expect []*models.Task 
	} {
		start: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		end: time.Date(2021, 2, 2, 0, 0, 0, 0, time.UTC),
		expect: []*models.Task{
			{
				Description: "Studera React Redux",
				Project: "Work",
				ProjectColor: "#990099",
				Start: time.Date(2021, 1, 25, 7, 49, 46, 0, stockholm),
				End: time.Date(2021, 1, 25, 8, 34, 31, 0, stockholm),
				Duration: int(math.Round(float64(2685661) / 1000)),
				Tags: []string{"core-study"},
			},
			{
				Description: "Add SqliteGateway tests #test",
				Project: "My Musical Repertoire",
				ProjectColor: "#465bb3",
				Start: time.Date(2021, 1, 24, 18, 39, 29, 0, stockholm),
				End: time.Date(2021, 1, 24, 20, 7, 43, 0, stockholm),
				Duration: 5294,
				Tags: []string{},
			},
			{
				Description: "Reddit helper",
				Project: "Personal Development",
				ProjectColor: "#2da608",
				Start: time.Date(2021, 1, 24, 11, 11, 38, 0, stockholm),
				End: time.Date(2021, 1, 24, 11, 53, 38, 0, stockholm),
				Duration: 2520,
				Tags: []string{"helper"},
			},
		},
	}

	result, _ := togglAPI.Tasks(1, test.start, test.end)

	if len(test.expect) != len(result) {
		t.Errorf("Expected length %v not equal to gotten length %v", len(test.expect), len(result))
	}

	// Check that they are equal (can't use deep equality here)
	for i := 0; i < len(result); i++ {
		expectedTask := test.expect[i]
		resultTask := result[i]

		if expectedTask.Description != resultTask.Description {
			t.Errorf("Expected description (%v) not equal to gotten (%v)", expectedTask.Description, resultTask.Description)
		}

		if expectedTask.Project != resultTask.Project{
			t.Errorf("Expected project (%v) not equal to gotten (%v)", expectedTask.Project, resultTask.Project)
		}

		if expectedTask.ProjectColor != resultTask.ProjectColor {
			t.Errorf("Expected project color (%v) not equal to gotten (%v)", expectedTask.ProjectColor, resultTask.ProjectColor)
		}

		if expectedTask.Duration != resultTask.Duration {
			t.Errorf("Expected duration (%v) not equal to gotten (%v)", expectedTask.Duration, resultTask.Duration)
		}

		if expectedTask.Start.Format(time.RFC3339) != resultTask.Start.Format(time.RFC3339) {
			t.Errorf("Expected start time (%v) not equal to gotten (%v)", expectedTask.Start, resultTask.Start)
		}

		if expectedTask.End.Format(time.RFC3339) != resultTask.End.Format(time.RFC3339) {
			t.Errorf("Expected start time (%v) not equal to gotten (%v)", expectedTask.Start, resultTask.Start)
		}

		if !reflect.DeepEqual(expectedTask.Tags, resultTask.Tags) {
			t.Errorf("Expected tags (%v) not equal to gotten (%v)", expectedTask.Tags, resultTask.Tags)
		}
	}
}

var tasksJSON string = `{
  "total_grand": 143528031,
  "total_billable": null,
  "total_currencies": [
    {
      "currency": null,
      "amount": null
    }
  ],
  "total_count": 24,
  "per_page": 50,
  "data": [
    {
      "id": 1854079901,
      "pid": 14921666,
      "tid": null,
      "uid": 301760,
      "description": "Studera React Redux",
      "start": "2021-01-25T07:49:46+01:00",
      "end": "2021-01-25T08:34:31+01:00",
      "updated": "2021-01-25T08:34:31+01:00",
      "dur": 2685661,
      "user": "Matteus Magnusson",
      "use_stop": true,
      "client": null,
      "project": "Work",
      "project_color": "0",
      "project_hex_color": "#990099",
      "task": null,
      "billable": null,
      "is_billable": false,
      "cur": null,
      "tags": ["core-study"]
    },
    {
      "id": 1853694547,
      "pid": 164527893,
      "tid": null,
      "uid": 301760,
      "description": "Add SqliteGateway tests #test",
      "start": "2021-01-24T18:39:29+01:00",
      "end": "2021-01-24T20:07:43+01:00",
      "updated": "2021-01-24T20:07:44+01:00",
      "dur": 5294000,
      "user": "Matteus Magnusson",
      "use_stop": true,
      "client": null,
      "project": "My Musical Repertoire",
      "project_color": "0",
      "project_hex_color": "#465bb3",
      "task": null,
      "billable": null,
      "is_billable": false,
      "cur": null,
      "tags": []
    },
    {
      "id": 1853539739,
      "pid": 10563913,
      "tid": null,
      "uid": 301760,
      "description": "Reddit helper",
      "start": "2021-01-24T11:11:38+01:00",
      "end": "2021-01-24T11:53:38+01:00",
      "updated": "2021-01-24T11:53:38+01:00",
      "dur": 2520000,
      "user": "Matteus Magnusson",
      "use_stop": true,
      "client": null,
      "project": "Personal Development",
      "project_color": "0",
      "project_hex_color": "#2da608",
      "task": null,
      "billable": null,
      "is_billable": false,
      "cur": null,
      "tags": ["helper"]
    }
  ]
}
`