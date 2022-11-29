package sonarcloud

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/form/v4"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/paging"
	"strings"
)

// Get returns the items that are found under the given items-key of the returned JSON object of the given page.
func Get[T any, R any, RR ~[]R](client *Client, path string, request T, itemsKey string, pagingParams paging.Params) (RR, *paging.Paging, error) {
	params := paramsFrom(request, pagingParams)

	req, err := client.GetRequest(API+path, params...)
	if err != nil {
		return nil, nil, fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return nil, nil, fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return nil, nil, errorResponse
		}
	}

	rawResponse := map[string]json.RawMessage{}
	if err = json.NewDecoder(resp.Body).Decode(&rawResponse); err != nil {
		return nil, nil, fmt.Errorf("could not decode response: %+v", err)
	}

	items := make([]R, 0)
	if err = json.Unmarshal(rawResponse[itemsKey], &items); err != nil {
		return nil, nil, fmt.Errorf("could not unmarshall items under item-key '%s': %+v", itemsKey, err)
	}

	pager := paging.Paging{}
	if err = json.Unmarshal(rawResponse["paging"], &pager); err != nil {
		return nil, nil, fmt.Errorf("could not unmarshall paging: %+v", err)
	}

	return RR(items), &pager, nil
}

// GetAll returns the items from all pages of a particular endpoint.
// The itemsKey should be the key under which the array of items is found in each individual response from the endpoint.
func GetAll[T any, R any, RR ~[]R](client *Client, path string, request T, itemsKey string) (RR, error) {
	p := paging.Params{
		P:  1,
		Ps: 100,
	}
	allItems := make([]R, 0)
	for {
		items, pager, err := Get[T, R, RR](client, path, request, itemsKey, p)
		if err != nil {
			return nil, fmt.Errorf("error during call to %s: , %+v", path, err)
		}

		allItems = append(allItems, []R(items)...)

		if pager.End() {
			break
		} else {
			p.P++
		}
	}
	return RR(allItems), nil
}

// Post sends a POST request without returning a response.
func Post[T any](client *Client, path string, request T) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(request)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := client.PostRequest(API+path, strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := client.Do(req)
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

// PostWithResponse sends a POST request and returns the unmarshalled JSON response.
func PostWithResponse[T any, R any](client *Client, path string, request T) (*R, error) {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(request)
	if err != nil {
		return nil, fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := client.PostRequest(API+path, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := client.Do(req)
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

	response := new(R)
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}
