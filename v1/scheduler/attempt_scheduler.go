package scheduler

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

type AttemptScheduler struct {
	DefaultSchedule Schedule
	Schedules       map[int]Schedule // attempt number > Schedule
}

const KindAttemptScheduler object.Kind = "AttemptScheduler"

func (i AttemptScheduler) GetKind() object.Kind { return KindAttemptScheduler }
func (i AttemptScheduler) GetVersion() string   { return "v1" }

func NewAttemptSchedulerInstance(i *object.Instance) (*AttemptSchedulerInstance, error) {
	if _, ok := i.Spec.(*AttemptScheduler); ok {
		return &AttemptSchedulerInstance{i: i}, nil
	}
	return nil, errors.New("invalid attempt scheduler object")
}

type AttemptSchedulerInstance struct{ i *object.Instance }

func (i *AttemptSchedulerInstance) MarshalJSON() ([]byte, error) { return json.Marshal(i.i) }
func (i *AttemptSchedulerInstance) AttemptScheduler() *AttemptScheduler {
	return i.i.Spec.(*AttemptScheduler)
}
