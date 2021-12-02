package sonarcloud

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/form/v4"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/project_tags"
	"strings"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type ProjectTags service

func (s *ProjectTags) Search(r project_tags.SearchRequest) (*project_tags.SearchResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.NewRequestWithParameters("GET", fmt.Sprintf("%s/project_tags/search", API), params...)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		// TODO: parse error message
		return nil, fmt.Errorf("received non 2xx status code: %d", resp.StatusCode)
	}

	response := &project_tags.SearchResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *ProjectTags) Set(r project_tags.SetRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/project_tags/set", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		// TODO: parse error message
		return fmt.Errorf("received non 2xx status code: %d", resp.StatusCode)
	}

	return nil
}
