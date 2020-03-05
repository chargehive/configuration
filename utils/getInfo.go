package utils

import (
	"errors"
	"github.com/chargehive/configuration/object"
	"strconv"
)

// Takes a json config, deserializes and returns map of config info
func GetInfo(rawJson []byte, version string) (map[string]string, error) {
	if version != "v1" {
		return nil, errors.New("invalid config version")
	}
	def, err := object.FromJson(rawJson)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"kind":        string(def.Kind),
		"version":     def.SpecVersion,
		"name":        def.MetaData.Name,
		"projectId":   def.MetaData.ProjectID,
		"disabled":    strconv.FormatBool(def.MetaData.Disabled),
		"displayName": def.MetaData.DisplayName,
	}, nil
}
