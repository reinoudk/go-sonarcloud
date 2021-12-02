package users

import paging "github.com/reinoudk/go-sonarcloud/sonarcloud/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// GroupsRequest Lists the groups a user belongs to. <br/>Requires the permission 'Administer' on the organization.
type GroupsRequest struct {
	Login    string `form:"login,omitempty"`    // A user login
	Q        string `form:"q,omitempty"`        // Limit search to group names that contain the supplied string.
	Selected string `form:"selected,omitempty"` // Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).
}

// GroupsResponse is the response for GroupsRequest
type GroupsResponse struct {
	Groups []struct {
		Default     bool    `json:"default,omitempty"`
		Description string  `json:"description,omitempty"`
		Id          float64 `json:"id,omitempty"`
		Name        string  `json:"name,omitempty"`
		Selected    bool    `json:"selected,omitempty"`
	} `json:"groups,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from GroupsResponse
func (r *GroupsResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// GroupsResponseAll is the collection for GroupsRequest
type GroupsResponseAll struct {
	Groups []struct {
		Default     bool    `json:"default,omitempty"`
		Description string  `json:"description,omitempty"`
		Id          float64 `json:"id,omitempty"`
		Name        string  `json:"name,omitempty"`
		Selected    bool    `json:"selected,omitempty"`
	} `json:"groups,omitempty"`
}

// SearchRequest Get a list of active users. <br/>The following fields are only returned when user has Administer System permission or for logged-in in user :<ul>   <li>'email'</li>   <li>'externalIdentity'</li>   <li>'externalProvider'</li>   <li>'groups'</li>   <li>'lastConnectionDate'</li>   <li>'tokensCount'</li></ul>Field 'lastConnectionDate' is only updated every hour, so it may not be accurate, for instance when a user authenticates many times in less than one hour.
type SearchRequest struct {
	Q string `form:"q,omitempty"` // Filter on login, name and email
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Paging paging.Paging `json:"paging,omitempty"`
	Users  []struct {
		Active           bool     `json:"active,omitempty"`
		Avatar           string   `json:"avatar,omitempty"`
		Email            string   `json:"email,omitempty"`
		ExternalIdentity string   `json:"externalIdentity,omitempty"`
		ExternalProvider string   `json:"externalProvider,omitempty"`
		Groups           []string `json:"groups,omitempty"`
		Local            bool     `json:"local,omitempty"`
		Login            string   `json:"login,omitempty"`
		Name             string   `json:"name,omitempty"`
		TokensCount      float64  `json:"tokensCount,omitempty"`
	} `json:"users,omitempty"`
}

// GetPaging extracts the paging from SearchResponse
func (r *SearchResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchResponseAll is the collection for SearchRequest
type SearchResponseAll struct {
	Users []struct {
		Active           bool     `json:"active,omitempty"`
		Avatar           string   `json:"avatar,omitempty"`
		Email            string   `json:"email,omitempty"`
		ExternalIdentity string   `json:"externalIdentity,omitempty"`
		ExternalProvider string   `json:"externalProvider,omitempty"`
		Groups           []string `json:"groups,omitempty"`
		Local            bool     `json:"local,omitempty"`
		Login            string   `json:"login,omitempty"`
		Name             string   `json:"name,omitempty"`
		TokensCount      float64  `json:"tokensCount,omitempty"`
	} `json:"users,omitempty"`
}
