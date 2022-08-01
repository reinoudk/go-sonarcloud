// Copyright 2013 The go-github AUTHORS. All rights reserved.
// Copyright 2021 Reinoud Kruithof
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sonarcloud

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"

	"github.com/iancoleman/strcase"
)

// API FIXME: this should be injected
const API = "https://sonarcloud.io/api"

type Client struct {
	client *http.Client
	org    string
	token  string

	common service

	Authentication      *Authentication
	Ce                  *Ce
	Components          *Components
	Favorites           *Favorites
	Issues              *Issues
	Languages           *Languages
	Measures            *Measures
	Metrics             *Metrics
	Notifications       *Notifications
	Permissions         *Permissions
	ProjectAnalyses     *ProjectAnalyses
	ProjectBranches     *ProjectBranches
	ProjectLinks        *ProjectLinks
	ProjectPullRequests *ProjectPullRequests
	ProjectTags         *ProjectTags
	Projects            *Projects
	Qualitygates        *Qualitygates
	Rules               *Rules
	Settings            *Settings
	Timemachine         *Timemachine
	UserGroups          *UserGroups
	UserProperties      *UserProperties
	UserTokens          *UserTokens
	Users               *Users
	Webhooks            *Webhooks
	Webservices         *Webservices
}

type service struct {
	client *Client
}

func NewClient(org string, token string, client *http.Client) *Client {
	if client == nil {
		client = &http.Client{}
	}

	c := &Client{
		client: client,
		org:    org,
		token:  token,
	}
	c.common.client = c
	c.Authentication = (*Authentication)(&c.common)
	c.Ce = (*Ce)(&c.common)
	c.Components = (*Components)(&c.common)
	c.Favorites = (*Favorites)(&c.common)
	c.Issues = (*Issues)(&c.common)
	c.Languages = (*Languages)(&c.common)
	c.Measures = (*Measures)(&c.common)
	c.Metrics = (*Metrics)(&c.common)
	c.Notifications = (*Notifications)(&c.common)
	c.Permissions = (*Permissions)(&c.common)
	c.ProjectAnalyses = (*ProjectAnalyses)(&c.common)
	c.ProjectBranches = (*ProjectBranches)(&c.common)
	c.ProjectLinks = (*ProjectLinks)(&c.common)
	c.ProjectPullRequests = (*ProjectPullRequests)(&c.common)
	c.ProjectTags = (*ProjectTags)(&c.common)
	c.Projects = (*Projects)(&c.common)
	c.Qualitygates = (*Qualitygates)(&c.common)
	c.Rules = (*Rules)(&c.common)
	c.Settings = (*Settings)(&c.common)
	c.Timemachine = (*Timemachine)(&c.common)
	c.UserGroups = (*UserGroups)(&c.common)
	c.UserProperties = (*UserProperties)(&c.common)
	c.UserTokens = (*UserTokens)(&c.common)
	c.Users = (*Users)(&c.common)
	c.Webhooks = (*Webhooks)(&c.common)
	c.Webservices = (*Webservices)(&c.common)

	return c
}

func (c *Client) PostRequest(url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.token, "")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *Client) GetRequest(url string, params ...string) (*http.Request, error) {
	if l := len(params); l%2 != 0 {
		return nil, fmt.Errorf("params must be an even number, %d given", l)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("organization", c.org)

	for i := 0; i < len(params); i++ {
		q.Add(params[i], params[i+1])
		i++
	}
	req.URL.RawQuery = q.Encode()

	req.SetBasicAuth(c.token, "")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

// paramsFrom creates a slice with interleaving param and value entries, i.e. ["key1", "value1", "key2, "value2"]
func paramsFrom(items ...interface{}) []string {
	allParams := make([]string, 0)

	for _, item := range items {
		v := reflect.ValueOf(item)
		t := v.Type()

		params := make([]string, 2*v.NumField())

		for i := 0; i < v.NumField(); i++ {
			j := i * 2
			k := j + 1

			if v.Field(i).IsZero() {
				continue
			}

			// Convert some basic types to strings for convenience.
			// Note: other types should not be used as parameter values.
			fieldValue := ""
			switch t.Field(i).Type.Name() {
			case "int":
				fieldValue = strconv.Itoa(v.Field(i).Interface().(int))
			case "string":
				fieldValue = v.Field(i).Interface().(string)
			case "bool":
				fieldValue = strconv.FormatBool(v.Field(i).Interface().(bool))
			}

			params[j] = strcase.ToCamel(t.Field(i).Name)
			params[k] = fieldValue
		}

		allParams = append(allParams, params...)
	}

	return allParams
}
