package sonarcloud

import (
	"encoding/json"
	"fmt"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/components"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/paging"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Components service

func (s *Components) Search(r components.SearchRequest, p paging.PagingParams) (*components.SearchResponse, error) {
	params := paramsFrom(r, p)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/components/search", API), params...)
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

	response := &components.SearchResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Components) SearchAll(r components.SearchRequest) (*components.SearchResponseAll, error) {
	p := paging.PagingParams{
		P:  1,
		Ps: 100,
	}
	response := &components.SearchResponseAll{}
	for {
		res, err := s.Search(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to components.Search: %+v", err)
		}
		response.Components = append(response.Components, res.Components...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *Components) Show(r components.ShowRequest) (*components.ShowResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/components/show", API), params...)
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

	response := &components.ShowResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Components) Tree(r components.TreeRequest, p paging.PagingParams) (*components.TreeResponse, error) {
	params := paramsFrom(r, p)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/components/tree", API), params...)
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

	response := &components.TreeResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Components) TreeAll(r components.TreeRequest) (*components.TreeResponseAll, error) {
	p := paging.PagingParams{
		P:  1,
		Ps: 100,
	}
	response := &components.TreeResponseAll{}
	for {
		res, err := s.Tree(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to components.Tree: %+v", err)
		}
		response.Components = append(response.Components, res.Components...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}
