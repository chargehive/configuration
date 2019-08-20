package object

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/selector"
)

type Kind string

const KindNone Kind = ""

type Definition struct {
	Kind        Kind              `json:"kind" yaml:"kind"`
	MetaData    MetaData          `json:"metadata" yaml:"metadata"`
	SpecVersion string            `json:"specVersion,omitempty" yaml:"specVersion,omitempty"`
	Selector    selector.Selector `json:"selector,omitempty" yaml:"selector,omitempty"`
	Spec        interface{}       `json:"spec" yaml:"spec"`
}

func (d *Definition) Definition() *Definition { return d }
func (d *Definition) GetID() string           { return d.MetaData.Name }

func FromJson(jsonData []byte) (*Definition, error) {
	var raw json.RawMessage
	obj := &Definition{Spec: &raw}

	if err := json.Unmarshal(jsonData, obj); err != nil {
		return nil, err
	}

	spec := SpecFromJson(obj.Kind, obj.SpecVersion, raw)
	if spec == nil {
		return nil, errors.New("kind " + string(obj.Kind) + ", version " + obj.SpecVersion + " has not yet been implemented")
	}
	obj.Spec = spec
	return obj, nil
}

func SpecFromJson(kind Kind, version string, jsonData []byte) Specification {
	if handler, ok := getKindHandler(kind, version); ok {
		spec := handler()
		if err := json.Unmarshal(jsonData, spec); err != nil {
			return nil
		}
		return spec
	}
	return nil
}

func DefinitionFromSpec(specification Specification) *Definition {
	return &Definition{Kind: specification.GetKind(), MetaData: MetaData{}, Spec: specification, SpecVersion: specification.GetVersion()}
}

type DefinitionHolder interface {
	Definition() *Definition
}
