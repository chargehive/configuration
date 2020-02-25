package scheduler

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

// KindSequentialScheduler is the identifier for an KindSequentialScheduler scheduler config
const KindSequentialScheduler object.Kind = "SchedulerSequential"

// Sequential scheduler is a schedule that is ran based on factors such as attempt number
type Sequential struct {
	// Schedules to use based on attempt number map[attempt number > Schedule]
	Schedules map[int]Schedule `json:"schedules" yaml:"schedules"`
}

// GetKind returns the Sequential kind
func (Sequential) GetKind() object.Kind { return KindSequentialScheduler }

// GetVersion returns the Sequential version
func (Sequential) GetVersion() string { return "v1" }

// NewSequentialDefinition creates a new SequentialDefinition
func NewSequentialDefinition(d *object.Definition) (*SequentialDefinition, error) {
	if _, ok := d.Spec.(*Sequential); ok {
		return &SequentialDefinition{def: d}, nil
	}
	return nil, errors.New("invalid sequential scheduler object")
}

// SequentialDefinition is the Sequential object definition
type SequentialDefinition struct{ def *object.Definition }

// Definition returns the SequentialDefinition structure
func (d *SequentialDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the SequentialDefinition
func (d *SequentialDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the Sequential structure contained within the SequentialDefinition
func (d *SequentialDefinition) Spec() *Sequential { return d.def.Spec.(*Sequential) }
