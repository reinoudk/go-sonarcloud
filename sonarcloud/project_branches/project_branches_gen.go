package project_branches

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// DeleteRequest Delete a non-main branch of a project.<br/>Requires 'Administer' rights on the specified project.
type DeleteRequest struct {
	Branch  string `form:"branch,omitempty"`  // Name of the branch
	Project string `form:"project,omitempty"` // Project key
}

// ListRequest List the branches of a project.<br/>Requires 'Browse' or 'Execute analysis' rights on the specified project.
type ListRequest struct {
	Project string `form:"project,omitempty"` // Project key
}

// ListResponse is the response for ListRequest
type ListResponse struct {
	Branches []struct {
		AnalysisDate string `json:"analysisDate,omitempty"`
		Commit       struct {
			Sha string `json:"sha,omitempty"`
		} `json:"commit,omitempty"`
		IsMain      bool   `json:"isMain,omitempty"`
		MergeBranch string `json:"mergeBranch,omitempty"`
		Name        string `json:"name,omitempty"`
		Status      struct {
			Bugs              float64 `json:"bugs,omitempty"`
			CodeSmells        float64 `json:"codeSmells,omitempty"`
			QualityGateStatus string  `json:"qualityGateStatus,omitempty"`
			Vulnerabilities   float64 `json:"vulnerabilities,omitempty"`
		} `json:"status,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"branches,omitempty"`
}

// RenameRequest Rename the main branch of a project.<br/>Requires 'Administer' permission on the specified project.
type RenameRequest struct {
	Name    string `form:"name,omitempty"`    // New name of the main branch
	Project string `form:"project,omitempty"` // Project key
}
