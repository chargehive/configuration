package utils

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/connectorconfig"
	"github.com/chargehive/configuration/object"
	"github.com/chargehive/configuration/v1/connector"
)

// Clean will load a json file in, merge fields where possible, removes old fields no longer in struct
func Clean(input []byte, version string, pretty bool) (modified bool, output []byte, err error) {
	if version != "v1" {
		err = errors.New("version mismatch")
		return
	}

	// standardise the input for comparison later
	var inputStd interface{}
	err = json.Unmarshal(input, &inputStd)
	if err != nil {
		return
	}

	// convert to config
	def, err := object.FromJson(input)
	if err != nil {
		return
	}

	// Parse and re-serialize if is connector
	if def.Kind == "Connector" {
		c, ok := def.Spec.(*connector.Connector)
		if !ok {
			err = errors.New("spec is not a connector")
			return
		}

		var cred connectorconfig.Credentials
		cred, err = connectorconfig.GetCredentials(c)
		if err != nil {
			return
		}

		c.Configuration = cred
		def.Spec = c
	}

	// convert to new json
	output, _ = json.Marshal(def)
	if pretty {
		output, _ = json.MarshalIndent(def, "", "  ")
	}

	// standardise the output
	var outputStd interface{}
	err = json.Unmarshal(output, &outputStd)
	if err != nil {
		return
	}

	// had to double unmarshal because the ordering could be different
	// compare contents
	a, _ := json.Marshal(inputStd)
	b, _ := json.Marshal(outputStd)
	modified = !bytes.Equal(a, b)

	return
}
