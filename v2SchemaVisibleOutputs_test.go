package v2

import (
	"testing"
)

func TestGenerateV2VisibleOutputs(t *testing.T) {
	customProperty := CustomPropertiesType{Key: "test", Value: "test", Reference: "test"}
	customProperties := []CustomPropertiesType{customProperty}

	startupScript := StartupScriptType{Type: "qwiklabs", Path: "tf", CustomProperties: customProperties}

	// Create a sample Resources struct
	resource := EnvironmentResourceType{
		Type:          "Lab",
		ID:            "gsp001",
		Variant:       "gcpd",
		StartupScript: startupScript,
	}
	resources := []EnvironmentResourceType{resource}

	projectOutput := StudentVisibleOutputsType{Label: "Project ID", Reference: "project_0.project_id"}
	studentVisibleOutput := []StudentVisibleOutputsType{projectOutput}

	environment := EnvironmentType{Resources: resources, StudentVisibleOutputs: studentVisibleOutput}

	// Create a sample SchemaV2 struct
	v2 := &SchemaV2{
		EntityType:    "Lab",
		SchemaVersion: 2,
		DefaultLocale: "en-US",
		Title:         "Sample Title",
		Description:   "Sample Description",
		Duration:      60,
		//    MaxDuration: 0,
		Credits:     2,
		Level:       "Intermediate",
		Environment: environment,
	}

	// Test the function with the sample struct
	generateV2SchemaVisibleOutputs(v2)
}
