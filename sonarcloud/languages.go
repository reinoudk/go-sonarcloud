package sonarcloud

import (
	"encoding/json"
	"fmt"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/languages"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Languages service

func (s *Languages) List(r languages.ListRequest) (*languages.ListResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/languages/list", API), params...)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return nil, fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return nil, errorResponse
		}
	}

	response := &languages.ListResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}
