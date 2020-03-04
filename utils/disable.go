package utils

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

// Takes a json config, deserializes and sets the disable flag in metadata
func Disable(setDisabled bool, rawJson []byte, version string) ([]byte, error) {
	if version != "v1" {
		return nil, errors.New("invalid config version")
	}
	def, err := object.FromJson(rawJson)
	if err != nil {
		return nil, err
	}
	def.MetaData.Disabled = setDisabled
	return json.Marshal(def)
}
