/*
Copyright © 2022 Rich Rose <richardrose@google.com>

*/
package v2

import (
	"fmt"
)

// Name: generateV2Instruction.
// Description: Handle the instruction definition in the schema.
func generateV2SchemaInstruction(v2 *SchemaV2) {
	fmt.Printf("instruction:\n")

	if v2.Instruction.Type != "" {
		fmt.Printf("  type: %s\n", v2.Instruction.Type)
	} else {
		fmt.Printf("  type: html\n")
	}

	if v2.Instruction.URI != "" {
		fmt.Printf("  uri: %s\n", v2.Instruction.URI)
	} else {
		fmt.Printf("  uri: instructions/en.md\n")
	}
}
