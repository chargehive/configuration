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

func (d *Definition) GetID() string { return d.MetaData.Name }

func FromJson(jsonData []byte) (*Definition, error) {
	var raw json.RawMessage
	obj := &Definition{Spec: &raw}

	if err := json.Unmarshal(jsonData, obj); err != nil {
		return nil, err
	}

	if handler, ok := getKindHandler(obj.Kind, obj.SpecVersion); ok {
		spec := handler()
		if err := json.Unmarshal(raw, spec); err != nil {
			return obj, err
		}
		obj.Spec = spec
		return obj, nil
	}

	return nil, errors.New("kind " + string(obj.Kind) + ", version " + obj.SpecVersion + " has not yet been implemented")
}

func DefinitionFromSpec(specification Specification) *Definition {
	return &Definition{Kind: specification.GetKind(), MetaData: MetaData{}, Spec: specification, SpecVersion: specification.GetVersion()}
}

type DefinitionHolder interface {
	Definition() *Definition
	MarshalJSON() ([]byte, error)
}
