package scheduler

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

// OnDemand is a schedule that is run on demand
type OnDemand struct {
	Schedule Schedule
}

// KindOnDemandScheduler is the identifier for an OnDemand scheduler config
const KindOnDemandScheduler object.Kind = "SchedulerOnDemand"

// GetKind returns the OnDemand kind
func (OnDemand) GetKind() object.Kind { return KindOnDemandScheduler }

// GetVersion returns the OnDemand version
func (OnDemand) GetVersion() string { return "v1" }

// NewOnDemandDefinition returns a new OnDemandDefinition
func NewOnDemandDefinition(d *object.Definition) (*OnDemandDefinition, error) {
	if _, ok := d.Spec.(*OnDemand); ok {
		return &OnDemandDefinition{def: d}, nil
	}
	return nil, errors.New("invalid on demand scheduler object")
}

// OnDemandDefinition is the OnDemand object definition
type OnDemandDefinition struct{ def *object.Definition }

// Definition returns the OnDemandDefinition structure
func (d *OnDemandDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the OnDemandDefinition
func (d *OnDemandDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the OnDemand structure contained within the OnDemandDefinition
func (d *OnDemandDefinition) Spec() *OnDemand { return d.def.Spec.(*OnDemand) }
