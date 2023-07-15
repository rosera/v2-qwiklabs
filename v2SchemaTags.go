/*
Copyright © 2022 Rich Rose <richardrose@google.com>

*/
package v2

import (
	"fmt"
)

// Name: generateV2SchemaTags.
// Description: Handle the Tags section of the schema.
func generateV2SchemaTags(v2 *SchemaV2) {
	if len(v2.Tags) > 0 {
		fmt.Printf("tags:\n")
		for _, tag := range v2.Tags {
			fmt.Printf("  - %s\n", tag)
		}
	}
}
