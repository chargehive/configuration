package object

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/chargehive/configuration/selector"
)

type Kind string

const KindNone Kind = ""

type Definition struct {
	Kind        Kind              `json:"kind" yaml:"kind"`
	MetaData    MetaData          `json:"metadata" yaml:"metadata" validate:"dive"`
	SpecVersion string            `json:"specVersion" yaml:"specVersion" validate:"oneof=v1"`
	Selector    selector.Selector `json:"selector" yaml:"selector" validate:"dive"`
	Spec        interface{}       `json:"spec" yaml:"spec" validate:"dive"`
}

func (d *Definition) Definition() *Definition { return d }
func (d *Definition) GetID() string           { return d.MetaData.Name }

func FromJson(jsonData []byte) (*Definition, error) {
	return jsonToObj(jsonData, false)
}

func FromJsonStrict(jsonData []byte) (*Definition, error) {
	return jsonToObj(jsonData, true)
}

func SpecFromJson(kind Kind, version string, jsonData []byte) (Specification, error) {
	return jsonSpecToObj(kind, version, jsonData, false)
}

func SpecFromJsonStrict(kind Kind, version string, jsonData []byte) (Specification, error) {
	return jsonSpecToObj(kind, version, jsonData, true)
}

func jsonToObj(jsonData []byte, strict bool) (*Definition, error) {
	var raw json.RawMessage
	obj := &Definition{Spec: &raw}

	reader := strings.NewReader(string(jsonData))
	dec := json.NewDecoder(reader)
	if strict {
		dec.DisallowUnknownFields()
	}
	if err := dec.Decode(obj); err != nil {
		return nil, err
	}

	spec, err := jsonSpecToObj(obj.Kind, obj.SpecVersion, raw, strict)
	if err != nil {
		return nil, err
	}
	if spec == nil {
		return nil, fmt.Errorf("kind:`%v`, version:`%v` has not been implemented", obj.Kind, obj.SpecVersion)
	}
	obj.Spec = spec
	return obj, nil
}

func jsonSpecToObj(kind Kind, version string, jsonData []byte, strict bool) (Specification, error) {
	if handler, ok := getKindHandlerFunc(kind, version); ok {
		spec := handler()
		reader := strings.NewReader(string(jsonData))
		dec := json.NewDecoder(reader)
		if strict {
			dec.DisallowUnknownFields()
		}
		if err := dec.Decode(spec); err != nil {
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
