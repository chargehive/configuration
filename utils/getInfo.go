package utils

import (
	"encoding/json"
	"github.com/chargehive/configuration/object"
)

type ConfigInfo struct {
	Kind        object.Kind
	MetaData    object.MetaData
	SpecVersion string
}

// Takes a json config, deserializes and returns summary config info
func GetInfo(rawJson []byte) (ConfigInfo, error) {
	obj := ConfigInfo{}
	err := json.Unmarshal(rawJson, &obj)
	return obj, err
}
