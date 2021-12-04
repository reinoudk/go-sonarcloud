// This small program is used to generate request structs and services for the SonarCloud API.
// It expects a JSON file with the same structure as returned by `https://sonarcloud.io/api/webservices/list`.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

type Api struct {
	Services []Service `json:"webServices"`
}

// These endpoints cannot/should not be generated
var skippedEndpoints = []string{
	"duplications", // numeric map keys cause parse errors
	"properties",   // unmarshall errors on already deprecated endpoint
	"favourites",   // deprecated in favour of favorites ;)
	"paging",       // non-existent, but there to prevent overwriting custom paging
}

// These fields don't need to be in each request struct
var skippedRequestFields = []string{
	"organization", // the organization is part of the client config
}

const input = "gen/webservices.json"
const packageName = "sonarcloud"

func main() {
	file, err := ioutil.ReadFile(input)
	guard(err)

	var api Api
	err = json.Unmarshal(file, &api)
	guard(err)

	wg := &sync.WaitGroup{}
	for _, service := range api.Services {
		s := service

		wg.Add(1)
		go func() {
			defer wg.Done()
			err := s.process(packageName)
			if err != nil {
				fmt.Printf("Error processing service at path %s: %+v\n", s.Path, err)
			}
		}()
	}

	wg.Wait()
}
