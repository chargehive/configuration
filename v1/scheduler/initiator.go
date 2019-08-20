package scheduler

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

const KindInitiator object.Kind = "Initiator"

type Initiator struct {
	Type             InitiatorType
	InitialConnector ConnectorSelector
	// AttemptConfig to be used when using ConnectorSelectorConfig
	AttemptConfig *AttemptConfig
}

func (Initiator) GetKind() object.Kind { return KindInitiator }
func (Initiator) GetVersion() string   { return "v1" }

func NewInitiatorDefinition(d *object.Definition) (*InitiatorDefinition, error) {
	if _, ok := d.Spec.(*Initiator); ok {
		return &InitiatorDefinition{def: d}, nil
	}
	return nil, errors.New("invalid initiator object")
}

type InitiatorDefinition struct{ def *object.Definition }

func (d *InitiatorDefinition) Definition() *object.Definition { return d.def }
func (d *InitiatorDefinition) MarshalJSON() ([]byte, error)   { return json.Marshal(d.def) }
func (d *InitiatorDefinition) Spec() *Initiator               { return d.def.Spec.(*Initiator) }
