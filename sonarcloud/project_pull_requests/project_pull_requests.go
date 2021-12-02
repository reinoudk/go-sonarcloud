package project_pull_requests

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// DeleteRequest Delete a pull request.<br/>Requires 'Administer' rights on the specified project.
type DeleteRequest struct {
	Project     string `form:"project,omitempty"`     // Project key
	PullRequest string `form:"pullRequest,omitempty"` // Pull request id
}

// ListRequest List the pull requests of a project.<br/>One of the following permissions is required: <ul><li>'Browse' rights on the specified project</li><li>'Execute Analysis' rights on the specified project</li></ul>
type ListRequest struct {
	Project string `form:"project,omitempty"` // Project key
}

// ListResponse is the response for ListRequest
type ListResponse struct {
	PullRequests []struct {
		AnalysisDate string `json:"analysisDate,omitempty"`
		Base         string `json:"base,omitempty"`
		Branch       string `json:"branch,omitempty"`
		Commit       struct {
			Sha string `json:"sha,omitempty"`
		} `json:"commit,omitempty"`
		Contributors []struct {
			Avatar string `json:"avatar,omitempty"`
			Login  string `json:"login,omitempty"`
			Name   string `json:"name,omitempty"`
		} `json:"contributors,omitempty"`
		Key    string `json:"key,omitempty"`
		Status struct {
			Bugs              float64 `json:"bugs,omitempty"`
			CodeSmells        float64 `json:"codeSmells,omitempty"`
			QualityGateStatus string  `json:"qualityGateStatus,omitempty"`
			Vulnerabilities   float64 `json:"vulnerabilities,omitempty"`
		} `json:"status,omitempty"`
		Target string `json:"target,omitempty"`
		Title  string `json:"title,omitempty"`
		Url    string `json:"url,omitempty"`
	} `json:"pullRequests,omitempty"`
}
