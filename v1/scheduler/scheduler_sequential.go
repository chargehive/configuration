package scheduler

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

type Sequential struct {
	DefaultSchedule Schedule
	Schedules       map[int]Schedule // attempt number > Schedule
}

const KindSequentialScheduler object.Kind = "SchedulerSequential"

func (Sequential) GetKind() object.Kind { return KindSequentialScheduler }
func (Sequential) GetVersion() string   { return "v1" }

func NewSequentialDefinition(d *object.Definition) (*SequentialDefinition, error) {
	if _, ok := d.Spec.(*Sequential); ok {
		return &SequentialDefinition{def: d}, nil
	}
	return nil, errors.New("invalid sequential scheduler object")
}

type SequentialDefinition struct{ def *object.Definition }

func (d *SequentialDefinition) Definition() *object.Definition { return d.def }
func (d *SequentialDefinition) MarshalJSON() ([]byte, error)   { return json.Marshal(d.def) }
func (d *SequentialDefinition) Spec() *Sequential              { return d.def.Spec.(*Sequential) }
