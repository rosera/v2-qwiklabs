package v2

import (
	"io/ioutil"
	"testing"
)

func TestNonExistantFileV2Schema(t *testing.T) {
	// Call ProcessV2Schema with a non-existent file and check for the correct error
	v2 := &SchemaV2{}
	badFileName := "noexistent.yaml"
	_, err := v2.ReadV2Schema(&badFileName)
	if err == nil {
		t.Errorf("ProcessV2Schema did not return an error for a non-existent file")
	}
}

func TestTempFileV2Schema(t *testing.T) {
	// Create a temporary file to test with
	fileData := "key: value"
	tmpfile, err := ioutil.TempFile("", "testfile.*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer tmpfile.Close()
	defer ioutil.TempDir("", "testfile.*.yaml")

	// Write test data to temporary file
	if _, err := tmpfile.Write([]byte(fileData)); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	// Call ProcessV2Schema and check the returned result
	v2 := &SchemaV2{}
	fileName := tmpfile.Name()
	result, err := v2.ReadV2Schema(&fileName)
	if err != nil {
		t.Fatalf("ProcessV2Schema returned an error: %v", err)
	}
	if result != v2 {
		t.Errorf("ProcessV2Schema returned incorrect result: expected %v, got %v", v2, result)
	}
}

func TestInvalidV2Schema(t *testing.T) {
	// Create a temporary file to test with
	fileData := "key: value"
	tmpfile, err := ioutil.TempFile("", "testfile.*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer tmpfile.Close()
	defer ioutil.TempDir("", "testfile.*.yaml")

	// Write test data to temporary file
	if _, err := tmpfile.Write([]byte(fileData)); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	// Call ProcessV2Schema and check the returned result
	v2 := &SchemaV2{}

	// Call ProcessV2Schema with an invalid YAML file and check for the correct error
	invalidYaml := "key:\n- value1\n- value2"
	if _, err := tmpfile.Write([]byte(invalidYaml)); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	invalidYamlFile := tmpfile.Name()
	_, err = v2.ReadV2Schema(&invalidYamlFile)
	if err == nil {
		t.Errorf("ProcessV2Schema did not return an error for an invalid YAML file")
	}
}
