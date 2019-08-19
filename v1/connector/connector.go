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

func NewConnectorInstance(i *object.Instance) (*ConnectorInstance, error) {
	if _, ok := i.Spec.(*Connector); ok {
		return &ConnectorInstance{i: i}, nil
	}
	return nil, errors.New("invalid connector object")
}

type ConnectorInstance struct{ i *object.Instance }

func (i *ConnectorInstance) MarshalJSON() ([]byte, error) { return json.Marshal(i.i) }
func (i *ConnectorInstance) Connector() *Connector {
	return i.i.Spec.(*Connector)
}
