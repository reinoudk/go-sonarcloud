package sonarcloud

import (
	"encoding/json"
	"fmt"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/metrics"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/paging"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Metrics service

func (s *Metrics) Domains(r metrics.DomainsRequest) (*metrics.DomainsResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/metrics/domains", API), params...)
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

	response := &metrics.DomainsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Metrics) Search(r metrics.SearchRequest, p paging.PagingParams) (*metrics.SearchResponse, error) {
	params := paramsFrom(r, p)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/metrics/search", API), params...)
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

	response := &metrics.SearchResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Metrics) SearchAll(r metrics.SearchRequest) (*metrics.SearchResponseAll, error) {
	p := paging.PagingParams{
		P:  1,
		Ps: 100,
	}
	response := &metrics.SearchResponseAll{}
	for {
		res, err := s.Search(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to metrics.Search: %+v", err)
		}
		response.Metrics = append(response.Metrics, res.Metrics...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *Metrics) Types(r metrics.TypesRequest) (*metrics.TypesResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/metrics/types", API), params...)
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

	response := &metrics.TypesResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}
