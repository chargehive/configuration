package object

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/selector"
)

type Kind string

const KindNone Kind = ""

type Definition struct {
	Kind        Kind              `json:"Kind" yaml:"Kind"`
	MetaData    MetaData          `json:"metadata" yaml:"metadata" validate:"dive"`
	SpecVersion string            `json:"specVersion" yaml:"specVersion"`
	Selector    selector.Selector `json:"selector" yaml:"selector" validate:"dive"`
	Spec        interface{}       `json:"spec" yaml:"spec" validate:"dive"`
}

func (d *Definition) Definition() *Definition { return d }
func (d *Definition) GetID() string           { return d.MetaData.Name }

func FromJson(jsonData []byte) (*Definition, error) {
	var raw json.RawMessage
	obj := &Definition{Spec: &raw}

	if err := json.Unmarshal(jsonData, obj); err != nil {
		return nil, err
	}

	spec, err := SpecFromJson(obj.Kind, obj.SpecVersion, raw)
	if err != nil {
		return nil, errors.New("invalid JSON format in configuration")
	}
	if spec == nil {
		return nil, errors.New("Kind " + string(obj.Kind) + ", Version " + obj.SpecVersion + " has not yet been implemented")
	}
	obj.Spec = spec
	return obj, nil
}

func SpecFromJson(kind Kind, version string, jsonData []byte) (Specification, error) {
	if handler, ok := getKindHandlerFunc(kind, version); ok {
		spec := handler()
		if err := json.Unmarshal(jsonData, spec); err != nil {
			return nil, err
		}
		return spec, nil
	}
	return nil, nil
}

func DefinitionFromSpec(specification Specification) *Definition {
	return &Definition{Kind: specification.GetKind(), MetaData: MetaData{}, Spec: specification, SpecVersion: specification.GetVersion()}
}

type DefinitionHolder interface {
	Definition() *Definition
}
