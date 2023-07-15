/*
Copyright © 2022 Rich Rose <richardrose@google.com>

*/
package v2

import (
	"fmt"
)

func generateV2SchemaVisibleOutputs(v2 *SchemaV2) {
	fmt.Printf("  student_visible_outputs:\n")

	// Print Student Visible Outputs list
	for _, visibleOutput := range v2.Environment.StudentVisibleOutputs {
		fmt.Printf("  - label: %s\n", visibleOutput.Label)
		fmt.Printf("    reference: %s\n", visibleOutput.Reference)
	}
}
