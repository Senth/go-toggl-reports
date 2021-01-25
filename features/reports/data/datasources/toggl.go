package datasources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Senth/go-toggl-reports/core/consts"
	"github.com/Senth/go-toggl-reports/features/reports/data/models"
)

// TogglAPI Instance to a new toggl api
type TogglAPI struct{
	token string
	userAgent string
	dateFormatter togglDate;
}

// NewTogglAPI creates a new client for Toggl API
func NewTogglAPI(token, email string) TogglAPI {
	return TogglAPI{
		token: token,
		userAgent: email,
		dateFormatter: newTogglDate(),
	}
}

// Tasks Get all tasks betwene specified dates
func (api TogglAPI) Tasks(workspaceID int, start, end time.Time) ([]*models.Task, error) {
	// TODO use start and end times

	obj, err := api.request("GET", "details")
	if err != nil {
		return nil, err
	}

	// Convert object to []*models.Task
	tasks := []*models.Task{}

	dataTasks := obj["data"].([]interface{})

	for _, taskInterface := range dataTasks {
		dataTask := taskInterface.(map[string]interface{})
		start, err := api.dateFormatter.fromTogglDate(dataTask["start"].(string))
		if err != nil {
			return nil, err
		}
		end, err := api.dateFormatter.fromTogglDate(dataTask["end"].(string))
		if err != nil {
			return nil, err
		}

		tags := []string{}
		for _, dataTag := range dataTask["tags"].([]interface{}) {
			tag := dataTag.(string)
			tags = append(tags, tag)
		}

		task := &models.Task{
			Description: dataTask["description"].(string),
			Project: dataTask["project"].(string),
			ProjectColor: dataTask["project_hex_color"].(string),
			Start: start,
			End: end,
			Duration: int(dataTask["dur"].(float64)),
			Tags: tags,
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// HTTPClient Interface for the Client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	// Client the HTTPClient that is used to make a call
	Client HTTPClient = &http.Client{}
)

func (api TogglAPI) request(method string, route string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%v/%v", consts.TogglReportAPIURL, route)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	// Set authentication
	req.SetBasicAuth(api.token, "api_token")

	// Execute request
	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, readError := ioutil.ReadAll(resp.Body)
	if readError != nil {
		return nil, fmt.Errorf("Failed to read response body")
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("The %s request against %s failed (%s): %s", req.Method, req.URL, resp.Status, content)
	}

	var v interface{}
	json.Unmarshal(content, &v)

	var obj map[string]interface{} = v.(map[string]interface{})

	return obj, nil
}