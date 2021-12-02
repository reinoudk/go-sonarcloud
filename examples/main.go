package main

import (
	"fmt"
	"github.com/reinoudk/go-sonarcloud/sonarcloud"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/projects"
	"log"
	"os"
)

func main() {
	org, ok := os.LookupEnv("SONARCLOUD_ORG")
	if !ok {
		log.Fatalf("missing SONARCLOUD_ORG environment variable")
	}

	token, ok := os.LookupEnv("SONARCLOUD_TOKEN")
	if !ok {
		log.Fatalf("mising SONARCLOUD_TOKEN environment variable")
	}

	client := sonarcloud.NewClient(org, token, nil)
	req := projects.SearchRequest{}

	res, err := client.Projects.SearchAll(req)
	if err != nil {
		log.Fatalf("could not search projects: %+v", err)
	}

	fmt.Printf("%+v\n", res)
}
