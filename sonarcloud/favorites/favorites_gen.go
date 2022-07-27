package favorites

import paging "github.com/reinoudk/go-sonarcloud/sonarcloud/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// AddRequest Add a project as favorite for the authenticated user.<br>Only 100 components can be added as favorite.<br>Requires authentication and the following permission: 'Browse' on the project.
type AddRequest struct {
	Component string `form:"component,omitempty"` // Component key. Only components with qualifier TRK are supported
}

// RemoveRequest Remove a component (project, directory, file etc.) as favorite for the authenticated user.<br>Requires authentication.
type RemoveRequest struct {
	Component string `form:"component,omitempty"` // Component key
}

// SearchRequest Search for the authenticated user favorites.<br>Requires authentication.
type SearchRequest struct{}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Favorites []struct {
		Key          string `json:"key,omitempty"`
		Name         string `json:"name,omitempty"`
		Organization string `json:"organization,omitempty"`
		Qualifier    string `json:"qualifier,omitempty"`
	} `json:"favorites,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from SearchResponse
func (r *SearchResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchResponseAll is the collection for SearchRequest
type SearchResponseAll struct {
	Favorites []struct {
		Key          string `json:"key,omitempty"`
		Name         string `json:"name,omitempty"`
		Organization string `json:"organization,omitempty"`
		Qualifier    string `json:"qualifier,omitempty"`
	} `json:"favorites,omitempty"`
}
