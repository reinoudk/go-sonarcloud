package webhooks

import paging "github.com/reinoudk/go-sonarcloud/sonarcloud/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// CreateRequest Create a Webhook.<br>Requires 'Administer' permission on the specified project.
type CreateRequest struct {
	Name         string `form:"name,omitempty"`         // Name displayed in the administration console of webhooks
	Organization string `form:"organization,omitempty"` // The key of the organization that will own the webhook
	Project      string `form:"project,omitempty"`      // The key of the project that will own the webhook
	Secret       string `form:"secret,omitempty"`       // If provided, secret will be used as the key to generate the HMAC hex (lowercase) digest value in the 'X-Sonar-Webhook-HMAC-SHA256' header
	Url          string `form:"url,omitempty"`          // Server endpoint that will receive the webhook payload, for example 'http://my_server/foo'. If HTTP Basic authentication is used, HTTPS is recommended to avoid man in the middle attacks. Example: 'https://myLogin:myPassword@my_server/foo'
}

// CreateResponse is the response for CreateRequest
type CreateResponse struct {
	Webhook struct {
		Key    string `json:"key,omitempty"`
		Name   string `json:"name,omitempty"`
		Secret string `json:"secret,omitempty"`
		Url    string `json:"url,omitempty"`
	} `json:"webhook,omitempty"`
}

// DeleteRequest Delete a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
type DeleteRequest struct {
	Webhook string `form:"webhook,omitempty"` // The key of the webhook to be deleted, auto-generated value can be obtained through api/webhooks/create or api/webhooks/list
}

// DeliveriesRequest Get the recent deliveries for a specified project or Compute Engine task.<br/>Require 'Administer' permission on the related project.<br/>Note that additional information are returned by api/webhooks/delivery.
type DeliveriesRequest struct {
	CeTaskId     string `form:"ceTaskId,omitempty"`     // Id of the Compute Engine task
	ComponentKey string `form:"componentKey,omitempty"` // Key of the project
	Webhook      string `form:"webhook,omitempty"`      // Key of the webhook that triggered those deliveries, auto-generated value that can be obtained through api/webhooks/create or api/webhooks/list
}

// DeliveriesResponse is the response for DeliveriesRequest
type DeliveriesResponse struct {
	Deliveries []struct {
		At           string  `json:"at,omitempty"`
		CeTaskId     string  `json:"ceTaskId,omitempty"`
		ComponentKey string  `json:"componentKey,omitempty"`
		DurationMs   float64 `json:"durationMs,omitempty"`
		HttpStatus   float64 `json:"httpStatus,omitempty"`
		Id           string  `json:"id,omitempty"`
		Name         string  `json:"name,omitempty"`
		Success      bool    `json:"success,omitempty"`
		Url          string  `json:"url,omitempty"`
	} `json:"deliveries,omitempty"`
	Paging paging.Paging `json:"paging,omitempty"`
}

// GetPaging extracts the paging from DeliveriesResponse
func (r *DeliveriesResponse) GetPaging() *paging.Paging {
	return &r.Paging
}

// DeliveriesResponseAll is the collection for DeliveriesRequest
type DeliveriesResponseAll struct {
	Deliveries []struct {
		At           string  `json:"at,omitempty"`
		CeTaskId     string  `json:"ceTaskId,omitempty"`
		ComponentKey string  `json:"componentKey,omitempty"`
		DurationMs   float64 `json:"durationMs,omitempty"`
		HttpStatus   float64 `json:"httpStatus,omitempty"`
		Id           string  `json:"id,omitempty"`
		Name         string  `json:"name,omitempty"`
		Success      bool    `json:"success,omitempty"`
		Url          string  `json:"url,omitempty"`
	} `json:"deliveries,omitempty"`
}

// DeliveryRequest Get a webhook delivery by its id.<br/>Note that additional information are returned by api/webhooks/delivery.
type DeliveryRequest struct {
	DeliveryId string `form:"deliveryId,omitempty"` // Id of delivery
}

// DeliveryResponse is the response for DeliveryRequest
type DeliveryResponse struct {
	Delivery struct {
		At           string  `json:"at,omitempty"`
		CeTaskId     string  `json:"ceTaskId,omitempty"`
		ComponentKey string  `json:"componentKey,omitempty"`
		DurationMs   float64 `json:"durationMs,omitempty"`
		HttpStatus   float64 `json:"httpStatus,omitempty"`
		Id           string  `json:"id,omitempty"`
		Name         string  `json:"name,omitempty"`
		Payload      string  `json:"payload,omitempty"`
		Success      bool    `json:"success,omitempty"`
		Url          string  `json:"url,omitempty"`
	} `json:"delivery,omitempty"`
}

// ListRequest Search for global webhooks or project webhooks. Webhooks are ordered by name.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
type ListRequest struct {
	Organization string `form:"organization,omitempty"` // Organization key
	Project      string `form:"project,omitempty"`      // Project key
}

// ListResponse is the response for ListRequest
type ListResponse struct {
	Webhooks []struct {
		Key    string `json:"key,omitempty"`
		Name   string `json:"name,omitempty"`
		Url    string `json:"url,omitempty"`
		Secret string `json:"secret,omitempty"`
	} `json:"webhooks,omitempty"`
}

// UpdateRequest Update a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
type UpdateRequest struct {
	Name    string `form:"name,omitempty"`    // new name of the webhook
	Secret  string `form:"secret,omitempty"`  // If provided, secret will be used as the key to generate the HMAC hex (lowercase) digest value in the 'X-Sonar-Webhook-HMAC-SHA256' header
	Url     string `form:"url,omitempty"`     // new url to be called by the webhook
	Webhook string `form:"webhook,omitempty"` // The key of the webhook to be updated, auto-generated value can be obtained through api/webhooks/create or api/webhooks/list
}
