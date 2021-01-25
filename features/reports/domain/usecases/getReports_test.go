package usecases_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/Senth/go-toggl-reports/features/reports/data/models"
	"github.com/Senth/go-toggl-reports/features/reports/domain/entities"
	"github.com/Senth/go-toggl-reports/features/reports/domain/usecases"
)


var stockholm, _ = time.LoadLocation("Europe/Stockholm")

type protoRepo struct {}

var modelTasks = &[]models.Task{
	{
		Description: "Coding",
		Duration:     12557,
		Start:        time.Date(2021, 1, 2, 0, 0, 0, 0, stockholm),
		End:          time.Date(2021, 1, 2, 12, 0, 0, 0, stockholm),
		Project:      "Work",
		ProjectColor: "#ffffff",
		Tags:         []string{
			"Blue",
			"Red",
		},
	},
	{
		Description:  "Administration",
		Duration:     5568,
		Start:        time.Date(2021, 1, 3, 0, 0, 0, 0, stockholm),
		End:          time.Date(2021, 1, 3, 12, 59, 0, 0, stockholm),
		Project:      "Work",
		ProjectColor: "#ffffff",
		Tags:         []string{
			"Red",
		},
	},
	{
		Description:  "Workout",
		Duration:     57889,
		Start:        time.Date(2021, 1, 4, 0, 0, 0, 0, stockholm),
		End:          time.Date(2021, 1, 4, 3, 0, 0, 0, stockholm),
		Project:      "Personal",
		ProjectColor: "#00ff00",
		Tags:         []string{},
	},
}
var projects = map[string]*entities.Project {
	"Work": {
		Name: "Work",
		Color: "#ffffff",
	},
	"Personal": {
		Name: "Personal",
		Color: "#00ff00",
	},
}
var tasks = []*entities.Task{}
var tags = map[string][]*entities.Task{}

func generateTasks() {
	for _, modelTask := range *modelTasks {
		project := projects[modelTask.Project]

		t := entities.Task{
			Project:	project,
			Name:     modelTask.Description,
			Start:    modelTask.Start,
			End:      modelTask.End,
			Duration: modelTask.Duration,
			Tags: 		modelTask.Tags,
		}
		tasks = append(tasks, &t)
		project.Tasks = append(project.Tasks, &t)
	
		for _, tag := range modelTask.Tags {
			if _, ok := tags[tag]; !ok {
				tags[tag] = make([]*entities.Task, 0, 1)
			}
			tags[tag] = append(tags[tag], &t)
		}
		
	}
}

func (protoRepo) Tasks(start, end time.Time) (*[]models.Task, error) {
	return modelTasks, nil
}

var (
	repo = protoRepo{}
	usecase = usecases.NewReports(&repo)
)

func TestBetween(t *testing.T) {
	generateTasks()

	type testData struct {
		start time.Time
		end time.Time
		expected entities.Report
	}

	data := testData {
		start: time.Date(2021, 1, 1, 0, 0, 0, 0, stockholm),
		end: time.Date(2021, 3, 1, 0, 0, 0, 0, stockholm),
		expected: entities.Report{
			Tasks:  tasks,
			ByProject: projects,
			ByTag:    tags,
		},
	}

	r, err := usecase.Between(data.start, data.end)
	if err != nil {
		t.Errorf("Error not nil %v", err)
	} else if r == nil {
		t.Error("Reports nil")
	} else if !reflect.DeepEqual(data.expected, *r) {
		t.Errorf("Expected: %v, Got: %v", data.expected, *r)
	}
}