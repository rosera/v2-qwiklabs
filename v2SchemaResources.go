/*
Copyright © 2022 Rich Rose <richardrose@google.com>

*/
package v2

import (
	"fmt"
)

// Name: generateV2Resources.
// Description: Handle the Resources of the schema.
// TODO: Add this functionality
func generateV2SchemaResources(v2 *SchemaV2) {

	if len(v2.Resources) > 0 {
		fmt.Printf("resources:\n")
		for _, resource := range v2.Resources {
			fmt.Printf("  - id: %s\n", resource.ID)
			fmt.Printf("    title: '%s'\n", resource.Title)
			fmt.Printf("    description: '%s'\n", resource.Description)
			fmt.Printf("    type: %s\n", resource.Type)
			fmt.Printf("    uri: %s\n", resource.URI)
		}
	} else {
		fmt.Printf("resources: []\n")
	}
}
