package metrics

import paging "github.com/reinoudk/go-sonarcloud/sonarcloud/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// DomainsRequest List all custom metric domains.
// Deprecated: this action has been deprecated since version 7.7
type DomainsRequest struct{}

// DomainsResponse is the response for DomainsRequest
type DomainsResponse struct {
	Domains []string `json:"domains,omitempty"`
}

// SearchRequest Search for metrics
type SearchRequest struct {
	F string `form:"f,omitempty"` // Comma-separated list of the fields to be returned in response. All the fields are returned by default.
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Metrics []struct {
		Custom      bool    `json:"custom,omitempty"`
		Description string  `json:"description,omitempty"`
		Direction   float64 `json:"direction,omitempty"`
		Domain      string  `json:"domain,omitempty"`
		Hidden      bool    `json:"hidden,omitempty"`
		Id          float64 `json:"id,omitempty"`
		Key         string  `json:"key,omitempty"`
		Name        string  `json:"name,omitempty"`
		Qualitative bool    `json:"qualitative,omitempty"`
		Type        string  `json:"type,omitempty"`
	} `json:"metrics,omitempty"`
	P     float64 `json:"p,omitempty"`
	Ps    float64 `json:"ps,omitempty"`
	Total float64 `json:"total,omitempty"`
}

// GetPaging extracts the paging from SearchResponse
func (r *SearchResponse) GetPaging() *paging.Paging {
	return &paging.Paging{

		PageIndex: int(r.P),
		PageSize:  int(r.Ps),
		Total:     int(r.Total),
	}
}

// SearchResponseAll is the collection for SearchRequest
type SearchResponseAll struct {
	Metrics []struct {
		Custom      bool    `json:"custom,omitempty"`
		Description string  `json:"description,omitempty"`
		Direction   float64 `json:"direction,omitempty"`
		Domain      string  `json:"domain,omitempty"`
		Hidden      bool    `json:"hidden,omitempty"`
		Id          float64 `json:"id,omitempty"`
		Key         string  `json:"key,omitempty"`
		Name        string  `json:"name,omitempty"`
		Qualitative bool    `json:"qualitative,omitempty"`
		Type        string  `json:"type,omitempty"`
	} `json:"metrics,omitempty"`
}

// TypesRequest List all available metric types.
type TypesRequest struct{}

// TypesResponse is the response for TypesRequest
type TypesResponse struct {
	Types []string `json:"types,omitempty"`
}
