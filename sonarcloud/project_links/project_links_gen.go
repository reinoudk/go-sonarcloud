package project_links

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// CreateRequest Create a new project link.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
type CreateRequest struct {
	Name       string `form:"name,omitempty"`       // Link name
	ProjectId  string `form:"projectId,omitempty"`  // Project id
	ProjectKey string `form:"projectKey,omitempty"` // Project key
	Url        string `form:"url,omitempty"`        // Link url
}

// CreateResponse is the response for CreateRequest
type CreateResponse struct {
	Link struct {
		Id   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Url  string `json:"url,omitempty"`
	} `json:"link,omitempty"`
}

// DeleteRequest Delete existing project link.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
type DeleteRequest struct {
	Id string `form:"id,omitempty"` // Link id
}

// SearchRequest List links of a project.<br>The 'projectId' or 'projectKey' must be provided.<br>Requires one of the following permissions:<ul><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>
type SearchRequest struct {
	ProjectId  string `form:"projectId,omitempty"`  // Project Id
	ProjectKey string `form:"projectKey,omitempty"` // Project Key
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Links []struct {
		Id   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Type string `json:"type,omitempty"`
		Url  string `json:"url,omitempty"`
	} `json:"links,omitempty"`
}
