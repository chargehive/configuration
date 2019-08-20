package scheduler

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

type OnDemand struct {
	Schedule Schedule
}

const KindOnDemandScheduler object.Kind = "SchedulerOnDemand"

func (OnDemand) GetKind() object.Kind { return KindOnDemandScheduler }
func (OnDemand) GetVersion() string   { return "v1" }

func NewOnDemandDefinition(d *object.Definition) (*OnDemandDefinition, error) {
	if _, ok := d.Spec.(*OnDemand); ok {
		return &OnDemandDefinition{def: d}, nil
	}
	return nil, errors.New("invalid on demand scheduler object")
}

type OnDemandDefinition struct{ def *object.Definition }

func (d *OnDemandDefinition) Definition() *object.Definition { return d.def }
func (d *OnDemandDefinition) MarshalJSON() ([]byte, error)   { return json.Marshal(d.def) }
func (d *OnDemandDefinition) Soec() *OnDemand                { return d.def.Spec.(*OnDemand) }
