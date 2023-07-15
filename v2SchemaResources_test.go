package v2

import (
	"testing"
)

func TestGenerateV2SchemaResources(t *testing.T) {

	instruction := InstructionType{Type: "html", URI: "instructions/en.md"}

	// Create a sample SchemaV2 struct
	v2 := &SchemaV2{
		Instruction: instruction,
	}

	// Test the function with the sample struct
	generateV2SchemaInstruction(v2)
}
