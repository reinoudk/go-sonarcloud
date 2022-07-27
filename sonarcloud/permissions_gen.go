package sonarcloud

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/form/v4"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/paging"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/permissions"
	"strings"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Permissions service

func (s *Permissions) AddGroup(r permissions.AddGroupRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/permissions/add_group", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Permissions) AddGroupToTemplate(r permissions.AddGroupToTemplateRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/permissions/add_group_to_template", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Permissions) AddProjectCreatorToTemplate(r permissions.AddProjectCreatorToTemplateRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/permissions/add_project_creator_to_template", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Permissions) AddUser(r permissions.AddUserRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/permissions/add_user", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Permissions) AddUserToTemplate(r permissions.AddUserToTemplateRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/permissions/add_user_to_template", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Permissions) ApplyTemplate(r permissions.ApplyTemplateRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/permissions/apply_template", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Permissions) BulkApplyTemplate(r permissions.BulkApplyTemplateRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/permissions/bulk_apply_template", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Permissions) CreateTemplate(r permissions.CreateTemplateRequest) (*permissions.CreateTemplateResponse, error) {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return nil, fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/permissions/create_template", API), strings.NewReader(values.Encode()))
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

	response := &permissions.CreateTemplateResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Permissions) DeleteTemplate(r permissions.DeleteTemplateRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/permissions/delete_template", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Permissions) RemoveGroup(r permissions.RemoveGroupRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/permissions/remove_group", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Permissions) RemoveGroupFromTemplate(r permissions.RemoveGroupFromTemplateRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/permissions/remove_group_from_template", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Permissions) RemoveProjectCreatorFromTemplate(r permissions.RemoveProjectCreatorFromTemplateRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/permissions/remove_project_creator_from_template", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Permissions) RemoveUser(r permissions.RemoveUserRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/permissions/remove_user", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Permissions) RemoveUserFromTemplate(r permissions.RemoveUserFromTemplateRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/permissions/remove_user_from_template", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Permissions) SearchGlobalPermissions(r permissions.SearchGlobalPermissionsRequest) (*permissions.SearchGlobalPermissionsResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/permissions/search_global_permissions", API), params...)
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

	response := &permissions.SearchGlobalPermissionsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Permissions) SearchProjectPermissions(r permissions.SearchProjectPermissionsRequest, p paging.PagingParams) (*permissions.SearchProjectPermissionsResponse, error) {
	params := paramsFrom(r, p)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/permissions/search_project_permissions", API), params...)
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

	response := &permissions.SearchProjectPermissionsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Permissions) SearchProjectPermissionsAll(r permissions.SearchProjectPermissionsRequest) (*permissions.SearchProjectPermissionsResponseAll, error) {
	p := paging.PagingParams{
		P:  1,
		Ps: 100,
	}
	response := &permissions.SearchProjectPermissionsResponseAll{}
	for {
		res, err := s.SearchProjectPermissions(r, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to permissions.SearchProjectPermissions: %+v", err)
		}
		response.Permissions = append(response.Permissions, res.Permissions...)
		response.Projects = append(response.Projects, res.Projects...)
		if res.GetPaging().End() {
			break
		} else {
			p.P++
		}
	}
	return response, nil
}

func (s *Permissions) SearchTemplates(r permissions.SearchTemplatesRequest) error {
	params := paramsFrom(r)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/permissions/search_templates", API), params...)
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Permissions) SetDefaultTemplate(r permissions.SetDefaultTemplateRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/permissions/set_default_template", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Permissions) UpdateTemplate(r permissions.UpdateTemplateRequest) (*permissions.UpdateTemplateResponse, error) {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return nil, fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/permissions/update_template", API), strings.NewReader(values.Encode()))
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

	response := &permissions.UpdateTemplateResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}
