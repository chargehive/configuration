package scheduler

import "github.com/chargehive/configuration/object"

type OnDemandScheduler struct {
	Schedule Schedule
}

const KindOnDemandScheduler object.Kind = "OnDemandScheduler"

func (i OnDemandScheduler) GetKind() object.Kind { return KindOnDemandScheduler }
func (i OnDemandScheduler) GetVersion() string   { return "v1" }
