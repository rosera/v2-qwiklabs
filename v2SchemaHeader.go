/*
Copyright © 2022 Rich Rose <richardrose@google.com>

*/
package v2

import (
	"fmt"
)

func generateV2SchemaHeader(v2 *SchemaV2) {

	if v2.EntityType != "" {
		fmt.Printf("entity_type: %s\n", v2.EntityType)
	} else {
		fmt.Printf("entity_type: Lab\n")
	}

	fmt.Printf("schema_version: %d\n", v2.SchemaVersion)
	fmt.Printf("default_locale: %s\n", v2.DefaultLocale)
	fmt.Printf("title: '%s'\n", v2.Title)
	fmt.Printf("description: '%s'\n", v2.Description)
	fmt.Printf("duration: %d\n", v2.Duration)

	if v2.MaxDuration != 0 {
		fmt.Printf("max_duration: %d\n", v2.MaxDuration)
	} else {
		fmt.Printf("max_duration: %d\n", v2.Duration)
	}

	if v2.Credits != 0 {
		fmt.Printf("credits: %d\n", v2.Credits)
	} else {
		fmt.Printf("credits: 1\n")
	}

	fmt.Printf("level: %s\n", v2.Level)
}
