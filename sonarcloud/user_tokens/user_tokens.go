package user_tokens

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// GenerateRequest Generate a user access token. <br />Please keep your tokens secret. They enable to authenticate and analyze projects.<br />It requires administration permissions to specify a 'login' and generate a token for another user. Otherwise, a token is generated for the current user.
type GenerateRequest struct {
	Login string `form:"login,omitempty"` // User login. If not set, the token is generated for the authenticated user.
	Name  string `form:"name,omitempty"`  // Token name
}

// GenerateResponse is the response for GenerateRequest
type GenerateResponse struct {
	CreatedAt string `json:"createdAt,omitempty"`
	Login     string `json:"login,omitempty"`
	Name      string `json:"name,omitempty"`
	Token     string `json:"Token,omitempty"`
}

// RevokeRequest Revoke a user access token. <br/>It requires administration permissions to specify a 'login' and revoke a token for another user. Otherwise, the token for the current user is revoked.
type RevokeRequest struct {
	Login string `form:"login,omitempty"` // User login
	Name  string `form:"name,omitempty"`  // Token name
}

// SearchRequest List the access tokens of a user.<br>The login must exist and active.<br>Field 'lastConnectionDate' is only updated every hour, so it may not be accurate, for instance when a user is using a token many times in less than one hour.<br/It requires administration permissions to specify a 'login' and list the tokens of another user. Otherwise, tokens for the current user are listed.
type SearchRequest struct {
	Login string `form:"login,omitempty"` // User login
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Login      string `json:"login,omitempty"`
	UserTokens []struct {
		CreatedAt string `json:"createdAt,omitempty"`
		Name      string `json:"name,omitempty"`
	} `json:"userTokens,omitempty"`
}
