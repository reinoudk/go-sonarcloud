package paging

// Params are the params used for paging the request
type Params struct {
	P  int // 1-based page number
	Ps int // Page size. Must be greater than 0 and less or equal than 500
}

// Paging is the part of the response that contains paging info
type Paging struct {
	PageIndex int `json:"pageIndex,omitempty"`
	PageSize  int `json:"pageSize,omitempty"`
	Total     int `json:"total,omitempty"`
}

// End returns whether the last page has been reached or not
func (p *Paging) End() bool {
	return p.PageIndex * p.PageSize >= p.Total
}
