/*
Copyright © 2022 Rich Rose <richardrose@google.com>

*/
package v2

import (
	"fmt"
)

func generateV2SchemaAssessment(v2 *SchemaV2) {

	if v2.Assessment != "" {
		fmt.Printf("assessment: %s\n", v2.Assessment)
	}
}
