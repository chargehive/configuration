package scheduler

import "time"

// ScheduleRefund is a single schedule item for a single refund attempt
type ScheduleRefund struct {
	// TimeDelay is the amount of time to wait since the last refund attempt
	TimeDelay time.Duration `json:"timeDelay" yaml:"timeDelay"`
}
