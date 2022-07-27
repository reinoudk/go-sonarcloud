package sonarcloud

import (
	"encoding/json"
	"fmt"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/measures"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/paging"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Measures service

func (s *Measures) Component(r measures.ComponentRequest) (*measures.ComponentResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/measures/component", API), params...)
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

	response := &measures.ComponentResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Measures) ComponentTree(r measures.ComponentTreeRequest, p paging.PagingParams) (*measures.ComponentTreeResponse, error) {
	params := paramsFrom(r, p)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/measures/component_tree", API), params...)
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

	response := &measures.ComponentTreeResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Measures) ComponentTreeAll(r measures.ComponentTreeRequest) (*measures.ComponentTreeResponseAll, error) {
	p := paging.PagingParams{
		P:  1,
		Ps: 100,
	}
	response := &measures.ComponentTreeResponseAll{}
	for {
		res, err := s.ComponentTree(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to measures.ComponentTree: %+v", err)
		}
		response.Components = append(response.Components, res.Components...)
		response.Metrics = append(response.Metrics, res.Metrics...)
		response.Periods = append(response.Periods, res.Periods...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *Measures) SearchHistory(r measures.SearchHistoryRequest, p paging.PagingParams) (*measures.SearchHistoryResponse, error) {
	params := paramsFrom(r, p)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/measures/search_history", API), params...)
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

	response := &measures.SearchHistoryResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Measures) SearchHistoryAll(r measures.SearchHistoryRequest) (*measures.SearchHistoryResponseAll, error) {
	p := paging.PagingParams{
		P:  1,
		Ps: 100,
	}
	response := &measures.SearchHistoryResponseAll{}
	for {
		res, err := s.SearchHistory(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to measures.SearchHistory: %+v", err)
		}
		response.Measures = append(response.Measures, res.Measures...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}
