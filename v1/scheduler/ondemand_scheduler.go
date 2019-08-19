package scheduler

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

type OnDemandScheduler struct {
	Schedule Schedule
}

const KindOnDemandScheduler object.Kind = "OnDemandScheduler"

func (i OnDemandScheduler) GetKind() object.Kind { return KindOnDemandScheduler }
func (i OnDemandScheduler) GetVersion() string   { return "v1" }

func NewOnDemandSchedulerInstance(i *object.Instance) (*OnDemandSchedulerInstance, error) {
	if _, ok := i.Spec.(*OnDemandScheduler); ok {
		return &OnDemandSchedulerInstance{i: i}, nil
	}
	return nil, errors.New("invalid on demand scheduler object")
}

type OnDemandSchedulerInstance struct{ i *object.Instance }

func (i *OnDemandSchedulerInstance) MarshalJSON() ([]byte, error) { return json.Marshal(i.i) }
func (i *OnDemandSchedulerInstance) OnDemandScheduler() *OnDemandScheduler {
	return i.i.Spec.(*OnDemandScheduler)
}
