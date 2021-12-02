package measures

import paging "github.com/reinoudk/go-sonarcloud/sonarcloud/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ComponentRequest Return component with specified measures. The componentId or the component parameter must be provided.<br>Requires the following permission: 'Browse' on the project of specified component.
type ComponentRequest struct {
	AdditionalFields string `form:"additionalFields,omitempty"` // Comma-separated list of additional fields that can be returned in the response.
	Branch           string `form:"branch,omitempty"`           // Branch key
	Component        string `form:"component,omitempty"`        // Component key
	ComponentId      string `form:"componentId,omitempty"`      // Component id
	DeveloperId      string `form:"developerId,omitempty"`      // Deprecated parameter, used previously with the Developer Cockpit plugin. No measures are returned if parameter is set.
	DeveloperKey     string `form:"developerKey,omitempty"`     // Deprecated parameter, used previously with the Developer Cockpit plugin. No measures are returned if parameter is set.
	MetricKeys       string `form:"metricKeys,omitempty"`       // Comma-separated list of metric keys
	PullRequest      string `form:"pullRequest,omitempty"`      // Pull request id
}

// ComponentResponse is the response for ComponentRequest
type ComponentResponse struct {
	Component struct {
		Key      string `json:"key,omitempty"`
		Language string `json:"language,omitempty"`
		Measures []struct {
			Metric string `json:"metric,omitempty"`
			Value  string `json:"value,omitempty"`
		} `json:"measures,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"component,omitempty"`
	Metrics []struct {
		Custom                bool   `json:"custom,omitempty"`
		Description           string `json:"description,omitempty"`
		Domain                string `json:"domain,omitempty"`
		Hidden                bool   `json:"hidden,omitempty"`
		HigherValuesAreBetter bool   `json:"higherValuesAreBetter,omitempty"`
		Key                   string `json:"key,omitempty"`
		Name                  string `json:"name,omitempty"`
		Qualitative           bool   `json:"qualitative,omitempty"`
		Type                  string `json:"type,omitempty"`
	} `json:"metrics,omitempty"`
	Periods []struct {
		Date      string  `json:"date,omitempty"`
		Index     float64 `json:"index,omitempty"`
		Mode      string  `json:"mode,omitempty"`
		Parameter string  `json:"parameter,omitempty"`
	} `json:"periods,omitempty"`
}

// ComponentTreeRequest Navigate through components based on the chosen strategy with specified measures. The baseComponentId or the component parameter must be provided.<br>Requires the following permission: 'Browse' on the specified project.<br>When limiting search with the q parameter, directories are not returned.
type ComponentTreeRequest struct {
	AdditionalFields string `form:"additionalFields,omitempty"` // Comma-separated list of additional fields that can be returned in the response.
	Asc              string `form:"asc,omitempty"`              // Ascending sort
	BaseComponentId  string `form:"baseComponentId,omitempty"`  // Base component id. The search is based on this component.
	Branch           string `form:"branch,omitempty"`           // Branch key
	Component        string `form:"component,omitempty"`        // Component key. The search is based on this component.
	DeveloperId      string `form:"developerId,omitempty"`      // Deprecated parameter, used previously with the Developer Cockpit plugin. No measures are returned if parameter is set.
	DeveloperKey     string `form:"developerKey,omitempty"`     // Deprecated parameter, used previously with the Developer Cockpit plugin. No measures are returned if parameter is set.
	MetricKeys       string `form:"metricKeys,omitempty"`       // Comma-separated list of metric keys. Types DISTRIB, DATA are not allowed.
	MetricPeriodSort string `form:"metricPeriodSort,omitempty"` // Sort measures by leak period or not ?. The 's' parameter must contain the 'metricPeriod' value.
	MetricSort       string `form:"metricSort,omitempty"`       // Metric key to sort by. The 's' parameter must contain the 'metric' or 'metricPeriod' value. It must be part of the 'metricKeys' parameter
	MetricSortFilter string `form:"metricSortFilter,omitempty"` // Filter components. Sort must be on a metric. Possible values are: <ul><li>all: return all components</li><li>withMeasuresOnly: filter out components that do not have a measure on the sorted metric</li></ul>
	PullRequest      string `form:"pullRequest,omitempty"`      // Pull request id
	Q                string `form:"q,omitempty"`                // Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that are exactly the same as the supplied string</li></ul>
	Qualifiers       string `form:"qualifiers,omitempty"`       // Comma-separated list of component qualifiers. Filter the results with the specified qualifiers. Possible values are:<ul><li>BRC - Sub-projects</li><li>DIR - Directories</li><li>FIL - Files</li><li>TRK - Projects</li><li>UTS - Test Files</li></ul>
	S                string `form:"s,omitempty"`                // Comma-separated list of sort fields
	Strategy         string `form:"strategy,omitempty"`         // Strategy to search for base component descendants:<ul><li>children: return the children components of the base component. Grandchildren components are not returned</li><li>all: return all the descendants components of the base component. Grandchildren are returned.</li><li>leaves: return all the descendant components (files, in general) which don't have other children. They are the leaves of the component tree.</li></ul>
}

// ComponentTreeResponse is the response for ComponentTreeRequest
type ComponentTreeResponse struct {
	BaseComponent struct {
		Key      string `json:"key,omitempty"`
		Measures []struct {
			Metric  string `json:"metric,omitempty"`
			Periods []struct {
				Index float64 `json:"index,omitempty"`
				Value string  `json:"value,omitempty"`
			} `json:"periods,omitempty"`
		} `json:"measures,omitempty"`
		Name      string `json:"name,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"baseComponent,omitempty"`
	Components []struct {
		Key      string `json:"key,omitempty"`
		Language string `json:"language,omitempty"`
		Measures []struct {
			Metric  string `json:"metric,omitempty"`
			Periods []struct {
				Index float64 `json:"index,omitempty"`
				Value string  `json:"value,omitempty"`
			} `json:"periods,omitempty"`
		} `json:"measures,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Metrics []struct {
		Custom                bool   `json:"custom,omitempty"`
		Description           string `json:"description,omitempty"`
		Domain                string `json:"domain,omitempty"`
		Hidden                bool   `json:"hidden,omitempty"`
		HigherValuesAreBetter bool   `json:"higherValuesAreBetter,omitempty"`
		Key                   string `json:"key,omitempty"`
		Name                  string `json:"name,omitempty"`
		Qualitative           bool   `json:"qualitative,omitempty"`
		Type                  string `json:"type,omitempty"`
	} `json:"metrics,omitempty"`
	Paging  paging.Paging `json:"paging,omitempty"`
	Periods []struct {
		Date      string  `json:"date,omitempty"`
		Index     float64 `json:"index,omitempty"`
		Mode      string  `json:"mode,omitempty"`
		Parameter string  `json:"parameter,omitempty"`
	} `json:"periods,omitempty"`
}

// GetPaging extracts the paging from ComponentTreeResponse
func (r *ComponentTreeResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// ComponentTreeResponseAll is the collection for ComponentTreeRequest
type ComponentTreeResponseAll struct {
	BaseComponent struct {
		Key      string `json:"key,omitempty"`
		Measures []struct {
			Metric  string `json:"metric,omitempty"`
			Periods []struct {
				Index float64 `json:"index,omitempty"`
				Value string  `json:"value,omitempty"`
			} `json:"periods,omitempty"`
		} `json:"measures,omitempty"`
		Name      string `json:"name,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"baseComponent,omitempty"`
	Components []struct {
		Key      string `json:"key,omitempty"`
		Language string `json:"language,omitempty"`
		Measures []struct {
			Metric  string `json:"metric,omitempty"`
			Periods []struct {
				Index float64 `json:"index,omitempty"`
				Value string  `json:"value,omitempty"`
			} `json:"periods,omitempty"`
		} `json:"measures,omitempty"`
		Name      string `json:"name,omitempty"`
		Path      string `json:"path,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Metrics []struct {
		Custom                bool   `json:"custom,omitempty"`
		Description           string `json:"description,omitempty"`
		Domain                string `json:"domain,omitempty"`
		Hidden                bool   `json:"hidden,omitempty"`
		HigherValuesAreBetter bool   `json:"higherValuesAreBetter,omitempty"`
		Key                   string `json:"key,omitempty"`
		Name                  string `json:"name,omitempty"`
		Qualitative           bool   `json:"qualitative,omitempty"`
		Type                  string `json:"type,omitempty"`
	} `json:"metrics,omitempty"`
	Periods []struct {
		Date      string  `json:"date,omitempty"`
		Index     float64 `json:"index,omitempty"`
		Mode      string  `json:"mode,omitempty"`
		Parameter string  `json:"parameter,omitempty"`
	} `json:"periods,omitempty"`
}

// SearchHistoryRequest Search measures history of a component.<br>Measures are ordered chronologically.<br>Pagination applies to the number of measures for each metric.<br>Requires the following permission: 'Browse' on the specified component
type SearchHistoryRequest struct {
	Branch      string `form:"branch,omitempty"`      // Branch key
	Component   string `form:"component,omitempty"`   // Component key
	From        string `form:"from,omitempty"`        // Filter measures created after the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided
	Metrics     string `form:"metrics,omitempty"`     // Comma-separated list of metric keys
	PullRequest string `form:"pullRequest,omitempty"` // Pull request id
	To          string `form:"to,omitempty"`          // Filter measures created before the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided
}

// SearchHistoryResponse is the response for SearchHistoryRequest
type SearchHistoryResponse struct {
	Measures []struct {
		History []struct {
			Date  string `json:"date,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"history,omitempty"`
		Metric string `json:"metric,omitempty"`
	} `json:"measures,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from SearchHistoryResponse
func (r *SearchHistoryResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// SearchHistoryResponseAll is the collection for SearchHistoryRequest
type SearchHistoryResponseAll struct {
	Measures []struct {
		History []struct {
			Date  string `json:"date,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"history,omitempty"`
		Metric string `json:"metric,omitempty"`
	} `json:"measures,omitempty"`
}
