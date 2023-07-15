package v2

import (
	"testing"
)

func TestGenerateV2GetStartupType(t *testing.T) {
	customProperty := CustomPropertiesType{Key: "test", Value: "test", Reference: "test"}
	customProperties := []CustomPropertiesType{customProperty}

	startupScript := StartupScriptType{Type: "qwiklabs", Path: "tf", CustomProperties: customProperties}

	// Create a sample Resources struct
	resources := EnvironmentResourceType{
		Type:          "Lab",
		ID:            "gsp001",
		Variant:       "gcpd",
		StartupScript: startupScript,
	}

	// Test the function with the sample struct
	getStartupType(resources)
}
