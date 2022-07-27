package qualitygates

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// CopyRequest Copy a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
type CopyRequest struct {
	Id           string `form:"id,omitempty"`           // The ID of the source quality gate
	Name         string `form:"name,omitempty"`         // The name of the quality gate to create
	Organization string `form:"organization,omitempty"` // Organization key.
}

// CreateRequest Create a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
type CreateRequest struct {
	Name         string `form:"name,omitempty"`         // The name of the quality gate to create
	Organization string `form:"organization,omitempty"` // Organization key.
}

// CreateResponse is the response for CreateRequest
type CreateResponse struct {
	Id   float64 `json:"id,omitempty"`
	Name string  `json:"name,omitempty"`
}

// CreateConditionRequest Add a new condition to a quality gate.<br>Requires the 'Administer Quality Gates' permission.
type CreateConditionRequest struct {
	Error        string `form:"error,omitempty"`        // Condition error threshold
	GateId       string `form:"gateId,omitempty"`       // ID of the quality gate
	Metric       string `form:"metric,omitempty"`       // Condition metric.<br/> Only metric of the following types are allowed:<ul><li>INT</li><li>MILLISEC</li><li>RATING</li><li>WORK_DUR</li><li>FLOAT</li><li>PERCENT</li><li>LEVEL</li></ul>Following metrics are forbidden:<ul><li>alert_status</li><li>security_hotspots</li><li>new_security_hotspots</li></ul>
	Op           string `form:"op,omitempty"`           // Condition operator:<br/><ul><li>LT = is lower than</li><li>GT = is greater than</li></ul>
	Organization string `form:"organization,omitempty"` // Organization key.
}

// CreateConditionResponse is the response for CreateConditionRequest
type CreateConditionResponse struct {
	Error   string  `json:"error,omitempty"`
	Id      float64 `json:"id,omitempty"`
	Metric  string  `json:"metric,omitempty"`
	Op      string  `json:"op,omitempty"`
	Warning string  `json:"warning,omitempty"`
}

// DeleteConditionRequest Delete a condition from a quality gate.<br>Requires the 'Administer Quality Gates' permission.
type DeleteConditionRequest struct {
	Id           string `form:"id,omitempty"`           // Condition ID
	Organization string `form:"organization,omitempty"` // Organization key.
}

// DeselectRequest Remove the association of a project from a quality gate.<br>Requires one of the following permissions:<ul><li>'Administer Quality Gates'</li><li>'Administer' rights on the project</li></ul>
type DeselectRequest struct {
	Organization string `form:"organization,omitempty"` // Organization key.
	ProjectId    string `form:"projectId,omitempty"`    // Project id
	ProjectKey   string `form:"projectKey,omitempty"`   // Project key
}

// DestroyRequest Delete a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
type DestroyRequest struct {
	Id           string `form:"id,omitempty"`           // ID of the quality gate to delete
	Organization string `form:"organization,omitempty"` // Organization key.
}

// GetByProjectRequest Get the quality gate of a project.<br />Requires one of the following permissions:<ul><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>
type GetByProjectRequest struct {
	Organization string `form:"organization,omitempty"` // Organization key.
	Project      string `form:"project,omitempty"`      // Project key
}

// GetByProjectResponse is the response for GetByProjectRequest
type GetByProjectResponse struct {
	QualityGate struct {
		Default bool    `json:"default,omitempty"`
		Id      float64 `json:"id,omitempty"`
		Name    string  `json:"name,omitempty"`
	} `json:"qualityGate,omitempty"`
}

// ListRequest Get a list of quality gates
type ListRequest struct {
	Organization string `form:"organization,omitempty"` // Organization key.
}

// ListResponse is the response for ListRequest
type ListResponse struct {
	Actions struct {
		Create bool `json:"create,omitempty"`
	} `json:"actions,omitempty"`
	Default      float64 `json:"default,omitempty"`
	Qualitygates []struct {
		Actions struct {
			AssociateProjects bool `json:"associateProjects,omitempty"`
			Copy              bool `json:"copy,omitempty"`
			Delete            bool `json:"delete,omitempty"`
			ManageConditions  bool `json:"manageConditions,omitempty"`
			Rename            bool `json:"rename,omitempty"`
			SetAsDefault      bool `json:"setAsDefault,omitempty"`
		} `json:"actions,omitempty"`
		Conditions []struct {
			Error  string  `json:"error,omitempty"`
			Id     float64 `json:"id,omitempty"`
			Metric string  `json:"metric,omitempty"`
			Op     string  `json:"op,omitempty"`
		} `json:"conditions,omitempty"`
		Id        float64 `json:"id,omitempty"`
		IsBuiltIn bool    `json:"isBuiltIn,omitempty"`
		IsDefault bool    `json:"isDefault,omitempty"`
		Name      string  `json:"name,omitempty"`
	} `json:"qualitygates,omitempty"`
}

// ProjectStatusRequest Get the quality gate status of a project or a Compute Engine task.<br />Either 'analysisId', 'projectId' or 'projectKey' must be provided<br />The different statuses returned are: OK, WARN, ERROR, NONE. The NONE status is returned when there is no quality gate associated with the analysis.<br />Returns an HTTP code 404 if the analysis associated with the task is not found or does not exist.<br />Requires one of the following permissions:<ul><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>
type ProjectStatusRequest struct {
	AnalysisId  string `form:"analysisId,omitempty"`  // Analysis id
	Branch      string `form:"branch,omitempty"`      // Branch key
	ProjectId   string `form:"projectId,omitempty"`   // Project id. Doesn't work with branches or pull requests
	ProjectKey  string `form:"projectKey,omitempty"`  // Project key
	PullRequest string `form:"pullRequest,omitempty"` // Pull request id
}

// ProjectStatusResponse is the response for ProjectStatusRequest
type ProjectStatusResponse struct {
	ProjectStatus struct {
		Conditions []struct {
			ActualValue    string  `json:"actualValue,omitempty"`
			Comparator     string  `json:"comparator,omitempty"`
			ErrorThreshold string  `json:"errorThreshold,omitempty"`
			MetricKey      string  `json:"metricKey,omitempty"`
			PeriodIndex    float64 `json:"periodIndex,omitempty"`
			Status         string  `json:"status,omitempty"`
		} `json:"conditions,omitempty"`
		IgnoredConditions bool `json:"ignoredConditions,omitempty"`
		Periods           []struct {
			Date      string  `json:"date,omitempty"`
			Index     float64 `json:"index,omitempty"`
			Mode      string  `json:"mode,omitempty"`
			Parameter string  `json:"parameter,omitempty"`
		} `json:"periods,omitempty"`
		Status string `json:"status,omitempty"`
	} `json:"projectStatus,omitempty"`
}

// RenameRequest Rename a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
type RenameRequest struct {
	Id           string `form:"id,omitempty"`           // ID of the quality gate to rename
	Name         string `form:"name,omitempty"`         // New name of the quality gate
	Organization string `form:"organization,omitempty"` // Organization key.
}

// SearchRequest Search for projects associated (or not) to a quality gate.<br/>Only authorized projects for current user will be returned.
type SearchRequest struct {
	GateId       string `form:"gateId,omitempty"`       // Quality Gate ID
	Organization string `form:"organization,omitempty"` // Organization key.
	Page         string `form:"page,omitempty"`         // Page number
	PageSize     string `form:"pageSize,omitempty"`     // Page size
	Query        string `form:"query,omitempty"`        // To search for projects containing this string. If this parameter is set, "selected" is set to "all".
	Selected     string `form:"selected,omitempty"`     // Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Paging struct {
		PageIndex float64 `json:"pageIndex,omitempty"`
		PageSize  float64 `json:"pageSize,omitempty"`
		Total     float64 `json:"total,omitempty"`
	} `json:"paging,omitempty"`
	Results []struct {
		Id       float64 `json:"id,omitempty"`
		Key      string  `json:"key,omitempty"`
		Name     string  `json:"name,omitempty"`
		Selected bool    `json:"selected,omitempty"`
	} `json:"results,omitempty"`
}

// SelectRequest Associate a project to a quality gate.<br>The 'projectId' or 'projectKey' must be provided.<br>Project id as a numeric value is deprecated since 6.1. Please use the id similar to 'AU-TpxcA-iU5OvuD2FLz'.<br>Requires the 'Administer Quality Gates' permission.
type SelectRequest struct {
	GateId       string `form:"gateId,omitempty"`       // Quality gate id
	Organization string `form:"organization,omitempty"` // Organization key.
	ProjectId    string `form:"projectId,omitempty"`    // Project id. Project id as an numeric value is deprecated since 6.1
	ProjectKey   string `form:"projectKey,omitempty"`   // Project key
}

// SetAsDefaultRequest Set a quality gate as the default quality gate.<br>Requires the 'Administer Quality Gates' permission.
type SetAsDefaultRequest struct {
	Id           string `form:"id,omitempty"`           // ID of the quality gate to set as default
	Organization string `form:"organization,omitempty"` // Organization key.
}

// ShowRequest Display the details of a quality gate
type ShowRequest struct {
	Id           string `form:"id,omitempty"`           // ID of the quality gate. Either id or name must be set
	Name         string `form:"name,omitempty"`         // Name of the quality gate. Either id or name must be set
	Organization string `form:"organization,omitempty"` // Organization key.
}

// ShowResponse is the response for ShowRequest
type ShowResponse struct {
	Actions struct {
		AssociateProjects bool `json:"associateProjects,omitempty"`
		Copy              bool `json:"copy,omitempty"`
		Delete            bool `json:"delete,omitempty"`
		ManageConditions  bool `json:"manageConditions,omitempty"`
		Rename            bool `json:"rename,omitempty"`
		SetAsDefault      bool `json:"setAsDefault,omitempty"`
	} `json:"actions,omitempty"`
	Conditions []struct {
		Error  string  `json:"error,omitempty"`
		Id     float64 `json:"id,omitempty"`
		Metric string  `json:"metric,omitempty"`
		Op     string  `json:"op,omitempty"`
	} `json:"conditions,omitempty"`
	Id        float64 `json:"id,omitempty"`
	IsBuiltIn bool    `json:"isBuiltIn,omitempty"`
	Name      string  `json:"name,omitempty"`
}

// UnsetDefaultRequest This webservice is no more available : a default quality gate is mandatory.
// Deprecated: this action has been deprecated since version 7.0
type UnsetDefaultRequest struct{}

// UnsetDefaultResponse is the response for UnsetDefaultRequest
type UnsetDefaultResponse struct {
	Errors []struct {
		Msg string `json:"msg,omitempty"`
	} `json:"errors,omitempty"`
}

// UpdateConditionRequest Update a condition attached to a quality gate.<br>Requires the 'Administer Quality Gates' permission.
type UpdateConditionRequest struct {
	Error        string `form:"error,omitempty"`        // Condition error threshold
	Id           string `form:"id,omitempty"`           // Condition ID
	Metric       string `form:"metric,omitempty"`       // Condition metric.<br/> Only metric of the following types are allowed:<ul><li>INT</li><li>MILLISEC</li><li>RATING</li><li>WORK_DUR</li><li>FLOAT</li><li>PERCENT</li><li>LEVEL</li></ul>Following metrics are forbidden:<ul><li>alert_status</li><li>security_hotspots</li><li>new_security_hotspots</li></ul>
	Op           string `form:"op,omitempty"`           // Condition operator:<br/><ul><li>LT = is lower than</li><li>GT = is greater than</li></ul>
	Organization string `form:"organization,omitempty"` // Organization key.
}
