/*
Copyright © 2022 Rich Rose <richardrose@google.com>

*/
package v2

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// Name: processV2Schema
// Description: Process a compatible schema
func (v2 *SchemaV2) ReadV2Schema(filename *string) (*SchemaV2, error) {
	file, err := ioutil.ReadFile(*filename)

	if err != nil {
		// error handling
		//      fmt.Println(err.Error())
		return v2, err
	}

	err = yaml.Unmarshal([]byte(file), &v2)
	if err != nil {
		// fmt.Printf("Unmarshal in file %q: %v\n", *filename, err)
		// log.Fatalf("Unmarshal in file %q: %v", *filename, err)
		return v2, err
	}

	return v2, err
}

func (v2 *SchemaV2) ProcessV2Schema(actionToken string) *SchemaV2 {
	generateV2SchemaHeader(v2)

	generateV2SchemaTags(v2)

	generateV2SchemaResources(v2)

	generateV2SchemaInstruction(v2)

	generateV2SchemaEnvironment(v2, actionToken)

	generateV2SchemaVisibleOutputs(v2)

	generateV2SchemaAssessment(v2)

	return v2
}

func (v2 *SchemaV2) WriteV2Schema(filename string, actionToken string) *SchemaV2 {

  // TODO: Write v2 Schema to a file

	return v2
}