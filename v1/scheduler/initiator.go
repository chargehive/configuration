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

func NewInitiatorInstance(i *object.Instance) (*InitiatorInstance, error) {
	if _, ok := i.Spec.(*Initiator); ok {
		return &InitiatorInstance{i: i}, nil
	}
	return nil, errors.New("invalid initiator object")
}

type InitiatorInstance struct{ i *object.Instance }

func (i *InitiatorInstance) MarshalJSON() ([]byte, error) { return json.Marshal(i.i) }
func (i *InitiatorInstance) Initiator() *Initiator        { return i.i.Spec.(*Initiator) }
