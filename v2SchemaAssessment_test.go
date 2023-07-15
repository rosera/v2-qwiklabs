package v2

import (
	"testing"
)

func TestGenerateV2SchemaAssessment(t *testing.T) {
	// Create a sample SchemaV2 struct
	v2 := &SchemaV2{
		Assessment: "assessment.yaml",
	}

	// Test the function with the sample struct
	generateV2SchemaAssessment(v2)
}
