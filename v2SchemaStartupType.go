/*
Copyright © 2022 Rich Rose <richardrose@google.com>

*/
package v2

import (
	"fmt"
)

func getStartupType(resource EnvironmentResourceType) {
	if resource.StartupScript.Type != "" {
		fmt.Printf("    startup_script:\n")
		fmt.Printf("      type: %s\n", resource.StartupScript.Type)
		fmt.Printf("      path: %s\n", resource.StartupScript.Path)

		// TODO: Add Custom Properties
		if resource.StartupScript.CustomProperties != nil {
			fmt.Printf("      custom_properties:\n")
			for _, customProperties := range resource.StartupScript.CustomProperties {
				fmt.Printf("      - key: %s\n", customProperties.Key)
				if customProperties.Value != "" {
					fmt.Printf("        value:  %s\n", customProperties.Value)
				} else if customProperties.Reference != "" {
					fmt.Printf("        reference: %s\n", customProperties.Reference)
				}
			}
		}
	}
}

func setStartupType(resource EnvironmentResourceType, actionToken string) {
	if resource.StartupScript.Type != "" {
		fmt.Printf("    startup_script:\n")
		fmt.Printf("      type: %s\n", resource.StartupScript.Type)
		fmt.Printf("      path: %s\n", resource.StartupScript.Path)
	} else {
		fmt.Printf("    startup_script:\n")
		fmt.Printf("      type: %s\n", actionToken)
		if actionToken == "qwiklabs" {
			fmt.Printf("      path: tf\n")
		} else {
			fmt.Printf("      path: dm\n")
		}
	}
}
