package v2

import (
	"testing"
)

func TestGenerateV2GetAllowedLocations(t *testing.T) {
	locations := []string{"us-central", "us-east1", "us-west1"}

	// Create a sample Resources struct
	resources := EnvironmentResourceType{
		Type:             "Lab",
		ID:               "gsp001",
		Variant:          "gcpd",
		AllowedLocations: locations,
	}

	// Test the function with the sample struct
	getAllowedLocations(resources)
}
