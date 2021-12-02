package project_analyses

import paging "github.com/reinoudk/go-sonarcloud/sonarcloud/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// CreateEventRequest Create a project analysis event.<br>Only event of category 'VERSION' and 'OTHER' can be created.<br>Requires the permission 'Administer' on the specified project.
type CreateEventRequest struct {
	Analysis string `form:"analysis,omitempty"` // Analysis key
	Category string `form:"category,omitempty"` // Category
	Name     string `form:"name,omitempty"`     // Name
}

// CreateEventResponse is the response for CreateEventRequest
type CreateEventResponse struct {
	Event struct {
		Analysis string `json:"analysis,omitempty"`
		Category string `json:"category,omitempty"`
		Key      string `json:"key,omitempty"`
		Name     string `json:"name,omitempty"`
	} `json:"event,omitempty"`
}

// DeleteRequest Delete a project analysis.<br>Requires the permission 'Administer' on the project of the specified analysis.
type DeleteRequest struct {
	Analysis string `form:"analysis,omitempty"` // Analysis key
}

// DeleteEventRequest Delete a project analysis event.<br>Only event of category 'VERSION' and 'OTHER' can be deleted.<br>Requires the permission 'Administer' on the specified project.
type DeleteEventRequest struct {
	Event string `form:"event,omitempty"` // Event key
}

// SearchRequest Search a project analyses and attached events.<br>Requires the following permission: 'Browse' on the specified project
type SearchRequest struct {
	Branch   string `form:"branch,omitempty"`   // Key of a long lived branch
	Category string `form:"category,omitempty"` // Event category. Filter analyses that have at least one event of the category specified.
	From     string `form:"from,omitempty"`     // Filter analyses created after the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided
	Project  string `form:"project,omitempty"`  // Project key
	To       string `form:"to,omitempty"`       // Filter analyses created before the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Analyses []struct {
		BuildString string `json:"buildString,omitempty"`
		Date        string `json:"date,omitempty"`
		Events      []struct {
			Category string `json:"category,omitempty"`
			Key      string `json:"key,omitempty"`
			Name     string `json:"name,omitempty"`
		} `json:"events,omitempty"`
		Key                         string `json:"key,omitempty"`
		ManualNewCodePeriodBaseline bool   `json:"manualNewCodePeriodBaseline,omitempty"`
		ProjectVersion              string `json:"projectVersion,omitempty"`
		Revision                    string `json:"revision,omitempty"`
	} `json:"analyses,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from SearchResponse
func (r *SearchResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchResponseAll is the collection for SearchRequest
type SearchResponseAll struct {
	Analyses []struct {
		BuildString string `json:"buildString,omitempty"`
		Date        string `json:"date,omitempty"`
		Events      []struct {
			Category string `json:"category,omitempty"`
			Key      string `json:"key,omitempty"`
			Name     string `json:"name,omitempty"`
		} `json:"events,omitempty"`
		Key                         string `json:"key,omitempty"`
		ManualNewCodePeriodBaseline bool   `json:"manualNewCodePeriodBaseline,omitempty"`
		ProjectVersion              string `json:"projectVersion,omitempty"`
		Revision                    string `json:"revision,omitempty"`
	} `json:"analyses,omitempty"`
}

// SetBaselineRequest Set an analysis as the baseline of the New Code Period on a project or a long-lived branch.<br/>This manually set baseline overrides the `sonar.leak.period` setting.<br/>Requires the permission 'Administer' on the specified project.
type SetBaselineRequest struct {
	Analysis string `form:"analysis,omitempty"` // Analysis key
	Branch   string `form:"branch,omitempty"`   // Branch key
	Project  string `form:"project,omitempty"`  // Project key
}

// UnsetBaselineRequest Unset any manually-set New Code Period baseline on a project or a long-lived branch.<br/>Unsetting a manual baseline restores the use of the `sonar.leak.period` setting.<br/>Requires the permission 'Administer' on the specified project.
type UnsetBaselineRequest struct {
	Branch  string `form:"branch,omitempty"`  // Branch key
	Project string `form:"project,omitempty"` // Project key
}

// UpdateEventRequest Update a project analysis event.<br>Only events of category 'VERSION' and 'OTHER' can be updated.<br>Requires the permission 'Administer' on the specified project.
type UpdateEventRequest struct {
	Event string `form:"event,omitempty"` // Event key
	Name  string `form:"name,omitempty"`  // New name
}

// UpdateEventResponse is the response for UpdateEventRequest
type UpdateEventResponse struct {
	Event struct {
		Analysis string `json:"analysis,omitempty"`
		Category string `json:"category,omitempty"`
		Key      string `json:"key,omitempty"`
		Name     string `json:"name,omitempty"`
	} `json:"event,omitempty"`
}
