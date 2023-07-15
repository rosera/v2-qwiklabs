package v2

import (
	"testing"
)

func TestGenerateV2SchemaHeader(t *testing.T) {
	// Create a sample SchemaV2 struct
	v2 := &SchemaV2{
		EntityType:    "Lab",
		SchemaVersion: 2,
		DefaultLocale: "en-US",
		Title:         "Sample Title",
		Description:   "Sample Description",
		Duration:      60,
		//    MaxDuration: 0,
		Credits: 2,
		Level:   "Intermediate",
	}

	// Test the function with the sample struct
	generateV2SchemaHeader(v2)
}
