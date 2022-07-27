package issues

import paging "github.com/reinoudk/go-sonarcloud/sonarcloud/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// AddCommentRequest Add a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue.
type AddCommentRequest struct {
	IsFeedback string `form:"isFeedback,omitempty"` // Define is the given comment is a feedback
	Issue      string `form:"issue,omitempty"`      // Issue key
	Text       string `form:"text,omitempty"`       // Comment text
}

// AddCommentResponse is the response for AddCommentRequest
type AddCommentResponse struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
		Uuid      string `json:"uuid,omitempty"`
	} `json:"components,omitempty"`
	Issue struct {
		Actions  []string `json:"actions,omitempty"`
		Assignee string   `json:"assignee,omitempty"`
		Author   string   `json:"author,omitempty"`
		Comments []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string   `json:"component,omitempty"`
		CreationDate string   `json:"creationDate,omitempty"`
		Debt         string   `json:"debt,omitempty"`
		Effort       string   `json:"effort,omitempty"`
		Flows        []string `json:"flows,omitempty"`
		Key          string   `json:"key,omitempty"`
		Line         float64  `json:"line,omitempty"`
		Message      string   `json:"message,omitempty"`
		Project      string   `json:"project,omitempty"`
		Rule         string   `json:"rule,omitempty"`
		Severity     string   `json:"severity,omitempty"`
		Status       string   `json:"status,omitempty"`
		Tags         []string `json:"tags,omitempty"`
		TextRange    struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		Type        string   `json:"type,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issue,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Email  string `json:"email,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// AssignRequest Assign/Unassign an issue. Requires authentication and Browse permission on project
type AssignRequest struct {
	Assignee string `form:"assignee,omitempty"` // Login of the assignee. When not set, it will unassign the issue. Use '_me' to assign to current user
	Issue    string `form:"issue,omitempty"`    // Issue key
}

// AssignResponse is the response for AssignRequest
type AssignResponse struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
		Uuid      string `json:"uuid,omitempty"`
	} `json:"components,omitempty"`
	Issue struct {
		Actions  []string `json:"actions,omitempty"`
		Assignee string   `json:"assignee,omitempty"`
		Author   string   `json:"author,omitempty"`
		Comments []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string   `json:"component,omitempty"`
		CreationDate string   `json:"creationDate,omitempty"`
		Debt         string   `json:"debt,omitempty"`
		Effort       string   `json:"effort,omitempty"`
		Flows        []string `json:"flows,omitempty"`
		Key          string   `json:"key,omitempty"`
		Line         float64  `json:"line,omitempty"`
		Message      string   `json:"message,omitempty"`
		Project      string   `json:"project,omitempty"`
		Rule         string   `json:"rule,omitempty"`
		Severity     string   `json:"severity,omitempty"`
		Status       string   `json:"status,omitempty"`
		Tags         []string `json:"tags,omitempty"`
		TextRange    struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		Type        string   `json:"type,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issue,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Email  string `json:"email,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// AuthorsRequest Search SCM accounts which match a given query.<br/>Requires authentication.
type AuthorsRequest struct {
	Organization string `form:"organization,omitempty"` // Organization key
	Project      string `form:"project,omitempty"`      // Project key
	Q            string `form:"q,omitempty"`            // Limit search to authors that contain the supplied string.
}

// AuthorsResponse is the response for AuthorsRequest
type AuthorsResponse struct {
	Authors []string `json:"authors,omitempty"`
}

// BulkChangeRequest Bulk change on issues.<br/>Requires authentication.
type BulkChangeRequest struct {
	AddTags           string `form:"add_tags,omitempty"`          // Add tags
	Assign            string `form:"assign,omitempty"`            // To assign the list of issues to a specific user (login), or un-assign all the issues
	Comment           string `form:"comment,omitempty"`           // To add a comment to a list of issues
	DoTransition      string `form:"do_transition,omitempty"`     // Transition
	Issues            string `form:"issues,omitempty"`            // Comma-separated list of issue keys
	RemoveTags        string `form:"remove_tags,omitempty"`       // Remove tags
	SendNotifications string `form:"sendNotifications,omitempty"` //
	SetSeverity       string `form:"set_severity,omitempty"`      // To change the severity of the list of issues
	SetType           string `form:"set_type,omitempty"`          // To change the type of the list of issues
}

// BulkChangeResponse is the response for BulkChangeRequest
type BulkChangeResponse struct {
	Failures float64 `json:"failures,omitempty"`
	Ignored  float64 `json:"ignored,omitempty"`
	Success  float64 `json:"success,omitempty"`
	Total    float64 `json:"total,omitempty"`
}

// ChangelogRequest Display changelog of an issue.<br/>Requires the 'Browse' permission on the project of the specified issue.
type ChangelogRequest struct {
	Issue string `form:"issue,omitempty"` // Issue key
}

// ChangelogResponse is the response for ChangelogRequest
type ChangelogResponse struct {
	Changelog []struct {
		Avatar       string `json:"avatar,omitempty"`
		CreationDate string `json:"creationDate,omitempty"`
		Diffs        []struct {
			Key      string `json:"key,omitempty"`
			NewValue string `json:"newValue,omitempty"`
			OldValue string `json:"oldValue,omitempty"`
		} `json:"diffs,omitempty"`
		IsUserActive bool   `json:"isUserActive,omitempty"`
		User         string `json:"user,omitempty"`
		UserName     string `json:"userName,omitempty"`
	} `json:"changelog,omitempty"`
}

// DeleteCommentRequest Delete a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue.
type DeleteCommentRequest struct {
	Comment string `form:"comment,omitempty"` // Comment key
}

// DeleteCommentResponse is the response for DeleteCommentRequest
type DeleteCommentResponse struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
		Uuid      string `json:"uuid,omitempty"`
	} `json:"components,omitempty"`
	Issue struct {
		Actions  []string `json:"actions,omitempty"`
		Assignee string   `json:"assignee,omitempty"`
		Author   string   `json:"author,omitempty"`
		Comments []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string   `json:"component,omitempty"`
		CreationDate string   `json:"creationDate,omitempty"`
		Debt         string   `json:"debt,omitempty"`
		Effort       string   `json:"effort,omitempty"`
		Flows        []string `json:"flows,omitempty"`
		Key          string   `json:"key,omitempty"`
		Line         float64  `json:"line,omitempty"`
		Message      string   `json:"message,omitempty"`
		Project      string   `json:"project,omitempty"`
		Rule         string   `json:"rule,omitempty"`
		Severity     string   `json:"severity,omitempty"`
		Status       string   `json:"status,omitempty"`
		Tags         []string `json:"tags,omitempty"`
		TextRange    struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		Type        string   `json:"type,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issue,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Email  string `json:"email,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// DoTransitionRequest Do workflow transition on an issue. Requires authentication and Browse permission on project.<br/>The transitions 'wontfix' and 'falsepositive' require the permission 'Administer Issues'.<br/>The transitions involving security hotspots require the permission 'Administer Security Hotspot'.
type DoTransitionRequest struct {
	Issue      string `form:"issue,omitempty"`      // Issue key
	Transition string `form:"transition,omitempty"` // Transition
}

// DoTransitionResponse is the response for DoTransitionRequest
type DoTransitionResponse struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
		Uuid      string `json:"uuid,omitempty"`
	} `json:"components,omitempty"`
	Issue struct {
		Actions  []string `json:"actions,omitempty"`
		Assignee string   `json:"assignee,omitempty"`
		Author   string   `json:"author,omitempty"`
		Comments []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string   `json:"component,omitempty"`
		CreationDate string   `json:"creationDate,omitempty"`
		Debt         string   `json:"debt,omitempty"`
		Effort       string   `json:"effort,omitempty"`
		Flows        []string `json:"flows,omitempty"`
		Key          string   `json:"key,omitempty"`
		Line         float64  `json:"line,omitempty"`
		Message      string   `json:"message,omitempty"`
		Project      string   `json:"project,omitempty"`
		Rule         string   `json:"rule,omitempty"`
		Severity     string   `json:"severity,omitempty"`
		Status       string   `json:"status,omitempty"`
		Tags         []string `json:"tags,omitempty"`
		TextRange    struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		Type        string   `json:"type,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issue,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Email  string `json:"email,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// EditCommentRequest Edit a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue.
type EditCommentRequest struct {
	Comment string `form:"comment,omitempty"` // Comment key
	Text    string `form:"text,omitempty"`    // Comment text
}

// EditCommentResponse is the response for EditCommentRequest
type EditCommentResponse struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
		Uuid      string `json:"uuid,omitempty"`
	} `json:"components,omitempty"`
	Issue struct {
		Actions  []string `json:"actions,omitempty"`
		Assignee string   `json:"assignee,omitempty"`
		Author   string   `json:"author,omitempty"`
		Comments []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string   `json:"component,omitempty"`
		CreationDate string   `json:"creationDate,omitempty"`
		Debt         string   `json:"debt,omitempty"`
		Effort       string   `json:"effort,omitempty"`
		Flows        []string `json:"flows,omitempty"`
		Key          string   `json:"key,omitempty"`
		Line         float64  `json:"line,omitempty"`
		Message      string   `json:"message,omitempty"`
		Project      string   `json:"project,omitempty"`
		Rule         string   `json:"rule,omitempty"`
		Severity     string   `json:"severity,omitempty"`
		Status       string   `json:"status,omitempty"`
		Tags         []string `json:"tags,omitempty"`
		TextRange    struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		Type        string   `json:"type,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issue,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Email  string `json:"email,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// SearchRequest Search for issues.<br>Requires the 'Browse' permission on the specified project(s).
type SearchRequest struct {
	AdditionalFields    string `form:"additionalFields,omitempty"`    // Comma-separated list of the optional fields to be returned in response. Action plans are dropped in 5.5, it is not returned in the response.
	Asc                 string `form:"asc,omitempty"`                 // Ascending sort
	Assigned            string `form:"assigned,omitempty"`            // To retrieve assigned or unassigned issues
	Assignees           string `form:"assignees,omitempty"`           // Comma-separated list of assignee logins. The value '__me__' can be used as a placeholder for user who performs the request
	Author              string `form:"author,omitempty"`              // SCM accounts. To set several values, the parameter must be called once for each value.
	Authors             string `form:"authors,omitempty"`             // This parameter is deprecated, please use 'author' instead
	Branch              string `form:"branch,omitempty"`              // Branch key
	ComponentKeys       string `form:"componentKeys,omitempty"`       // Comma-separated list of component keys. Retrieve issues associated to a specific list of components (and all its descendants). A component can be a project, directory or file.
	CreatedAfter        string `form:"createdAfter,omitempty"`        // To retrieve issues created after the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided. <br>If this parameter is set, createdSince must not be set
	CreatedAt           string `form:"createdAt,omitempty"`           // Datetime to retrieve issues created during a specific analysis
	CreatedBefore       string `form:"createdBefore,omitempty"`       // To retrieve issues created before the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided.
	CreatedInLast       string `form:"createdInLast,omitempty"`       // To retrieve issues created during a time span before the current time (exclusive). Accepted units are 'y' for year, 'm' for month, 'w' for week and 'd' for day. If this parameter is set, createdAfter must not be set
	Cwe                 string `form:"cwe,omitempty"`                 // Comma-separated list of CWE identifiers. Use 'unknown' to select issues not associated to any CWE.
	FacetMode           string `form:"facetMode,omitempty"`           // Choose the returned value for facet items, either count of issues or sum of remediation effort.
	Facets              string `form:"facets,omitempty"`              // Comma-separated list of the facets to be computed. No facet is computed by default.
	Issues              string `form:"issues,omitempty"`              // Comma-separated list of issue keys
	Languages           string `form:"languages,omitempty"`           // Comma-separated list of languages. Available since 4.4
	OnComponentOnly     string `form:"onComponentOnly,omitempty"`     // Return only issues at a component's level, not on its descendants (modules, directories, files, etc). This parameter is only considered when componentKeys or componentUuids is set.
	Organization        string `form:"organization,omitempty"`        // Organization key
	OwaspTop10          string `form:"owaspTop10,omitempty"`          // Comma-separated list of OWASP Top 10 lowercase categories.
	PullRequest         string `form:"pullRequest,omitempty"`         // Pull request id
	Resolutions         string `form:"resolutions,omitempty"`         // Comma-separated list of resolutions
	Resolved            string `form:"resolved,omitempty"`            // To match resolved or unresolved issues
	Rules               string `form:"rules,omitempty"`               // Comma-separated list of coding rule keys. Format is &lt;repository&gt;:&lt;rule&gt;
	S                   string `form:"s,omitempty"`                   // Sort field
	SansTop25           string `form:"sansTop25,omitempty"`           // Comma-separated list of SANS Top 25 categories.
	Severities          string `form:"severities,omitempty"`          // Comma-separated list of severities
	SinceLeakPeriod     string `form:"sinceLeakPeriod,omitempty"`     // To retrieve issues created since the leak period.<br>If this parameter is set to a truthy value, createdAfter must not be set and one component id or key must be provided.
	SonarsourceSecurity string `form:"sonarsourceSecurity,omitempty"` // Comma-separated list of SonarSource security categories. Use 'others' to select issues not associated with any category
	Statuses            string `form:"statuses,omitempty"`            // Comma-separated list of statuses
	Tags                string `form:"tags,omitempty"`                // Comma-separated list of tags.
	Types               string `form:"types,omitempty"`               // Comma-separated list of types.
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Issues []struct {
		Actions []string `json:"actions,omitempty"`
		Attr    struct {
			JiraIssueKey string `json:"jira-issue-key,omitempty"`
		} `json:"attr,omitempty"`
		Author   string `json:"author,omitempty"`
		Comments []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string `json:"component,omitempty"`
		CreationDate string `json:"creationDate,omitempty"`
		Effort       string `json:"effort,omitempty"`
		Flows        []struct {
			Locations []struct {
				Msg       string `json:"msg,omitempty"`
				TextRange struct {
					EndLine     float64 `json:"endLine,omitempty"`
					EndOffset   float64 `json:"endOffset,omitempty"`
					StartLine   float64 `json:"startLine,omitempty"`
					StartOffset float64 `json:"startOffset,omitempty"`
				} `json:"textRange,omitempty"`
			} `json:"locations,omitempty"`
		} `json:"flows,omitempty"`
		Hash       string   `json:"hash,omitempty"`
		Key        string   `json:"key,omitempty"`
		Line       float64  `json:"line,omitempty"`
		Message    string   `json:"message,omitempty"`
		Project    string   `json:"project,omitempty"`
		Resolution string   `json:"resolution,omitempty"`
		Rule       string   `json:"rule,omitempty"`
		Severity   string   `json:"severity,omitempty"`
		Status     string   `json:"status,omitempty"`
		Tags       []string `json:"tags,omitempty"`
		TextRange  struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		Type        string   `json:"type,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issues,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
	Rules  []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Avatar string `json:"avatar,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// GetPaging extracts the paging from SearchResponse
func (r *SearchResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchResponseAll is the collection for SearchRequest
type SearchResponseAll struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Issues []struct {
		Actions []string `json:"actions,omitempty"`
		Attr    struct {
			JiraIssueKey string `json:"jira-issue-key,omitempty"`
		} `json:"attr,omitempty"`
		Author   string `json:"author,omitempty"`
		Comments []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string `json:"component,omitempty"`
		CreationDate string `json:"creationDate,omitempty"`
		Effort       string `json:"effort,omitempty"`
		Flows        []struct {
			Locations []struct {
				Msg       string `json:"msg,omitempty"`
				TextRange struct {
					EndLine     float64 `json:"endLine,omitempty"`
					EndOffset   float64 `json:"endOffset,omitempty"`
					StartLine   float64 `json:"startLine,omitempty"`
					StartOffset float64 `json:"startOffset,omitempty"`
				} `json:"textRange,omitempty"`
			} `json:"locations,omitempty"`
		} `json:"flows,omitempty"`
		Hash       string   `json:"hash,omitempty"`
		Key        string   `json:"key,omitempty"`
		Line       float64  `json:"line,omitempty"`
		Message    string   `json:"message,omitempty"`
		Project    string   `json:"project,omitempty"`
		Resolution string   `json:"resolution,omitempty"`
		Rule       string   `json:"rule,omitempty"`
		Severity   string   `json:"severity,omitempty"`
		Status     string   `json:"status,omitempty"`
		Tags       []string `json:"tags,omitempty"`
		TextRange  struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		Type        string   `json:"type,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issues,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Avatar string `json:"avatar,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// SetSeverityRequest Change severity.<br/>Requires the following permissions:<ul>  <li>'Authentication'</li>  <li>'Browse' rights on project of the specified issue</li>  <li>'Administer Issues' rights on project of the specified issue</li></ul>
type SetSeverityRequest struct {
	Issue    string `form:"issue,omitempty"`    // Issue key
	Severity string `form:"severity,omitempty"` // New severity
}

// SetSeverityResponse is the response for SetSeverityRequest
type SetSeverityResponse struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
		Uuid      string `json:"uuid,omitempty"`
	} `json:"components,omitempty"`
	Issue struct {
		Actions  []string `json:"actions,omitempty"`
		Assignee string   `json:"assignee,omitempty"`
		Author   string   `json:"author,omitempty"`
		Comments []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string   `json:"component,omitempty"`
		CreationDate string   `json:"creationDate,omitempty"`
		Debt         string   `json:"debt,omitempty"`
		Effort       string   `json:"effort,omitempty"`
		Flows        []string `json:"flows,omitempty"`
		Key          string   `json:"key,omitempty"`
		Line         float64  `json:"line,omitempty"`
		Message      string   `json:"message,omitempty"`
		Project      string   `json:"project,omitempty"`
		Rule         string   `json:"rule,omitempty"`
		Severity     string   `json:"severity,omitempty"`
		Status       string   `json:"status,omitempty"`
		Tags         []string `json:"tags,omitempty"`
		TextRange    struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		Type        string   `json:"type,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issue,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Email  string `json:"email,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// SetTagsRequest Set tags on an issue. <br/>Requires authentication and Browse permission on project
type SetTagsRequest struct {
	Issue string `form:"issue,omitempty"` // Issue key
	Tags  string `form:"tags,omitempty"`  // Comma-separated list of tags. All tags are removed if parameter is empty or not set.
}

// SetTagsResponse is the response for SetTagsRequest
type SetTagsResponse struct {
	Components []struct {
		Enabled      bool    `json:"enabled,omitempty"`
		Key          string  `json:"key,omitempty"`
		LongName     string  `json:"longName,omitempty"`
		Name         string  `json:"name,omitempty"`
		Path         string  `json:"path,omitempty"`
		ProjectId    float64 `json:"projectId,omitempty"`
		Qualifier    string  `json:"qualifier,omitempty"`
		SubProjectId float64 `json:"subProjectId,omitempty"`
		Uuid         string  `json:"uuid,omitempty"`
	} `json:"components,omitempty"`
	Issue struct {
		Actions  []string `json:"actions,omitempty"`
		Assignee string   `json:"assignee,omitempty"`
		Author   string   `json:"author,omitempty"`
		Comments []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string   `json:"component,omitempty"`
		CreationDate string   `json:"creationDate,omitempty"`
		Debt         string   `json:"debt,omitempty"`
		Effort       string   `json:"effort,omitempty"`
		Flows        []string `json:"flows,omitempty"`
		Key          string   `json:"key,omitempty"`
		Line         float64  `json:"line,omitempty"`
		Message      string   `json:"message,omitempty"`
		Project      string   `json:"project,omitempty"`
		Rule         string   `json:"rule,omitempty"`
		Severity     string   `json:"severity,omitempty"`
		Status       string   `json:"status,omitempty"`
		Tags         []string `json:"tags,omitempty"`
		TextRange    struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		Type        string   `json:"type,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issue,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Email  string `json:"email,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// SetTypeRequest Change type of issue, for instance from 'code smell' to 'bug'.<br/>Requires the following permissions:<ul>  <li>'Authentication'</li>  <li>'Browse' rights on project of the specified issue</li>  <li>'Administer Issues' rights on project of the specified issue</li></ul>
type SetTypeRequest struct {
	Issue string `form:"issue,omitempty"` // Issue key
	Type  string `form:"type,omitempty"`  // New type
}

// SetTypeResponse is the response for SetTypeRequest
type SetTypeResponse struct {
	Components []struct {
		Enabled   bool   `json:"enabled,omitempty"`
		Key       string `json:"key,omitempty"`
		LongName  string `json:"longName,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
		Uuid      string `json:"uuid,omitempty"`
	} `json:"components,omitempty"`
	Issue struct {
		Actions  []string `json:"actions,omitempty"`
		Assignee string   `json:"assignee,omitempty"`
		Author   string   `json:"author,omitempty"`
		Comments []struct {
			CreatedAt string `json:"createdAt,omitempty"`
			HtmlText  string `json:"htmlText,omitempty"`
			Key       string `json:"key,omitempty"`
			Login     string `json:"login,omitempty"`
			Markdown  string `json:"markdown,omitempty"`
			Updatable bool   `json:"updatable,omitempty"`
		} `json:"comments,omitempty"`
		Component    string   `json:"component,omitempty"`
		CreationDate string   `json:"creationDate,omitempty"`
		Debt         string   `json:"debt,omitempty"`
		Effort       string   `json:"effort,omitempty"`
		Flows        []string `json:"flows,omitempty"`
		Key          string   `json:"key,omitempty"`
		Line         float64  `json:"line,omitempty"`
		Message      string   `json:"message,omitempty"`
		Project      string   `json:"project,omitempty"`
		Rule         string   `json:"rule,omitempty"`
		Severity     string   `json:"severity,omitempty"`
		Status       string   `json:"status,omitempty"`
		Tags         []string `json:"tags,omitempty"`
		TextRange    struct {
			EndLine     float64 `json:"endLine,omitempty"`
			EndOffset   float64 `json:"endOffset,omitempty"`
			StartLine   float64 `json:"startLine,omitempty"`
			StartOffset float64 `json:"startOffset,omitempty"`
		} `json:"textRange,omitempty"`
		Transitions []string `json:"transitions,omitempty"`
		Type        string   `json:"type,omitempty"`
		UpdateDate  string   `json:"updateDate,omitempty"`
	} `json:"issue,omitempty"`
	Rules []struct {
		Key      string `json:"key,omitempty"`
		Lang     string `json:"lang,omitempty"`
		LangName string `json:"langName,omitempty"`
		Name     string `json:"name,omitempty"`
		Status   string `json:"status,omitempty"`
	} `json:"rules,omitempty"`
	Users []struct {
		Active bool   `json:"active,omitempty"`
		Email  string `json:"email,omitempty"`
		Login  string `json:"login,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"users,omitempty"`
}

// TagsRequest List tags matching a given query
type TagsRequest struct {
	Organization string `form:"organization,omitempty"` // Organization key
	Project      string `form:"project,omitempty"`      // Project key
	Q            string `form:"q,omitempty"`            // Limit search to tags that contain the supplied string.
}

// TagsResponse is the response for TagsRequest
type TagsResponse struct {
	Tags []string `json:"tags,omitempty"`
}
