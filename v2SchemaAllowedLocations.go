/*
Copyright © 2022 Rich Rose <richardrose@google.com>

*/
package v2

import (
	"fmt"
)

func getAllowedLocations(resource EnvironmentResourceType) {
	if resource.AllowedLocations != nil {
		fmt.Printf("    allowed_locations:\n")
		for _, allowedLocations := range resource.AllowedLocations {
			fmt.Printf("      - %s\n", allowedLocations)
		}
	}
}

func setAllowedLocations(resource EnvironmentResourceType, actionToken string) {
	if resource.AllowedLocations != nil {
		fmt.Printf("    allowed_locations:\n")
		for _, allowedLocations := range resource.AllowedLocations {
			fmt.Printf("      - %s\n", allowedLocations)
		}
	} else if actionToken == "general" {
		fmt.Printf("    allowed_locations:\n")
		fmt.Printf("      - us-central1\n")
		fmt.Printf("      - us-west1\n")
		fmt.Printf("      - us-west2\n")
		fmt.Printf("      - us-west3\n")
		fmt.Printf("      - us-west4\n")
		fmt.Printf("      - us-east1\n")
		fmt.Printf("      - us-east4\n")
		fmt.Printf("      - us-east5\n")
	} else if actionToken == "appengine" {
		fmt.Printf("    allowed_locations:\n")
		fmt.Printf("      - us-central1\n")
		fmt.Printf("      - us-west1\n")
		fmt.Printf("      - us-west2\n")
		fmt.Printf("      - us-west3\n")
		fmt.Printf("      - us-west4\n")
		fmt.Printf("      - us-east1\n")
		fmt.Printf("      - us-east4\n")
	} else if actionToken == "compute" {
		fmt.Printf("    allowed_locations:\n")
		fmt.Printf("      - us-central1\n")
		fmt.Printf("      - us-west1\n")
		fmt.Printf("      - us-west3\n")
		fmt.Printf("      - us-west4\n")
		fmt.Printf("      - us-east1\n")
		fmt.Printf("      - us-east4\n")
		fmt.Printf("      - us-east5\n")
	} else if actionToken == "bigdata" {
		fmt.Printf("    allowed_locations:\n")
		fmt.Printf("      - us-central1\n")
		fmt.Printf("      - us-west1\n")
		fmt.Printf("      - us-west3\n")
		fmt.Printf("      - us-west4\n")
		fmt.Printf("      - us-east1\n")
		fmt.Printf("      - us-east4\n")
		fmt.Printf("      - us-east5\n")
	} else if actionToken == "storage" {
		fmt.Printf("    allowed_locations:\n")
		fmt.Printf("      - us-central1\n")
		fmt.Printf("      - us-west1\n")
		fmt.Printf("      - us-west3\n")
		fmt.Printf("      - us-west4\n")
		fmt.Printf("      - us-east1\n")
		fmt.Printf("      - us-east4\n")
		fmt.Printf("      - us-east5\n")
	} else if actionToken == "vertex" {
		fmt.Printf("    allowed_locations:\n")
		fmt.Printf("      - us-central1-a\n")
		fmt.Printf("      - us-central1-b\n")
		fmt.Printf("      - us-central1-c\n")
		fmt.Printf("      - us-west4-a\n")
		fmt.Printf("      - us-west4-b\n")
		fmt.Printf("      - us-west4-c\n")
		fmt.Printf("      - us-west2-a\n")
		fmt.Printf("      - us-west2-b\n")
		fmt.Printf("      - us-west2-c\n")
	} else {
		fmt.Printf("    allowed_locations:\n")
		fmt.Printf("      - us-central1\n")
		fmt.Printf("      - us-west1\n")
		fmt.Printf("      - us-west3\n")
		fmt.Printf("      - us-west4\n")
		fmt.Printf("      - us-east1\n")
		fmt.Printf("      - us-east4\n")
		fmt.Printf("      - us-east5\n")
	}
}
