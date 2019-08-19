package scheduler

import "github.com/chargehive/configuration/object"

type AttemptScheduler struct {
	DefaultSchedule Schedule
	Schedules       map[int]Schedule // attempt number > Schedule
}

const KindAttemptScheduler object.Kind = "AttemptScheduler"

func (i AttemptScheduler) GetKind() object.Kind { return KindAttemptScheduler }
func (i AttemptScheduler) GetVersion() string   { return "v1" }
