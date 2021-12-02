package sonarcloud

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/form/v4"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/qualitygates"
	"strings"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Qualitygates service

func (s *Qualitygates) Copy(r qualitygates.CopyRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/qualitygates/copy", API), strings.NewReader(values.Encode()))
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

func (s *Qualitygates) Create(r qualitygates.CreateRequest) (*qualitygates.CreateResponse, error) {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return nil, fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/qualitygates/create", API), strings.NewReader(values.Encode()))
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

	response := &qualitygates.CreateResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Qualitygates) CreateCondition(r qualitygates.CreateConditionRequest) (*qualitygates.CreateConditionResponse, error) {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return nil, fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/qualitygates/create_condition", API), strings.NewReader(values.Encode()))
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

	response := &qualitygates.CreateConditionResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Qualitygates) DeleteCondition(r qualitygates.DeleteConditionRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/qualitygates/delete_condition", API), strings.NewReader(values.Encode()))
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

func (s *Qualitygates) Deselect(r qualitygates.DeselectRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/qualitygates/deselect", API), strings.NewReader(values.Encode()))
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

func (s *Qualitygates) Destroy(r qualitygates.DestroyRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/qualitygates/destroy", API), strings.NewReader(values.Encode()))
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

func (s *Qualitygates) GetByProject(r qualitygates.GetByProjectRequest) (*qualitygates.GetByProjectResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.NewRequestWithParameters("GET", fmt.Sprintf("%s/qualitygates/get_by_project", API), params...)
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

	response := &qualitygates.GetByProjectResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Qualitygates) List(r qualitygates.ListRequest) (*qualitygates.ListResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.NewRequestWithParameters("GET", fmt.Sprintf("%s/qualitygates/list", API), params...)
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

	response := &qualitygates.ListResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Qualitygates) ProjectStatus(r qualitygates.ProjectStatusRequest) (*qualitygates.ProjectStatusResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.NewRequestWithParameters("GET", fmt.Sprintf("%s/qualitygates/project_status", API), params...)
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

	response := &qualitygates.ProjectStatusResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Qualitygates) Rename(r qualitygates.RenameRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/qualitygates/rename", API), strings.NewReader(values.Encode()))
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

func (s *Qualitygates) Search(r qualitygates.SearchRequest) (*qualitygates.SearchResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.NewRequestWithParameters("GET", fmt.Sprintf("%s/qualitygates/search", API), params...)
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

	response := &qualitygates.SearchResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Qualitygates) Select(r qualitygates.SelectRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/qualitygates/select", API), strings.NewReader(values.Encode()))
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

func (s *Qualitygates) SetAsDefault(r qualitygates.SetAsDefaultRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/qualitygates/set_as_default", API), strings.NewReader(values.Encode()))
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

func (s *Qualitygates) Show(r qualitygates.ShowRequest) (*qualitygates.ShowResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.NewRequestWithParameters("GET", fmt.Sprintf("%s/qualitygates/show", API), params...)
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

	response := &qualitygates.ShowResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Qualitygates) UnsetDefault(r qualitygates.UnsetDefaultRequest) (*qualitygates.UnsetDefaultResponse, error) {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return nil, fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/qualitygates/unset_default", API), strings.NewReader(values.Encode()))
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

	response := &qualitygates.UnsetDefaultResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Qualitygates) UpdateCondition(r qualitygates.UpdateConditionRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("%s/qualitygates/update_condition", API), strings.NewReader(values.Encode()))
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
