package scheduler

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/object"
)

// KindConnectorScheduler is the identifier for an Connector scheduler config
const KindConnectorScheduler object.Kind = "SchedulerConnector"

// Connector is a schedule that can determine the next run time based on a connector
type Connector struct {
	ConnectorID string `json:"connectorID" yaml:"connectorID" validate:"required"`
}

// GetKind returns the Connector kind
func (Connector) GetKind() object.Kind { return KindConnectorScheduler }

// GetVersion returns the Connector version
func (Connector) GetVersion() string { return "v1" }

// NewConnectorDefinition returns a new ConnectorDefinition
func NewConnectorDefinition(d *object.Definition) (*ConnectorDefinition, error) {
	if _, ok := d.Spec.(*Connector); ok {
		return &ConnectorDefinition{def: d}, nil
	}
	return nil, errors.New("invalid connector scheduler object")
}

// ConnectorDefinition is the Connector object definition
type ConnectorDefinition struct{ def *object.Definition }

// Definition returns the ConnectorDefinition structure
func (d *ConnectorDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the ConnectorDefinition
func (d *ConnectorDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the Connector structure contained within the ConnectorDefinition
func (d *ConnectorDefinition) Spec() *Connector { return d.def.Spec.(*Connector) }
