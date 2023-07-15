/*
Copyright © 2022 Rich Rose <richardrose@google.com>

*/
package v2

import (
	"fmt"
)

// Name: generateV2SchemaEnviroment.
// Description: Handle the Tags section of the schema.
func generateV2SchemaEnvironment(v2 *SchemaV2, actionToken string) {
	fmt.Printf("environment:\n")
	fmt.Printf("  resources:\n")

	for _, resource := range v2.Environment.Resources {
		fmt.Printf("  - type: %s\n", resource.Type)
		fmt.Printf("    id: %s\n", resource.ID)

		// GCP User settings
		if resource.Type == "gcp_user" {
			fmt.Printf("    permissions:\n")

			for _, permission := range resource.Permissions {
				fmt.Printf("     - project: %s\n", permission.Project)
				fmt.Printf("       roles:\n")

				for _, iam := range permission.Roles {
					fmt.Printf("       - %s\n", iam)
				}
			}
		}

		// GCP Project settings
		if resource.Type == "gcp_project" {

			// VARIANT
			if actionToken == "gcpd" {
				fmt.Printf("    variant: %s\n", actionToken)
			} else if actionToken == "gcpfree" {
				fmt.Printf("    variant: %s\n", actionToken)
			} else if actionToken == "gcp_low_extra" {
				fmt.Printf("    variant: %s\n", actionToken)
			} else if actionToken == "gcp_very_low_base" {
				fmt.Printf("    variant: %s\n", actionToken)
			} else if actionToken != "variant" {
				fmt.Printf("    variant: %s\n", resource.Variant)
			} else if resource.Variant != "" {
				fmt.Printf("    variant: %s\n", resource.Variant)
			}

			fmt.Printf("    ssh_key_user: %s\n", resource.SshKeyUser)

			// STARTUP SCRIPT
			if actionToken == "qwiklabs" {
				setStartupType(resource, actionToken)
			} else if actionToken == "deployment" {
				setStartupType(resource, actionToken)
			} else if actionToken != "startup_script" {
				getStartupType(resource)
			}

			// TODO: Add CleanupScript

			// ALLOWED_LOCATIONS
			if actionToken == "general" {
				setAllowedLocations(resource, actionToken)
			} else if actionToken == "appengine" {
				setAllowedLocations(resource, actionToken)
			} else if actionToken == "compute" {
				setAllowedLocations(resource, actionToken)
			} else if actionToken == "bigdata" {
				setAllowedLocations(resource, actionToken)
			} else if actionToken == "storage" {
				setAllowedLocations(resource, actionToken)
			} else if actionToken == "vertex" {
				setAllowedLocations(resource, actionToken)
			} else if actionToken != "allowed_location" {
				getAllowedLocations(resource)
			}

		}
	}
}
