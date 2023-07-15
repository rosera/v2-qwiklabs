package v2

import (
	"testing"
)

func TestGenerateV2SchemaEnvironment(t *testing.T) {

	project := EnvironmentResourceType{Type: "gcp_project", ID: "gsp001", Variant: "gcpd"}
	user := EnvironmentResourceType{Type: "gcp_user", ID: "gsp001", Variant: "gcpd"}

	projectId := StudentVisibleOutputsType{Label: "Google Cloud project", Reference: "project_0.project_id"}
	projectPassword := StudentVisibleOutputsType{Label: "Google Cloud password", Reference: "project_0.password"}

	resourceTypeList := []EnvironmentResourceType{project, user}
	studentVisibleOutputs := []StudentVisibleOutputsType{projectId, projectPassword}

	environment := EnvironmentType{resourceTypeList, studentVisibleOutputs}

	v2 := &SchemaV2{
		EntityType:    "Lab",
		SchemaVersion: 2,
		DefaultLocale: "en-US",
		Title:         "Sample Title",
		Description:   "Sample Description",
		Duration:      60,
		Credits:       2,
		Level:         "Intermediate",
		Environment:   environment,
	}

	// Test the function with the sample struct
	generateV2SchemaEnvironment(v2, "")
}
