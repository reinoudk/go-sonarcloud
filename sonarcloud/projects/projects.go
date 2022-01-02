package projects

import paging "github.com/reinoudk/go-sonarcloud/sonarcloud/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// BulkDeleteRequest Delete one or several projects.<br />Only the 1'000 first items in project filters are taken into account.<br />Requires 'Administer System' permission.<br />At least one parameter is required among analyzedBefore, projects and q
type BulkDeleteRequest struct {
	AnalyzedBefore    string `form:"analyzedBefore,omitempty"`    // Filter the projects for which last analysis is older than the given date (exclusive).<br> Either a date (server timezone) or datetime can be provided.
	OnProvisionedOnly string `form:"onProvisionedOnly,omitempty"` // Filter the projects that are provisioned
	Organization      string `form:"organization,omitempty"`      // The key of the organization
	Projects          string `form:"projects,omitempty"`          // Comma-separated list of project keys
	Q                 string `form:"q,omitempty"`                 // Limit to: <ul><li>component names that contain the supplied string</li><li>component keys that contain the supplied string</li></ul>
	Qualifiers        string `form:"qualifiers,omitempty"`        // No longer used
}

// BulkUpdateKeyRequest Bulk update a project or module key and all its sub-components keys. The bulk update allows to replace a part of the current key by another string on the current project and all its sub-modules.<br>It's possible to simulate the bulk update by setting the parameter 'dryRun' at true. No key is updated with a dry run.<br>Ex: to rename a project with key 'my_project' to 'my_new_project' and all its sub-components keys, call the WS with parameters:<ul>  <li>project: my_project</li>  <li>from: my_</li>  <li>to: my_new_</li></ul>Requires the permission 'Administer' on the specified project.
// Deprecated: this action has been deprecated since version 7.6
type BulkUpdateKeyRequest struct {
	DryRun  string `form:"dryRun,omitempty"`  // Simulate bulk update. No component key is updated.
	From    string `form:"from,omitempty"`    // String to match in components keys
	Project string `form:"project,omitempty"` // Project or module key
	To      string `form:"to,omitempty"`      // String replacement in components keys
}

// BulkUpdateKeyResponse is the response for BulkUpdateKeyRequest
type BulkUpdateKeyResponse struct {
	Keys []struct {
		Duplicate bool   `json:"duplicate,omitempty"`
		Key       string `json:"key,omitempty"`
		NewKey    string `json:"newKey,omitempty"`
	} `json:"keys,omitempty"`
}

// CreateRequest Create a project.<br/>Requires 'Create Projects' permission
type CreateRequest struct {
	Branch       string `form:"branch,omitempty"`       // SCM Branch of the project. The key of the project will become key:branch, for instance 'SonarQube:branch-5.0'
	Name         string `form:"name,omitempty"`         // Name of the project. If name is longer than 500, it is abbreviated.
	Organization string `form:"organization,omitempty"` // The key of the organization
	Project      string `form:"project,omitempty"`      // Key of the project
	Visibility   string `form:"visibility,omitempty"`   // Whether the created project should be visible to everyone, or only specific user/groups.<br/>If no visibility is specified, the default project visibility of the organization will be used.
}

// CreateResponse is the response for CreateRequest
type CreateResponse struct {
	Project struct {
		Key       string `json:"key,omitempty"`
		Name      string `json:"name,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"project,omitempty"`
}

// DeleteRequest Delete a project.<br> Requires 'Administer System' permission or 'Administer' permission on the project.
type DeleteRequest struct {
	Project string `form:"project,omitempty"` // Project key
}

// SearchRequest Search for projects to administrate them.<br>Requires 'System Administrator' permission
type SearchRequest struct {
	AnalyzedBefore    string `form:"analyzedBefore,omitempty"`    // Filter the projects for which last analysis is older than the given date (exclusive).<br> Either a date (server timezone) or datetime can be provided.
	OnProvisionedOnly string `form:"onProvisionedOnly,omitempty"` // Filter the projects that are provisioned
	Organization      string `form:"organization,omitempty"`      // The key of the organization
	Projects          string `form:"projects,omitempty"`          // Comma-separated list of project keys
	Q                 string `form:"q,omitempty"`                 // Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that contain the supplied string</li></ul>
	Qualifiers        string `form:"qualifiers,omitempty"`        // No longer used
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Components []struct {
		Key              string `json:"key,omitempty"`
		LastAnalysisDate string `json:"lastAnalysisDate,omitempty"`
		Name             string `json:"name,omitempty"`
		Organization     string `json:"organization,omitempty"`
		Qualifier        string `json:"qualifier,omitempty"`
		Revision         string `json:"revision,omitempty"`
		Visibility       string `json:"visibility,omitempty"`
	} `json:"components,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from SearchResponse
func (r *SearchResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchResponseAll is the collection for SearchRequest
type SearchResponseAll struct {
	Components []struct {
		Key              string `json:"key,omitempty"`
		LastAnalysisDate string `json:"lastAnalysisDate,omitempty"`
		Name             string `json:"name,omitempty"`
		Organization     string `json:"organization,omitempty"`
		Qualifier        string `json:"qualifier,omitempty"`
		Revision         string `json:"revision,omitempty"`
		Visibility       string `json:"visibility,omitempty"`
	} `json:"components,omitempty"`
}

// UpdateKeyRequest Update a project or module key and all its sub-components keys.<br>Requires the permission 'Administer' on the specified project.
type UpdateKeyRequest struct {
	From string `form:"from,omitempty"` // Project or module key
	To   string `form:"to,omitempty"`   // New component key
}

// UpdateVisibilityRequest Updates visibility of a project.<br>Requires 'Project administer' permission on the specified project
type UpdateVisibilityRequest struct {
	Project    string `form:"project,omitempty"`    // Project key
	Visibility string `form:"visibility,omitempty"` // New visibility
}
