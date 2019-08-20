package connector

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

const KindConnector object.Kind = "Connector"

type Connector struct {
	Library       string
	Configuration []byte
}

func (Connector) GetKind() object.Kind { return KindConnector }
func (Connector) GetVersion() string   { return "v1" }

func NewDefinition(d *object.Definition) (*Definition, error) {
	if _, ok := d.Spec.(*Connector); ok {
		return &Definition{def: d}, nil
	}
	return nil, errors.New("invalid connector object")
}

type Definition struct{ def *object.Definition }

func (d *Definition) Definition() *object.Definition { return d.def }
func (d *Definition) MarshalJSON() ([]byte, error)   { return json.Marshal(d.def) }
func (d *Definition) Spec() *Connector               { return d.def.Spec.(*Connector) }
