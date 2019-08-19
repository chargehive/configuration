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

func NewInstance(i *object.Instance) (*Instance, error) {
	if _, ok := i.Spec.(*Connector); ok {
		return &Instance{i: i}, nil
	}
	return nil, errors.New("invalid connector object")
}

type Instance struct{ i *object.Instance }

func (i *Instance) Instance() *object.Instance   { return i.i }
func (i *Instance) MarshalJSON() ([]byte, error) { return json.Marshal(i.i) }
func (i *Instance) Connector() *Connector {
	return i.i.Spec.(*Connector)
}
