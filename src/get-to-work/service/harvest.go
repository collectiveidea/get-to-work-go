package service

import (
	"strconv"
	"time"

	"github.com/adlio/harvest"
)

// HarvestService defines a harvest service
type HarvestService struct {
	Name string
	Service
	User *harvest.User
	API  *harvest.API
}

type ProjectAssignment struct {
	ID               int64                     `json:"id,omitempty"`
	IsProjectManager bool                      `json:"is_project_manager"`
	IsActive         bool                      `json:"is_active"`
	Project          *harvest.Project          `json:"project"`
	Client           *harvest.Client           `json:"client"`
	TaskAsignments   []*harvest.TaskAssignment `json:"task_assignments"`
}

type UserAssignmentsResponse struct {
	ProjectAssignments []*ProjectAssignment `json:"project_assignments"`
	PerPage            int64                `json:"per_page"`
	TotalPages         int64                `json:"total_pages"`
	TotalEntries       int64                `json:"total_entries"`
	NextPage           *int64               `json:"next_page"`
	PreviousPage       *int64               `json:"previous_page"`
	Page               int64                `json:"page"`
}

type TimeEntry struct {
	ID             int    `json:"id"`
	ProjectID      int    `json:"project_id"`
	TaskID         int    `json:"task_id"`
	SpentDate      string `json:"spent_date"`
	Notes          string `json:"notes"`
	TimerStartedAt string `json:"timer_started_at,omitempty"`
}

// WhoAmIResponse defines the response from the /account/who_am_i endpoint
type WhoAmIResponse struct {
	User *harvest.User `json:"user"`
}

// ProjectsResponse is a collection of projects returned from /daily
type ProjectsResponse struct {
	Projects []*harvest.Project `json:"projects"`
}

// NewHarvestService creates a HarvestService instance
func NewHarvestService() (harvestService *HarvestService) {
	harvestService = &HarvestService{Name: "Harvest"}
	return
}

// GetName returns the name value
func (hs *HarvestService) GetName() (name string) {
	name = hs.Name
	return
}

// SignIn signs a harvest user in
func (hs *HarvestService) SignIn(account_id string, token string) error {
	hs.API = harvest.NewTokenAPI(account_id, token)
	// Get the user
	user := harvest.User{}
	err := hs.API.Get("/users/me", harvest.Defaults(), &user)

	return err
}

// GetProjects returns projects
func (hs *HarvestService) GetProjects() (projectAssignments []*ProjectAssignment) {
	res := UserAssignmentsResponse{}
	err := hs.API.Get("/users/me/project_assignments", harvest.Defaults(), &res)

	if err != nil {
		return
	}

	projectAssignments = res.ProjectAssignments
	return
}

func (hs *HarvestService) GetTasks(projectAssignment *ProjectAssignment) (tasks []*harvest.TaskAssignment) {
	tasks = projectAssignment.TaskAsignments
	return
}

func (hs *HarvestService) StartTimer(projectID string, taskID string, notes string) (timerID int, err error) {
	args := harvest.Defaults()

	timeEntry := TimeEntry{}
	timeEntry.ProjectID, _ = strconv.Atoi(projectID)
	timeEntry.TaskID, _ = strconv.Atoi(taskID)
	timeEntry.SpentDate = time.Now().UTC().Format("2006-01-02")
	timeEntry.Notes = notes

	err = hs.API.PostWithoutRedirect("/time_entries", args, timeEntry, &timeEntry)

	if err != nil {
		return
	}

	timerID = timeEntry.ID
	return
}
