package scheduler

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/object"
)

// KindOnTriggerScheduler is the identifier for an OnTrigger scheduler config
const KindOnTriggerScheduler object.Kind = "SchedulerOnTrigger"

// OnTrigger is a schedule that is run on demand
type OnTrigger struct {
	Schedule Schedule `json:"schedule" yaml:"schedule" validate:"required,dive"`
}

// GetKind returns the OnTrigger kind
func (OnTrigger) GetKind() object.Kind { return KindOnTriggerScheduler }

// GetVersion returns the OnTrigger version
func (OnTrigger) GetVersion() string { return "v1" }

// NewOnTriggerDefinition returns a new OnTriggerDefinition
func NewOnTriggerDefinition(d *object.Definition) (*OnTriggerDefinition, error) {
	if _, ok := d.Spec.(*OnTrigger); ok {
		return &OnTriggerDefinition{def: d}, nil
	}
	return nil, errors.New("invalid on demand scheduler object")
}

// OnTriggerDefinition is the OnTrigger object definition
type OnTriggerDefinition struct{ def *object.Definition }

// Definition returns the OnTriggerDefinition structure
func (d *OnTriggerDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the OnTriggerDefinition
func (d *OnTriggerDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the OnTrigger structure contained within the OnTriggerDefinition
func (d *OnTriggerDefinition) Spec() *OnTrigger { return d.def.Spec.(*OnTrigger) }
