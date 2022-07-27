package project_tags

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// SearchRequest Search tags
type SearchRequest struct {
	Q string `form:"q,omitempty"` // Limit search to tags that contain the supplied string.
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Tags []string `json:"tags,omitempty"`
}

// SetRequest Set tags on a project.<br>Requires the following permission: 'Administer' rights on the specified project
type SetRequest struct {
	Project string `form:"project,omitempty"` // Project key
	Tags    string `form:"tags,omitempty"`    // Comma-separated list of tags
}
