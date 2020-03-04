package scheduler

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
	"time"
)

// KindRefundScheduler is the identifier for an KindRefundScheduler scheduler config
const KindRefundScheduler object.Kind = "SchedulerRefund"

// Refund scheduler is a schedule that defines when refunds should be attempted
type Refund struct {
	// Schedules to use based on attempt number map[attempt number > Schedule]
	Schedules map[int]ScheduleRefund `json:"schedules" yaml:"schedules" validate:"min=1,dive"`
}

// GetKind returns the schedule kind
func (Refund) GetKind() object.Kind { return KindRefundScheduler }

// GetVersion returns the schedule version
func (Refund) GetVersion() string { return "v1" }

// NewRefundDefinition creates a new RefundDefinition
func NewRefundDefinition(d *object.Definition) (*RefundDefinition, error) {
	if _, ok := d.Spec.(*Refund); ok {
		return &RefundDefinition{def: d}, nil
	}
	return nil, errors.New("invalid refund scheduler object")
}

// RefundDefinition is the refund object definition
type RefundDefinition struct{ def *object.Definition }

// Definition returns the RefundDefinition structure
func (d *RefundDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the RefundDefinition
func (d *RefundDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the Refund structure contained within the RefundDefinition
func (d *RefundDefinition) Spec() *Refund { return d.def.Spec.(*Refund) }

// ScheduleRefund is a single schedule item for a single refund attempt
type ScheduleRefund struct {
	// TimeDelay is the amount of time to wait since the last refund attempt
	TimeDelay time.Duration `json:"timeDelay" yaml:"timeDelay" validate:"gte=0"`
}
