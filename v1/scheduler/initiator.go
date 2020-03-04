package scheduler

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

// KindInitiator is the identifier for an initiator config
const KindInitiator object.Kind = "Initiator"

// Initiator is a config that specifies actions for an initial transaction
type Initiator struct {
	// Type indicates the type of this initiator
	Type InitiatorType `json:"type" yaml:"type" validate:"required,oneof=auth renewal"`

	// InitialConnector is the selection method for the initial selector (only used on renewal)
	InitialConnector ConnectorSelector `json:"initialConnector" yaml:"initialConnector" validate:"required,oneof=none sticky-first sticky-last sticky-any sticky-verified config"`

	// AttemptConfig to be used when ConnectorSelectorConfig is set as InitialConnector
	AttemptConfig *AttemptConfig `json:"attemptConfig" yaml:"attemptConfig" validate:"required,dive"`
}

// GetKind returns the Initiator kind
func (Initiator) GetKind() object.Kind { return KindInitiator }

// GetVersion returns the Initiator version
func (Initiator) GetVersion() string { return "v1" }

// NewInitiatorDefinition creates a new InitiatorDefinition
func NewInitiatorDefinition(d *object.Definition) (*InitiatorDefinition, error) {
	if _, ok := d.Spec.(*Initiator); ok {
		return &InitiatorDefinition{def: d}, nil
	}
	return nil, errors.New("invalid initiator object")
}

// InitiatorDefinition is the Initiator config object definition
type InitiatorDefinition struct{ def *object.Definition }

// Definition returns the InitiatorDefinition structure
func (d *InitiatorDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the InitiatorDefinition
func (d *InitiatorDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the CascadePolicy contained within the InitiatorDefinition
func (d *InitiatorDefinition) Spec() *Initiator { return d.def.Spec.(*Initiator) }
