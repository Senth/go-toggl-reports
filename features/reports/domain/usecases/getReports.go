package usecases

import (
	"time"

	coretime "github.com/Senth/go-toggl-reports/core/time"
	"github.com/Senth/go-toggl-reports/features/reports/data/models"
	"github.com/Senth/go-toggl-reports/features/reports/data/repositories"
	"github.com/Senth/go-toggl-reports/features/reports/domain/entities"
)

// Reports contains the access repositories
type Reports struct {
	repo repositories.TaskFetcher
}

// NewReports Creates a new reports function
func NewReports(r repositories.TaskFetcher) *Reports {
	return &Reports{
		repo: r,
	}
}

// ThisWeek Get the report (tasks) for this week
func (reports *Reports) ThisWeek(workspaceID int) (r *entities.Report, err error) {
	now := time.Now()
	return reports.Between(workspaceID, coretime.StartOfWeek(now), coretime.EndOfWeek(now))
}

// Between Get the report (tasks) between two dates
func (reports *Reports) Between(workspaceID int, start, end time.Time) (r *entities.Report, err error) {
	modelTasks, _ := reports.repo.Tasks(workspaceID, start, end)

	r = entities.NewEmptyReport()
	for _, m := range modelTasks {
		task := newTaskFromModel(m)
		addTaskToReport(task, r)
		project := getProject(m.Project, m.ProjectColor, r)
		bindTaskAndProject(task, project)
		addTaskToTags(task, r)
	}

	return
}

func newTaskFromModel(m *models.Task) (*entities.Task) {
	task := &entities.Task{
			Name: m.Description,
			Start: m.Start,
			End: m.End,
			Duration: m.Duration,
			Tags: m.Tags,
		}
	return task
}

// getProject get an existing project from the report, or create a new one
func getProject(name string, color string, r *entities.Report) (*entities.Project) {
	var project *entities.Project
		if p, ok := r.ByProject[name]; ok {
			project = p
		} else {
			project = &entities.Project{
				Name: name,
				Color: color,
			}
			r.ByProject[name] = project
		}
	return project
}

// bindTaskAndProject Bind task to project, and project to task
func bindTaskAndProject(t* entities.Task, p* entities.Project) {
	t.Project = p
	p.Tasks = append(p.Tasks, t)
}

func addTaskToReport(t* entities.Task, r* entities.Report) {
	r.Tasks = append(r.Tasks, t)
}

func addTaskToTags(t* entities.Task, r *entities.Report) {
	for _, tagName := range t.Tags {
		// Create tag array if not exists
		if _, ok := r.ByTag[tagName]; !ok {
			r.ByTag[tagName] = make([]*entities.Task, 0, 1)
		}

		// Add task to tag
		r.ByTag[tagName] = append(r.ByTag[tagName], t)
	}
}