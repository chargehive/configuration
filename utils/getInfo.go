package utils

import (
	"errors"
	"github.com/chargehive/configuration/object"
)

type ConfigInfo struct {
	Kind        string
	Version     string
	Name        string
	ProjectId   string
	Disabled    bool
	DisplayName string
}

// Takes a json config, deserializes and returns summary config info
func GetInfo(rawJson []byte, version string) (ConfigInfo, error) {
	if version != "v1" {
		return ConfigInfo{}, errors.New("invalid config version")
	}
	def, err := object.FromJson(rawJson)
	if err != nil {
		return ConfigInfo{}, err
	}
	return ConfigInfo{
		Kind:        string(def.Kind),
		Version:     def.SpecVersion,
		Name:        def.MetaData.Name,
		ProjectId:   def.MetaData.ProjectID,
		Disabled:    def.MetaData.Disabled,
		DisplayName: def.MetaData.DisplayName,
	}, nil
}
