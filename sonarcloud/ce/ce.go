package ce

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ActivityRequest Search for tasks.<br> Either componentId or component can be provided, but not both.<br> Requires the project administration permission if componentId or component is set.
type ActivityRequest struct {
	Component      string `form:"component,omitempty"`      // Key of the component (project) to filter on
	ComponentId    string `form:"componentId,omitempty"`    // Id of the component (project) to filter on
	MaxExecutedAt  string `form:"maxExecutedAt,omitempty"`  // Maximum date of end of task processing (inclusive)
	MinSubmittedAt string `form:"minSubmittedAt,omitempty"` // Minimum date of task submission (inclusive)
	OnlyCurrents   string `form:"onlyCurrents,omitempty"`   // Filter on the last tasks (only the most recent finished task by project)
	Q              string `form:"q,omitempty"`              // Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that are exactly the same as the supplied string</li><li>task ids that are exactly the same as the supplied string</li></ul>Must not be set together with componentId
	Status         string `form:"status,omitempty"`         // Comma separated list of task statuses
	Type           string `form:"type,omitempty"`           // Task type
}

// ActivityResponse is the response for ActivityRequest
type ActivityResponse struct {
	Tasks []struct {
		AnalysisId         string  `json:"analysisId,omitempty"`
		ComponentId        string  `json:"componentId,omitempty"`
		ComponentKey       string  `json:"componentKey,omitempty"`
		ComponentName      string  `json:"componentName,omitempty"`
		ComponentQualifier string  `json:"componentQualifier,omitempty"`
		ExecutedAt         string  `json:"executedAt,omitempty"`
		ExecutionTimeMs    float64 `json:"executionTimeMs,omitempty"`
		HasErrorStacktrace bool    `json:"hasErrorStacktrace,omitempty"`
		HasScannerContext  bool    `json:"hasScannerContext,omitempty"`
		Id                 string  `json:"id,omitempty"`
		Logs               bool    `json:"logs,omitempty"`
		Organization       string  `json:"organization,omitempty"`
		StartedAt          string  `json:"startedAt,omitempty"`
		Status             string  `json:"status,omitempty"`
		SubmittedAt        string  `json:"submittedAt,omitempty"`
		SubmitterLogin     string  `json:"submitterLogin,omitempty"`
		Type               string  `json:"type,omitempty"`
	} `json:"tasks,omitempty"`
}

// ActivityStatusRequest Returns CE activity related metrics.<br>Requires 'Administer' permission on the specified project.
type ActivityStatusRequest struct {
	ComponentId  string `form:"componentId,omitempty"`  // Id of the component (project) to filter on
	ComponentKey string `form:"componentKey,omitempty"` // Key of the component (project) to filter on
}

// ActivityStatusResponse is the response for ActivityStatusRequest
type ActivityStatusResponse struct {
	Failing     float64 `json:"failing,omitempty"`
	InProgress  float64 `json:"inProgress,omitempty"`
	Pending     float64 `json:"pending,omitempty"`
	PendingTime float64 `json:"pendingTime,omitempty"`
}

// ComponentRequest Get the pending tasks, in-progress tasks and the last executed task of a given component (usually a project).<br>Requires the following permission: 'Browse' on the specified component.<br>Either 'componentId' or 'component' must be provided.
type ComponentRequest struct {
	Component   string `form:"component,omitempty"`   //
	ComponentId string `form:"componentId,omitempty"` //
}

// ComponentResponse is the response for ComponentRequest
type ComponentResponse struct {
	Current struct {
		AnalysisId         float64 `json:"analysisId,omitempty"`
		ComponentId        string  `json:"componentId,omitempty"`
		ComponentKey       string  `json:"componentKey,omitempty"`
		ComponentName      string  `json:"componentName,omitempty"`
		ComponentQualifier string  `json:"componentQualifier,omitempty"`
		ErrorMessage       string  `json:"errorMessage,omitempty"`
		ErrorType          string  `json:"errorType,omitempty"`
		ExecutionTimeMs    float64 `json:"executionTimeMs,omitempty"`
		FinishedAt         string  `json:"finishedAt,omitempty"`
		HasErrorStacktrace bool    `json:"hasErrorStacktrace,omitempty"`
		HasScannerContext  bool    `json:"hasScannerContext,omitempty"`
		Id                 string  `json:"id,omitempty"`
		Logs               bool    `json:"logs,omitempty"`
		Organization       string  `json:"organization,omitempty"`
		StartedAt          string  `json:"startedAt,omitempty"`
		Status             string  `json:"status,omitempty"`
		SubmittedAt        string  `json:"submittedAt,omitempty"`
		Type               string  `json:"type,omitempty"`
	} `json:"current,omitempty"`
	Queue []struct {
		ComponentId        string `json:"componentId,omitempty"`
		ComponentKey       string `json:"componentKey,omitempty"`
		ComponentName      string `json:"componentName,omitempty"`
		ComponentQualifier string `json:"componentQualifier,omitempty"`
		Id                 string `json:"id,omitempty"`
		Logs               bool   `json:"logs,omitempty"`
		Organization       string `json:"organization,omitempty"`
		Status             string `json:"status,omitempty"`
		SubmittedAt        string `json:"submittedAt,omitempty"`
		Type               string `json:"type,omitempty"`
	} `json:"queue,omitempty"`
}

// TaskRequest Give Compute Engine task details such as type, status, duration and associated component.<br />Requires 'Execute Analysis' permission.
type TaskRequest struct {
	AdditionalFields string `form:"additionalFields,omitempty"` // Comma-separated list of the optional fields to be returned in response.
	Id               string `form:"id,omitempty"`               // Id of task
}

// TaskResponse is the response for TaskRequest
type TaskResponse struct {
	Task struct {
		AnalysisId         float64 `json:"analysisId,omitempty"`
		ComponentId        string  `json:"componentId,omitempty"`
		ComponentKey       string  `json:"componentKey,omitempty"`
		ComponentName      string  `json:"componentName,omitempty"`
		ComponentQualifier string  `json:"componentQualifier,omitempty"`
		ErrorMessage       string  `json:"errorMessage,omitempty"`
		ErrorStacktrace    string  `json:"errorStacktrace,omitempty"`
		ExecutedAt         string  `json:"executedAt,omitempty"`
		ExecutionTimeMs    float64 `json:"executionTimeMs,omitempty"`
		HasErrorStacktrace bool    `json:"hasErrorStacktrace,omitempty"`
		HasScannerContext  bool    `json:"hasScannerContext,omitempty"`
		Id                 string  `json:"id,omitempty"`
		Logs               bool    `json:"logs,omitempty"`
		Organization       string  `json:"organization,omitempty"`
		ScannerContext     string  `json:"scannerContext,omitempty"`
		StartedAt          string  `json:"startedAt,omitempty"`
		Status             string  `json:"status,omitempty"`
		SubmittedAt        string  `json:"submittedAt,omitempty"`
		Type               string  `json:"type,omitempty"`
	} `json:"task,omitempty"`
}
