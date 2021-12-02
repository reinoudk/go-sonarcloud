package permissions

import paging "github.com/reinoudk/go-sonarcloud/sonarcloud/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// AddGroupRequest Add permission to a group.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> The group name or group id must be provided. <br />Requires the permission 'Administer' on the specified project.
type AddGroupRequest struct {
	GroupId    string `form:"groupId,omitempty"`    // Group id
	GroupName  string `form:"groupName,omitempty"`  // Group name or 'anyone' (case insensitive)
	Permission string `form:"permission,omitempty"` // Permission<ul><li>Possible values for global permissions: admin, profileadmin, gateadmin, scan, provisioning</li><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	ProjectId  string `form:"projectId,omitempty"`  // Project id
	ProjectKey string `form:"projectKey,omitempty"` // Project key
}

// AddGroupToTemplateRequest Add a group to a permission template.<br /> The group id or group name must be provided. <br />Requires the permission 'Administer' on the organization.
type AddGroupToTemplateRequest struct {
	GroupId      string `form:"groupId,omitempty"`      // Group id
	GroupName    string `form:"groupName,omitempty"`    // Group name or 'anyone' (case insensitive)
	Permission   string `form:"permission,omitempty"`   // Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	TemplateId   string `form:"templateId,omitempty"`   // Template id
	TemplateName string `form:"templateName,omitempty"` // Template name
}

// AddProjectCreatorToTemplateRequest Add a project creator to a permission template.<br>Requires the permission 'Administer' on the organization.
type AddProjectCreatorToTemplateRequest struct {
	Permission   string `form:"permission,omitempty"`   // Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	TemplateId   string `form:"templateId,omitempty"`   // Template id
	TemplateName string `form:"templateName,omitempty"` // Template name
}

// AddUserRequest Add permission to a user.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br />Requires the permission 'Administer' on the specified project.
type AddUserRequest struct {
	Login      string `form:"login,omitempty"`      // User login
	Permission string `form:"permission,omitempty"` // Permission<ul><li>Possible values for global permissions: admin, profileadmin, gateadmin, scan, provisioning</li><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	ProjectId  string `form:"projectId,omitempty"`  // Project id
	ProjectKey string `form:"projectKey,omitempty"` // Project key
}

// AddUserToTemplateRequest Add a user to a permission template.<br /> Requires the permission 'Administer' on the organization.
type AddUserToTemplateRequest struct {
	Login        string `form:"login,omitempty"`        // User login
	Permission   string `form:"permission,omitempty"`   // Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	TemplateId   string `form:"templateId,omitempty"`   // Template id
	TemplateName string `form:"templateName,omitempty"` // Template name
}

// ApplyTemplateRequest Apply a permission template to one project.<br>The project id or project key must be provided.<br>The template id or name must be provided.<br>Requires the permission 'Administer' on the organization.
type ApplyTemplateRequest struct {
	ProjectId    string `form:"projectId,omitempty"`    // Project id
	ProjectKey   string `form:"projectKey,omitempty"`   // Project key
	TemplateId   string `form:"templateId,omitempty"`   // Template id
	TemplateName string `form:"templateName,omitempty"` // Template name
}

// BulkApplyTemplateRequest Apply a permission template to several projects.<br />The template id or name must be provided.<br />Requires the permission 'Administer' on the organization.
type BulkApplyTemplateRequest struct {
	AnalyzedBefore    string `form:"analyzedBefore,omitempty"`    // Filter the projects for which last analysis is older than the given date (exclusive).<br> Either a date (server timezone) or datetime can be provided.
	OnProvisionedOnly string `form:"onProvisionedOnly,omitempty"` // Filter the projects that are provisioned
	Projects          string `form:"projects,omitempty"`          // Comma-separated list of project keys
	Q                 string `form:"q,omitempty"`                 // Limit search to: <ul><li>project names that contain the supplied string</li><li>project keys that are exactly the same as the supplied string</li></ul>
	Qualifiers        string `form:"qualifiers,omitempty"`        // Comma-separated list of component qualifiers. Filter the results with the specified qualifiers. Possible values are:<ul><li>TRK - Projects</li></ul>
	TemplateId        string `form:"templateId,omitempty"`        // Template id
	TemplateName      string `form:"templateName,omitempty"`      // Template name
}

// CreateTemplateRequest Create a permission template.<br />Requires the permission 'Administer' on the organization.
type CreateTemplateRequest struct {
	Description       string `form:"description,omitempty"`       // Description
	Name              string `form:"name,omitempty"`              // Name
	ProjectKeyPattern string `form:"projectKeyPattern,omitempty"` // Project key pattern. Must be a valid Java regular expression
}

// CreateTemplateResponse is the response for CreateTemplateRequest
type CreateTemplateResponse struct {
	PermissionTemplate struct {
		Description       string `json:"description,omitempty"`
		Name              string `json:"name,omitempty"`
		ProjectKeyPattern string `json:"projectKeyPattern,omitempty"`
	} `json:"permissionTemplate,omitempty"`
}

// DeleteTemplateRequest Delete a permission template.<br />Requires the permission 'Administer' on the organization.
type DeleteTemplateRequest struct {
	TemplateId   string `form:"templateId,omitempty"`   // Template id
	TemplateName string `form:"templateName,omitempty"` // Template name
}

// RemoveGroupRequest Remove a permission from a group.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> The group id or group name must be provided, not both.<br />Requires the permission 'Administer' on the specified project.
type RemoveGroupRequest struct {
	GroupId    string `form:"groupId,omitempty"`    // Group id
	GroupName  string `form:"groupName,omitempty"`  // Group name or 'anyone' (case insensitive)
	Permission string `form:"permission,omitempty"` // Permission<ul><li>Possible values for global permissions: admin, profileadmin, gateadmin, scan, provisioning</li><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	ProjectId  string `form:"projectId,omitempty"`  // Project id
	ProjectKey string `form:"projectKey,omitempty"` // Project key
}

// RemoveGroupFromTemplateRequest Remove a group from a permission template.<br /> The group id or group name must be provided. <br />Requires the permission 'Administer' on the organization.
type RemoveGroupFromTemplateRequest struct {
	GroupId      string `form:"groupId,omitempty"`      // Group id
	GroupName    string `form:"groupName,omitempty"`    // Group name or 'anyone' (case insensitive)
	Permission   string `form:"permission,omitempty"`   // Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	TemplateId   string `form:"templateId,omitempty"`   // Template id
	TemplateName string `form:"templateName,omitempty"` // Template name
}

// RemoveProjectCreatorFromTemplateRequest Remove a project creator from a permission template.<br>Requires the permission 'Administer' on the organization.
type RemoveProjectCreatorFromTemplateRequest struct {
	Permission   string `form:"permission,omitempty"`   // Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	TemplateId   string `form:"templateId,omitempty"`   // Template id
	TemplateName string `form:"templateName,omitempty"` // Template name
}

// RemoveUserRequest Remove permission from a user.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> Requires the permission 'Administer' on the specified project.
type RemoveUserRequest struct {
	Login      string `form:"login,omitempty"`      // User login
	Permission string `form:"permission,omitempty"` // Permission<ul><li>Possible values for global permissions: admin, profileadmin, gateadmin, scan, provisioning</li><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	ProjectId  string `form:"projectId,omitempty"`  // Project id
	ProjectKey string `form:"projectKey,omitempty"` // Project key
}

// RemoveUserFromTemplateRequest Remove a user from a permission template.<br /> Requires the permission 'Administer' on the organization.
type RemoveUserFromTemplateRequest struct {
	Login        string `form:"login,omitempty"`        // User login
	Permission   string `form:"permission,omitempty"`   // Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, securityhotspotadmin, scan, user</li></ul>
	TemplateId   string `form:"templateId,omitempty"`   // Template id
	TemplateName string `form:"templateName,omitempty"` // Template name
}

// SearchGlobalPermissionsRequest List global permissions. <br />Requires the following permission: 'Administer System'
// Deprecated: this action has been deprecated since version 6.5
type SearchGlobalPermissionsRequest struct{}

// SearchGlobalPermissionsResponse is the response for SearchGlobalPermissionsRequest
type SearchGlobalPermissionsResponse struct {
	Permissions []struct {
		Description string  `json:"description,omitempty"`
		GroupsCount float64 `json:"groupsCount,omitempty"`
		Key         string  `json:"key,omitempty"`
		Name        string  `json:"name,omitempty"`
		UsersCount  float64 `json:"usersCount,omitempty"`
	} `json:"permissions,omitempty"`
}

// SearchProjectPermissionsRequest List project permissions. A project can be a technical project, a view or a developer.<br />Requires the permission 'Administer' on the specified project.
// Deprecated: this action has been deprecated since version 6.5
type SearchProjectPermissionsRequest struct {
	ProjectId  string `form:"projectId,omitempty"`  // Project id
	ProjectKey string `form:"projectKey,omitempty"` // Project key
	Q          string `form:"q,omitempty"`          // Limit search to: <ul><li>project names that contain the supplied string</li><li>project keys that are exactly the same as the supplied string</li></ul>
	Qualifier  string `form:"qualifier,omitempty"`  // Project qualifier. Filter the results with the specified qualifier. Possible values are:<ul><li>TRK - Projects</li></ul>
}

// SearchProjectPermissionsResponse is the response for SearchProjectPermissionsRequest
type SearchProjectPermissionsResponse struct {
	Paging      paging.Paging `json:"paging,omitempty"`
	Permissions []struct {
		Description string `json:"description,omitempty"`
		Key         string `json:"key,omitempty"`
		Name        string `json:"name,omitempty"`
	} `json:"permissions,omitempty"`
	Projects []struct {
		Id          string `json:"id,omitempty"`
		Key         string `json:"key,omitempty"`
		Name        string `json:"name,omitempty"`
		Permissions []struct {
			GroupsCount float64 `json:"groupsCount,omitempty"`
			Key         string  `json:"key,omitempty"`
			UsersCount  float64 `json:"usersCount,omitempty"`
		} `json:"permissions,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"projects,omitempty"`
}

// GetPaging extracts the paging from SearchProjectPermissionsResponse
func (r *SearchProjectPermissionsResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchProjectPermissionsResponseAll is the collection for SearchProjectPermissionsRequest
type SearchProjectPermissionsResponseAll struct {
	Permissions []struct {
		Description string `json:"description,omitempty"`
		Key         string `json:"key,omitempty"`
		Name        string `json:"name,omitempty"`
	} `json:"permissions,omitempty"`
	Projects []struct {
		Id          string `json:"id,omitempty"`
		Key         string `json:"key,omitempty"`
		Name        string `json:"name,omitempty"`
		Permissions []struct {
			GroupsCount float64 `json:"groupsCount,omitempty"`
			Key         string  `json:"key,omitempty"`
			UsersCount  float64 `json:"usersCount,omitempty"`
		} `json:"permissions,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"projects,omitempty"`
}

// SearchTemplatesRequest List permission templates.<br />Requires the permission 'Administer' on the organization.
type SearchTemplatesRequest struct {
	Q string `form:"q,omitempty"` // Limit search to permission template names that contain the supplied string.
}

// SetDefaultTemplateRequest Set a permission template as default.<br />Requires the permission 'Administer' on the organization.
type SetDefaultTemplateRequest struct {
	Qualifier    string `form:"qualifier,omitempty"`    // Project qualifier. Filter the results with the specified qualifier. Possible values are:<ul><li>TRK - Projects</li></ul>
	TemplateId   string `form:"templateId,omitempty"`   // Template id
	TemplateName string `form:"templateName,omitempty"` // Template name
}

// UpdateTemplateRequest Update a permission template.<br />Requires the permission 'Administer' on the organization.
type UpdateTemplateRequest struct {
	Description       string `form:"description,omitempty"`       // Description
	Id                string `form:"id,omitempty"`                // Id
	Name              string `form:"name,omitempty"`              // Name
	ProjectKeyPattern string `form:"projectKeyPattern,omitempty"` // Project key pattern. Must be a valid Java regular expression
}

// UpdateTemplateResponse is the response for UpdateTemplateRequest
type UpdateTemplateResponse struct {
	PermissionTemplate struct {
		CreatedAt         string `json:"createdAt,omitempty"`
		Description       string `json:"description,omitempty"`
		Id                string `json:"id,omitempty"`
		Name              string `json:"name,omitempty"`
		ProjectKeyPattern string `json:"projectKeyPattern,omitempty"`
		UpdatedAt         string `json:"updatedAt,omitempty"`
	} `json:"permissionTemplate,omitempty"`
}
