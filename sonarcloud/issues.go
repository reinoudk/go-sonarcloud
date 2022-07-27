package sonarcloud

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/form/v4"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/issues"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/paging"
	"strings"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Issues service

func (s *Issues) AddComment(r issues.AddCommentRequest) (*issues.AddCommentResponse, error) {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return nil, fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/issues/add_comment", API), strings.NewReader(values.Encode()))
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

	response := &issues.AddCommentResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Issues) Assign(r issues.AssignRequest) (*issues.AssignResponse, error) {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return nil, fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/issues/assign", API), strings.NewReader(values.Encode()))
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

	response := &issues.AssignResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Issues) Authors(r issues.AuthorsRequest) (*issues.AuthorsResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/issues/authors", API), params...)
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

	response := &issues.AuthorsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Issues) BulkChange(r issues.BulkChangeRequest) (*issues.BulkChangeResponse, error) {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return nil, fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/issues/bulk_change", API), strings.NewReader(values.Encode()))
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

	response := &issues.BulkChangeResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Issues) Changelog(r issues.ChangelogRequest) (*issues.ChangelogResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/issues/changelog", API), params...)
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

	response := &issues.ChangelogResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Issues) DeleteComment(r issues.DeleteCommentRequest) (*issues.DeleteCommentResponse, error) {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return nil, fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/issues/delete_comment", API), strings.NewReader(values.Encode()))
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

	response := &issues.DeleteCommentResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Issues) DoTransition(r issues.DoTransitionRequest) (*issues.DoTransitionResponse, error) {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return nil, fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/issues/do_transition", API), strings.NewReader(values.Encode()))
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

	response := &issues.DoTransitionResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Issues) EditComment(r issues.EditCommentRequest) (*issues.EditCommentResponse, error) {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return nil, fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/issues/edit_comment", API), strings.NewReader(values.Encode()))
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

	response := &issues.EditCommentResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Issues) Search(r issues.SearchRequest, p paging.PagingParams) (*issues.SearchResponse, error) {
	params := paramsFrom(r, p)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/issues/search", API), params...)
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

	response := &issues.SearchResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Issues) SearchAll(r issues.SearchRequest) (*issues.SearchResponseAll, error) {
	p := paging.PagingParams{
		P:  1,
		Ps: 100,
	}
	response := &issues.SearchResponseAll{}
	for {
		res, err := s.Search(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to issues.Search: %+v", err)
		}
		response.Components = append(response.Components, res.Components...)
		response.Issues = append(response.Issues, res.Issues...)
		response.Rules = append(response.Rules, res.Rules...)
		response.Users = append(response.Users, res.Users...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *Issues) SetSeverity(r issues.SetSeverityRequest) (*issues.SetSeverityResponse, error) {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return nil, fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/issues/set_severity", API), strings.NewReader(values.Encode()))
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

	response := &issues.SetSeverityResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Issues) SetTags(r issues.SetTagsRequest) (*issues.SetTagsResponse, error) {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return nil, fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/issues/set_tags", API), strings.NewReader(values.Encode()))
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

	response := &issues.SetTagsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Issues) SetType(r issues.SetTypeRequest) (*issues.SetTypeResponse, error) {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return nil, fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/issues/set_type", API), strings.NewReader(values.Encode()))
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

	response := &issues.SetTypeResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Issues) Tags(r issues.TagsRequest) (*issues.TagsResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/issues/tags", API), params...)
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

	response := &issues.TagsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}
