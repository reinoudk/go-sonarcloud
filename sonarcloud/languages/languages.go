package languages

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ListRequest List supported programming languages
type ListRequest struct {
	Q string `form:"q,omitempty"` // A pattern to match language keys/names against
}

// ListResponse is the response for ListRequest
type ListResponse struct {
	Languages []struct {
		Key  string `json:"key,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"languages,omitempty"`
}
