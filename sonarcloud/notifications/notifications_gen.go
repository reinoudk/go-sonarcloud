package notifications

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// AddRequest Add a notification for the authenticated user.<br>Requires one of the following permissions:<ul> <li>Authentication if no login is provided. If a project is provided, requires the 'Browse' permission on the specified project.</li> <li>If a project is provided, requires the 'Browse' permission on the specified project.</li></ul>
type AddRequest struct {
	Channel string `form:"channel,omitempty"` // Channel through which the notification is sent. For example, notifications can be sent by email.
	Login   string `form:"login,omitempty"`   // User login
	Project string `form:"project,omitempty"` // Project key
	Type    string `form:"type,omitempty"`    // Notification type. Possible values are for:<ul>  <li>Global notifications: CeReportTaskFailure, ChangesOnMyIssue, SQ-MyNewIssues</li>  <li>Per project notifications: CeReportTaskFailure, ChangesOnMyIssue, NewAlerts, NewFalsePositiveIssue, NewIssues, SQ-MyNewIssues</li></ul>
}

// ListRequest List notifications of the authenticated user
type ListRequest struct {
	Login string `form:"login,omitempty"` // User login
}

// ListResponse is the response for ListRequest
type ListResponse struct {
	Channels    []string `json:"channels,omitempty"`
	GlobalTypes []struct {
		IsVisibleOnlyForOrgMembers bool   `json:"isVisibleOnlyForOrgMembers,omitempty"`
		Key                        string `json:"key,omitempty"`
	} `json:"globalTypes,omitempty"`
	Notifications []struct {
		Channel      string `json:"channel,omitempty"`
		Type         string `json:"type,omitempty"`
		Organization string `json:"organization,omitempty"`
		Project      string `json:"project,omitempty"`
		ProjectName  string `json:"projectName,omitempty"`
	} `json:"notifications,omitempty"`
	Organizations []struct {
		IsMember bool   `json:"isMember,omitempty"`
		Key      string `json:"key,omitempty"`
		Name     string `json:"name,omitempty"`
	} `json:"organizations,omitempty"`
	PerProjectTypes []struct {
		IsVisibleOnlyForOrgMembers bool   `json:"isVisibleOnlyForOrgMembers,omitempty"`
		Key                        string `json:"key,omitempty"`
	} `json:"perProjectTypes,omitempty"`
}

// RemoveRequest Remove a notification for the authenticated user
type RemoveRequest struct {
	Channel string `form:"channel,omitempty"` // Channel through which the notification is sent. For example, notifications can be sent by email.
	Login   string `form:"login,omitempty"`   // User login
	Project string `form:"project,omitempty"` // Project key
	Type    string `form:"type,omitempty"`    // Notification type. Possible values are for:<ul>  <li>Global notifications: CeReportTaskFailure, ChangesOnMyIssue, SQ-MyNewIssues</li>  <li>Per project notifications: CeReportTaskFailure, ChangesOnMyIssue, NewAlerts, NewFalsePositiveIssue, NewIssues, SQ-MyNewIssues</li></ul>
}
